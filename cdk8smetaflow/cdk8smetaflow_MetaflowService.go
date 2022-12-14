// cdk8s-metaflow
package cdk8smetaflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/bcgalvin/cdk8s-metaflow-go/cdk8smetaflow/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/bcgalvin/cdk8s-metaflow-go/cdk8smetaflow/internal"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus21/v2"
)

// Experimental.
type MetaflowService interface {
	cdk8s.Chart
	// Labels applied to all resources in this chart.
	//
	// This is an immutable copy.
	// Experimental.
	Labels() *map[string]*string
	// The default namespace for all objects in this chart.
	// Experimental.
	Namespace() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	ServiceAccount() cdk8splus21.IServiceAccount
	// Create a dependency between this Chart and other constructs.
	//
	// These can be other ApiObjects, Charts, or custom.
	// Experimental.
	AddDependency(dependencies ...constructs.IConstruct)
	// Generates a app-unique name for an object given it's construct node path.
	//
	// Different resource types may have different constraints on names
	// (`metadata.name`). The previous version of the name generator was
	// compatible with DNS_SUBDOMAIN but not with DNS_LABEL.
	//
	// For example, `Deployment` names must comply with DNS_SUBDOMAIN while
	// `Service` names must comply with DNS_LABEL.
	//
	// Since there is no formal specification for this, the default name
	// generation scheme for kubernetes objects in cdk8s was changed to DNS_LABEL,
	// since it’s the common denominator for all kubernetes resources
	// (supposedly).
	//
	// You can override this method if you wish to customize object names at the
	// chart level.
	// Experimental.
	GenerateObjectName(apiObject cdk8s.ApiObject) *string
	// Renders this chart to a set of Kubernetes JSON resources.
	//
	// Returns: array of resource manifests.
	// Experimental.
	ToJson() *[]interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MetaflowService
type jsiiProxy_MetaflowService struct {
	internal.Type__cdk8sChart
}

func (j *jsiiProxy_MetaflowService) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetaflowService) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
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

func (j *jsiiProxy_MetaflowService) ServiceAccount() cdk8splus21.IServiceAccount {
	var returns cdk8splus21.IServiceAccount
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}


// Experimental.
func NewMetaflowService(scope constructs.Construct, name *string, props *MetaflowChartProps) MetaflowService {
	_init_.Initialize()

	j := jsiiProxy_MetaflowService{}

	_jsii_.Create(
		"cdk8s-metaflow.MetaflowService",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

// Experimental.
func NewMetaflowService_Override(m MetaflowService, scope constructs.Construct, name *string, props *MetaflowChartProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-metaflow.MetaflowService",
		[]interface{}{scope, name, props},
		m,
	)
}

// Return whether the given object is a Chart.
//
// We do attribute detection since we can't reliably use 'instanceof'.
// Experimental.
func MetaflowService_IsChart(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-metaflow.MetaflowService",
		"isChart",
		[]interface{}{x},
		&returns,
	)

	return returns
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

// Finds the chart in which a node is defined.
// Experimental.
func MetaflowService_Of(c constructs.IConstruct) cdk8s.Chart {
	_init_.Initialize()

	var returns cdk8s.Chart

	_jsii_.StaticInvoke(
		"cdk8s-metaflow.MetaflowService",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetaflowService) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		m,
		"addDependency",
		args,
	)
}

func (m *jsiiProxy_MetaflowService) GenerateObjectName(apiObject cdk8s.ApiObject) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"generateObjectName",
		[]interface{}{apiObject},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetaflowService) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		m,
		"toJson",
		nil, // no parameters
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

