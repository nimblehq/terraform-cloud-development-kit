package alb

import (
	terraformelb "cdk.tf/go/stack/generated/hashicorp/aws/elb"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func CreateAlb(stack cdktf.TerraformStack) cdktf.TerraformStack {
	subnetsIds := []string{"subnet-dd3c94bb", "subnet-d505dd9d", "subnet-9fb2d5c6"}
	terraformelb.NewAlb(stack, jsii.String("aws_lb"), &terraformelb.AlbConfig{
		Name:             jsii.String("cencuri-load-balancer"),
		LoadBalancerType: jsii.String("application"),
		Subnets:          jsii.Strings(subnetsIds...),
	})

	terraformelb.NewAlbTargetGroup(stack, jsii.String("aws_lb_target_group"), &terraformelb.AlbTargetGroupConfig{
		Name:       jsii.String("cencuri-load-balancer-tg"),
		Port:       jsii.Number(4000),
		Protocol:   jsii.String("HTTP"),
		VpcId:      jsii.String("vpc-68d2150e"),
		TargetType: jsii.String("ip"),
		HealthCheck: &terraformelb.AlbTargetGroupHealthCheck{
			Interval:           jsii.Number(70),
			Path:               jsii.String("/"),
			Port:               jsii.String("4000"),
			HealthyThreshold:   jsii.Number(2),
			UnhealthyThreshold: jsii.Number(2),
			Timeout:            jsii.Number(65),
			Protocol:           jsii.String("HTTP"),
			Matcher:            jsii.String("200"),
		},
	})

	return stack
}
