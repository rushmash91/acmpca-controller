// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type AccessMethodType string

const (
	AccessMethodType_CA_REPOSITORY         AccessMethodType = "CA_REPOSITORY"
	AccessMethodType_RESOURCE_PKI_MANIFEST AccessMethodType = "RESOURCE_PKI_MANIFEST"
	AccessMethodType_RESOURCE_PKI_NOTIFY   AccessMethodType = "RESOURCE_PKI_NOTIFY"
)

type ActionType string

const (
	ActionType_GetCertificate   ActionType = "GetCertificate"
	ActionType_IssueCertificate ActionType = "IssueCertificate"
	ActionType_ListPermissions  ActionType = "ListPermissions"
)

type AuditReportResponseFormat string

const (
	AuditReportResponseFormat_CSV  AuditReportResponseFormat = "CSV"
	AuditReportResponseFormat_JSON AuditReportResponseFormat = "JSON"
)

type AuditReportStatus string

const (
	AuditReportStatus_CREATING AuditReportStatus = "CREATING"
	AuditReportStatus_FAILED   AuditReportStatus = "FAILED"
	AuditReportStatus_SUCCESS  AuditReportStatus = "SUCCESS"
)

type CertificateAuthorityStatus_SDK string

const (
	CertificateAuthorityStatus_SDK_ACTIVE              CertificateAuthorityStatus_SDK = "ACTIVE"
	CertificateAuthorityStatus_SDK_CREATING            CertificateAuthorityStatus_SDK = "CREATING"
	CertificateAuthorityStatus_SDK_DELETED             CertificateAuthorityStatus_SDK = "DELETED"
	CertificateAuthorityStatus_SDK_DISABLED            CertificateAuthorityStatus_SDK = "DISABLED"
	CertificateAuthorityStatus_SDK_EXPIRED             CertificateAuthorityStatus_SDK = "EXPIRED"
	CertificateAuthorityStatus_SDK_FAILED              CertificateAuthorityStatus_SDK = "FAILED"
	CertificateAuthorityStatus_SDK_PENDING_CERTIFICATE CertificateAuthorityStatus_SDK = "PENDING_CERTIFICATE"
)

type CertificateAuthorityType string

const (
	CertificateAuthorityType_ROOT        CertificateAuthorityType = "ROOT"
	CertificateAuthorityType_SUBORDINATE CertificateAuthorityType = "SUBORDINATE"
)

type CertificateAuthorityUsageMode string

const (
	CertificateAuthorityUsageMode_GENERAL_PURPOSE         CertificateAuthorityUsageMode = "GENERAL_PURPOSE"
	CertificateAuthorityUsageMode_SHORT_LIVED_CERTIFICATE CertificateAuthorityUsageMode = "SHORT_LIVED_CERTIFICATE"
)

type ExtendedKeyUsageType string

const (
	ExtendedKeyUsageType_CERTIFICATE_TRANSPARENCY ExtendedKeyUsageType = "CERTIFICATE_TRANSPARENCY"
	ExtendedKeyUsageType_CLIENT_AUTH              ExtendedKeyUsageType = "CLIENT_AUTH"
	ExtendedKeyUsageType_CODE_SIGNING             ExtendedKeyUsageType = "CODE_SIGNING"
	ExtendedKeyUsageType_DOCUMENT_SIGNING         ExtendedKeyUsageType = "DOCUMENT_SIGNING"
	ExtendedKeyUsageType_EMAIL_PROTECTION         ExtendedKeyUsageType = "EMAIL_PROTECTION"
	ExtendedKeyUsageType_OCSP_SIGNING             ExtendedKeyUsageType = "OCSP_SIGNING"
	ExtendedKeyUsageType_SERVER_AUTH              ExtendedKeyUsageType = "SERVER_AUTH"
	ExtendedKeyUsageType_SMART_CARD_LOGIN         ExtendedKeyUsageType = "SMART_CARD_LOGIN"
	ExtendedKeyUsageType_TIME_STAMPING            ExtendedKeyUsageType = "TIME_STAMPING"
)

type FailureReason string

const (
	FailureReason_OTHER                 FailureReason = "OTHER"
	FailureReason_REQUEST_TIMED_OUT     FailureReason = "REQUEST_TIMED_OUT"
	FailureReason_UNSUPPORTED_ALGORITHM FailureReason = "UNSUPPORTED_ALGORITHM"
)

type KeyAlgorithm string

const (
	KeyAlgorithm_EC_prime256v1 KeyAlgorithm = "EC_prime256v1"
	KeyAlgorithm_EC_secp384r1  KeyAlgorithm = "EC_secp384r1"
	KeyAlgorithm_RSA_2048      KeyAlgorithm = "RSA_2048"
	KeyAlgorithm_RSA_4096      KeyAlgorithm = "RSA_4096"
	KeyAlgorithm_SM2           KeyAlgorithm = "SM2"
)

type KeyStorageSecurityStandard string

const (
	KeyStorageSecurityStandard_CCPC_LEVEL_1_OR_HIGHER       KeyStorageSecurityStandard = "CCPC_LEVEL_1_OR_HIGHER"
	KeyStorageSecurityStandard_FIPS_140_2_LEVEL_2_OR_HIGHER KeyStorageSecurityStandard = "FIPS_140_2_LEVEL_2_OR_HIGHER"
	KeyStorageSecurityStandard_FIPS_140_2_LEVEL_3_OR_HIGHER KeyStorageSecurityStandard = "FIPS_140_2_LEVEL_3_OR_HIGHER"
)

type PolicyQualifierID string

const (
	PolicyQualifierID_CPS PolicyQualifierID = "CPS"
)

type ResourceOwner string

const (
	ResourceOwner_OTHER_ACCOUNTS ResourceOwner = "OTHER_ACCOUNTS"
	ResourceOwner_SELF           ResourceOwner = "SELF"
)

type RevocationReason string

const (
	RevocationReason_AFFILIATION_CHANGED              RevocationReason = "AFFILIATION_CHANGED"
	RevocationReason_A_A_COMPROMISE                   RevocationReason = "A_A_COMPROMISE"
	RevocationReason_CERTIFICATE_AUTHORITY_COMPROMISE RevocationReason = "CERTIFICATE_AUTHORITY_COMPROMISE"
	RevocationReason_CESSATION_OF_OPERATION           RevocationReason = "CESSATION_OF_OPERATION"
	RevocationReason_KEY_COMPROMISE                   RevocationReason = "KEY_COMPROMISE"
	RevocationReason_PRIVILEGE_WITHDRAWN              RevocationReason = "PRIVILEGE_WITHDRAWN"
	RevocationReason_SUPERSEDED                       RevocationReason = "SUPERSEDED"
	RevocationReason_UNSPECIFIED                      RevocationReason = "UNSPECIFIED"
)

type S3ObjectACL string

const (
	S3ObjectACL_BUCKET_OWNER_FULL_CONTROL S3ObjectACL = "BUCKET_OWNER_FULL_CONTROL"
	S3ObjectACL_PUBLIC_READ               S3ObjectACL = "PUBLIC_READ"
)

type SigningAlgorithm string

const (
	SigningAlgorithm_SHA256WITHECDSA SigningAlgorithm = "SHA256WITHECDSA"
	SigningAlgorithm_SHA256WITHRSA   SigningAlgorithm = "SHA256WITHRSA"
	SigningAlgorithm_SHA384WITHECDSA SigningAlgorithm = "SHA384WITHECDSA"
	SigningAlgorithm_SHA384WITHRSA   SigningAlgorithm = "SHA384WITHRSA"
	SigningAlgorithm_SHA512WITHECDSA SigningAlgorithm = "SHA512WITHECDSA"
	SigningAlgorithm_SHA512WITHRSA   SigningAlgorithm = "SHA512WITHRSA"
	SigningAlgorithm_SM3WITHSM2      SigningAlgorithm = "SM3WITHSM2"
)

type ValidityPeriodType string

const (
	ValidityPeriodType_ABSOLUTE ValidityPeriodType = "ABSOLUTE"
	ValidityPeriodType_DAYS     ValidityPeriodType = "DAYS"
	ValidityPeriodType_END_DATE ValidityPeriodType = "END_DATE"
	ValidityPeriodType_MONTHS   ValidityPeriodType = "MONTHS"
	ValidityPeriodType_YEARS    ValidityPeriodType = "YEARS"
)
