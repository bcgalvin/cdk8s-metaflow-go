package cdk8smetaflow

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.IngressOptions",
		reflect.TypeOf((*IngressOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-metaflow.MetaflowService",
		reflect.TypeOf((*MetaflowService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deployment", GoGetter: "Deployment"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MetaflowService{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.MetaflowServiceProps",
		reflect.TypeOf((*MetaflowServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-metaflow.PostgresOptions",
		reflect.TypeOf((*PostgresOptions)(nil)).Elem(),
	)
}
