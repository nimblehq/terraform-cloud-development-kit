resource "aws_lb" "app" {
  name               = var.name
  load_balancer_type = "application"

  subnets = var.subnet_ids
}

resource "aws_lb_target_group" "app" {
  name                 = var.name
  port                 = 4000
  protocol             = "HTTP"
  vpc_id               = var.vpc_id
  deregistration_delay = 100 # Given the instance 100 seconds to finish the queued requests before removing out the ALB
  target_type          = "ip"

  health_check {
    interval            = 70
    path                = "/"
    port                = 4000
    healthy_threshold   = 2
    unhealthy_threshold = 2
    timeout             = 65
    protocol            = "HTTP"
    matcher             = "200"
  }
}

resource "aws_lb_listener" "app_http" {
  load_balancer_arn = aws_lb.app.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.app.arn
  }
}
