// cdk8s-metaflow
package cdk8smetaflow


// Experimental.
type MetadataDatabaseOptions struct {
	// Experimental.
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// Experimental.
	Auth *DatabaseAuthOptions `field:"optional" json:"auth" yaml:"auth"`
	// Experimental.
	Metrics *DatabaseMetricsOptions `field:"optional" json:"metrics" yaml:"metrics"`
	// Experimental.
	Replication *DatabaseReplicationOptions `field:"optional" json:"replication" yaml:"replication"`
	// Experimental.
	Resources *DatabaseResourcesOptions `field:"optional" json:"resources" yaml:"resources"`
	// Experimental.
	VolumePermissions *DatabaseVolumePermissionsOptions `field:"optional" json:"volumePermissions" yaml:"volumePermissions"`
}

