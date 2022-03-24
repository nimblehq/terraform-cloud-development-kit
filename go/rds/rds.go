package rds

import (
	"fmt"

	terraformrds "cdk.tf/go/stack/generated/hashicorp/aws/rds"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func CreateRdsInstance(stack cdktf.TerraformStack) (cdktf.TerraformStack, terraformrds.DbInstance) {
	username := cdktf.NewTerraformVariable(stack, jsii.String("DB_USERNAME"), &cdktf.TerraformVariableConfig{
		Default: jsii.String("somessname"),
		Type:    jsii.String("string"),
	})

	test := cdktf.NewTerraformVariable(stack, jsii.String("TEST"), &cdktf.TerraformVariableConfig{
		Default: jsii.String("swss"),
		Type:    jsii.String("string"),
	})

	fmt.Println("***************")
	fmt.Printf("Result is: %+v", username)
	fmt.Printf("Result is: %s", *username.StringValue())
	fmt.Printf("\nResult is: %s", *test.StringValue())
	fmt.Println("\n***************")

	db := terraformrds.NewDbInstance(stack, jsii.String("aws_db_instance"), &terraformrds.DbInstanceConfig{
		Name:               jsii.String("cencuri_dev_db"),
		InstanceClass:      jsii.String("db.t3.micro"),
		AllocatedStorage:   jsii.Number(5),
		Engine:             jsii.String("postgres"),
		Username:           jsii.String(*username.StringValue()),
		Password:           jsii.String("junan123x"),
		PubliclyAccessible: jsii.Bool(true),
		SkipFinalSnapshot:  jsii.Bool(true),
	})

	return stack, db
}
