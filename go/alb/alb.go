package alb

import (
	terraformelb "cdk.tf/go/stack/generated/hashicorp/aws/elb"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func CreateAlb(stack cdktf.TerraformStack) cdktf.TerraformStack {
	subnetsIds := []string{"subnet-10621875", "subnet-1c9cb65a"}

	newAlb := terraformelb.NewAlb(stack, jsii.String("aws_lb"), &terraformelb.AlbConfig{
		Name:             jsii.String("cencuri-load-balancer"),
		LoadBalancerType: jsii.String("application"),
		Subnets:          jsii.Strings(subnetsIds...),
	})

	newAlbTargetGroup := terraformelb.NewAlbTargetGroup(stack, jsii.String("aws_lb_target_group"), &terraformelb.AlbTargetGroupConfig{
		Name:       jsii.String("cencuri-load-balancer-tg"),
		Port:       jsii.Number(4000),
		Protocol:   jsii.String("HTTP"),
		VpcId:      jsii.String("vpc-6f9b120a"),
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

	terraformelb.NewLbListener(stack, jsii.String("aws_lb_listener"), &terraformelb.LbListenerConfig{
		Port:            jsii.Number(80),
		Protocol:        jsii.String("HTTP"),
		LoadBalancerArn: newAlb.Arn(),
		DefaultAction: []interface{}{
			map[string]interface{}{
				"type":           "forward",
				"targetGroupArn": *newAlbTargetGroup.Arn(),
			},
		},
	})

	// Add cluster, add task definition, in that task definition add the container, then run the tasks

	return stack
}
