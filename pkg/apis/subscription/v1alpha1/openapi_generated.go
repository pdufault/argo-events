// +build !ignore_autogenerated

/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.HTTPSubscription":   schema_pkg_apis_subscription_v1alpha1_HTTPSubscription(ref),
		"github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.NATSSubscription":   schema_pkg_apis_subscription_v1alpha1_NATSSubscription(ref),
		"github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.Subscription":       schema_pkg_apis_subscription_v1alpha1_Subscription(ref),
		"github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.SubscriptionList":   schema_pkg_apis_subscription_v1alpha1_SubscriptionList(ref),
		"github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.SubscriptionSpec":   schema_pkg_apis_subscription_v1alpha1_SubscriptionSpec(ref),
		"github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.SubscriptionStatus": schema_pkg_apis_subscription_v1alpha1_SubscriptionStatus(ref),
	}
}

func schema_pkg_apis_subscription_v1alpha1_HTTPSubscription(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HTTPSubscription describes the subscription details over HTTP",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"url": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name", "url"},
			},
		},
	}
}

func schema_pkg_apis_subscription_v1alpha1_NATSSubscription(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NATSSubscription describes the subscription details over NATS protocol",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the subscription",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"serverURL": {
						SchemaProps: spec.SchemaProps{
							Description: "ServerURL is NATS server URL",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"subject": {
						SchemaProps: spec.SchemaProps{
							Description: "Subject is the name of the NATS subject",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name", "serverURL", "subject"},
			},
		},
	}
}

func schema_pkg_apis_subscription_v1alpha1_Subscription(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Subscription is the definition of a subscription resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.SubscriptionSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.SubscriptionStatus"),
						},
					},
				},
				Required: []string{"metadata", "spec", "status"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.SubscriptionSpec", "github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.SubscriptionStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_subscription_v1alpha1_SubscriptionList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SubscriptionList is the list of subscription resources",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "items",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.Subscription"),
									},
								},
							},
						},
					},
				},
				Required: []string{"metadata", "items"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.Subscription", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_subscription_v1alpha1_SubscriptionSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SubscriptionSpec describes the specification of the subscription resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"http": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "subscriptions",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "HTTP refers to list of subscriptions over HTTP protocol",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.HTTPSubscription"),
									},
								},
							},
						},
					},
					"nats": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "subscriptions",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "NATS refers to list of subscriptions over NATS protocol",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.NATSSubscription"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.HTTPSubscription", "github.com/argoproj/argo-events/pkg/apis/subscription/v1alpha1.NATSSubscription"},
	}
}

func schema_pkg_apis_subscription_v1alpha1_SubscriptionStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SubscriptionStatus describes the status of the subscription resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"createdAt": {
						SchemaProps: spec.SchemaProps{
							Description: "CreatedAt refers to creation time",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"updatedAt": {
						SchemaProps: spec.SchemaProps{
							Description: "UpdatedAt refers to time at the resource was updated",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
				Required: []string{"createdAt", "updatedAt"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}
