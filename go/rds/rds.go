package rds

import (
	terraformvariable "cdk.tf/go/stack/variable"

	terraformrds "cdk.tf/go/stack/generated/hashicorp/aws/rds"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func CreateRdsInstance(stack cdktf.TerraformStack) (cdktf.TerraformStack, terraformrds.DbInstance) {
	db := terraformrds.NewDbInstance(stack, jsii.String("aws_db_instance"), &terraformrds.DbInstanceConfig{
		Name:               jsii.String(*terraformvariable.GetDbName(stack)),
		InstanceClass:      jsii.String("db.t3.micro"),
		AllocatedStorage:   jsii.Number(5),
		Engine:             jsii.String("postgres"),
		Username:           jsii.String(*terraformvariable.GetDbUsername(stack)),
		Password:           jsii.String(*terraformvariable.GetDbPassword(stack)),
		PubliclyAccessible: jsii.Bool(true),
		SkipFinalSnapshot:  jsii.Bool(true),
	})

	return stack, db
}
