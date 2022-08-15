// cdk8s-metaflow
package cdk8smetaflow

import (
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus22/v2"
)

// Experimental.
type MetaflowChartProps struct {
	// Labels to apply to all resources in this chart.
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The default namespace for all objects defined in this chart (directly or indirectly).
	//
	// This namespace will only apply to objects that don't have a
	// `namespace` explicitly defined for them.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Experimental.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Experimental.
	ImageTag *string `field:"required" json:"imageTag" yaml:"imageTag"`
	// Experimental.
	ServiceAccount cdk8splus22.ServiceAccount `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
	// Experimental.
	ServiceType cdk8splus22.ServiceType `field:"required" json:"serviceType" yaml:"serviceType"`
	// Experimental.
	EnvVars *map[string]*string `field:"optional" json:"envVars" yaml:"envVars"`
	// Experimental.
	InitImage *string `field:"optional" json:"initImage" yaml:"initImage"`
	// Experimental.
	InitImageTag *string `field:"optional" json:"initImageTag" yaml:"initImageTag"`
}

