import { Construct } from 'constructs'
import { ecs, iam } from '@cdktf/provider-aws'
import { EcsCluster } from '@cdktf/provider-aws/lib/ecs'

export interface EcsConfig {
  readonly name: string
  readonly image: string
  readonly logGroup: string
  readonly logRegion: string
  readonly databaseUsername: string
  readonly databasePassword: string
  readonly databaseHost: string
  readonly databasePort: string
  readonly databaseName: string
  readonly albDns: string
  readonly albTargetGroupArn: string
  readonly appSecretKeybase: string
  readonly appPort: string
  readonly subnetIds: string[]
}

export class Ecs extends Construct {
  private readonly ecsCluster: ecs.EcsCluster
  private taskExecutionIamPolicy: iam.DataAwsIamPolicyDocument
  private taskExecutionRole: iam.IamRole
  private taskDefinition: ecs.EcsTaskDefinition

  constructor(scope: Construct, name: string, config: EcsConfig) {
    super(scope, name)

    this.taskExecutionIamPolicy = new iam.DataAwsIamPolicyDocument(this, "ecs_task_execution_policy", {
      statement: [
        {
          actions: ['sts:AssumeRole'],
    
          principals: [
            {
              type: 'Service',
              identifiers: ['ecs-tasks.amazonaws.com']
            }
          ]
        }
      ]
    })

    this.taskExecutionRole = new iam.IamRole(this, "ecs_task_execution_role", {
      name: config.name,
      path: '/',
      assumeRolePolicy: this.taskExecutionIamPolicy.json
    })

    new iam.IamRolePolicyAttachment(this, 'ecs_task_execution_role_policy_attachment', {
      role: this.taskExecutionRole.name,
      policyArn: 'arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy'
    })

    this.ecsCluster = new EcsCluster(this, "ecs_cluster", {
      name: config.name
    })

    this.taskDefinition = new ecs.EcsTaskDefinition(this, 'ecs_task_definition', {
      family: config.name,
      networkMode: 'awsvpc',
      requiresCompatibilities: ['FARGATE'],
      cpu: '256',
      memory: '512',

      executionRoleArn: this.taskExecutionRole.arn,
      taskRoleArn: this.taskExecutionRole.arn,

      containerDefinitions: JSON.stringify([
        {
          name: config.name,
          image: config.image,

          essential: true,

          portMappings: [
            {
              protocal: 'tcp',
              containerPort: 4000,
              hostPort: 4000,
            },
          ],

          logConfiguration: {
            logDriver: 'awslogs',
            options: {
              "awslogs-group": config.logGroup,
              "awslogs-region": config.logRegion,
              "awslogs-stream-prefix": config.name,
            },
          },
          
          environment: [
            {
              name: 'DATABASE_URL',
              value: `postgres://${config.databaseUsername}:${config.databasePassword}@${config.databaseHost}:${config.databasePort}/${config.databaseName}`
            },
            {
              name: 'HOST',
              value: config.albDns,
            },
            {
              name: 'SECRET_KEY_BASE',
              value: config.appSecretKeybase
            },
            {
              name: 'PORT',
              value: config.appPort
            }
          ]
        },
      ])
    })

    new ecs.EcsService(this, 'ecs_service', {
      name: config.name,
      cluster: this.ecsCluster.id,
      taskDefinition: this.taskDefinition.arn,
      desiredCount: 1,
      deploymentMinimumHealthyPercent: 0,
      deploymentMaximumPercent: 100,
      launchType: 'FARGATE',
      schedulingStrategy: 'REPLICA',

      networkConfiguration: {
        subnets: config.subnetIds,
        assignPublicIp: true
      },

      loadBalancer: [
        {
          targetGroupArn: config.albTargetGroupArn,
          containerName: config.name,
          containerPort: 4000
        }
      ]
    })
  }
}