SHELL           := /usr/bin/env bash -Eeu -o pipefail
GITROOT         := $(shell git rev-parse --show-toplevel || pwd || echo '.')
MAKEFILE_DIR    := $(subst /,,$(dir $(lastword ${MAKEFILE_LIST})))
PRE_PUSH        := ${GITROOT}/.git/hooks/pre-push
GOMODULE        := github.com/kunitsuinc/certcounter
VERSION         := $(shell git describe --tags --abbrev=0 --always)
REVISION        := $(shell git log -1 --format='%H')
BRANCH          := $(shell git rev-parse --abbrev-ref HEAD)
TIMESTAMP       := $(shell git log -1 --format='%cI')
GO_BUILD_OUTPUT := ./.local/bin/certcounter
GO_BUILD        := go build -o ${GO_BUILD_OUTPUT} -ldflags "-X ${GOMODULE}/pkg/config.version=${VERSION} -X ${GOMODULE}/pkg/config.revision=${REVISION} -X ${GOMODULE}/pkg/config.branch=${BRANCH} -X ${GOMODULE}/pkg/config.timestamp=${TIMESTAMP}" ./cmd/certcounter/...
IMAGE_TAG       := ${REVISION}
LOCAL_CR        := ${GOMODULE}

export PATH := ${GITROOT}/.local/bin:${GITROOT}/.bin:${PATH}

.DEFAULT_GOAL := help
.PHONY: help
help: githooks ## display this help documents
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: githooks ## Setup tools for development
	# == SETUP =====================================================
	# direnv
	direnv allow .
	# buf
	buf --version
	# golangci-lint
	golangci-lint --version
	# install-protoc-gen
	make install-protoc-gen
	# --------------------------------------------------------------

.PHONY: githooks
githooks:
	@[[ -f "${PRE_PUSH}" ]] || cp -aiv "${GITROOT}/.githooks/pre-push" "${PRE_PUSH}"

.PHONY: install-protoc-gen
install-protoc-gen:
	GOBIN="${GITROOT}/.local/bin" go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/envoyproxy/protoc-gen-validate \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

.PHONY: buf-mod-update
buf-mod-update: ## Run buf mod update
	buf --debug --verbose mod update

.PHONY: buf
buf: ## Run buf generate
	buf --debug --verbose generate

.PHONY: gogenerate
gogenerate:  ## Run go generate
	go generate ./...

.PHONY: generate
generate: buf gogenerate ## Generate files

.PHONY: credits
credits:  ## Generate CREDITS file.
	command -v gocredits || go install github.com/Songmu/gocredits/cmd/gocredits@latest
	gocredits . > CREDITS
	git diff --exit-code

.PHONY: clean
clean:  ## Clean up cache, etc
	# go build cache
	go env GOCACHE
	go clean -x -cache -testcache -modcache -fuzzcache
	# golangci-lint cache
	golangci-lint cache status
	golangci-lint cache clean

.PHONY: lint
lint:  ## Run secretlint, go mod tidy, golangci-lint
	# tidy
	go mod tidy
	git diff --exit-code go.mod go.sum
	# buf
	buf lint
	# lint
	# ref. https://golangci-lint.run/usage/linters/
	golangci-lint run --fix --sort-results
	git diff --exit-code
	# ref. https://github.com/secretlint/secretlint
	docker run -v "`pwd`:`pwd`" -w "`pwd`" --rm secretlint/secretlint secretlint "**/*"

.PHONY: test
test: githooks ## Run go test and display coverage
	# test
	go test -v -race -p=4 -parallel=8 -timeout=300s -cover -coverprofile=./coverage.txt ./...
	go tool cover -func=./coverage.txt

.PHONY: ci
ci: credits lint test ## CI command set

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
	${GO_BUILD}

.PHONY: run
run: gobuild ## Run go build and exec
	direnv exec . ${GO_BUILD_OUTPUT}

.PHONY: runjq
runjq: gobuild ## Run go build and exec with jq
	make run | jqlog

.PHONY: air
air:  ## Run air
	@[[ -x "${GITROOT}/.local/bin/air" ]] || bash -cx "curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b ${GITROOT}/.local/bin && chmod +x ${GITROOT}/.local/bin/air"
	air -tmp_dir ./.tmp -log.time true -build.exclude_regex "_test\.go" -build.cmd '${GO_BUILD}' -build.bin 'direnv exec . ${GO_BUILD_OUTPUT}'

.PHONY: airjq
airjq:  ## Run air with jq
	make air | ./.bin/jqlog

.PHONY: build
build:  ## docker build -t ${LOCAL_CR}:${IMAGE_TAG}
	docker build --platform linux/amd64 -t ${LOCAL_CR}:${IMAGE_TAG} --build-arg VERSION=${VERSION} --build-arg REVISION=${REVISION} --build-arg BRANCH=${BRANCH} --build-arg TIMESTAMP=${TIMESTAMP} .

.PHONY: push
push:  ## docker push ${LOCAL_CR}:${IMAGE_TAG} as ${REMOTE_CR}:${IMAGE_TAG}
	docker tag ${LOCAL_CR}:${IMAGE_TAG} ${REMOTE_CR}:${IMAGE_TAG}
	docker push ${REMOTE_CR}:${IMAGE_TAG}

.PHONY: build-push
build-push: build push ## docker build and docker push
