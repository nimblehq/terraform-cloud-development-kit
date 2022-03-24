package alb

import (
	constant "cdk.tf/go/stack/constant"

	terraformelb "cdk.tf/go/stack/generated/hashicorp/aws/elb"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func CreateAlb(stack cdktf.TerraformStack) (cdktf.TerraformStack, terraformelb.Alb, terraformelb.AlbTargetGroup) {
	subnetsIds := []string{"subnet-10621875", "subnet-1c9cb65a"}

	newAlb := terraformelb.NewAlb(stack, jsii.String("aws_lb"), &terraformelb.AlbConfig{
		Name:             jsii.String(constant.Name),
		LoadBalancerType: jsii.String("application"),
		Subnets:          jsii.Strings(subnetsIds...),
	})

	newAlbTargetGroup := terraformelb.NewAlbTargetGroup(stack, jsii.String("aws_lb_target_group"), &terraformelb.AlbTargetGroupConfig{
		Name:       jsii.String(constant.Name),
		Port:       jsii.Number(4000),
		Protocol:   jsii.String(constant.Http),
		VpcId:      jsii.String("vpc-6f9b120a"),
		TargetType: jsii.String("ip"),
		HealthCheck: &terraformelb.AlbTargetGroupHealthCheck{
			Interval:           jsii.Number(70),
			Path:               jsii.String("/"),
			Port:               jsii.String("4000"),
			HealthyThreshold:   jsii.Number(2),
			UnhealthyThreshold: jsii.Number(2),
			Timeout:            jsii.Number(65),
			Protocol:           jsii.String(constant.Http),
			Matcher:            jsii.String("200"),
		},
	})

	terraformelb.NewLbListener(stack, jsii.String("aws_lb_listener"), &terraformelb.LbListenerConfig{
		Port:            jsii.Number(80),
		Protocol:        jsii.String(constant.Http),
		LoadBalancerArn: newAlb.Arn(),
		DefaultAction: []interface{}{
			map[string]interface{}{
				"type":           "forward",
				"targetGroupArn": *newAlbTargetGroup.Arn(),
			},
		},
	})

	return stack, newAlb, newAlbTargetGroup
}
