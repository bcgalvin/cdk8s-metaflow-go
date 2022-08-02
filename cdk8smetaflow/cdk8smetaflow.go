// cdk8s-metaflow
package cdk8smetaflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/bcgalvin/cdk8s-metaflow-go/cdk8smetaflow/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/bcgalvin/cdk8s-metaflow-go/cdk8smetaflow/internal"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus22/v2"
)

// Experimental.
type IngressOptions struct {
	// Experimental.
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
}

// Experimental.
type MetaflowService interface {
	constructs.Construct
	// Experimental.
	Deployment() cdk8splus22.Deployment
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Service() cdk8splus22.Service
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MetaflowService
type jsiiProxy_MetaflowService struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_MetaflowService) Deployment() cdk8splus22.Deployment {
	var returns cdk8splus22.Deployment
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetaflowService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetaflowService) Service() cdk8splus22.Service {
	var returns cdk8splus22.Service
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}


// Experimental.
func NewMetaflowService(scope constructs.Construct, id *string, props *MetaflowServiceProps) MetaflowService {
	_init_.Initialize()

	j := jsiiProxy_MetaflowService{}

	_jsii_.Create(
		"cdk8s-metaflow.MetaflowService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewMetaflowService_Override(m MetaflowService, scope constructs.Construct, id *string, props *MetaflowServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-metaflow.MetaflowService",
		[]interface{}{scope, id, props},
		m,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func MetaflowService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-metaflow.MetaflowService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetaflowService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type MetaflowServiceProps struct {
	// Experimental.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Experimental.
	IngressEnabled *bool `field:"optional" json:"ingressEnabled" yaml:"ingressEnabled"`
	// Experimental.
	IngressOptions *IngressOptions `field:"optional" json:"ingressOptions" yaml:"ingressOptions"`
	// Experimental.
	MetadataServicePort *float64 `field:"optional" json:"metadataServicePort" yaml:"metadataServicePort"`
	// Experimental.
	PostgresEnabled *bool `field:"optional" json:"postgresEnabled" yaml:"postgresEnabled"`
	// Experimental.
	PostgresOptions *PostgresOptions `field:"optional" json:"postgresOptions" yaml:"postgresOptions"`
	// Experimental.
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// Experimental.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Experimental.
	UpgradesServicePort *float64 `field:"optional" json:"upgradesServicePort" yaml:"upgradesServicePort"`
}

// Experimental.
type PostgresOptions struct {
	// Experimental.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Experimental.
	DatabasePassword *string `field:"optional" json:"databasePassword" yaml:"databasePassword"`
	// Experimental.
	DatabaseUser *string `field:"optional" json:"databaseUser" yaml:"databaseUser"`
}

