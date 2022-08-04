// cdk8s-metaflow
package cdk8smetaflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/bcgalvin/cdk8s-metaflow-go/cdk8smetaflow/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/bcgalvin/cdk8s-metaflow-go/cdk8smetaflow/internal"
	"github.com/bcgalvin/cdk8s-metaflow-go/cdk8smetaflow/k8s"
)

// Experimental.
type AutoscalingOptions struct {
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Experimental.
	MaxReplicas *float64 `field:"optional" json:"maxReplicas" yaml:"maxReplicas"`
	// Experimental.
	MinReplicas *float64 `field:"optional" json:"minReplicas" yaml:"minReplicas"`
	// Experimental.
	TargetCPUUtilizationPercentage *float64 `field:"optional" json:"targetCPUUtilizationPercentage" yaml:"targetCPUUtilizationPercentage"`
}

// Experimental.
type DatabaseAuthOptions struct {
	// Experimental.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Experimental.
	EnablePostgresUser *bool `field:"optional" json:"enablePostgresUser" yaml:"enablePostgresUser"`
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Experimental.
	PostgresPassword *string `field:"optional" json:"postgresPassword" yaml:"postgresPassword"`
	// Experimental.
	ReplicationPassword *string `field:"optional" json:"replicationPassword" yaml:"replicationPassword"`
	// Experimental.
	ReplicationUsername *string `field:"optional" json:"replicationUsername" yaml:"replicationUsername"`
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// Experimental.
type DatabaseMetricsOptions struct {
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

// Experimental.
type DatabaseReplicationOptions struct {
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Experimental.
	ReadReplicas *float64 `field:"optional" json:"readReplicas" yaml:"readReplicas"`
}

// Experimental.
type DatabaseResourceRequestOptions struct {
	// Experimental.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Experimental.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

// Experimental.
type DatabaseResourcesOptions struct {
	// Experimental.
	Requests *DatabaseResourceRequestOptions `field:"optional" json:"requests" yaml:"requests"`
}

// Experimental.
type DatabaseVolumePermissionsOptions struct {
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

// Experimental.
type DbMigrationOptions struct {
	// Experimental.
	OnlyIfDbEmpty *bool `field:"optional" json:"onlyIfDbEmpty" yaml:"onlyIfDbEmpty"`
	// Experimental.
	RunOnStart *bool `field:"optional" json:"runOnStart" yaml:"runOnStart"`
}

// Experimental.
type HostOptions struct {
	// Experimental.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Experimental.
	Paths *[]*HttpIngressPath `field:"optional" json:"paths" yaml:"paths"`
}

// Experimental.
type HttpIngressPath struct {
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Experimental.
	PathType HttpIngressPathType `field:"optional" json:"pathType" yaml:"pathType"`
}

// Experimental.
type HttpIngressPathType string

const (
	// Experimental.
	HttpIngressPathType_PREFIX HttpIngressPathType = "PREFIX"
	// Experimental.
	HttpIngressPathType_EXACT HttpIngressPathType = "EXACT"
	// Experimental.
	HttpIngressPathType_IMPLEMENTATION_SPECIFIC HttpIngressPathType = "IMPLEMENTATION_SPECIFIC"
)

// Experimental.
type IApiResource interface {
	// Experimental.
	ApiGroup() *string
	// Experimental.
	ResourceName() *string
	// Experimental.
	ResourceType() *string
}

// The jsii proxy for IApiResource
type jsiiProxy_IApiResource struct {
	_ byte // padding
}

func (j *jsiiProxy_IApiResource) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiResource) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiResource) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

// Experimental.
type ImageOptions struct {
	// Experimental.
	PullPolicy ImagePullPolicy `field:"optional" json:"pullPolicy" yaml:"pullPolicy"`
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// Experimental.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

// Experimental.
type ImagePullPolicy string

const (
	// Experimental.
	ImagePullPolicy_ALWAYS ImagePullPolicy = "ALWAYS"
	// Experimental.
	ImagePullPolicy_IF_NOT_PRESENT ImagePullPolicy = "IF_NOT_PRESENT"
	// Experimental.
	ImagePullPolicy_NEVER ImagePullPolicy = "NEVER"
)

// Experimental.
type IngressOptions struct {
	// Experimental.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Experimental.
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Experimental.
	Hosts *[]*HostOptions `field:"optional" json:"hosts" yaml:"hosts"`
	// Experimental.
	Tls *k8s.IngressTls `field:"optional" json:"tls" yaml:"tls"`
}

// Experimental.
type MetadataDatabase interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MetadataDatabase
type jsiiProxy_MetadataDatabase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_MetadataDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewMetadataDatabase(scope constructs.Construct, id *string, props *MetadataDatabaseProps) MetadataDatabase {
	_init_.Initialize()

	j := jsiiProxy_MetadataDatabase{}

	_jsii_.Create(
		"cdk8s-metaflow.MetadataDatabase",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewMetadataDatabase_Override(m MetadataDatabase, scope constructs.Construct, id *string, props *MetadataDatabaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-metaflow.MetadataDatabase",
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
func MetadataDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-metaflow.MetadataDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetadataDatabase) ToString() *string {
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
type MetadataDatabaseOptions struct {
	// Experimental.
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// Experimental.
	Auth *DatabaseAuthOptions `field:"optional" json:"auth" yaml:"auth"`
	// Experimental.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Experimental.
	Metrics *DatabaseMetricsOptions `field:"optional" json:"metrics" yaml:"metrics"`
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Experimental.
	Replication *DatabaseReplicationOptions `field:"optional" json:"replication" yaml:"replication"`
	// Experimental.
	Resources *DatabaseResourcesOptions `field:"optional" json:"resources" yaml:"resources"`
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
	// Experimental.
	VolumePermissions *DatabaseVolumePermissionsOptions `field:"optional" json:"volumePermissions" yaml:"volumePermissions"`
}

// Experimental.
type MetadataDatabaseProps struct {
	// Experimental.
	ChartVersion *string `field:"required" json:"chartVersion" yaml:"chartVersion"`
	// Experimental.
	ChartValues *MetadataDatabaseOptions `field:"optional" json:"chartValues" yaml:"chartValues"`
}

// Experimental.
type MetaflowService interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MetaflowService
type jsiiProxy_MetaflowService struct {
	internal.Type__constructsConstruct
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
type MetaflowServiceOptions struct {
	// Experimental.
	Affinity *k8s.Affinity `field:"optional" json:"affinity" yaml:"affinity"`
	// Experimental.
	Autoscaling *AutoscalingOptions `field:"optional" json:"autoscaling" yaml:"autoscaling"`
	// Experimental.
	DbMigrations *DbMigrationOptions `field:"optional" json:"dbMigrations" yaml:"dbMigrations"`
	// Experimental.
	EnvFrom *[]*k8s.EnvFromSource `field:"optional" json:"envFrom" yaml:"envFrom"`
	// Experimental.
	FullnameOverride *string `field:"optional" json:"fullnameOverride" yaml:"fullnameOverride"`
	// Experimental.
	Image *ImageOptions `field:"optional" json:"image" yaml:"image"`
	// Experimental.
	ImagePullSecrets *[]*string `field:"optional" json:"imagePullSecrets" yaml:"imagePullSecrets"`
	// Experimental.
	Ingress *IngressOptions `field:"optional" json:"ingress" yaml:"ingress"`
	// Experimental.
	Metadatadb *MetadataDatabaseOptions `field:"optional" json:"metadatadb" yaml:"metadatadb"`
	// Experimental.
	NameOverride *string `field:"optional" json:"nameOverride" yaml:"nameOverride"`
	// Experimental.
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	// Experimental.
	PodAnnotations *map[string]*string `field:"optional" json:"podAnnotations" yaml:"podAnnotations"`
	// Experimental.
	PodSecurityContext *map[string]*string `field:"optional" json:"podSecurityContext" yaml:"podSecurityContext"`
	// Experimental.
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// Experimental.
	Resources *[]IApiResource `field:"optional" json:"resources" yaml:"resources"`
	// Experimental.
	SecurityContext *k8s.SecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	// Experimental.
	Service *ServiceOptions `field:"optional" json:"service" yaml:"service"`
	// Experimental.
	ServiceAccount *ServiceAccountOptions `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Experimental.
	Tolerations *[]*k8s.Toleration `field:"optional" json:"tolerations" yaml:"tolerations"`
}

// Experimental.
type MetaflowServiceProps struct {
	// Experimental.
	ChartVersion *string `field:"required" json:"chartVersion" yaml:"chartVersion"`
	// Experimental.
	ChartValues *MetaflowServiceOptions `field:"optional" json:"chartValues" yaml:"chartValues"`
}

// Experimental.
type MetaflowUI interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MetaflowUI
type jsiiProxy_MetaflowUI struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_MetaflowUI) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewMetaflowUI(scope constructs.Construct, id *string, props *MetaflowUIProps) MetaflowUI {
	_init_.Initialize()

	j := jsiiProxy_MetaflowUI{}

	_jsii_.Create(
		"cdk8s-metaflow.MetaflowUI",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewMetaflowUI_Override(m MetaflowUI, scope constructs.Construct, id *string, props *MetaflowUIProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-metaflow.MetaflowUI",
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
func MetaflowUI_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-metaflow.MetaflowUI",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetaflowUI) ToString() *string {
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
type MetaflowUIOptions struct {
	// Experimental.
	Env *map[string]*string `field:"required" json:"env" yaml:"env"`
	// Experimental.
	MetaflowDatastoreSysrootS3 *string `field:"required" json:"metaflowDatastoreSysrootS3" yaml:"metaflowDatastoreSysrootS3"`
	// Experimental.
	Affinity *k8s.Affinity `field:"optional" json:"affinity" yaml:"affinity"`
	// Experimental.
	Autoscaling *AutoscalingOptions `field:"optional" json:"autoscaling" yaml:"autoscaling"`
	// Experimental.
	EnvFrom *[]*k8s.EnvFromSource `field:"optional" json:"envFrom" yaml:"envFrom"`
	// Experimental.
	FullnameOverride *string `field:"optional" json:"fullnameOverride" yaml:"fullnameOverride"`
	// Experimental.
	Image *ImageOptions `field:"optional" json:"image" yaml:"image"`
	// Experimental.
	ImagePullSecrets *[]*string `field:"optional" json:"imagePullSecrets" yaml:"imagePullSecrets"`
	// Experimental.
	Ingress *IngressOptions `field:"optional" json:"ingress" yaml:"ingress"`
	// Experimental.
	Metadatadb *MetadataDatabaseOptions `field:"optional" json:"metadatadb" yaml:"metadatadb"`
	// Experimental.
	NameOverride *string `field:"optional" json:"nameOverride" yaml:"nameOverride"`
	// Experimental.
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	// Experimental.
	PodAnnotations *map[string]*string `field:"optional" json:"podAnnotations" yaml:"podAnnotations"`
	// Experimental.
	PodSecurityContext *map[string]*string `field:"optional" json:"podSecurityContext" yaml:"podSecurityContext"`
	// Experimental.
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// Experimental.
	Resources *[]IApiResource `field:"optional" json:"resources" yaml:"resources"`
	// Experimental.
	SecurityContext *k8s.SecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	// Experimental.
	Service *ServiceOptions `field:"optional" json:"service" yaml:"service"`
	// Experimental.
	ServiceAccount *ServiceAccountOptions `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Experimental.
	ServiceStatic *ServiceOptions `field:"optional" json:"serviceStatic" yaml:"serviceStatic"`
	// Experimental.
	Tolerations *[]*k8s.Toleration `field:"optional" json:"tolerations" yaml:"tolerations"`
	// Experimental.
	UiImage *ImageOptions `field:"optional" json:"uiImage" yaml:"uiImage"`
}

// Experimental.
type MetaflowUIProps struct {
	// Experimental.
	ChartVersion *string `field:"required" json:"chartVersion" yaml:"chartVersion"`
	// Experimental.
	ChartValues *MetaflowUIOptions `field:"optional" json:"chartValues" yaml:"chartValues"`
}

// Experimental.
type ServiceAccountOptions struct {
	// Experimental.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Experimental.
	Create *bool `field:"optional" json:"create" yaml:"create"`
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// Experimental.
type ServiceOptions struct {
	// Experimental.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Experimental.
	Type ServiceType `field:"optional" json:"type" yaml:"type"`
}

// Experimental.
type ServiceType string

const (
	// Experimental.
	ServiceType_CLUSTER_IP ServiceType = "CLUSTER_IP"
	// Experimental.
	ServiceType_NODE_PORT ServiceType = "NODE_PORT"
	// Experimental.
	ServiceType_LOAD_BALANCER ServiceType = "LOAD_BALANCER"
	// Experimental.
	ServiceType_EXTERNAL_NAME ServiceType = "EXTERNAL_NAME"
)

