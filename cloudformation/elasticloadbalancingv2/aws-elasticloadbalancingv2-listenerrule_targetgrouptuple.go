package elasticloadbalancingv2

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// ListenerRule_TargetGroupTuple AWS CloudFormation Resource (AWS::ElasticLoadBalancingV2::ListenerRule.TargetGroupTuple)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-targetgrouptuple.html
type ListenerRule_TargetGroupTuple struct {

	// TargetGroupArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-targetgrouptuple.html#cfn-elasticloadbalancingv2-listenerrule-targetgrouptuple-targetgrouparn
	TargetGroupArn string `json:"TargetGroupArn,omitempty"`

	// Weight AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-targetgrouptuple.html#cfn-elasticloadbalancingv2-listenerrule-targetgrouptuple-weight
	Weight int `json:"Weight,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *ListenerRule_TargetGroupTuple) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::ListenerRule.TargetGroupTuple"
}
