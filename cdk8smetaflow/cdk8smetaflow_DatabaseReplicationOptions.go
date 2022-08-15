// cdk8s-metaflow
package cdk8smetaflow


// Experimental.
type DatabaseReplicationOptions struct {
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Experimental.
	ReadReplicas *float64 `field:"optional" json:"readReplicas" yaml:"readReplicas"`
}

