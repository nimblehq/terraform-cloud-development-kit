package rds

import (
	terraformrds "cdk.tf/go/stack/generated/hashicorp/aws/rds"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func CreateRdsInstance(stack cdktf.TerraformStack) cdktf.TerraformStack {
	terraformrds.NewDbInstance(stack, jsii.String("aws_db_instance"), &terraformrds.DbInstanceConfig{
		Name:               jsii.String("cencuri_dev_db"),
		InstanceClass:      jsii.String("db.t3.micro"),
		AllocatedStorage:   jsii.Number(5),
		Engine:             jsii.String("postgres"),
		Username:           jsii.String("junan"),
		Password:           jsii.String("junan123x"),
		PubliclyAccessible: jsii.Bool(true),
		SkipFinalSnapshot:  jsii.Bool(true),
	})

	return stack
}
