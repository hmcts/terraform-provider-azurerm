package timeseriesinsights

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessPolicyRole enumerates the values for access policy role.
type AccessPolicyRole string

const (
	// Contributor ...
	Contributor AccessPolicyRole = "Contributor"
	// Reader ...
	Reader AccessPolicyRole = "Reader"
)

// PossibleAccessPolicyRoleValues returns an array of possible values for the AccessPolicyRole const type.
func PossibleAccessPolicyRoleValues() []AccessPolicyRole {
	return []AccessPolicyRole{Contributor, Reader}
}

// DataStringComparisonBehavior enumerates the values for data string comparison behavior.
type DataStringComparisonBehavior string

const (
	// Ordinal ...
	Ordinal DataStringComparisonBehavior = "Ordinal"
	// OrdinalIgnoreCase ...
	OrdinalIgnoreCase DataStringComparisonBehavior = "OrdinalIgnoreCase"
)

// PossibleDataStringComparisonBehaviorValues returns an array of possible values for the DataStringComparisonBehavior const type.
func PossibleDataStringComparisonBehaviorValues() []DataStringComparisonBehavior {
	return []DataStringComparisonBehavior{Ordinal, OrdinalIgnoreCase}
}

// IngressState enumerates the values for ingress state.
type IngressState string

const (
	// Disabled ...
	Disabled IngressState = "Disabled"
	// Paused ...
	Paused IngressState = "Paused"
	// Ready ...
	Ready IngressState = "Ready"
	// Running ...
	Running IngressState = "Running"
	// Unknown ...
	Unknown IngressState = "Unknown"
)

// PossibleIngressStateValues returns an array of possible values for the IngressState const type.
func PossibleIngressStateValues() []IngressState {
	return []IngressState{Disabled, Paused, Ready, Running, Unknown}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// KindEnvironmentCreateOrUpdateParameters ...
	KindEnvironmentCreateOrUpdateParameters Kind = "EnvironmentCreateOrUpdateParameters"
	// KindGen1 ...
	KindGen1 Kind = "Gen1"
	// KindGen2 ...
	KindGen2 Kind = "Gen2"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{KindEnvironmentCreateOrUpdateParameters, KindGen1, KindGen2}
}

// KindBasicEnvironmentResource enumerates the values for kind basic environment resource.
type KindBasicEnvironmentResource string

const (
	// KindBasicEnvironmentResourceKindEnvironmentResource ...
	KindBasicEnvironmentResourceKindEnvironmentResource KindBasicEnvironmentResource = "EnvironmentResource"
	// KindBasicEnvironmentResourceKindGen1 ...
	KindBasicEnvironmentResourceKindGen1 KindBasicEnvironmentResource = "Gen1"
	// KindBasicEnvironmentResourceKindGen2 ...
	KindBasicEnvironmentResourceKindGen2 KindBasicEnvironmentResource = "Gen2"
)

// PossibleKindBasicEnvironmentResourceValues returns an array of possible values for the KindBasicEnvironmentResource const type.
func PossibleKindBasicEnvironmentResourceValues() []KindBasicEnvironmentResource {
	return []KindBasicEnvironmentResource{KindBasicEnvironmentResourceKindEnvironmentResource, KindBasicEnvironmentResourceKindGen1, KindBasicEnvironmentResourceKindGen2}
}

// KindBasicEventSourceCreateOrUpdateParameters enumerates the values for kind basic event source create or
// update parameters.
type KindBasicEventSourceCreateOrUpdateParameters string

const (
	// KindEventSourceCreateOrUpdateParameters ...
	KindEventSourceCreateOrUpdateParameters KindBasicEventSourceCreateOrUpdateParameters = "EventSourceCreateOrUpdateParameters"
	// KindMicrosoftEventHub ...
	KindMicrosoftEventHub KindBasicEventSourceCreateOrUpdateParameters = "Microsoft.EventHub"
	// KindMicrosoftIoTHub ...
	KindMicrosoftIoTHub KindBasicEventSourceCreateOrUpdateParameters = "Microsoft.IoTHub"
)

// PossibleKindBasicEventSourceCreateOrUpdateParametersValues returns an array of possible values for the KindBasicEventSourceCreateOrUpdateParameters const type.
func PossibleKindBasicEventSourceCreateOrUpdateParametersValues() []KindBasicEventSourceCreateOrUpdateParameters {
	return []KindBasicEventSourceCreateOrUpdateParameters{KindEventSourceCreateOrUpdateParameters, KindMicrosoftEventHub, KindMicrosoftIoTHub}
}

// KindBasicEventSourceResource enumerates the values for kind basic event source resource.
type KindBasicEventSourceResource string

const (
	// KindBasicEventSourceResourceKindEventSourceResource ...
	KindBasicEventSourceResourceKindEventSourceResource KindBasicEventSourceResource = "EventSourceResource"
	// KindBasicEventSourceResourceKindMicrosoftEventHub ...
	KindBasicEventSourceResourceKindMicrosoftEventHub KindBasicEventSourceResource = "Microsoft.EventHub"
	// KindBasicEventSourceResourceKindMicrosoftIoTHub ...
	KindBasicEventSourceResourceKindMicrosoftIoTHub KindBasicEventSourceResource = "Microsoft.IoTHub"
)

// PossibleKindBasicEventSourceResourceValues returns an array of possible values for the KindBasicEventSourceResource const type.
func PossibleKindBasicEventSourceResourceValues() []KindBasicEventSourceResource {
	return []KindBasicEventSourceResource{KindBasicEventSourceResourceKindEventSourceResource, KindBasicEventSourceResourceKindMicrosoftEventHub, KindBasicEventSourceResourceKindMicrosoftIoTHub}
}

// LocalTimestampFormat enumerates the values for local timestamp format.
type LocalTimestampFormat string

const (
	// Embedded ...
	Embedded LocalTimestampFormat = "Embedded"
)

// PossibleLocalTimestampFormatValues returns an array of possible values for the LocalTimestampFormat const type.
func PossibleLocalTimestampFormatValues() []LocalTimestampFormat {
	return []LocalTimestampFormat{Embedded}
}

// PropertyType enumerates the values for property type.
type PropertyType string

const (
	// String ...
	String PropertyType = "String"
)

// PossiblePropertyTypeValues returns an array of possible values for the PropertyType const type.
func PossiblePropertyTypeValues() []PropertyType {
	return []PropertyType{String}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Accepted ...
	Accepted ProvisioningState = "Accepted"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Accepted, Creating, Deleting, Failed, Succeeded, Updating}
}

// ReferenceDataKeyPropertyType enumerates the values for reference data key property type.
type ReferenceDataKeyPropertyType string

const (
	// ReferenceDataKeyPropertyTypeBool ...
	ReferenceDataKeyPropertyTypeBool ReferenceDataKeyPropertyType = "Bool"
	// ReferenceDataKeyPropertyTypeDateTime ...
	ReferenceDataKeyPropertyTypeDateTime ReferenceDataKeyPropertyType = "DateTime"
	// ReferenceDataKeyPropertyTypeDouble ...
	ReferenceDataKeyPropertyTypeDouble ReferenceDataKeyPropertyType = "Double"
	// ReferenceDataKeyPropertyTypeString ...
	ReferenceDataKeyPropertyTypeString ReferenceDataKeyPropertyType = "String"
)

// PossibleReferenceDataKeyPropertyTypeValues returns an array of possible values for the ReferenceDataKeyPropertyType const type.
func PossibleReferenceDataKeyPropertyTypeValues() []ReferenceDataKeyPropertyType {
	return []ReferenceDataKeyPropertyType{ReferenceDataKeyPropertyTypeBool, ReferenceDataKeyPropertyTypeDateTime, ReferenceDataKeyPropertyTypeDouble, ReferenceDataKeyPropertyTypeString}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// L1 ...
	L1 SkuName = "L1"
	// P1 ...
	P1 SkuName = "P1"
	// S1 ...
	S1 SkuName = "S1"
	// S2 ...
	S2 SkuName = "S2"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{L1, P1, S1, S2}
}

// StorageLimitExceededBehavior enumerates the values for storage limit exceeded behavior.
type StorageLimitExceededBehavior string

const (
	// PauseIngress ...
	PauseIngress StorageLimitExceededBehavior = "PauseIngress"
	// PurgeOldData ...
	PurgeOldData StorageLimitExceededBehavior = "PurgeOldData"
)

// PossibleStorageLimitExceededBehaviorValues returns an array of possible values for the StorageLimitExceededBehavior const type.
func PossibleStorageLimitExceededBehaviorValues() []StorageLimitExceededBehavior {
	return []StorageLimitExceededBehavior{PauseIngress, PurgeOldData}
}

// WarmStoragePropertiesState enumerates the values for warm storage properties state.
type WarmStoragePropertiesState string

const (
	// WarmStoragePropertiesStateError ...
	WarmStoragePropertiesStateError WarmStoragePropertiesState = "Error"
	// WarmStoragePropertiesStateOk ...
	WarmStoragePropertiesStateOk WarmStoragePropertiesState = "Ok"
	// WarmStoragePropertiesStateUnknown ...
	WarmStoragePropertiesStateUnknown WarmStoragePropertiesState = "Unknown"
)

// PossibleWarmStoragePropertiesStateValues returns an array of possible values for the WarmStoragePropertiesState const type.
func PossibleWarmStoragePropertiesStateValues() []WarmStoragePropertiesState {
	return []WarmStoragePropertiesState{WarmStoragePropertiesStateError, WarmStoragePropertiesStateOk, WarmStoragePropertiesStateUnknown}
}
