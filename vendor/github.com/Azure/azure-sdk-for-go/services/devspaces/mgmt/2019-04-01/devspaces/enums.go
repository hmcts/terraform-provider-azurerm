package devspaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// InstanceType enumerates the values for instance type.
type InstanceType string

const (
	// InstanceTypeKubernetes ...
	InstanceTypeKubernetes InstanceType = "Kubernetes"
	// InstanceTypeOrchestratorSpecificConnectionDetails ...
	InstanceTypeOrchestratorSpecificConnectionDetails InstanceType = "OrchestratorSpecificConnectionDetails"
)

// PossibleInstanceTypeValues returns an array of possible values for the InstanceType const type.
func PossibleInstanceTypeValues() []InstanceType {
	return []InstanceType{InstanceTypeKubernetes, InstanceTypeOrchestratorSpecificConnectionDetails}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleted ...
	Deleted ProvisioningState = "Deleted"
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
	return []ProvisioningState{Canceled, Creating, Deleted, Deleting, Failed, Succeeded, Updating}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Standard ...
	Standard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{Standard}
}
