import { Construct } from 'constructs'
import { TerraformOutput } from 'cdktf'
import { rds } from '@cdktf/provider-aws'
import { Password } from '@cdktf/provider-random'

export interface RdsConfig {
  readonly identifier: string
  readonly engine: string
  readonly engineVersion: string
  readonly instanceClass: string
  readonly allocateStorage: number
  readonly name: string
  readonly username: string
}

export class Rds extends Construct {
  public readonly instance: rds.DbInstance

  constructor(scope: Construct, name: string, config: RdsConfig) {
    super(scope, name)

    // Create a password stored in the TF State on the fly
    const password = new Password(this, `db_password`, {
      length: 16,
      special: false,
    });

    const db = new rds.DbInstance(this, 'db', {
      identifier: config.identifier,

      engine: config.engine,
      engineVersion: config.engineVersion,

      instanceClass: config.instanceClass,
      allocatedStorage: config.allocateStorage,

      name: config.name,
      username: config.username,
      password: password.result,

      skipFinalSnapshot: true
    })

    new TerraformOutput(this, 'aws_db_instance_host', {
      value: db.address
    })

    new TerraformOutput(this, 'aws_db_instance_port', {
      value: db.port
    })

    this.instance = db
  }
}
