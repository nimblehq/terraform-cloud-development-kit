import { Construct } from "constructs";
import { App, TerraformStack } from "cdktf";

class CentauriStack extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);

    // define resources here
  }
}

const app = new App();
new CentauriStack(app, "tmp");
app.synth();
