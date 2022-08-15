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
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.MetadataDatabaseOptions",
		reflect.TypeOf((*MetadataDatabaseOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.MetaflowChartProps",
		reflect.TypeOf((*MetaflowChartProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.MetaflowServiceChart",
		reflect.TypeOf((*MetaflowServiceChart)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberProperty{JsiiProperty: "deployment", GoGetter: "Deployment"},
			_jsii_.MemberMethod{JsiiMethod: "generateObjectName", GoMethod: "GenerateObjectName"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MetaflowServiceChart{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sChart)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.MetaflowUIChart",
		reflect.TypeOf((*MetaflowUIChart)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberProperty{JsiiProperty: "deployment", GoGetter: "Deployment"},
			_jsii_.MemberMethod{JsiiMethod: "generateObjectName", GoMethod: "GenerateObjectName"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MetaflowUIChart{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sChart)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.MetaflowUIStaticChart",
		reflect.TypeOf((*MetaflowUIStaticChart)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberProperty{JsiiProperty: "deployment", GoGetter: "Deployment"},
			_jsii_.MemberMethod{JsiiMethod: "generateObjectName", GoMethod: "GenerateObjectName"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MetaflowUIStaticChart{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sChart)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.PostgresAddon",
		reflect.TypeOf((*PostgresAddon)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "install", GoMethod: "Install"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_PostgresAddon{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAddon)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.PostgresAddonProps",
		reflect.TypeOf((*PostgresAddonProps)(nil)).Elem(),
	)
}
