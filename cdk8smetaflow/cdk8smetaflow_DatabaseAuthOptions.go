// cdk8s-metaflow
package cdk8smetaflow


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

