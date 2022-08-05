package cdk8smetaflow

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.AutoscalingOptions",
		reflect.TypeOf((*AutoscalingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.DatabaseAuthOptions",
		reflect.TypeOf((*DatabaseAuthOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.DatabaseMetricsOptions",
		reflect.TypeOf((*DatabaseMetricsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.DatabaseReplicationOptions",
		reflect.TypeOf((*DatabaseReplicationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.DatabaseResourceRequestOptions",
		reflect.TypeOf((*DatabaseResourceRequestOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.DatabaseResourcesOptions",
		reflect.TypeOf((*DatabaseResourcesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.DatabaseVolumePermissionsOptions",
		reflect.TypeOf((*DatabaseVolumePermissionsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.DbMigrationOptions",
		reflect.TypeOf((*DbMigrationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.HostOptions",
		reflect.TypeOf((*HostOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.HttpIngressPath",
		reflect.TypeOf((*HttpIngressPath)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-metaflow.HttpIngressPathType",
		reflect.TypeOf((*HttpIngressPathType)(nil)).Elem(),
		map[string]interface{}{
			"PREFIX": HttpIngressPathType_PREFIX,
			"EXACT": HttpIngressPathType_EXACT,
			"IMPLEMENTATION_SPECIFIC": HttpIngressPathType_IMPLEMENTATION_SPECIFIC,
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-metaflow.IApiResource",
		reflect.TypeOf((*IApiResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "resourceName", GoGetter: "ResourceName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
		},
		func() interface{} {
			return &jsiiProxy_IApiResource{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.ImageOptions",
		reflect.TypeOf((*ImageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-metaflow.ImagePullPolicy",
		reflect.TypeOf((*ImagePullPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": ImagePullPolicy_ALWAYS,
			"IF_NOT_PRESENT": ImagePullPolicy_IF_NOT_PRESENT,
			"NEVER": ImagePullPolicy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.IngressOptions",
		reflect.TypeOf((*IngressOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.MetadataDatabase",
		reflect.TypeOf((*MetadataDatabase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MetadataDatabase{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.MetadataDatabaseEnvOptions",
		reflect.TypeOf((*MetadataDatabaseEnvOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.MetadataDatabaseOptions",
		reflect.TypeOf((*MetadataDatabaseOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.MetadataDatabaseProps",
		reflect.TypeOf((*MetadataDatabaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.MetaflowService",
		reflect.TypeOf((*MetaflowService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MetaflowService{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.MetaflowServiceOptions",
		reflect.TypeOf((*MetaflowServiceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.MetaflowServiceProps",
		reflect.TypeOf((*MetaflowServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.MetaflowUI",
		reflect.TypeOf((*MetaflowUI)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MetaflowUI{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.MetaflowUIOptions",
		reflect.TypeOf((*MetaflowUIOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.MetaflowUIProps",
		reflect.TypeOf((*MetaflowUIProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.ServiceAccountOptions",
		reflect.TypeOf((*ServiceAccountOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.ServiceOptions",
		reflect.TypeOf((*ServiceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-metaflow.ServiceType",
		reflect.TypeOf((*ServiceType)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER_IP": ServiceType_CLUSTER_IP,
			"NODE_PORT": ServiceType_NODE_PORT,
			"LOAD_BALANCER": ServiceType_LOAD_BALANCER,
			"EXTERNAL_NAME": ServiceType_EXTERNAL_NAME,
		},
	)
}
