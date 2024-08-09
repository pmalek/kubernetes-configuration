<!-- This document is generated by KIC's 'generate.docs' make target, DO NOT EDIT -->

## Packages
- [konnect.konghq.com/v1alpha1](#konnectkonghqcomv1alpha1)


## konnect.konghq.com/v1alpha1

Package v1alpha1 contains API Schema definitions for the konnect.konghq.com v1alpha1 API group.

- [KonnectAPIAuthConfiguration](#konnectapiauthconfiguration)
- [KonnectControlPlane](#konnectcontrolplane)
### KonnectAPIAuthConfiguration


KonnectAPIAuthConfiguration is the Schema for the Konnect configuration type.

<!-- konnect_api_auth_configuration description placeholder -->

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `konnect.konghq.com/v1alpha1`
| `kind` _string_ | `KonnectAPIAuthConfiguration`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KonnectAPIAuthConfigurationSpec](#konnectapiauthconfigurationspec)_ | Spec is the specification of the KonnectAPIAuthConfiguration resource. |



### KonnectControlPlane


KonnectControlPlane is the Schema for the KonnectControlplanes API.

<!-- konnect_control_plane description placeholder -->

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `konnect.konghq.com/v1alpha1`
| `kind` _string_ | `KonnectControlPlane`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KonnectControlPlaneSpec](#konnectcontrolplanespec)_ |  |



### Types

In this section you will find types that the CRDs rely on.
#### KonnectAPIAuthConfigurationRef






| Field | Description |
| --- | --- |
| `name` _string_ | Name is the name of the KonnectAPIAuthConfiguration resource. |


_Appears in:_
- [KonnectConfiguration](#konnectconfiguration)

#### KonnectAPIAuthConfigurationSpec


KonnectAPIAuthConfigurationSpec is the specification of the KonnectAPIAuthConfiguration resource.



| Field | Description |
| --- | --- |
| `type` _[KonnectAPIAuthType](#konnectapiauthtype)_ |  |
| `token` _string_ | Token is the Konnect token used to authenticate with the Konnect API. |
| `secretRef` _[SecretReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#secretreference-v1-core)_ | SecretRef is a reference to a Kubernetes Secret containing the Konnect token. This secret is required to has the konghq.com/credential label set to "konnect". |
| `serverURL` _string_ | ServerURL is the URL of the Konnect server. |


_Appears in:_
- [KonnectAPIAuthConfiguration](#konnectapiauthconfiguration)



#### KonnectAPIAuthType
_Underlying type:_ `string`







_Appears in:_
- [KonnectAPIAuthConfigurationSpec](#konnectapiauthconfigurationspec)

#### KonnectConfiguration






| Field | Description |
| --- | --- |
| `authRef` _[KonnectAPIAuthConfigurationRef](#konnectapiauthconfigurationref)_ | APIAuthConfigurationRef is the reference to the API Auth Configuration that should be used for this Konnect Configuration. |


_Appears in:_
- [KonnectControlPlaneSpec](#konnectcontrolplanespec)

#### KonnectControlPlaneSpec


KonnectControlPlaneSpec defines the desired state of KonnectControlPlane.



| Field | Description |
| --- | --- |
| `name` _string_ | The name of the control plane. |
| `description` _string_ | The description of the control plane in Konnect. |
| `cluster_type` _[ClusterType](#clustertype)_ | The ClusterType value of the cluster associated with the Control Plane. |
| `auth_type` _[AuthType](#authtype)_ | The auth type value of the cluster associated with the Runtime Group. |
| `cloud_gateway` _boolean_ | Whether this control-plane can be used for cloud-gateways. |
| `proxy_urls` _[ProxyURL](#proxyurl) array_ | Array of proxy URLs associated with reaching the data-planes connected to a control-plane. |
| `labels` _object (keys:string, values:string)_ | Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.<br /><br /> Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_". |
| `konnect` _[KonnectConfiguration](#konnectconfiguration)_ |  |


_Appears in:_
- [KonnectControlPlane](#konnectcontrolplane)



#### KonnectEntityStatus






| Field | Description |
| --- | --- |
| `id` _string_ | ID is the unique identifier of the Konnect entity as assigned by Konnect API. If it's unset (empty string), it means the Konnect entity hasn't been created yet. |
| `serverURL` _string_ | ServerURL is the URL of the Konnect server in which the entity exists. |
| `organizationID` _string_ | OrgID is ID of Konnect Org that this entity has been created in. |


_Appears in:_
- [KonnectControlPlaneStatus](#konnectcontrolplanestatus)
- [KonnectEntityStatusWithControlPlaneAndServiceRefs](#konnectentitystatuswithcontrolplaneandservicerefs)
- [KonnectEntityStatusWithControlPlaneRef](#konnectentitystatuswithcontrolplaneref)




