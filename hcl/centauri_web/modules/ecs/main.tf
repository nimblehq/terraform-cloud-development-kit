data "aws_iam_policy_document" "ecs_task_execution" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_ecs_cluster" "main" {
  name = var.name
}

resource "aws_iam_role" "ecs_task_execution" {
  name               = var.name
  path               = "/"
  assume_role_policy = data.aws_iam_policy_document.ecs_task_execution.json
}

resource "aws_iam_role_policy_attachment" "ecs-task-execution-role-policy-attachment" {
  role       = aws_iam_role.ecs_task_execution.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_ecs_task_definition" "main" {
  family                   = var.name
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = 256
  memory                   = 512

  execution_role_arn = aws_iam_role.ecs_task_execution.arn
  task_role_arn      = aws_iam_role.ecs_task_execution.arn

  container_definitions = jsonencode([{
    name  = var.name
    image = var.image

    essential = true

    portMappings = [{
      protocol      = "tcp"
      containerPort = 4000
      hostPort      = 4000
    }]

    logConfiguration = {
      logDriver = "awslogs",
      options = {
        awslogs-group         = var.name
        awslogs-region        = var.region
        awslogs-stream-prefix = var.name
      }
    }

    environment = [
      {
        name  = "DATABASE_URL"
        value = "postgres://${var.database_username}:${var.database_password}@${var.database_host}:${var.database_port}/${var.database_name}"
      },
      {
        name  = "HOST"
        value = var.aws_alb_dns
      },
      {
        name  = "SECRET_KEY_BASE"
        value = "odb/2tIZ9+mBFI8rBMK/OWCRjLuDB+PfKtzv6Td73OfyjDPAjWWLdOmhRozk50rV"
      },
      {
        name  = "PORT"
        value = "4000"
      }
    ]

  }])
}

resource "aws_ecs_service" "main" {
  name                               = var.name
  cluster                            = aws_ecs_cluster.main.id
  task_definition                    = aws_ecs_task_definition.main.arn
  desired_count                      = 1
  deployment_minimum_healthy_percent = 0
  deployment_maximum_percent         = 100
  launch_type                        = "FARGATE"
  scheduling_strategy                = "REPLICA"

  network_configuration {
    subnets          = var.subnet_ids
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = var.aws_alb_target_group_arn
    container_name   = var.name
    container_port   = 4000
  }
}
