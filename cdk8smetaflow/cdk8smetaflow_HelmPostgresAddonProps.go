// cdk8s-metaflow
package cdk8smetaflow


// Experimental.
type HelmPostgresAddonProps struct {
	// Experimental.
	ChartVersion *string `field:"required" json:"chartVersion" yaml:"chartVersion"`
	// Experimental.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// Experimental.
	ChartValues *MetadataDatabaseOptions `field:"optional" json:"chartValues" yaml:"chartValues"`
}

