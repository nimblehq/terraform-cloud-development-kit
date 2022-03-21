import { Construct } from 'constructs'
import { TerraformOutput } from 'cdktf'
import { elb } from '@cdktf/provider-aws'

export interface AlbConfig {
  readonly name: string,
  readonly subnetIds: string[],
  readonly vpcId: string
}

export class Alb extends Construct {
  public readonly lb: elb.Lb
  public readonly lbl: elb.LbListener
  public readonly lbTargetGroup: elb.LbTargetGroup

  constructor(scope: Construct, name: string, config: AlbConfig) {
    super(scope, name)

    this.lb = new elb.Lb(this, "appalb", {
      name: config.name,
      loadBalancerType: 'application',
      subnets: config.subnetIds
    })

    this.lbTargetGroup = new elb.LbTargetGroup(this, "app_target_group", {
      name: config.name,
      port: 4000,
      protocol: 'HTTP',
      vpcId: config.vpcId,
      deregistrationDelay: '100',
      targetType: 'ip',

      healthCheck: {
        interval: 70,
        path: '/',
        port: '4000',
        healthyThreshold: 2,
        timeout: 65,
        protocol: 'HTTP',
        matcher: '200'
      }
    })

    this.lbl = new elb.LbListener(this, "app_http_listener", {
      loadBalancerArn: this.lb.arn,
      port: 80,
      protocol: 'HTTP',

      defaultAction: [
        {
          type: 'forward',
          targetGroupArn: this.lbTargetGroup.arn
        }
      ]
    })

    new TerraformOutput(this, 'alb_dns', {
      value: this.lb.dnsName
    })
  }
}
