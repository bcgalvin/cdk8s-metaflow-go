package k8s

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Affinity",
		reflect.TypeOf((*Affinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.AggregationRule",
		reflect.TypeOf((*AggregationRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.AggregationRuleV1Alpha1",
		reflect.TypeOf((*AggregationRuleV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.AllowedCsiDriverV1Beta1",
		reflect.TypeOf((*AllowedCsiDriverV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.AllowedFlexVolumeV1Beta1",
		reflect.TypeOf((*AllowedFlexVolumeV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.AllowedHostPathV1Beta1",
		reflect.TypeOf((*AllowedHostPathV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ApiService",
		reflect.TypeOf((*ApiService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApiService{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ApiServiceList",
		reflect.TypeOf((*ApiServiceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApiServiceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ApiServiceListProps",
		reflect.TypeOf((*ApiServiceListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ApiServiceProps",
		reflect.TypeOf((*ApiServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ApiServiceSpec",
		reflect.TypeOf((*ApiServiceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.AwsElasticBlockStoreVolumeSource",
		reflect.TypeOf((*AwsElasticBlockStoreVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.AzureDiskVolumeSource",
		reflect.TypeOf((*AzureDiskVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.AzureFilePersistentVolumeSource",
		reflect.TypeOf((*AzureFilePersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.AzureFileVolumeSource",
		reflect.TypeOf((*AzureFileVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Binding",
		reflect.TypeOf((*Binding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Binding{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.BindingProps",
		reflect.TypeOf((*BindingProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.BoundObjectReference",
		reflect.TypeOf((*BoundObjectReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Capabilities",
		reflect.TypeOf((*Capabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CephFsPersistentVolumeSource",
		reflect.TypeOf((*CephFsPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CephFsVolumeSource",
		reflect.TypeOf((*CephFsVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CertificateSigningRequest",
		reflect.TypeOf((*CertificateSigningRequest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CertificateSigningRequest{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CertificateSigningRequestList",
		reflect.TypeOf((*CertificateSigningRequestList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CertificateSigningRequestList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CertificateSigningRequestListProps",
		reflect.TypeOf((*CertificateSigningRequestListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CertificateSigningRequestProps",
		reflect.TypeOf((*CertificateSigningRequestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CertificateSigningRequestSpec",
		reflect.TypeOf((*CertificateSigningRequestSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CinderPersistentVolumeSource",
		reflect.TypeOf((*CinderPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CinderVolumeSource",
		reflect.TypeOf((*CinderVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ClientIpConfig",
		reflect.TypeOf((*ClientIpConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ClusterRole",
		reflect.TypeOf((*ClusterRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterRole{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ClusterRoleBinding",
		reflect.TypeOf((*ClusterRoleBinding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterRoleBinding{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ClusterRoleBindingList",
		reflect.TypeOf((*ClusterRoleBindingList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterRoleBindingList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ClusterRoleBindingListProps",
		reflect.TypeOf((*ClusterRoleBindingListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ClusterRoleBindingListV1Alpha1",
		reflect.TypeOf((*ClusterRoleBindingListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterRoleBindingListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ClusterRoleBindingListV1Alpha1Props",
		reflect.TypeOf((*ClusterRoleBindingListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ClusterRoleBindingProps",
		reflect.TypeOf((*ClusterRoleBindingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ClusterRoleBindingV1Alpha1",
		reflect.TypeOf((*ClusterRoleBindingV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterRoleBindingV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ClusterRoleBindingV1Alpha1Props",
		reflect.TypeOf((*ClusterRoleBindingV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ClusterRoleList",
		reflect.TypeOf((*ClusterRoleList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterRoleList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ClusterRoleListProps",
		reflect.TypeOf((*ClusterRoleListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ClusterRoleListV1Alpha1",
		reflect.TypeOf((*ClusterRoleListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterRoleListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ClusterRoleListV1Alpha1Props",
		reflect.TypeOf((*ClusterRoleListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ClusterRoleProps",
		reflect.TypeOf((*ClusterRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ClusterRoleV1Alpha1",
		reflect.TypeOf((*ClusterRoleV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterRoleV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ClusterRoleV1Alpha1Props",
		reflect.TypeOf((*ClusterRoleV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ComponentCondition",
		reflect.TypeOf((*ComponentCondition)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ComponentStatus",
		reflect.TypeOf((*ComponentStatus)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ComponentStatus{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ComponentStatusList",
		reflect.TypeOf((*ComponentStatusList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ComponentStatusList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ComponentStatusListProps",
		reflect.TypeOf((*ComponentStatusListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ComponentStatusProps",
		reflect.TypeOf((*ComponentStatusProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ConfigMap",
		reflect.TypeOf((*ConfigMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConfigMap{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ConfigMapEnvSource",
		reflect.TypeOf((*ConfigMapEnvSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ConfigMapKeySelector",
		reflect.TypeOf((*ConfigMapKeySelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ConfigMapList",
		reflect.TypeOf((*ConfigMapList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConfigMapList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ConfigMapListProps",
		reflect.TypeOf((*ConfigMapListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ConfigMapNodeConfigSource",
		reflect.TypeOf((*ConfigMapNodeConfigSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ConfigMapProjection",
		reflect.TypeOf((*ConfigMapProjection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ConfigMapProps",
		reflect.TypeOf((*ConfigMapProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ConfigMapVolumeSource",
		reflect.TypeOf((*ConfigMapVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Container",
		reflect.TypeOf((*Container)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ContainerPort",
		reflect.TypeOf((*ContainerPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ContainerResourceMetricSourceV2Beta1",
		reflect.TypeOf((*ContainerResourceMetricSourceV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ContainerResourceMetricSourceV2Beta2",
		reflect.TypeOf((*ContainerResourceMetricSourceV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ControllerRevision",
		reflect.TypeOf((*ControllerRevision)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ControllerRevision{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ControllerRevisionList",
		reflect.TypeOf((*ControllerRevisionList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ControllerRevisionList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ControllerRevisionListProps",
		reflect.TypeOf((*ControllerRevisionListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ControllerRevisionProps",
		reflect.TypeOf((*ControllerRevisionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CronJob",
		reflect.TypeOf((*CronJob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CronJob{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CronJobList",
		reflect.TypeOf((*CronJobList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CronJobList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CronJobListProps",
		reflect.TypeOf((*CronJobListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CronJobListV1Beta1",
		reflect.TypeOf((*CronJobListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CronJobListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CronJobListV1Beta1Props",
		reflect.TypeOf((*CronJobListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CronJobProps",
		reflect.TypeOf((*CronJobProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CronJobSpec",
		reflect.TypeOf((*CronJobSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CronJobSpecV1Beta1",
		reflect.TypeOf((*CronJobSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CronJobV1Beta1",
		reflect.TypeOf((*CronJobV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CronJobV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CronJobV1Beta1Props",
		reflect.TypeOf((*CronJobV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CrossVersionObjectReference",
		reflect.TypeOf((*CrossVersionObjectReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CrossVersionObjectReferenceV2Beta1",
		reflect.TypeOf((*CrossVersionObjectReferenceV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CrossVersionObjectReferenceV2Beta2",
		reflect.TypeOf((*CrossVersionObjectReferenceV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CsiDriver",
		reflect.TypeOf((*CsiDriver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CsiDriver{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CsiDriverList",
		reflect.TypeOf((*CsiDriverList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CsiDriverList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CsiDriverListProps",
		reflect.TypeOf((*CsiDriverListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CsiDriverProps",
		reflect.TypeOf((*CsiDriverProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CsiDriverSpec",
		reflect.TypeOf((*CsiDriverSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CsiNode",
		reflect.TypeOf((*CsiNode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CsiNode{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CsiNodeDriver",
		reflect.TypeOf((*CsiNodeDriver)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CsiNodeList",
		reflect.TypeOf((*CsiNodeList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CsiNodeList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CsiNodeListProps",
		reflect.TypeOf((*CsiNodeListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CsiNodeProps",
		reflect.TypeOf((*CsiNodeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CsiNodeSpec",
		reflect.TypeOf((*CsiNodeSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CsiPersistentVolumeSource",
		reflect.TypeOf((*CsiPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CsiStorageCapacityListV1Alpha1",
		reflect.TypeOf((*CsiStorageCapacityListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CsiStorageCapacityListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CsiStorageCapacityListV1Alpha1Props",
		reflect.TypeOf((*CsiStorageCapacityListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CsiStorageCapacityListV1Beta1",
		reflect.TypeOf((*CsiStorageCapacityListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CsiStorageCapacityListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CsiStorageCapacityListV1Beta1Props",
		reflect.TypeOf((*CsiStorageCapacityListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CsiStorageCapacityV1Alpha1",
		reflect.TypeOf((*CsiStorageCapacityV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CsiStorageCapacityV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CsiStorageCapacityV1Alpha1Props",
		reflect.TypeOf((*CsiStorageCapacityV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CsiStorageCapacityV1Beta1",
		reflect.TypeOf((*CsiStorageCapacityV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CsiStorageCapacityV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CsiStorageCapacityV1Beta1Props",
		reflect.TypeOf((*CsiStorageCapacityV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CsiVolumeSource",
		reflect.TypeOf((*CsiVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CustomResourceColumnDefinition",
		reflect.TypeOf((*CustomResourceColumnDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CustomResourceConversion",
		reflect.TypeOf((*CustomResourceConversion)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CustomResourceDefinition",
		reflect.TypeOf((*CustomResourceDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CustomResourceDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.CustomResourceDefinitionList",
		reflect.TypeOf((*CustomResourceDefinitionList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CustomResourceDefinitionList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CustomResourceDefinitionListProps",
		reflect.TypeOf((*CustomResourceDefinitionListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CustomResourceDefinitionNames",
		reflect.TypeOf((*CustomResourceDefinitionNames)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CustomResourceDefinitionProps",
		reflect.TypeOf((*CustomResourceDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CustomResourceDefinitionSpec",
		reflect.TypeOf((*CustomResourceDefinitionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CustomResourceDefinitionVersion",
		reflect.TypeOf((*CustomResourceDefinitionVersion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CustomResourceSubresourceScale",
		reflect.TypeOf((*CustomResourceSubresourceScale)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CustomResourceSubresources",
		reflect.TypeOf((*CustomResourceSubresources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.CustomResourceValidation",
		reflect.TypeOf((*CustomResourceValidation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.DaemonSet",
		reflect.TypeOf((*DaemonSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DaemonSet{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.DaemonSetList",
		reflect.TypeOf((*DaemonSetList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DaemonSetList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.DaemonSetListProps",
		reflect.TypeOf((*DaemonSetListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.DaemonSetProps",
		reflect.TypeOf((*DaemonSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.DaemonSetSpec",
		reflect.TypeOf((*DaemonSetSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.DaemonSetUpdateStrategy",
		reflect.TypeOf((*DaemonSetUpdateStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.DeleteOptions",
		reflect.TypeOf((*DeleteOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Deployment",
		reflect.TypeOf((*Deployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Deployment{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.DeploymentList",
		reflect.TypeOf((*DeploymentList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DeploymentList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.DeploymentListProps",
		reflect.TypeOf((*DeploymentListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.DeploymentProps",
		reflect.TypeOf((*DeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.DeploymentSpec",
		reflect.TypeOf((*DeploymentSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.DeploymentStrategy",
		reflect.TypeOf((*DeploymentStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.DownwardApiProjection",
		reflect.TypeOf((*DownwardApiProjection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.DownwardApiVolumeFile",
		reflect.TypeOf((*DownwardApiVolumeFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.DownwardApiVolumeSource",
		reflect.TypeOf((*DownwardApiVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EmptyDirVolumeSource",
		reflect.TypeOf((*EmptyDirVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Endpoint",
		reflect.TypeOf((*Endpoint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointAddress",
		reflect.TypeOf((*EndpointAddress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointConditions",
		reflect.TypeOf((*EndpointConditions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointConditionsV1Beta1",
		reflect.TypeOf((*EndpointConditionsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointHints",
		reflect.TypeOf((*EndpointHints)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointHintsV1Beta1",
		reflect.TypeOf((*EndpointHintsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointPort",
		reflect.TypeOf((*EndpointPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointPortV1Beta1",
		reflect.TypeOf((*EndpointPortV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.EndpointSlice",
		reflect.TypeOf((*EndpointSlice)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EndpointSlice{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.EndpointSliceList",
		reflect.TypeOf((*EndpointSliceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EndpointSliceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointSliceListProps",
		reflect.TypeOf((*EndpointSliceListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.EndpointSliceListV1Beta1",
		reflect.TypeOf((*EndpointSliceListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EndpointSliceListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointSliceListV1Beta1Props",
		reflect.TypeOf((*EndpointSliceListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointSliceProps",
		reflect.TypeOf((*EndpointSliceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.EndpointSliceV1Beta1",
		reflect.TypeOf((*EndpointSliceV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EndpointSliceV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointSliceV1Beta1Props",
		reflect.TypeOf((*EndpointSliceV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointSubset",
		reflect.TypeOf((*EndpointSubset)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointV1Beta1",
		reflect.TypeOf((*EndpointV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Endpoints",
		reflect.TypeOf((*Endpoints)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Endpoints{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.EndpointsList",
		reflect.TypeOf((*EndpointsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EndpointsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointsListProps",
		reflect.TypeOf((*EndpointsListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EndpointsProps",
		reflect.TypeOf((*EndpointsProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EnvFromSource",
		reflect.TypeOf((*EnvFromSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EnvVar",
		reflect.TypeOf((*EnvVar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EnvVarSource",
		reflect.TypeOf((*EnvVarSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EphemeralContainer",
		reflect.TypeOf((*EphemeralContainer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EphemeralVolumeSource",
		reflect.TypeOf((*EphemeralVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Event",
		reflect.TypeOf((*Event)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Event{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.EventList",
		reflect.TypeOf((*EventList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EventList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EventListProps",
		reflect.TypeOf((*EventListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.EventListV1Beta1",
		reflect.TypeOf((*EventListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EventListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EventListV1Beta1Props",
		reflect.TypeOf((*EventListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EventProps",
		reflect.TypeOf((*EventProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EventSeries",
		reflect.TypeOf((*EventSeries)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EventSeriesV1Beta1",
		reflect.TypeOf((*EventSeriesV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EventSource",
		reflect.TypeOf((*EventSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.EventV1Beta1",
		reflect.TypeOf((*EventV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EventV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EventV1Beta1Props",
		reflect.TypeOf((*EventV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Eviction",
		reflect.TypeOf((*Eviction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Eviction{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.EvictionProps",
		reflect.TypeOf((*EvictionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ExecAction",
		reflect.TypeOf((*ExecAction)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ExternalDocumentation",
		reflect.TypeOf((*ExternalDocumentation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ExternalMetricSourceV2Beta1",
		reflect.TypeOf((*ExternalMetricSourceV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ExternalMetricSourceV2Beta2",
		reflect.TypeOf((*ExternalMetricSourceV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.FcVolumeSource",
		reflect.TypeOf((*FcVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.FlexPersistentVolumeSource",
		reflect.TypeOf((*FlexPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.FlexVolumeSource",
		reflect.TypeOf((*FlexVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.FlockerVolumeSource",
		reflect.TypeOf((*FlockerVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.FlowDistinguisherMethodV1Beta1",
		reflect.TypeOf((*FlowDistinguisherMethodV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.FlowSchemaListV1Beta1",
		reflect.TypeOf((*FlowSchemaListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FlowSchemaListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.FlowSchemaListV1Beta1Props",
		reflect.TypeOf((*FlowSchemaListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.FlowSchemaSpecV1Beta1",
		reflect.TypeOf((*FlowSchemaSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.FlowSchemaV1Beta1",
		reflect.TypeOf((*FlowSchemaV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FlowSchemaV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.FlowSchemaV1Beta1Props",
		reflect.TypeOf((*FlowSchemaV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ForZone",
		reflect.TypeOf((*ForZone)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ForZoneV1Beta1",
		reflect.TypeOf((*ForZoneV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.FsGroupStrategyOptionsV1Beta1",
		reflect.TypeOf((*FsGroupStrategyOptionsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.GcePersistentDiskVolumeSource",
		reflect.TypeOf((*GcePersistentDiskVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.GitRepoVolumeSource",
		reflect.TypeOf((*GitRepoVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.GlusterfsPersistentVolumeSource",
		reflect.TypeOf((*GlusterfsPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.GlusterfsVolumeSource",
		reflect.TypeOf((*GlusterfsVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.GroupSubjectV1Beta1",
		reflect.TypeOf((*GroupSubjectV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Handler",
		reflect.TypeOf((*Handler)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscaler",
		reflect.TypeOf((*HorizontalPodAutoscaler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_HorizontalPodAutoscaler{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerBehaviorV2Beta2",
		reflect.TypeOf((*HorizontalPodAutoscalerBehaviorV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerList",
		reflect.TypeOf((*HorizontalPodAutoscalerList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_HorizontalPodAutoscalerList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerListProps",
		reflect.TypeOf((*HorizontalPodAutoscalerListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerListV2Beta1",
		reflect.TypeOf((*HorizontalPodAutoscalerListV2Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_HorizontalPodAutoscalerListV2Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerListV2Beta1Props",
		reflect.TypeOf((*HorizontalPodAutoscalerListV2Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerListV2Beta2",
		reflect.TypeOf((*HorizontalPodAutoscalerListV2Beta2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_HorizontalPodAutoscalerListV2Beta2{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerListV2Beta2Props",
		reflect.TypeOf((*HorizontalPodAutoscalerListV2Beta2Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerProps",
		reflect.TypeOf((*HorizontalPodAutoscalerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerSpec",
		reflect.TypeOf((*HorizontalPodAutoscalerSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerSpecV2Beta1",
		reflect.TypeOf((*HorizontalPodAutoscalerSpecV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerSpecV2Beta2",
		reflect.TypeOf((*HorizontalPodAutoscalerSpecV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerV2Beta1",
		reflect.TypeOf((*HorizontalPodAutoscalerV2Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_HorizontalPodAutoscalerV2Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerV2Beta1Props",
		reflect.TypeOf((*HorizontalPodAutoscalerV2Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerV2Beta2",
		reflect.TypeOf((*HorizontalPodAutoscalerV2Beta2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_HorizontalPodAutoscalerV2Beta2{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HorizontalPodAutoscalerV2Beta2Props",
		reflect.TypeOf((*HorizontalPodAutoscalerV2Beta2Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HostAlias",
		reflect.TypeOf((*HostAlias)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HostPathVolumeSource",
		reflect.TypeOf((*HostPathVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HostPortRangeV1Beta1",
		reflect.TypeOf((*HostPortRangeV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HpaScalingPolicyV2Beta2",
		reflect.TypeOf((*HpaScalingPolicyV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HpaScalingRulesV2Beta2",
		reflect.TypeOf((*HpaScalingRulesV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HttpGetAction",
		reflect.TypeOf((*HttpGetAction)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HttpHeader",
		reflect.TypeOf((*HttpHeader)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HttpIngressPath",
		reflect.TypeOf((*HttpIngressPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.HttpIngressRuleValue",
		reflect.TypeOf((*HttpIngressRuleValue)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IdRangeV1Beta1",
		reflect.TypeOf((*IdRangeV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Ingress",
		reflect.TypeOf((*Ingress)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ingress{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IngressBackend",
		reflect.TypeOf((*IngressBackend)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.IngressClass",
		reflect.TypeOf((*IngressClass)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IngressClass{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.IngressClassList",
		reflect.TypeOf((*IngressClassList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IngressClassList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IngressClassListProps",
		reflect.TypeOf((*IngressClassListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IngressClassParametersReference",
		reflect.TypeOf((*IngressClassParametersReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IngressClassProps",
		reflect.TypeOf((*IngressClassProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IngressClassSpec",
		reflect.TypeOf((*IngressClassSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.IngressList",
		reflect.TypeOf((*IngressList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IngressList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IngressListProps",
		reflect.TypeOf((*IngressListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IngressProps",
		reflect.TypeOf((*IngressProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IngressRule",
		reflect.TypeOf((*IngressRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IngressServiceBackend",
		reflect.TypeOf((*IngressServiceBackend)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IngressSpec",
		reflect.TypeOf((*IngressSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IngressTls",
		reflect.TypeOf((*IngressTls)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.IntOrString",
		reflect.TypeOf((*IntOrString)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IntOrString{}
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-metaflow.k8s.IoK8SApimachineryPkgApisMetaV1DeleteOptionsKind",
		reflect.TypeOf((*IoK8SApimachineryPkgApisMetaV1DeleteOptionsKind)(nil)).Elem(),
		map[string]interface{}{
			"DELETE_OPTIONS": IoK8SApimachineryPkgApisMetaV1DeleteOptionsKind_DELETE_OPTIONS,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IpBlock",
		reflect.TypeOf((*IpBlock)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IscsiPersistentVolumeSource",
		reflect.TypeOf((*IscsiPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.IscsiVolumeSource",
		reflect.TypeOf((*IscsiVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Job",
		reflect.TypeOf((*Job)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Job{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.JobList",
		reflect.TypeOf((*JobList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_JobList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.JobListProps",
		reflect.TypeOf((*JobListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.JobProps",
		reflect.TypeOf((*JobProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.JobSpec",
		reflect.TypeOf((*JobSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.JobTemplateSpec",
		reflect.TypeOf((*JobTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.JobTemplateSpecV1Beta1",
		reflect.TypeOf((*JobTemplateSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.JsonSchemaProps",
		reflect.TypeOf((*JsonSchemaProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.KeyToPath",
		reflect.TypeOf((*KeyToPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.LabelSelector",
		reflect.TypeOf((*LabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.LabelSelectorRequirement",
		reflect.TypeOf((*LabelSelectorRequirement)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Lease",
		reflect.TypeOf((*Lease)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Lease{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.LeaseList",
		reflect.TypeOf((*LeaseList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LeaseList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.LeaseListProps",
		reflect.TypeOf((*LeaseListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.LeaseProps",
		reflect.TypeOf((*LeaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.LeaseSpec",
		reflect.TypeOf((*LeaseSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Lifecycle",
		reflect.TypeOf((*Lifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.LimitRange",
		reflect.TypeOf((*LimitRange)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LimitRange{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.LimitRangeItem",
		reflect.TypeOf((*LimitRangeItem)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.LimitRangeList",
		reflect.TypeOf((*LimitRangeList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LimitRangeList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.LimitRangeListProps",
		reflect.TypeOf((*LimitRangeListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.LimitRangeProps",
		reflect.TypeOf((*LimitRangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.LimitRangeSpec",
		reflect.TypeOf((*LimitRangeSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.LimitResponseV1Beta1",
		reflect.TypeOf((*LimitResponseV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.LimitedPriorityLevelConfigurationV1Beta1",
		reflect.TypeOf((*LimitedPriorityLevelConfigurationV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ListMeta",
		reflect.TypeOf((*ListMeta)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.LocalObjectReference",
		reflect.TypeOf((*LocalObjectReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.LocalSubjectAccessReview",
		reflect.TypeOf((*LocalSubjectAccessReview)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LocalSubjectAccessReview{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.LocalSubjectAccessReviewProps",
		reflect.TypeOf((*LocalSubjectAccessReviewProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.LocalVolumeSource",
		reflect.TypeOf((*LocalVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ManagedFieldsEntry",
		reflect.TypeOf((*ManagedFieldsEntry)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.MetricIdentifierV2Beta2",
		reflect.TypeOf((*MetricIdentifierV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.MetricSpecV2Beta1",
		reflect.TypeOf((*MetricSpecV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.MetricSpecV2Beta2",
		reflect.TypeOf((*MetricSpecV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.MetricTargetV2Beta2",
		reflect.TypeOf((*MetricTargetV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.MutatingWebhook",
		reflect.TypeOf((*MutatingWebhook)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.MutatingWebhookConfiguration",
		reflect.TypeOf((*MutatingWebhookConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MutatingWebhookConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.MutatingWebhookConfigurationList",
		reflect.TypeOf((*MutatingWebhookConfigurationList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MutatingWebhookConfigurationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.MutatingWebhookConfigurationListProps",
		reflect.TypeOf((*MutatingWebhookConfigurationListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.MutatingWebhookConfigurationProps",
		reflect.TypeOf((*MutatingWebhookConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Namespace",
		reflect.TypeOf((*Namespace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Namespace{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.NamespaceList",
		reflect.TypeOf((*NamespaceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NamespaceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NamespaceListProps",
		reflect.TypeOf((*NamespaceListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NamespaceProps",
		reflect.TypeOf((*NamespaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NamespaceSpec",
		reflect.TypeOf((*NamespaceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.NetworkPolicy",
		reflect.TypeOf((*NetworkPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NetworkPolicyEgressRule",
		reflect.TypeOf((*NetworkPolicyEgressRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NetworkPolicyIngressRule",
		reflect.TypeOf((*NetworkPolicyIngressRule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.NetworkPolicyList",
		reflect.TypeOf((*NetworkPolicyList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkPolicyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NetworkPolicyListProps",
		reflect.TypeOf((*NetworkPolicyListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NetworkPolicyPeer",
		reflect.TypeOf((*NetworkPolicyPeer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NetworkPolicyPort",
		reflect.TypeOf((*NetworkPolicyPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NetworkPolicyProps",
		reflect.TypeOf((*NetworkPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NetworkPolicySpec",
		reflect.TypeOf((*NetworkPolicySpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NfsVolumeSource",
		reflect.TypeOf((*NfsVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Node",
		reflect.TypeOf((*Node)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Node{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NodeAffinity",
		reflect.TypeOf((*NodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NodeConfigSource",
		reflect.TypeOf((*NodeConfigSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.NodeList",
		reflect.TypeOf((*NodeList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NodeList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NodeListProps",
		reflect.TypeOf((*NodeListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NodeProps",
		reflect.TypeOf((*NodeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NodeSelector",
		reflect.TypeOf((*NodeSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NodeSelectorRequirement",
		reflect.TypeOf((*NodeSelectorRequirement)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NodeSelectorTerm",
		reflect.TypeOf((*NodeSelectorTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NodeSpec",
		reflect.TypeOf((*NodeSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NonResourceAttributes",
		reflect.TypeOf((*NonResourceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.NonResourcePolicyRuleV1Beta1",
		reflect.TypeOf((*NonResourcePolicyRuleV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ObjectFieldSelector",
		reflect.TypeOf((*ObjectFieldSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ObjectMeta",
		reflect.TypeOf((*ObjectMeta)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ObjectMetricSourceV2Beta1",
		reflect.TypeOf((*ObjectMetricSourceV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ObjectMetricSourceV2Beta2",
		reflect.TypeOf((*ObjectMetricSourceV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ObjectReference",
		reflect.TypeOf((*ObjectReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Overhead",
		reflect.TypeOf((*Overhead)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.OverheadV1Alpha1",
		reflect.TypeOf((*OverheadV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.OverheadV1Beta1",
		reflect.TypeOf((*OverheadV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.OwnerReference",
		reflect.TypeOf((*OwnerReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PersistentVolume",
		reflect.TypeOf((*PersistentVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PersistentVolume{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PersistentVolumeClaim",
		reflect.TypeOf((*PersistentVolumeClaim)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PersistentVolumeClaim{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PersistentVolumeClaimList",
		reflect.TypeOf((*PersistentVolumeClaimList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PersistentVolumeClaimList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PersistentVolumeClaimListProps",
		reflect.TypeOf((*PersistentVolumeClaimListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PersistentVolumeClaimProps",
		reflect.TypeOf((*PersistentVolumeClaimProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PersistentVolumeClaimSpec",
		reflect.TypeOf((*PersistentVolumeClaimSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PersistentVolumeClaimTemplate",
		reflect.TypeOf((*PersistentVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PersistentVolumeClaimVolumeSource",
		reflect.TypeOf((*PersistentVolumeClaimVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PersistentVolumeList",
		reflect.TypeOf((*PersistentVolumeList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PersistentVolumeList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PersistentVolumeListProps",
		reflect.TypeOf((*PersistentVolumeListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PersistentVolumeProps",
		reflect.TypeOf((*PersistentVolumeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PersistentVolumeSpec",
		reflect.TypeOf((*PersistentVolumeSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PhotonPersistentDiskVolumeSource",
		reflect.TypeOf((*PhotonPersistentDiskVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Pod",
		reflect.TypeOf((*Pod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Pod{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodAffinity",
		reflect.TypeOf((*PodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodAffinityTerm",
		reflect.TypeOf((*PodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodAntiAffinity",
		reflect.TypeOf((*PodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PodDisruptionBudget",
		reflect.TypeOf((*PodDisruptionBudget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PodDisruptionBudget{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PodDisruptionBudgetList",
		reflect.TypeOf((*PodDisruptionBudgetList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PodDisruptionBudgetList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodDisruptionBudgetListProps",
		reflect.TypeOf((*PodDisruptionBudgetListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PodDisruptionBudgetListV1Beta1",
		reflect.TypeOf((*PodDisruptionBudgetListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PodDisruptionBudgetListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodDisruptionBudgetListV1Beta1Props",
		reflect.TypeOf((*PodDisruptionBudgetListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodDisruptionBudgetProps",
		reflect.TypeOf((*PodDisruptionBudgetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodDisruptionBudgetSpec",
		reflect.TypeOf((*PodDisruptionBudgetSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodDisruptionBudgetSpecV1Beta1",
		reflect.TypeOf((*PodDisruptionBudgetSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PodDisruptionBudgetV1Beta1",
		reflect.TypeOf((*PodDisruptionBudgetV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PodDisruptionBudgetV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodDisruptionBudgetV1Beta1Props",
		reflect.TypeOf((*PodDisruptionBudgetV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodDnsConfig",
		reflect.TypeOf((*PodDnsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodDnsConfigOption",
		reflect.TypeOf((*PodDnsConfigOption)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PodList",
		reflect.TypeOf((*PodList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PodList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodListProps",
		reflect.TypeOf((*PodListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodProps",
		reflect.TypeOf((*PodProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodReadinessGate",
		reflect.TypeOf((*PodReadinessGate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodSecurityContext",
		reflect.TypeOf((*PodSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PodSecurityPolicyListV1Beta1",
		reflect.TypeOf((*PodSecurityPolicyListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PodSecurityPolicyListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodSecurityPolicyListV1Beta1Props",
		reflect.TypeOf((*PodSecurityPolicyListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodSecurityPolicySpecV1Beta1",
		reflect.TypeOf((*PodSecurityPolicySpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PodSecurityPolicyV1Beta1",
		reflect.TypeOf((*PodSecurityPolicyV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PodSecurityPolicyV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodSecurityPolicyV1Beta1Props",
		reflect.TypeOf((*PodSecurityPolicyV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodSpec",
		reflect.TypeOf((*PodSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PodTemplate",
		reflect.TypeOf((*PodTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PodTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PodTemplateList",
		reflect.TypeOf((*PodTemplateList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PodTemplateList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodTemplateListProps",
		reflect.TypeOf((*PodTemplateListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodTemplateProps",
		reflect.TypeOf((*PodTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodTemplateSpec",
		reflect.TypeOf((*PodTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodsMetricSourceV2Beta1",
		reflect.TypeOf((*PodsMetricSourceV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PodsMetricSourceV2Beta2",
		reflect.TypeOf((*PodsMetricSourceV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PolicyRule",
		reflect.TypeOf((*PolicyRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PolicyRuleV1Alpha1",
		reflect.TypeOf((*PolicyRuleV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PolicyRulesWithSubjectsV1Beta1",
		reflect.TypeOf((*PolicyRulesWithSubjectsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PortworxVolumeSource",
		reflect.TypeOf((*PortworxVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Preconditions",
		reflect.TypeOf((*Preconditions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PreferredSchedulingTerm",
		reflect.TypeOf((*PreferredSchedulingTerm)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PriorityClass",
		reflect.TypeOf((*PriorityClass)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PriorityClass{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PriorityClassList",
		reflect.TypeOf((*PriorityClassList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PriorityClassList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PriorityClassListProps",
		reflect.TypeOf((*PriorityClassListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PriorityClassListV1Alpha1",
		reflect.TypeOf((*PriorityClassListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PriorityClassListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PriorityClassListV1Alpha1Props",
		reflect.TypeOf((*PriorityClassListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PriorityClassProps",
		reflect.TypeOf((*PriorityClassProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PriorityClassV1Alpha1",
		reflect.TypeOf((*PriorityClassV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PriorityClassV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PriorityClassV1Alpha1Props",
		reflect.TypeOf((*PriorityClassV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PriorityLevelConfigurationListV1Beta1",
		reflect.TypeOf((*PriorityLevelConfigurationListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PriorityLevelConfigurationListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PriorityLevelConfigurationListV1Beta1Props",
		reflect.TypeOf((*PriorityLevelConfigurationListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PriorityLevelConfigurationReferenceV1Beta1",
		reflect.TypeOf((*PriorityLevelConfigurationReferenceV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PriorityLevelConfigurationSpecV1Beta1",
		reflect.TypeOf((*PriorityLevelConfigurationSpecV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.PriorityLevelConfigurationV1Beta1",
		reflect.TypeOf((*PriorityLevelConfigurationV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PriorityLevelConfigurationV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.PriorityLevelConfigurationV1Beta1Props",
		reflect.TypeOf((*PriorityLevelConfigurationV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Probe",
		reflect.TypeOf((*Probe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ProjectedVolumeSource",
		reflect.TypeOf((*ProjectedVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Quantity",
		reflect.TypeOf((*Quantity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_Quantity{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.QueuingConfigurationV1Beta1",
		reflect.TypeOf((*QueuingConfigurationV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.QuobyteVolumeSource",
		reflect.TypeOf((*QuobyteVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RbdPersistentVolumeSource",
		reflect.TypeOf((*RbdPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RbdVolumeSource",
		reflect.TypeOf((*RbdVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ReplicaSet",
		reflect.TypeOf((*ReplicaSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ReplicaSet{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ReplicaSetList",
		reflect.TypeOf((*ReplicaSetList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ReplicaSetList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ReplicaSetListProps",
		reflect.TypeOf((*ReplicaSetListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ReplicaSetProps",
		reflect.TypeOf((*ReplicaSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ReplicaSetSpec",
		reflect.TypeOf((*ReplicaSetSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ReplicationController",
		reflect.TypeOf((*ReplicationController)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ReplicationController{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ReplicationControllerList",
		reflect.TypeOf((*ReplicationControllerList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ReplicationControllerList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ReplicationControllerListProps",
		reflect.TypeOf((*ReplicationControllerListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ReplicationControllerProps",
		reflect.TypeOf((*ReplicationControllerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ReplicationControllerSpec",
		reflect.TypeOf((*ReplicationControllerSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ResourceAttributes",
		reflect.TypeOf((*ResourceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ResourceFieldSelector",
		reflect.TypeOf((*ResourceFieldSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ResourceMetricSourceV2Beta1",
		reflect.TypeOf((*ResourceMetricSourceV2Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ResourceMetricSourceV2Beta2",
		reflect.TypeOf((*ResourceMetricSourceV2Beta2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ResourcePolicyRuleV1Beta1",
		reflect.TypeOf((*ResourcePolicyRuleV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ResourceQuota",
		reflect.TypeOf((*ResourceQuota)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ResourceQuota{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ResourceQuotaList",
		reflect.TypeOf((*ResourceQuotaList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ResourceQuotaList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ResourceQuotaListProps",
		reflect.TypeOf((*ResourceQuotaListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ResourceQuotaProps",
		reflect.TypeOf((*ResourceQuotaProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ResourceQuotaSpec",
		reflect.TypeOf((*ResourceQuotaSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ResourceRequirements",
		reflect.TypeOf((*ResourceRequirements)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Role",
		reflect.TypeOf((*Role)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Role{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.RoleBinding",
		reflect.TypeOf((*RoleBinding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RoleBinding{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.RoleBindingList",
		reflect.TypeOf((*RoleBindingList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RoleBindingList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RoleBindingListProps",
		reflect.TypeOf((*RoleBindingListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.RoleBindingListV1Alpha1",
		reflect.TypeOf((*RoleBindingListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RoleBindingListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RoleBindingListV1Alpha1Props",
		reflect.TypeOf((*RoleBindingListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RoleBindingProps",
		reflect.TypeOf((*RoleBindingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.RoleBindingV1Alpha1",
		reflect.TypeOf((*RoleBindingV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RoleBindingV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RoleBindingV1Alpha1Props",
		reflect.TypeOf((*RoleBindingV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.RoleList",
		reflect.TypeOf((*RoleList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RoleList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RoleListProps",
		reflect.TypeOf((*RoleListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.RoleListV1Alpha1",
		reflect.TypeOf((*RoleListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RoleListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RoleListV1Alpha1Props",
		reflect.TypeOf((*RoleListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RoleProps",
		reflect.TypeOf((*RoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RoleRef",
		reflect.TypeOf((*RoleRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RoleRefV1Alpha1",
		reflect.TypeOf((*RoleRefV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.RoleV1Alpha1",
		reflect.TypeOf((*RoleV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RoleV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RoleV1Alpha1Props",
		reflect.TypeOf((*RoleV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RollingUpdateDaemonSet",
		reflect.TypeOf((*RollingUpdateDaemonSet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RollingUpdateDeployment",
		reflect.TypeOf((*RollingUpdateDeployment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RollingUpdateStatefulSetStrategy",
		reflect.TypeOf((*RollingUpdateStatefulSetStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RuleWithOperations",
		reflect.TypeOf((*RuleWithOperations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RunAsGroupStrategyOptionsV1Beta1",
		reflect.TypeOf((*RunAsGroupStrategyOptionsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RunAsUserStrategyOptionsV1Beta1",
		reflect.TypeOf((*RunAsUserStrategyOptionsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.RuntimeClass",
		reflect.TypeOf((*RuntimeClass)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RuntimeClass{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.RuntimeClassList",
		reflect.TypeOf((*RuntimeClassList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RuntimeClassList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RuntimeClassListProps",
		reflect.TypeOf((*RuntimeClassListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.RuntimeClassListV1Alpha1",
		reflect.TypeOf((*RuntimeClassListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RuntimeClassListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RuntimeClassListV1Alpha1Props",
		reflect.TypeOf((*RuntimeClassListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.RuntimeClassListV1Beta1",
		reflect.TypeOf((*RuntimeClassListV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RuntimeClassListV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RuntimeClassListV1Beta1Props",
		reflect.TypeOf((*RuntimeClassListV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RuntimeClassProps",
		reflect.TypeOf((*RuntimeClassProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RuntimeClassSpecV1Alpha1",
		reflect.TypeOf((*RuntimeClassSpecV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RuntimeClassStrategyOptionsV1Beta1",
		reflect.TypeOf((*RuntimeClassStrategyOptionsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.RuntimeClassV1Alpha1",
		reflect.TypeOf((*RuntimeClassV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RuntimeClassV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RuntimeClassV1Alpha1Props",
		reflect.TypeOf((*RuntimeClassV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.RuntimeClassV1Beta1",
		reflect.TypeOf((*RuntimeClassV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RuntimeClassV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.RuntimeClassV1Beta1Props",
		reflect.TypeOf((*RuntimeClassV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Scale",
		reflect.TypeOf((*Scale)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Scale{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ScaleIoPersistentVolumeSource",
		reflect.TypeOf((*ScaleIoPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ScaleIoVolumeSource",
		reflect.TypeOf((*ScaleIoVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ScaleProps",
		reflect.TypeOf((*ScaleProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ScaleSpec",
		reflect.TypeOf((*ScaleSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Scheduling",
		reflect.TypeOf((*Scheduling)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SchedulingV1Alpha1",
		reflect.TypeOf((*SchedulingV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SchedulingV1Beta1",
		reflect.TypeOf((*SchedulingV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ScopeSelector",
		reflect.TypeOf((*ScopeSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ScopedResourceSelectorRequirement",
		reflect.TypeOf((*ScopedResourceSelectorRequirement)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SeLinuxOptions",
		reflect.TypeOf((*SeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SeLinuxStrategyOptionsV1Beta1",
		reflect.TypeOf((*SeLinuxStrategyOptionsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SeccompProfile",
		reflect.TypeOf((*SeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Secret",
		reflect.TypeOf((*Secret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Secret{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SecretEnvSource",
		reflect.TypeOf((*SecretEnvSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SecretKeySelector",
		reflect.TypeOf((*SecretKeySelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.SecretList",
		reflect.TypeOf((*SecretList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SecretList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SecretListProps",
		reflect.TypeOf((*SecretListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SecretProjection",
		reflect.TypeOf((*SecretProjection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SecretProps",
		reflect.TypeOf((*SecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SecretReference",
		reflect.TypeOf((*SecretReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SecretVolumeSource",
		reflect.TypeOf((*SecretVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SecurityContext",
		reflect.TypeOf((*SecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.SelfSubjectAccessReview",
		reflect.TypeOf((*SelfSubjectAccessReview)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SelfSubjectAccessReview{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SelfSubjectAccessReviewProps",
		reflect.TypeOf((*SelfSubjectAccessReviewProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SelfSubjectAccessReviewSpec",
		reflect.TypeOf((*SelfSubjectAccessReviewSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.SelfSubjectRulesReview",
		reflect.TypeOf((*SelfSubjectRulesReview)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SelfSubjectRulesReview{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SelfSubjectRulesReviewProps",
		reflect.TypeOf((*SelfSubjectRulesReviewProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SelfSubjectRulesReviewSpec",
		reflect.TypeOf((*SelfSubjectRulesReviewSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Service",
		reflect.TypeOf((*Service)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Service{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ServiceAccount",
		reflect.TypeOf((*ServiceAccount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceAccount{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ServiceAccountList",
		reflect.TypeOf((*ServiceAccountList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceAccountList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ServiceAccountListProps",
		reflect.TypeOf((*ServiceAccountListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ServiceAccountProps",
		reflect.TypeOf((*ServiceAccountProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ServiceAccountSubjectV1Beta1",
		reflect.TypeOf((*ServiceAccountSubjectV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ServiceAccountTokenProjection",
		reflect.TypeOf((*ServiceAccountTokenProjection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ServiceBackendPort",
		reflect.TypeOf((*ServiceBackendPort)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ServiceList",
		reflect.TypeOf((*ServiceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ServiceListProps",
		reflect.TypeOf((*ServiceListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ServicePort",
		reflect.TypeOf((*ServicePort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ServiceProps",
		reflect.TypeOf((*ServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ServiceReference",
		reflect.TypeOf((*ServiceReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ServiceSpec",
		reflect.TypeOf((*ServiceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SessionAffinityConfig",
		reflect.TypeOf((*SessionAffinityConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.StatefulSet",
		reflect.TypeOf((*StatefulSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StatefulSet{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.StatefulSetList",
		reflect.TypeOf((*StatefulSetList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StatefulSetList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.StatefulSetListProps",
		reflect.TypeOf((*StatefulSetListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.StatefulSetProps",
		reflect.TypeOf((*StatefulSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.StatefulSetSpec",
		reflect.TypeOf((*StatefulSetSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.StatefulSetUpdateStrategy",
		reflect.TypeOf((*StatefulSetUpdateStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.Status",
		reflect.TypeOf((*Status)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Status{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.StatusCause",
		reflect.TypeOf((*StatusCause)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.StatusDetails",
		reflect.TypeOf((*StatusDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.StatusProps",
		reflect.TypeOf((*StatusProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.StorageClass",
		reflect.TypeOf((*StorageClass)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StorageClass{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.StorageClassList",
		reflect.TypeOf((*StorageClassList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StorageClassList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.StorageClassListProps",
		reflect.TypeOf((*StorageClassListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.StorageClassProps",
		reflect.TypeOf((*StorageClassProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.StorageOsPersistentVolumeSource",
		reflect.TypeOf((*StorageOsPersistentVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.StorageOsVolumeSource",
		reflect.TypeOf((*StorageOsVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.StorageVersionListV1Alpha1",
		reflect.TypeOf((*StorageVersionListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StorageVersionListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.StorageVersionListV1Alpha1Props",
		reflect.TypeOf((*StorageVersionListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.StorageVersionV1Alpha1",
		reflect.TypeOf((*StorageVersionV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StorageVersionV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.StorageVersionV1Alpha1Props",
		reflect.TypeOf((*StorageVersionV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Subject",
		reflect.TypeOf((*Subject)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.SubjectAccessReview",
		reflect.TypeOf((*SubjectAccessReview)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SubjectAccessReview{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SubjectAccessReviewProps",
		reflect.TypeOf((*SubjectAccessReviewProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SubjectAccessReviewSpec",
		reflect.TypeOf((*SubjectAccessReviewSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SubjectV1Alpha1",
		reflect.TypeOf((*SubjectV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SubjectV1Beta1",
		reflect.TypeOf((*SubjectV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.SupplementalGroupsStrategyOptionsV1Beta1",
		reflect.TypeOf((*SupplementalGroupsStrategyOptionsV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Sysctl",
		reflect.TypeOf((*Sysctl)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Taint",
		reflect.TypeOf((*Taint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.TcpSocketAction",
		reflect.TypeOf((*TcpSocketAction)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.TokenRequest",
		reflect.TypeOf((*TokenRequest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TokenRequest{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.TokenRequestProps",
		reflect.TypeOf((*TokenRequestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.TokenRequestSpec",
		reflect.TypeOf((*TokenRequestSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.TokenReview",
		reflect.TypeOf((*TokenReview)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TokenReview{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.TokenReviewProps",
		reflect.TypeOf((*TokenReviewProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.TokenReviewSpec",
		reflect.TypeOf((*TokenReviewSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Toleration",
		reflect.TypeOf((*Toleration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.TopologySelectorLabelRequirement",
		reflect.TypeOf((*TopologySelectorLabelRequirement)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.TopologySelectorTerm",
		reflect.TypeOf((*TopologySelectorTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.TopologySpreadConstraint",
		reflect.TypeOf((*TopologySpreadConstraint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.TypedLocalObjectReference",
		reflect.TypeOf((*TypedLocalObjectReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.UserSubjectV1Beta1",
		reflect.TypeOf((*UserSubjectV1Beta1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ValidatingWebhook",
		reflect.TypeOf((*ValidatingWebhook)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ValidatingWebhookConfiguration",
		reflect.TypeOf((*ValidatingWebhookConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ValidatingWebhookConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.ValidatingWebhookConfigurationList",
		reflect.TypeOf((*ValidatingWebhookConfigurationList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ValidatingWebhookConfigurationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ValidatingWebhookConfigurationListProps",
		reflect.TypeOf((*ValidatingWebhookConfigurationListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.ValidatingWebhookConfigurationProps",
		reflect.TypeOf((*ValidatingWebhookConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.Volume",
		reflect.TypeOf((*Volume)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.VolumeAttachment",
		reflect.TypeOf((*VolumeAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VolumeAttachment{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.VolumeAttachmentList",
		reflect.TypeOf((*VolumeAttachmentList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VolumeAttachmentList{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.VolumeAttachmentListProps",
		reflect.TypeOf((*VolumeAttachmentListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.VolumeAttachmentListV1Alpha1",
		reflect.TypeOf((*VolumeAttachmentListV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VolumeAttachmentListV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.VolumeAttachmentListV1Alpha1Props",
		reflect.TypeOf((*VolumeAttachmentListV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.VolumeAttachmentProps",
		reflect.TypeOf((*VolumeAttachmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.VolumeAttachmentSource",
		reflect.TypeOf((*VolumeAttachmentSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.VolumeAttachmentSourceV1Alpha1",
		reflect.TypeOf((*VolumeAttachmentSourceV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.VolumeAttachmentSpec",
		reflect.TypeOf((*VolumeAttachmentSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.VolumeAttachmentSpecV1Alpha1",
		reflect.TypeOf((*VolumeAttachmentSpecV1Alpha1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.k8s.VolumeAttachmentV1Alpha1",
		reflect.TypeOf((*VolumeAttachmentV1Alpha1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VolumeAttachmentV1Alpha1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.VolumeAttachmentV1Alpha1Props",
		reflect.TypeOf((*VolumeAttachmentV1Alpha1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.VolumeDevice",
		reflect.TypeOf((*VolumeDevice)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.VolumeMount",
		reflect.TypeOf((*VolumeMount)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.VolumeNodeAffinity",
		reflect.TypeOf((*VolumeNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.VolumeNodeResources",
		reflect.TypeOf((*VolumeNodeResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.VolumeProjection",
		reflect.TypeOf((*VolumeProjection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.VsphereVirtualDiskVolumeSource",
		reflect.TypeOf((*VsphereVirtualDiskVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.WebhookClientConfig",
		reflect.TypeOf((*WebhookClientConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.WebhookConversion",
		reflect.TypeOf((*WebhookConversion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.WeightedPodAffinityTerm",
		reflect.TypeOf((*WeightedPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.k8s.WindowsSecurityContextOptions",
		reflect.TypeOf((*WindowsSecurityContextOptions)(nil)).Elem(),
	)
}
