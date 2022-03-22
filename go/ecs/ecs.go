package ecs

import (
	"encoding/json"

	terraformecs "cdk.tf/go/stack/generated/hashicorp/aws/ecs"
	"cdk.tf/go/stack/generated/hashicorp/aws/iam"
	terraformiam "cdk.tf/go/stack/generated/hashicorp/aws/iam"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func InitEcs(stack cdktf.TerraformStack) cdktf.TerraformStack {
	// New task execution Policy
	taskExecutionPolicy := terraformiam.NewDataAwsIamPolicyDocument(stack, jsii.String("ecs_task_execution_policy"), &terraformiam.DataAwsIamPolicyDocumentConfig{
		Statement: []interface{}{
			map[string]interface{}{
				"actions": []interface{}{
					"sts:AssumeRole",
				},
				"principals": []interface{}{
					map[string]interface{}{
						"type": "Service",
						"identifiers": []interface{}{
							"ecs-tasks.amazonaws.com",
						},
					},
				},
			},
		},
	})

	// New task execution Role
	taskExecutionRole := iam.NewIamRole(stack, jsii.String("ecs_task_execution_policy"), &iam.IamRoleConfig{
		Name:             jsii.String("go-task-execution-role"),
		Path:             jsii.String("/"),
		AssumeRolePolicy: taskExecutionPolicy.Json(),
	})

	// Attach the policy to the role
	iam.NewIamRolePolicyAttachment(stack, jsii.String("ecs_task_execution_role_policy_attachment"), &iam.IamRolePolicyAttachmentConfig{
		Role:      taskExecutionRole.Name(),
		PolicyArn: jsii.String("arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"),
	})

	// New ECS cluster
	terraformecs.NewEcsCluster(stack, jsii.String("aws_ecs_cluster"), &terraformecs.EcsClusterConfig{
		Name: jsii.String("go-ecs-cluster"),
	})

	definitions := []interface{}{
		map[string]interface{}{
			"name":      "Go container",
			"image":     "301618631622.dkr.ecr.ap-southeast-1.amazonaws.com/centauri:0.4.4-without-basic-auth",
			"essential": true,
			"portMappings": []interface{}{
				map[string]interface{}{
					"protocal":      "tcp",
					"containerPort": 4000,
					"hostPort":      4000,
				},
			},
			"logConfiguration": map[string]interface{}{
				"logDriver": "awslogs",
				"options": map[string]interface{}{
					"awslogs-group":         "config.logGroup",
					"awslogs-region":        "config.logRegion",
					"awslogs-stream-prefix": "config.name",
				},
			},
			"environment": []interface{}{
				map[string]interface{}{
					"name":  "DATABASE_URL",
					"value": "postgres://${config.databaseUsername}:${config.databasePassword}@${config.databaseHost}:${config.databasePort}/${config.databaseName}",
				},
				map[string]interface{}{
					"name":  "HOST",
					"value": "config.albDns",
				},
				map[string]interface{}{
					"name":  "SECRET_KEY_BASE",
					"value": "config.appSecretKeybase",
				},
				map[string]interface{}{
					"name":  "PORT",
					"value": "config.appPort",
				},
			},
		},
	}

	data, _ := json.Marshal(&definitions)
	stringData := string(data)

	// New ECS task definition
	terraformecs.NewEcsTaskDefinition(stack, jsii.String("aws_ecs_task_definition"), &terraformecs.EcsTaskDefinitionConfig{
		Family:                  jsii.String("go-ecs-cluster-definition"),
		NetworkMode:             jsii.String("awsvpc"),
		RequiresCompatibilities: jsii.Strings("FARGATE"),
		Cpu:                     jsii.String("256"),
		Memory:                  jsii.String("512"),
		ExecutionRoleArn:        taskExecutionRole.Arn(),
		ContainerDefinitions:    &stringData,
	})

	return stack
}
