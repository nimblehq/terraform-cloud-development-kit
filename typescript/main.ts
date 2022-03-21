import { Construct } from 'constructs';
import { App, RemoteBackend, TerraformStack } from 'cdktf';
import { AwsProvider } from '@cdktf/provider-aws'
import { RandomProvider } from '@cdktf/provider-random'

import { Rds, Alb } from './construct'

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

    new Rds(this, 'centauri-typescript-db', {
      identifier: 'centauri-typescript-db',
      engine: 'postgres',
      engineVersion: '14.2',
      instanceClass: 'db.t3.small',
      allocateStorage: 5,
      name: 'centauritypescript',
      username: 'postgres'
    })

    new Alb(this, 'centauri-typescript-alb', {
      name: 'centauri-typescript-alb',
      subnetIds: ['subnet-10621875', 'subnet-1c9cb65a'],
      vpcId: 'vpc-6f9b120a'
    })
  }
}

const app = new App();
new CentauriStack(app, "centauri-typescript");
app.synth();
