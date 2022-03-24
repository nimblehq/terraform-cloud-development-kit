import { Construct } from 'constructs';
import { App, Fn, RemoteBackend, TerraformStack } from 'cdktf';
import { AwsProvider } from '@cdktf/provider-aws'
import { RandomProvider } from '@cdktf/provider-random'

import { Rds, Alb, Log, Ecs } from './construct'

class CentauriStack extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);

    new RemoteBackend(this, {
      hostname: 'app.terraform.io',
      organization: 'nimble',
      workspaces: {
        name: 'nimble-growth-37-centauri-web-typescript',
      },
    });

    new RandomProvider(this, "random", {});

    new AwsProvider(this, 'AWS', {
      region: 'ap-southeast-1'
    })

    const db = new Rds(this, 'centauri-typescript-db', {
      identifier: 'centauri-typescript-db',
      engine: 'postgres',
      engineVersion: '14.2',
      instanceClass: 'db.t3.small',
      allocateStorage: 5,
      name: 'centauritypescript',
      username: 'postgres'
    })

    const alb = new Alb(this, 'centauri-typescript-alb', {
      name: 'centauri-typescript-alb',
      subnetIds: ['subnet-10621875', 'subnet-1c9cb65a'],
      vpcId: 'vpc-6f9b120a'
    })

    const log = new Log(this, 'centauri-typescript-log', {
      name: 'centauri-typescript-log'
    })

    new Ecs(this, 'centauri-typescript-ecs', {
      name: 'centauri-typescript-ecs',
      image: '301618631622.dkr.ecr.ap-southeast-1.amazonaws.com/centauri:0.4.4-without-basic-auth',
      logGroup: log.logGroup.name,
      logRegion: 'ap-southeast-1',
      databaseHost: Fn.tostring(db.instance.address),
      databasePort: Fn.tostring(db.instance.port),
      databaseName: db.instance.name,
      databaseUsername: db.instance.username,
      databasePassword: db.instance.password,
      albDns: alb.lb.dnsName,
      albTargetGroupArn: alb.lbTargetGroup.arn,
      appSecretKeybase: "eD7Se7buNmNTWKhQrm2BFP3zlasjrQtg85ORlQ0PKqxtWq9FQetj8vOp1GIOZnyj",
      appPort: '4000',
      subnetIds: ['subnet-10621875', 'subnet-1c9cb65a']
    })
  }
}

const app = new App();
new CentauriStack(app, "centauri-typescript");
app.synth();
