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
