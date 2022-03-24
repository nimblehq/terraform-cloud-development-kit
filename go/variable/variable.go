package variable

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func GetDbUsername(stack cdktf.TerraformStack) *string {
	username := cdktf.NewTerraformVariable(stack, jsii.String("DB_USERNAME"), &cdktf.TerraformVariableConfig{
		Type:      jsii.String("string"),
		Sensitive: jsii.Bool(true),
	})

	return username.StringValue()
}

func GetDbPassword(stack cdktf.TerraformStack) *string {
	password := cdktf.NewTerraformVariable(stack, jsii.String("DB_PASSWORD"), &cdktf.TerraformVariableConfig{
		Type:      jsii.String("string"),
		Sensitive: jsii.Bool(true),
	})

	return password.StringValue()
}

func GetDbName(stack cdktf.TerraformStack) *string {
	dbName := cdktf.NewTerraformVariable(stack, jsii.String("DB_NAME"), &cdktf.TerraformVariableConfig{
		Type:      jsii.String("string"),
		Sensitive: jsii.Bool(true),
	})

	return dbName.StringValue()
}

func GetSecretKeyBase(stack cdktf.TerraformStack) *string {
	secret := cdktf.NewTerraformVariable(stack, jsii.String("SECRET_KEY_BASE"), &cdktf.TerraformVariableConfig{
		Type:      jsii.String("string"),
		Sensitive: jsii.Bool(true),
	})

	return secret.StringValue()
}

func GetDockerImage(stack cdktf.TerraformStack) *string {
	image := cdktf.NewTerraformVariable(stack, jsii.String("DOCKER_IMAGE"), &cdktf.TerraformVariableConfig{
		Type:    jsii.String("string"),
		Default: jsii.String("301618631622.dkr.ecr.ap-southeast-1.amazonaws.com/centauri:0.4.4-without-basic-auth"),
	})

	return image.StringValue()
}

func VpcId(stack cdktf.TerraformStack) *string {
	vpcId := cdktf.NewTerraformVariable(stack, jsii.String("VPC_ID"), &cdktf.TerraformVariableConfig{
		Type:    jsii.String("string"),
		Default: jsii.String("vpc-6f9b120a"),
	})

	return vpcId.StringValue()
}

func GetSubnetIds(stack cdktf.TerraformStack) *string {
	subnetIds := cdktf.NewTerraformVariable(stack, jsii.String("SUBNET_IDS"), &cdktf.TerraformVariableConfig{
		Type:    jsii.String("list(string)"),
		Default: []*string{jsii.String("subnet-10621875"), jsii.String("subnet-1c9cb65a")},
	})

	return subnetIds.StringValue()
}
