package crdsvalidation_test

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"
)

func TestKongCredentialBasicAuth(t *testing.T) {
	t.Run("updates not allowed for status conditions", func(t *testing.T) {
		CRDValidationTestCasesGroup[*configurationv1alpha1.KongCredentialBasicAuth]{
			{
				Name: "consumerRef change is not allowed for Programmed=True",
				TestObject: &configurationv1alpha1.KongCredentialBasicAuth{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongCredentialBasicAuthSpec{
						ConsumerRef: corev1.LocalObjectReference{
							Name: "test-kong-consumer",
						},
						KongCredentialBasicAuthAPISpec: configurationv1alpha1.KongCredentialBasicAuthAPISpec{
							Password: "password",
							Username: "username",
						},
					},
					Status: configurationv1alpha1.KongCredentialBasicAuthStatus{
						Konnect: &konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs{},
						Conditions: []metav1.Condition{
							{
								Type:               "Programmed",
								Status:             metav1.ConditionTrue,
								Reason:             "Valid",
								LastTransitionTime: metav1.Now(),
							},
						},
					},
				},
				Update: func(c *configurationv1alpha1.KongCredentialBasicAuth) {
					c.Spec.ConsumerRef.Name = "new-consumer"
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("spec.consumerRef is immutable when an entity is already Programmed"),
			},
			{
				Name: "consumerRef change is allowed when consumer is not Programmed=True nor APIAuthValid=True",
				TestObject: &configurationv1alpha1.KongCredentialBasicAuth{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongCredentialBasicAuthSpec{
						ConsumerRef: corev1.LocalObjectReference{
							Name: "test-kong-consumer",
						},
						KongCredentialBasicAuthAPISpec: configurationv1alpha1.KongCredentialBasicAuthAPISpec{
							Password: "password",
							Username: "username",
						},
					},
					Status: configurationv1alpha1.KongCredentialBasicAuthStatus{
						Konnect: &konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs{},
						Conditions: []metav1.Condition{
							{
								Type:               "Programmed",
								Status:             metav1.ConditionFalse,
								Reason:             "Invalid",
								LastTransitionTime: metav1.Now(),
							},
						},
					},
				},
				Update: func(c *configurationv1alpha1.KongCredentialBasicAuth) {
					c.Spec.ConsumerRef.Name = "new-consumer"
				},
			},
		}.Run(t)
	})

	t.Run("required fields validation", func(t *testing.T) {
		CRDValidationTestCasesGroup[*configurationv1alpha1.KongCredentialBasicAuth]{
			{
				Name: "password is required",
				TestObject: &configurationv1alpha1.KongCredentialBasicAuth{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongCredentialBasicAuthSpec{
						ConsumerRef: corev1.LocalObjectReference{
							Name: "test-kong-consumer",
						},
						KongCredentialBasicAuthAPISpec: configurationv1alpha1.KongCredentialBasicAuthAPISpec{
							Username: "username",
						},
					},
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("spec.consumerREf is immutable when an entity is already Programmed"),
			},
			{
				Name: "username is required",
				TestObject: &configurationv1alpha1.KongCredentialBasicAuth{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongCredentialBasicAuthSpec{
						ConsumerRef: corev1.LocalObjectReference{
							Name: "test-kong-consumer",
						},
						KongCredentialBasicAuthAPISpec: configurationv1alpha1.KongCredentialBasicAuthAPISpec{
							Password: "password",
						},
					},
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("spec.consumerREf is immutable when an entity is already Programmed"),
			},
			{
				Name: "password and username are required",
				TestObject: &configurationv1alpha1.KongCredentialBasicAuth{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongCredentialBasicAuthSpec{
						ConsumerRef: corev1.LocalObjectReference{
							Name: "test-kong-consumer",
						},
						KongCredentialBasicAuthAPISpec: configurationv1alpha1.KongCredentialBasicAuthAPISpec{
							Username: "username",
							Password: "password",
						},
					},
				},
			},
		}.Run(t)
	})

	t.Run("tags validation", func(t *testing.T) {
		CRDValidationTestCasesGroup[*configurationv1alpha1.KongCredentialBasicAuth]{
			{
				Name: "up to 20 tags are allowed",
				TestObject: &configurationv1alpha1.KongCredentialBasicAuth{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongCredentialBasicAuthSpec{
						ConsumerRef: corev1.LocalObjectReference{
							Name: "test-kong-consumer",
						},
						KongCredentialBasicAuthAPISpec: configurationv1alpha1.KongCredentialBasicAuthAPISpec{
							Password: "pass",
							Username: "user",
							Tags: func() []string {
								var tags []string
								for i := range 20 {
									tags = append(tags, fmt.Sprintf("tag-%d", i))
								}
								return tags
							}(),
						},
					},
				},
			},
			{
				Name: "more than 20 tags are not allowed",
				TestObject: &configurationv1alpha1.KongCredentialBasicAuth{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongCredentialBasicAuthSpec{
						ConsumerRef: corev1.LocalObjectReference{
							Name: "test-kong-consumer",
						},
						KongCredentialBasicAuthAPISpec: configurationv1alpha1.KongCredentialBasicAuthAPISpec{
							Password: "pass",
							Username: "user",
							Tags: func() []string {
								var tags []string
								for i := range 21 {
									tags = append(tags, fmt.Sprintf("tag-%d", i))
								}
								return tags
							}(),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("spec.tags: Too many: 21: must have at most 20 items"),
			},
			{
				Name: "tags entries must not be longer than 128 characters",
				TestObject: &configurationv1alpha1.KongCredentialBasicAuth{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongCredentialBasicAuthSpec{
						ConsumerRef: corev1.LocalObjectReference{
							Name: "test-kong-consumer",
						},
						KongCredentialBasicAuthAPISpec: configurationv1alpha1.KongCredentialBasicAuthAPISpec{
							Password: "pass",
							Username: "user",
							Tags: []string{
								lo.RandomString(129, lo.AlphanumericCharset),
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("tags entries must not be longer than 128 characters"),
			},
		}.Run(t)
	})
}