// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_autoscaling_lifecycle_hook", lifecycleHookDataSource)
}

// lifecycleHookDataSource returns the Terraform awscc_autoscaling_lifecycle_hook data source.
// This Terraform data source corresponds to the CloudFormation AWS::AutoScaling::LifecycleHook resource.
func lifecycleHookDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AutoScalingGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Auto Scaling group for the lifecycle hook.",
		//	  "type": "string"
		//	}
		"auto_scaling_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Auto Scaling group for the lifecycle hook.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DefaultResult
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON (default).",
		//	  "type": "string"
		//	}
		"default_result": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON (default).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: HeartbeatTimeout
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the DefaultResult property.",
		//	  "type": "integer"
		//	}
		"heartbeat_timeout": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the DefaultResult property.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LifecycleHookName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the lifecycle hook.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"lifecycle_hook_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the lifecycle hook.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LifecycleTransition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The instance state to which you want to attach the lifecycle hook.",
		//	  "type": "string"
		//	}
		"lifecycle_transition": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The instance state to which you want to attach the lifecycle hook.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NotificationMetadata
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.",
		//	  "maxLength": 1023,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"notification_metadata": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NotificationTargetARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata.",
		//	  "type": "string"
		//	}
		"notification_target_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::AutoScaling::LifecycleHook",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AutoScaling::LifecycleHook").WithTerraformTypeName("awscc_autoscaling_lifecycle_hook")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_scaling_group_name": "AutoScalingGroupName",
		"default_result":          "DefaultResult",
		"heartbeat_timeout":       "HeartbeatTimeout",
		"lifecycle_hook_name":     "LifecycleHookName",
		"lifecycle_transition":    "LifecycleTransition",
		"notification_metadata":   "NotificationMetadata",
		"notification_target_arn": "NotificationTargetARN",
		"role_arn":                "RoleARN",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}