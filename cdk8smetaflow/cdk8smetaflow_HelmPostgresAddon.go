// cdk8s-metaflow
package cdk8smetaflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/bcgalvin/cdk8s-metaflow-go/cdk8smetaflow/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type HelmPostgresAddon interface {
	IAddon
	// Experimental.
	Name() *string
	// Experimental.
	Install(scope constructs.Construct) constructs.Construct
}

// The jsii proxy struct for HelmPostgresAddon
type jsiiProxy_HelmPostgresAddon struct {
	jsiiProxy_IAddon
}

func (j *jsiiProxy_HelmPostgresAddon) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewHelmPostgresAddon(props *HelmPostgresAddonProps) HelmPostgresAddon {
	_init_.Initialize()

	j := jsiiProxy_HelmPostgresAddon{}

	_jsii_.Create(
		"cdk8s-metaflow.HelmPostgresAddon",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewHelmPostgresAddon_Override(h HelmPostgresAddon, props *HelmPostgresAddonProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-metaflow.HelmPostgresAddon",
		[]interface{}{props},
		h,
	)
}

func HelmPostgresAddon_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk8s-metaflow.HelmPostgresAddon",
		"NAME",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HelmPostgresAddon) Install(scope constructs.Construct) constructs.Construct {
	var returns constructs.Construct

	_jsii_.Invoke(
		h,
		"install",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

