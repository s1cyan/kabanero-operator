// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha2

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.Kabanero":       schema_pkg_apis_kabanero_v1alpha2_Kabanero(ref),
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KabaneroSpec":   schema_pkg_apis_kabanero_v1alpha2_KabaneroSpec(ref),
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KabaneroStatus": schema_pkg_apis_kabanero_v1alpha2_KabaneroStatus(ref),
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.Stack":          schema_pkg_apis_kabanero_v1alpha2_Stack(ref),
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.StackSpec":      schema_pkg_apis_kabanero_v1alpha2_StackSpec(ref),
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.StackStatus":    schema_pkg_apis_kabanero_v1alpha2_StackStatus(ref),
	}
}

func schema_pkg_apis_kabanero_v1alpha2_Kabanero(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
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
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KabaneroSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KabaneroStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KabaneroSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KabaneroStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_kabanero_v1alpha2_KabaneroSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KabaneroSpec defines the desired state of Kabanero",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"targetNamespaces": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"github": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.GithubConfig"),
						},
					},
					"stacks": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.InstanceStackConfig"),
						},
					},
					"triggers": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.TriggerSpec"),
									},
								},
							},
						},
					},
					"cliServices": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KabaneroCliServicesCustomizationSpec"),
						},
					},
					"landing": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KabaneroLandingCustomizationSpec"),
						},
					},
					"che": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.CheCustomizationSpec"),
						},
					},
					"events": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.EventsCustomizationSpec"),
						},
					},
					"collectionController": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.CollectionControllerSpec"),
						},
					},
					"stackController": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.StackControllerSpec"),
						},
					},
					"admissionControllerWebhook": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.AdmissionControllerWebhookCustomizationSpec"),
						},
					},
					"sso": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.SsoCustomizationSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.AdmissionControllerWebhookCustomizationSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.CheCustomizationSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.CollectionControllerSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.EventsCustomizationSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.GithubConfig", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.InstanceStackConfig", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KabaneroCliServicesCustomizationSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KabaneroLandingCustomizationSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.SsoCustomizationSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.StackControllerSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.TriggerSpec"},
	}
}

func schema_pkg_apis_kabanero_v1alpha2_KabaneroStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KabaneroStatus defines the observed state of the Kabanero instance.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kabaneroInstance": {
						SchemaProps: spec.SchemaProps{
							Description: "Kabanero operator instance readiness status. The status is directly correlated to the availability of resources dependencies.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KabaneroInstanceStatus"),
						},
					},
					"serverless": {
						SchemaProps: spec.SchemaProps{
							Description: "OpenShift serverless operator status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.ServerlessStatus"),
						},
					},
					"tekton": {
						SchemaProps: spec.SchemaProps{
							Description: "Tekton instance readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.TektonStatus"),
						},
					},
					"cli": {
						SchemaProps: spec.SchemaProps{
							Description: "CLI readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.CliStatus"),
						},
					},
					"landing": {
						SchemaProps: spec.SchemaProps{
							Description: "Kabanero Landing page readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KabaneroLandingPageStatus"),
						},
					},
					"appsody": {
						SchemaProps: spec.SchemaProps{
							Description: "Appsody instance readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.AppsodyStatus"),
						},
					},
					"kappnav": {
						SchemaProps: spec.SchemaProps{
							Description: "Kabanero Application Navigator instance readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KappnavStatus"),
						},
					},
					"che": {
						SchemaProps: spec.SchemaProps{
							Description: "Che instance readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.CheStatus"),
						},
					},
					"events": {
						SchemaProps: spec.SchemaProps{
							Description: "Events instance status",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.EventsStatus"),
						},
					},
					"collectionController": {
						SchemaProps: spec.SchemaProps{
							Description: "Kabanero collection controller readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.CollectionControllerStatus"),
						},
					},
					"stackController": {
						SchemaProps: spec.SchemaProps{
							Description: "Kabanero stack controller readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.StackControllerStatus"),
						},
					},
					"admissionControllerWebhook": {
						SchemaProps: spec.SchemaProps{
							Description: "Admission webhook instance status",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.AdmissionControllerWebhookStatus"),
						},
					},
					"sso": {
						SchemaProps: spec.SchemaProps{
							Description: "SSO server status",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.SsoStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.AdmissionControllerWebhookStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.AppsodyStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.CheStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.CliStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.CollectionControllerStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.EventsStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KabaneroInstanceStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KabaneroLandingPageStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.KappnavStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.ServerlessStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.SsoStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.StackControllerStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.TektonStatus"},
	}
}

func schema_pkg_apis_kabanero_v1alpha2_Stack(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Stack is the Schema for the stack API",
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
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.StackSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.StackStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.StackSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.StackStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_kabanero_v1alpha2_StackSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StackSpec defines the desired composition of a Stack",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"versions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.StackVersion"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.StackVersion"},
	}
}

func schema_pkg_apis_kabanero_v1alpha2_StackStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StackStatus defines the observed state of a stack",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"statusMessage": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"versions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.StackVersionStatus"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha2.StackVersionStatus"},
	}
}
