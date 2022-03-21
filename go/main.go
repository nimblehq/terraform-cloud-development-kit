package main

import (
	alb "cdk.tf/go/stack/alb"
	rds "cdk.tf/go/stack/rds"

	"cdk.tf/go/stack/generated/hashicorp/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	aws.NewAwsProvider(stack, jsii.String("AWS"), &aws.AwsProviderConfig{
		Region: jsii.String("ap-southeast-1"),
	})

	// Creating a new RDS instance
	stack = rds.CreateRdsInstance(stack)
	// Creating a new ELB
	stack = alb.CreateAlb(stack)

	return stack
}

func main() {
	app := cdktf.NewApp(nil)
	stack := NewMyStack(app, "aws_instance")
	cdktf.NewRemoteBackend(stack, &cdktf.RemoteBackendProps{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("nimble"),
		Workspaces:   cdktf.NewNamedRemoteWorkspace(jsii.String("nimble-growth-37-centauri-web-go")),
	})

	app.Synth()
}