package cdk8smetaflow

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
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
	_jsii_.RegisterClass(
		"cdk8s-metaflow.HelmPostgresAddon",
		reflect.TypeOf((*HelmPostgresAddon)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "install", GoMethod: "Install"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_HelmPostgresAddon{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAddon)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.HelmPostgresAddonProps",
		reflect.TypeOf((*HelmPostgresAddonProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk8s-metaflow.IAddon",
		reflect.TypeOf((*IAddon)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "install", GoMethod: "Install"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_IAddon{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.LocalPostgresAddon",
		reflect.TypeOf((*LocalPostgresAddon)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "install", GoMethod: "Install"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_LocalPostgresAddon{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAddon)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.LocalPostgresAddonProps",
		reflect.TypeOf((*LocalPostgresAddonProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.MetadataDatabaseOptions",
		reflect.TypeOf((*MetadataDatabaseOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.MetaflowChartProps",
		reflect.TypeOf((*MetaflowChartProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.MetaflowService",
		reflect.TypeOf((*MetaflowService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "generateObjectName", GoMethod: "GenerateObjectName"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MetaflowService{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sChart)
			return &j
		},
	)
}
