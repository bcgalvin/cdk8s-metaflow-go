// cdk8s-metaflow
package cdk8smetaflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/bcgalvin/cdk8s-metaflow-go/cdk8smetaflow/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Experimental.
type PostgresAddon interface {
	IAddon
	// Experimental.
	Name() *string
	// Experimental.
	Install(scope constructs.Construct) cdk8s.Helm
}

// The jsii proxy struct for PostgresAddon
type jsiiProxy_PostgresAddon struct {
	jsiiProxy_IAddon
}

func (j *jsiiProxy_PostgresAddon) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewPostgresAddon(props *PostgresAddonProps) PostgresAddon {
	_init_.Initialize()

	j := jsiiProxy_PostgresAddon{}

	_jsii_.Create(
		"cdk8s-metaflow.PostgresAddon",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewPostgresAddon_Override(p PostgresAddon, props *PostgresAddonProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-metaflow.PostgresAddon",
		[]interface{}{props},
		p,
	)
}

func PostgresAddon_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk8s-metaflow.PostgresAddon",
		"NAME",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PostgresAddon) Install(scope constructs.Construct) cdk8s.Helm {
	var returns cdk8s.Helm

	_jsii_.Invoke(
		p,
		"install",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

