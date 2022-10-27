// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: certcounter/v1/certificates.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// buf:lint:ignore ENUM_VALUE_UPPER_SNAKE_CASE
// buf:lint:ignore ENUM_VALUE_PREFIX
type DNSProvider_DNSProvider int32

const (
	DNSProvider_DNS_PROVIDER_UNSPECIFIED DNSProvider_DNSProvider = 0
	DNSProvider_gcloud                   DNSProvider_DNSProvider = 1
)

// Enum value maps for DNSProvider_DNSProvider.
var (
	DNSProvider_DNSProvider_name = map[int32]string{
		0: "DNS_PROVIDER_UNSPECIFIED",
		1: "gcloud",
	}
	DNSProvider_DNSProvider_value = map[string]int32{
		"DNS_PROVIDER_UNSPECIFIED": 0,
		"gcloud":                   1,
	}
)

func (x DNSProvider_DNSProvider) Enum() *DNSProvider_DNSProvider {
	p := new(DNSProvider_DNSProvider)
	*p = x
	return p
}

func (x DNSProvider_DNSProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DNSProvider_DNSProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_certcounter_v1_certificates_proto_enumTypes[0].Descriptor()
}

func (DNSProvider_DNSProvider) Type() protoreflect.EnumType {
	return &file_certcounter_v1_certificates_proto_enumTypes[0]
}

func (x DNSProvider_DNSProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DNSProvider_DNSProvider.Descriptor instead.
func (DNSProvider_DNSProvider) EnumDescriptor() ([]byte, []int) {
	return file_certcounter_v1_certificates_proto_rawDescGZIP(), []int{0, 0}
}

// buf:lint:ignore ENUM_VALUE_UPPER_SNAKE_CASE
// buf:lint:ignore ENUM_VALUE_PREFIX
type CertificatesServiceIssueRequest_DNSProvider int32

const (
	CertificatesServiceIssueRequest_DNS_PROVIDER_UNSPECIFIED CertificatesServiceIssueRequest_DNSProvider = 0
	CertificatesServiceIssueRequest_gcloud                   CertificatesServiceIssueRequest_DNSProvider = 1
)

// Enum value maps for CertificatesServiceIssueRequest_DNSProvider.
var (
	CertificatesServiceIssueRequest_DNSProvider_name = map[int32]string{
		0: "DNS_PROVIDER_UNSPECIFIED",
		1: "gcloud",
	}
	CertificatesServiceIssueRequest_DNSProvider_value = map[string]int32{
		"DNS_PROVIDER_UNSPECIFIED": 0,
		"gcloud":                   1,
	}
)

func (x CertificatesServiceIssueRequest_DNSProvider) Enum() *CertificatesServiceIssueRequest_DNSProvider {
	p := new(CertificatesServiceIssueRequest_DNSProvider)
	*p = x
	return p
}

func (x CertificatesServiceIssueRequest_DNSProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CertificatesServiceIssueRequest_DNSProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_certcounter_v1_certificates_proto_enumTypes[1].Descriptor()
}

func (CertificatesServiceIssueRequest_DNSProvider) Type() protoreflect.EnumType {
	return &file_certcounter_v1_certificates_proto_enumTypes[1]
}

func (x CertificatesServiceIssueRequest_DNSProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CertificatesServiceIssueRequest_DNSProvider.Descriptor instead.
func (CertificatesServiceIssueRequest_DNSProvider) EnumDescriptor() ([]byte, []int) {
	return file_certcounter_v1_certificates_proto_rawDescGZIP(), []int{1, 0}
}

type DNSProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DNSProvider) Reset() {
	*x = DNSProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certcounter_v1_certificates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSProvider) ProtoMessage() {}

func (x *DNSProvider) ProtoReflect() protoreflect.Message {
	mi := &file_certcounter_v1_certificates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSProvider.ProtoReflect.Descriptor instead.
func (*DNSProvider) Descriptor() ([]byte, []int) {
	return file_certcounter_v1_certificates_proto_rawDescGZIP(), []int{0}
}

type CertificatesServiceIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VaultProvider               CertificatesServiceIssueRequest_DNSProvider `protobuf:"varint,1,opt,name=vault_provider,json=vaultProvider,proto3,enum=certcounter.v1.CertificatesServiceIssueRequest_DNSProvider" json:"vault_provider,omitempty"`
	AcmeAccountKeyVaultResource string                                      `protobuf:"bytes,2,opt,name=acme_account_key_vault_resource,json=acmeAccountKeyVaultResource,proto3" json:"acme_account_key_vault_resource,omitempty"`
	PrivateKeyVaultResource     string                                      `protobuf:"bytes,3,opt,name=private_key_vault_resource,json=privateKeyVaultResource,proto3" json:"private_key_vault_resource,omitempty"`
	CertificateVaultResource    string                                      `protobuf:"bytes,4,opt,name=certificate_vault_resource,json=certificateVaultResource,proto3" json:"certificate_vault_resource,omitempty"`
	RenewPrivateKey             bool                                        `protobuf:"varint,5,opt,name=renew_private_key,json=renewPrivateKey,proto3" json:"renew_private_key,omitempty"`
	KeyAlgorithm                string                                      `protobuf:"bytes,6,opt,name=key_algorithm,json=keyAlgorithm,proto3" json:"key_algorithm,omitempty"`
	DnsProvider                 string                                      `protobuf:"bytes,7,opt,name=dns_provider,json=dnsProvider,proto3" json:"dns_provider,omitempty"`
	DnsProviderId               string                                      `protobuf:"bytes,8,opt,name=dns_provider_id,json=dnsProviderId,proto3" json:"dns_provider_id,omitempty"`
	TermsOfServiceAgreed        bool                                        `protobuf:"varint,9,opt,name=terms_of_service_agreed,json=termsOfServiceAgreed,proto3" json:"terms_of_service_agreed,omitempty"`
	Email                       string                                      `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	ThresholdOfDaysToExpire     int64                                       `protobuf:"varint,11,opt,name=threshold_of_days_to_expire,json=thresholdOfDaysToExpire,proto3" json:"threshold_of_days_to_expire,omitempty"`
	Domains                     []string                                    `protobuf:"bytes,12,rep,name=domains,proto3" json:"domains,omitempty"`
	Staging                     bool                                        `protobuf:"varint,13,opt,name=staging,proto3" json:"staging,omitempty"`
}

func (x *CertificatesServiceIssueRequest) Reset() {
	*x = CertificatesServiceIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certcounter_v1_certificates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificatesServiceIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificatesServiceIssueRequest) ProtoMessage() {}

func (x *CertificatesServiceIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_certcounter_v1_certificates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificatesServiceIssueRequest.ProtoReflect.Descriptor instead.
func (*CertificatesServiceIssueRequest) Descriptor() ([]byte, []int) {
	return file_certcounter_v1_certificates_proto_rawDescGZIP(), []int{1}
}

func (x *CertificatesServiceIssueRequest) GetVaultProvider() CertificatesServiceIssueRequest_DNSProvider {
	if x != nil {
		return x.VaultProvider
	}
	return CertificatesServiceIssueRequest_DNS_PROVIDER_UNSPECIFIED
}

func (x *CertificatesServiceIssueRequest) GetAcmeAccountKeyVaultResource() string {
	if x != nil {
		return x.AcmeAccountKeyVaultResource
	}
	return ""
}

func (x *CertificatesServiceIssueRequest) GetPrivateKeyVaultResource() string {
	if x != nil {
		return x.PrivateKeyVaultResource
	}
	return ""
}

func (x *CertificatesServiceIssueRequest) GetCertificateVaultResource() string {
	if x != nil {
		return x.CertificateVaultResource
	}
	return ""
}

func (x *CertificatesServiceIssueRequest) GetRenewPrivateKey() bool {
	if x != nil {
		return x.RenewPrivateKey
	}
	return false
}

func (x *CertificatesServiceIssueRequest) GetKeyAlgorithm() string {
	if x != nil {
		return x.KeyAlgorithm
	}
	return ""
}

func (x *CertificatesServiceIssueRequest) GetDnsProvider() string {
	if x != nil {
		return x.DnsProvider
	}
	return ""
}

func (x *CertificatesServiceIssueRequest) GetDnsProviderId() string {
	if x != nil {
		return x.DnsProviderId
	}
	return ""
}

func (x *CertificatesServiceIssueRequest) GetTermsOfServiceAgreed() bool {
	if x != nil {
		return x.TermsOfServiceAgreed
	}
	return false
}

func (x *CertificatesServiceIssueRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CertificatesServiceIssueRequest) GetThresholdOfDaysToExpire() int64 {
	if x != nil {
		return x.ThresholdOfDaysToExpire
	}
	return 0
}

func (x *CertificatesServiceIssueRequest) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *CertificatesServiceIssueRequest) GetStaging() bool {
	if x != nil {
		return x.Staging
	}
	return false
}

type CertificatesServiceIssueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateKeyVaultVersionResource  string `protobuf:"bytes,1,opt,name=private_key_vault_version_resource,json=privateKeyVaultVersionResource,proto3" json:"private_key_vault_version_resource,omitempty"`
	CertificateVaultVersionResource string `protobuf:"bytes,2,opt,name=certificate_vault_version_resource,json=certificateVaultVersionResource,proto3" json:"certificate_vault_version_resource,omitempty"`
}

func (x *CertificatesServiceIssueResponse) Reset() {
	*x = CertificatesServiceIssueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certcounter_v1_certificates_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificatesServiceIssueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificatesServiceIssueResponse) ProtoMessage() {}

func (x *CertificatesServiceIssueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_certcounter_v1_certificates_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificatesServiceIssueResponse.ProtoReflect.Descriptor instead.
func (*CertificatesServiceIssueResponse) Descriptor() ([]byte, []int) {
	return file_certcounter_v1_certificates_proto_rawDescGZIP(), []int{2}
}

func (x *CertificatesServiceIssueResponse) GetPrivateKeyVaultVersionResource() string {
	if x != nil {
		return x.PrivateKeyVaultVersionResource
	}
	return ""
}

func (x *CertificatesServiceIssueResponse) GetCertificateVaultVersionResource() string {
	if x != nil {
		return x.CertificateVaultVersionResource
	}
	return ""
}

var File_certcounter_v1_certificates_proto protoreflect.FileDescriptor

var file_certcounter_v1_certificates_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x65, 0x72, 0x74, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x65, 0x72, 0x74, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0b,
	0x44, 0x4e, 0x53, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x0b, 0x44,
	0x4e, 0x53, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x4e,
	0x53, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x67, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x10, 0x01, 0x22, 0xe2, 0x06, 0x0a, 0x1f, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x6c, 0x0a, 0x0e, 0x76, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x3b, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x44, 0x4e, 0x53, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x18, 0x01, 0x52, 0x0d, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x1f, 0x61, 0x63, 0x6d, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x1b, 0x61, 0x63, 0x6d, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x1a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x17, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x1a, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x18, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x72,
	0x65, 0x6e, 0x65, 0x77, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x5b,
	0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xfa, 0x42, 0x33, 0x72, 0x31, 0x52, 0x00, 0x52, 0x07,
	0x72, 0x73, 0x61, 0x32, 0x30, 0x34, 0x38, 0x52, 0x07, 0x72, 0x73, 0x61, 0x34, 0x30, 0x39, 0x36,
	0x52, 0x07, 0x72, 0x73, 0x61, 0x38, 0x31, 0x39, 0x32, 0x52, 0x08, 0x65, 0x63, 0x64, 0x73, 0x61,
	0x32, 0x35, 0x36, 0x52, 0x08, 0x65, 0x63, 0x64, 0x73, 0x61, 0x33, 0x38, 0x34, 0x52, 0x0c, 0x6b,
	0x65, 0x79, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x30, 0x0a, 0x0c, 0x64,
	0x6e, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x52, 0x06, 0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x52, 0x0b, 0x64, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a,
	0x0f, 0x64, 0x6e, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x0d, 0x64, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x35,
	0x0a, 0x17, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x14, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x67, 0x72, 0x65, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3c, 0x0a, 0x1b, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x5f, 0x6f, 0x66, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x74, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x73, 0x54, 0x6f, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x07, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x22, 0x37, 0x0a, 0x0b, 0x44, 0x4e, 0x53, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x18, 0x44, 0x4e, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x10, 0x01, 0x22, 0xbb, 0x01, 0x0a, 0x20, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a,
	0x0a, 0x22, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1e, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x22, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x32, 0xa9, 0x01, 0x0a, 0x13, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x91, 0x01, 0x0a, 0x05, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x2f, 0x2e, 0x63, 0x65, 0x72, 0x74,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x65, 0x72,
	0x74, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_certcounter_v1_certificates_proto_rawDescOnce sync.Once
	file_certcounter_v1_certificates_proto_rawDescData = file_certcounter_v1_certificates_proto_rawDesc
)

func file_certcounter_v1_certificates_proto_rawDescGZIP() []byte {
	file_certcounter_v1_certificates_proto_rawDescOnce.Do(func() {
		file_certcounter_v1_certificates_proto_rawDescData = protoimpl.X.CompressGZIP(file_certcounter_v1_certificates_proto_rawDescData)
	})
	return file_certcounter_v1_certificates_proto_rawDescData
}

var file_certcounter_v1_certificates_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_certcounter_v1_certificates_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_certcounter_v1_certificates_proto_goTypes = []interface{}{
	(DNSProvider_DNSProvider)(0),                     // 0: certcounter.v1.DNSProvider.DNSProvider
	(CertificatesServiceIssueRequest_DNSProvider)(0), // 1: certcounter.v1.CertificatesServiceIssueRequest.DNSProvider
	(*DNSProvider)(nil),                              // 2: certcounter.v1.DNSProvider
	(*CertificatesServiceIssueRequest)(nil),          // 3: certcounter.v1.CertificatesServiceIssueRequest
	(*CertificatesServiceIssueResponse)(nil),         // 4: certcounter.v1.CertificatesServiceIssueResponse
}
var file_certcounter_v1_certificates_proto_depIdxs = []int32{
	1, // 0: certcounter.v1.CertificatesServiceIssueRequest.vault_provider:type_name -> certcounter.v1.CertificatesServiceIssueRequest.DNSProvider
	3, // 1: certcounter.v1.CertificatesService.Issue:input_type -> certcounter.v1.CertificatesServiceIssueRequest
	4, // 2: certcounter.v1.CertificatesService.Issue:output_type -> certcounter.v1.CertificatesServiceIssueResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_certcounter_v1_certificates_proto_init() }
func file_certcounter_v1_certificates_proto_init() {
	if File_certcounter_v1_certificates_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_certcounter_v1_certificates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSProvider); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certcounter_v1_certificates_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificatesServiceIssueRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certcounter_v1_certificates_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificatesServiceIssueResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_certcounter_v1_certificates_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_certcounter_v1_certificates_proto_goTypes,
		DependencyIndexes: file_certcounter_v1_certificates_proto_depIdxs,
		EnumInfos:         file_certcounter_v1_certificates_proto_enumTypes,
		MessageInfos:      file_certcounter_v1_certificates_proto_msgTypes,
	}.Build()
	File_certcounter_v1_certificates_proto = out.File
	file_certcounter_v1_certificates_proto_rawDesc = nil
	file_certcounter_v1_certificates_proto_goTypes = nil
	file_certcounter_v1_certificates_proto_depIdxs = nil
}
