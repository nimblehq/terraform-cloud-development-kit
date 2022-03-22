package ecs

import (
	"encoding/json"
	"fmt"

	terraformcloudwatch "cdk.tf/go/stack/generated/hashicorp/aws/cloudwatch"
	"cdk.tf/go/stack/generated/hashicorp/aws/ecs"
	terraformecs "cdk.tf/go/stack/generated/hashicorp/aws/ecs"
	terraformalb "cdk.tf/go/stack/generated/hashicorp/aws/elb"
	terraformrds "cdk.tf/go/stack/generated/hashicorp/aws/rds"

	"cdk.tf/go/stack/generated/hashicorp/aws/iam"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func InitEcs(stack cdktf.TerraformStack, db terraformrds.DbInstance, loadBalancer terraformalb.Alb, targetGroup terraformalb.AlbTargetGroup) cdktf.TerraformStack {
	// New task execution Policy
	taskExecutionPolicy := iam.NewDataAwsIamPolicyDocument(stack, jsii.String("ecs_task_execution_policy"), &iam.DataAwsIamPolicyDocumentConfig{
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
	taskExecutionRole := iam.NewIamRole(stack, jsii.String("ecs_task_execution_role"), &iam.IamRoleConfig{
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
	cluster := terraformecs.NewEcsCluster(stack, jsii.String("aws_ecs_cluster"), &terraformecs.EcsClusterConfig{
		Name: jsii.String("go-ecs-cluster"),
	})

	cloudWatchLogGroup := terraformcloudwatch.NewCloudwatchLogGroup(stack, jsii.String("aws_cloudwatch_log_group"), &terraformcloudwatch.CloudwatchLogGroupConfig{
		Name:            jsii.String("go-cloud-watch-log-group"),
		RetentionInDays: jsii.Number(14),
	})

	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", *db.Username(), *db.Password(), *db.Address(), 5432, *db.Name())

	definitions := []interface{}{
		map[string]interface{}{
			"name":      "Go-container",
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
					"awslogs-group":         cloudWatchLogGroup.Name(),
					"awslogs-region":        "ap-southeast-1",
					"awslogs-stream-prefix": "go-cencuri-ptrefix",
				},
			},
			"environment": []interface{}{
				map[string]interface{}{
					"name":  "DATABASE_URL",
					"value": databaseUrl,
				},
				map[string]interface{}{
					"name":  "HOST",
					"value": loadBalancer.DnsName(),
				},
				map[string]interface{}{
					"name":  "SECRET_KEY_BASE",
					"value": "eD7Se7buNmNTWKhQrm2BFP3zlasjrQtg85ORlQ0PKqxtWq9FQetj8vOp1GIOZnyj",
				},
				map[string]interface{}{
					"name":  "PORT",
					"value": "4000",
				},
			},
		},
	}

	data, _ := json.Marshal(&definitions)
	stringData := string(data)
	subnetsIds := []*string{jsii.String("subnet-10621875"), jsii.String("subnet-1c9cb65a")}

	// New ECS task definition
	taskDefinition := terraformecs.NewEcsTaskDefinition(stack, jsii.String("aws_ecs_task_definition"), &terraformecs.EcsTaskDefinitionConfig{
		Family:                  jsii.String("go-ecs-cluster-definition"),
		NetworkMode:             jsii.String("awsvpc"),
		RequiresCompatibilities: jsii.Strings("FARGATE"),
		Cpu:                     jsii.String("256"),
		Memory:                  jsii.String("512"),
		ExecutionRoleArn:        taskExecutionRole.Arn(),
		ContainerDefinitions:    &stringData,
	})

	ecs.NewEcsService(stack, jsii.String("ecs_service"), &terraformecs.EcsServiceConfig{
		Name:                            jsii.String("go-ecs-service-name"),
		Cluster:                         cluster.Id(),
		TaskDefinition:                  taskDefinition.Arn(),
		DesiredCount:                    jsii.Number(1),
		DeploymentMinimumHealthyPercent: jsii.Number(0),
		DeploymentMaximumPercent:        jsii.Number(100),
		LaunchType:                      jsii.String("FARGATE"),
		SchedulingStrategy:              jsii.String("REPLICA"),

		NetworkConfiguration: &terraformecs.EcsServiceNetworkConfiguration{
			Subnets:        &subnetsIds,
			AssignPublicIp: jsii.Bool(true),
		},

		LoadBalancer: []interface{}{
			map[string]interface{}{
				"targetGroupArn": targetGroup.Arn(),
				"containerName":  jsii.String("Go-container"),
				"containerPort":  4000,
			},
		},
	})

	return stack
}
