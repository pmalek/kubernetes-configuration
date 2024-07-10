package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	SchemeBuilder.Register(&KonnectAPIAuthConfiguration{}, &KonnectAPIAuthConfigurationList{})
}

// KonnectAPIAuthConfiguration is the Schema for the Konnect configuration type.
//
// +genclient
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
// +kubebuilder:object:generate=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Valid",description="The API authentication information is valid",type=string,JSONPath=`.status.conditions[?(@.type=='Valid')].status`
// +kubebuilder:printcolumn:name="OrgID",description="Konnect Organization ID this API authentication configuration belongs to.",type=string,JSONPath=`.status.organizationID`
// +kubebuilder:printcolumn:name="ServerURL",description="Configured server URL.",type=string,JSONPath=`.status.serverURL`
// +kubebuilder:validation:XValidation:rule="self.spec.token.startsWith('spat_') || self.spec.token.startsWith('kpat_')", message="Konnect tokens have to start with spat_ or kpat_"
// +kubebuilder:validation:XValidation:rule="!has(oldSelf.spec.token) || has(self.spec.token)", message="Token is required once set"
// +kubebuilder:validation:XValidation:rule="!has(oldSelf.spec.serverURL) || has(self.spec.serverURL)", message="Server URL is required once set"
type KonnectAPIAuthConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the specification of the KonnectAPIAuthConfiguration resource.
	Spec KonnectAPIAuthConfigurationSpec `json:"spec,omitempty"`

	// Status is the status of the KonnectAPIAuthConfiguration resource.
	Status KonnectAPIAuthConfigurationStatus `json:"status,omitempty"`
}

// KonnectAPIAuthConfigurationSpec is the specification of the KonnectAPIAuthConfiguration resource.
type KonnectAPIAuthConfigurationSpec struct {
	// Token is the Konnect token used to authenticate with the Konnect API.
	// +kubebuilder:validation:Required
	Token string `json:"token"`

	// ServerURL is the URL of the Konnect server.
	// TODO(pmalek): do we want this validation?
	// +Xkubebuilder:validation:Enum=us.api.konghq.com;eu.api.konghq.com;au.api.konghq.com
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Server URL is immutable"
	ServerURL string `json:"serverURL"`
}

type KonnectAPIAuthConfigurationStatus struct {
	// OrganizationID is the unique identifier of the organization in Konnect.
	OrganizationID string `json:"organizationID,omitempty"`

	// ServerURL is configured server URL.
	ServerURL string `json:"serverURL,omitempty"`

	// Conditions describe the status of the Konnect configuration.
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=8
	// +kubebuilder:default={{type: "Valid", status: "Unknown", reason:"Pending", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"}}
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// GetConditions returns the Status Conditions
func (in *KonnectAPIAuthConfigurationStatus) GetConditions() []metav1.Condition {
	return in.Conditions
}

// SetConditions sets the Status Conditions
func (in *KonnectAPIAuthConfigurationStatus) SetConditions(conditions []metav1.Condition) {
	in.Conditions = conditions
}

type KonnectAPIAuthConfigurationRef struct {
	// Name is the name of the KonnectAPIAuthConfiguration resource.
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Namespace is the namespace of the KonnectAPIAuthConfiguration resource.
	// Namespace string `json:"namespace,omitempty"`
}

// +kubebuilder:object:root=true
type KonnectAPIAuthConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []KonnectAPIAuthConfiguration `json:"items"`
}
