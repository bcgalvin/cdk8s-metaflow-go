// cdk8s-metaflow
package cdk8smetaflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/bcgalvin/cdk8s-metaflow-go/cdk8smetaflow/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type LocalPostgresAddon interface {
	IAddon
	// Experimental.
	Name() *string
	// Experimental.
	Install(scope constructs.Construct) constructs.Construct
}

// The jsii proxy struct for LocalPostgresAddon
type jsiiProxy_LocalPostgresAddon struct {
	jsiiProxy_IAddon
}

func (j *jsiiProxy_LocalPostgresAddon) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewLocalPostgresAddon(props *LocalPostgresAddonProps) LocalPostgresAddon {
	_init_.Initialize()

	j := jsiiProxy_LocalPostgresAddon{}

	_jsii_.Create(
		"cdk8s-metaflow.LocalPostgresAddon",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewLocalPostgresAddon_Override(l LocalPostgresAddon, props *LocalPostgresAddonProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-metaflow.LocalPostgresAddon",
		[]interface{}{props},
		l,
	)
}

func LocalPostgresAddon_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk8s-metaflow.LocalPostgresAddon",
		"NAME",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LocalPostgresAddon) Install(scope constructs.Construct) constructs.Construct {
	var returns constructs.Construct

	_jsii_.Invoke(
		l,
		"install",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

