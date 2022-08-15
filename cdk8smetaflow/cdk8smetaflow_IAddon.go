// cdk8s-metaflow
package cdk8smetaflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Experimental.
type IAddon interface {
	// Experimental.
	Install(scope constructs.Construct) cdk8s.Helm
	// Experimental.
	Name() *string
}

// The jsii proxy for IAddon
type jsiiProxy_IAddon struct {
	_ byte // padding
}

func (i *jsiiProxy_IAddon) Install(scope constructs.Construct) cdk8s.Helm {
	var returns cdk8s.Helm

	_jsii_.Invoke(
		i,
		"install",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAddon) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

