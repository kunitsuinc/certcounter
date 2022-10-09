SHELL        := /usr/bin/env bash -Eeu -o pipefail
GITROOT      := $(shell git rev-parse --show-toplevel || pwd || echo '.')
MAKEFILE_DIR := $(subst /,,$(dir $(lastword ${MAKEFILE_LIST})))
PRE_PUSH     := ${GITROOT}/.git/hooks/pre-push
GOMODULE     := github.com/kunitsuinc/certcounter
VERSION      := $(shell git describe --tags --abbrev=0 --always)
REVISION     := $(shell git log -1 --format='%H')
BRANCH       := $(shell git rev-parse --abbrev-ref HEAD)
TIMESTAMP    := $(shell git log -1 --format='%cI')
IMAGE_TAG    := ${REVISION}
LOCAL_CR     := ${GOMODULE}

.DEFAULT_GOAL := help
.PHONY: help
help: githooks ## display this help documents
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'

.PHONY: githooks
githooks:
	@test -f "${PRE_PUSH}" || cp -aiv "${GITROOT}/.githooks/pre-push" "${PRE_PUSH}"

.PHONY: protocgens
protocgens:
	GOBIN="${GITROOT}/.local/bin" go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/envoyproxy/protoc-gen-validate

.PHONY: buf-mod-update
buf-mod-update: ## Run buf mod update
	cd proto && ${GITROOT}/.bin/buf --debug --verbose mod update

.PHONY: buf
buf: ## Run buf generate
	cd proto && ${GITROOT}/.bin/buf --debug --verbose generate

.PHONY: lint
lint:  ## Run golangci-lint after go mod tidy.
	# tidy
	go mod tidy
	git diff --exit-code go.mod go.sum
	# buf
	buf lint ./proto
	# lint
	# cf. https://golangci-lint.run/usage/linters/
	./.bin/golangci-lint run --fix --sort-results
	git diff --exit-code

.PHONY: setup
setup: githooks protocgens ## Setup tools for development

.PHONY: credits
credits:  ## Generate CREDITS file.
	command -v gocredits || go install github.com/Songmu/gocredits/cmd/gocredits@latest
	gocredits . > CREDITS
	git diff --exit-code

.PHONY: test
test: githooks ## Run go test and display coverage
	# test
	go test -v -race -p=4 -parallel=8 -timeout=300s -cover -coverprofile=./coverage.txt ./...
	go tool cover -func=./coverage.txt

.PHONY: ci
ci: lint credits test ## CI command set

.PHONY: goxz
goxz: ci ## Run goxz for release files
	command -v goxz || go install github.com/Songmu/goxz/cmd/goxz@latest
	goxz -d ./.tmp -os=linux,darwin,windows -arch=amd64,arm64 -pv ${VERSION} -build-ldflags "-X ${GOMODULE}/pkg/config.version=${VERSION} -X ${GOMODULE}/pkg/config.revision=${REVISION} -X ${GOMODULE}/pkg/config.branch=${BRANCH} -X ${GOMODULE}/pkg/config.timestamp=${TIMESTAMP}" ./cmd/certcounter

.PHONY: up
up:  ## docker compose up background
	docker compose up -d

.PHONY: down
down:  ## docker compose down and remove image and volume
	docker compose down --rmi all --volumes --remove-orphans

.PHONY: restart
restart: down up ## docker compose restart

.PHONY: logs
logs:  ## docker compose logs -f
	@printf '[\033[36mNOTICE\033[0m] %s\n' "If back prompt, enter Ctrl+C"
	docker compose logs -f

.PHONY: gobuild
gobuild: ## Run go build
	go build -o .local/bin/certcounter -ldflags "-X ${GOMODULE}/pkg/config.version=${VERSION} -X ${GOMODULE}/pkg/config.revision=${REVISION} -X ${GOMODULE}/pkg/config.branch=${BRANCH} -X ${GOMODULE}/pkg/config.timestamp=${TIMESTAMP}" ./cmd/certcounter/...

.PHONY: run
run: gobuild ## Run go build and exec
	# ./.local/bin/certcounter | ./.bin/jqlog
	./.local/bin/certcounter

.PHONY: air
air:  ## Run air
	air --build.cmd 'go build -o .local/bin/certcounter -ldflags "-X ${GOMODULE}/pkg/config.version=${VERSION} -X ${GOMODULE}/pkg/config.revision=${REVISION} -X ${GOMODULE}/pkg/config.branch=${BRANCH} -X ${GOMODULE}/pkg/config.timestamp=${TIMESTAMP}" ./cmd/certcounter/...' --build.bin './.local/bin'

.PHONY: build
build:  ## docker build -t ${LOCAL_CR}:${IMAGE_TAG}
	docker build --platform linux/amd64 -t ${LOCAL_CR}:${IMAGE_TAG} --build-arg VERSION=${VERSION} --build-arg REVISION=${REVISION} --build-arg BRANCH=${BRANCH} --build-arg TIMESTAMP=${TIMESTAMP} .

.PHONY: push
push:  ## docker push ${LOCAL_CR}:${IMAGE_TAG} as ${REMOTE_CR}:${IMAGE_TAG}
	docker tag ${LOCAL_CR}:${IMAGE_TAG} ${REMOTE_CR}:${IMAGE_TAG}
	docker push ${REMOTE_CR}:${IMAGE_TAG}

.PHONY: build-push
build-push: build push ## docker build and docker push
