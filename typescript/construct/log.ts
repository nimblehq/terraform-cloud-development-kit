import { Construct } from 'constructs'
import { TerraformOutput } from 'cdktf'
import { cloudwatch } from '@cdktf/provider-aws'

export interface LogConfig {
  readonly name: string
}

export class Log extends Construct {
  public readonly logGroup: cloudwatch.CloudwatchLogGroup

  constructor(scope: Construct, id: string, config: LogConfig) {
    super(scope, id)

    this.logGroup = new cloudwatch.CloudwatchLogGroup(this, id, {
      name: config.name,
      retentionInDays: 14
    })

    new TerraformOutput(this, 'cloudwatch_log_group_arn', {
      value: this.logGroup.arn
    })

    new TerraformOutput(this, 'cloudwatch_log_group_name', {
      value: this.logGroup.name
    })
  }
}
