import { Construct } from "constructs";
import { App, RemoteBackend, TerraformStack } from "cdktf";

class CentauriStack extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);

    new RemoteBackend(this, {
      hostname: "app.terraform.io",
      organization: "nimble",
      workspaces: {
        name: "nimble-growth-37-centauri-web-typescript",
      },
    });
  }
}

const app = new App();
new CentauriStack(app, "staging");
app.synth();
