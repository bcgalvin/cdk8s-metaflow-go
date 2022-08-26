// cdk8s-metaflow
package cdk8smetaflow


// Experimental.
type LocalPostgresAddonProps struct {
	// Experimental.
	EndpointIp *string `field:"required" json:"endpointIp" yaml:"endpointIp"`
	// Experimental.
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
}

