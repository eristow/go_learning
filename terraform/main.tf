resource "aws_iam_role" "ecs_execution_role" {
  name = "ecs-execution-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      }
    ]
  })
}

data "aws_secretsmanager_secret" "gitlab_deploy_token" {
  name = "gitlab_deploy_token_go_learning"
}

resource "aws_iam_policy" "secrets_access" {
  name        = "ecs-secrets-access"
  description = "Allow ECS to access GitLab deploy token in Secrets Manager"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "secretsmanager:GetSecretValue",
          "secretsmanager:DescribeSecret"
        ],
        Resource = data.aws_secretsmanager_secret.gitlab_deploy_token.arn
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ecs_execution_secrets_access" {
  role       = aws_iam_role.ecs_execution_role.name
  policy_arn = aws_iam_policy.secrets_access.arn
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_role_policy" {
  role       = aws_iam_role.ecs_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_iam_role_policy_attachment" "ecs_logs" {
  role       = aws_iam_role.ecs_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/CloudWatchLogsFullAccess"
}

resource "aws_security_group" "lb_sg" {
  name        = "lb-security-group"
  description = "Security group for the load balancer"
  vpc_id      = aws_vpc.app_vpc.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 3000
    to_port     = 3000
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "app_sg" {
  name        = "app-security-group"
  description = "Security group for the application"
  vpc_id      = aws_vpc.app_vpc.id

  # RDB PostgreSQL access
  ingress {
    from_port = 5432
    to_port   = 5432
    protocol  = "tcp"
    self      = true
  }

  # Container-to-container access
  ingress {
    from_port = 8080
    to_port   = 8080
    protocol  = "tcp"
    self      = true
  }

  # Frontend access
  ingress {
    from_port       = 3000
    to_port         = 3000
    protocol        = "tcp"
    security_groups = [aws_security_group.lb_sg.id]
  }

  # Backend access
  ingress {
    from_port       = 8080
    to_port         = 8080
    protocol        = "tcp"
    security_groups = [aws_security_group.lb_sg.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_vpc" "app_vpc" {
  cidr_block = "10.0.0.0/16"

  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name = "app-vpc"
  }
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.app_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.app_igw.id
  }

  tags = {
    Name = "app-route-table"
  }
}

resource "aws_route_table_association" "public_subnet" {
  count = 2

  subnet_id      = aws_subnet.app_subnets[count.index].id
  route_table_id = aws_route_table.public.id
}

resource "aws_internet_gateway" "app_igw" {
  vpc_id = aws_vpc.app_vpc.id

  tags = {
    Name = "app-igw"
  }
}

# TODO: switch to private subnets and NAT Gateway
resource "aws_subnet" "app_subnets" {
  count = 2

  vpc_id                  = aws_vpc.app_vpc.id
  cidr_block              = "10.0.${count.index + 1}.0/24"
  availability_zone       = "us-east-1${["a", "b"][count.index]}"
  map_public_ip_on_launch = true

  tags = {
    Name = "app-subnets-${count.index + 1}"
  }
}

resource "aws_db_subnet_group" "app_db_subnet_group" {
  name       = "app-db-subnet-group"
  subnet_ids = aws_subnet.app_subnets[*].id

  tags = {
    Name = "app-db-subnet-group"
  }
}

resource "aws_service_discovery_private_dns_namespace" "app" {
  name        = "internal.app"
  description = "Internal service discovery"
  vpc         = aws_vpc.app_vpc.id
}

resource "aws_service_discovery_service" "backend" {
  name = "backend"

  dns_config {
    namespace_id = aws_service_discovery_private_dns_namespace.app.id

    dns_records {
      type = "A"
      ttl  = 10
    }
    routing_policy = "MULTIVALUE"
  }

  health_check_custom_config {
    failure_threshold = 1
  }
}


resource "aws_db_instance" "app_db" {
  allocated_storage      = 10
  storage_type           = "gp2"
  engine                 = "postgres"
  engine_version         = "17.4"
  instance_class         = "db.t4g.micro"
  db_name                = "golearningdb"
  identifier             = "golearningdb"
  username               = var.db_username
  password               = var.db_password
  vpc_security_group_ids = [aws_security_group.app_sg.id]
  db_subnet_group_name   = aws_db_subnet_group.app_db_subnet_group.name

  skip_final_snapshot       = false
  final_snapshot_identifier = "go-learning-final-snapshot-${formatdate("YYYYMMDDhhmmss", timestamp())}"
}

// Only run the DB setup if the database instance is for dev
# resource "null_resource" "db_setup" {
#   depends_on = [aws_db_instance.app_db]

#   provisioner "local-exec" {
#     command = "psql -h ${aws_db_instance.app_db.address} -p ${aws_db_instance.app_db.port} -U ${var.db_username} -d ${aws_db_instance.app_db.db_name} -f ${path.module}/../db/seed.sql"

#     environment = {
#       PGPASSWORD = var.db_password
#     }
#   }
# }


locals {
  db_connection_string = "postgres://${var.db_username}:${var.db_password}@${aws_db_instance.app_db.endpoint}/${aws_db_instance.app_db.db_name}"
}

resource "aws_ecs_cluster" "go_learning_cluster" {
  name = "go-learning-cluster"
}

resource "aws_lb" "app_lb" {
  name               = "app-lb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.lb_sg.id]
  subnets            = aws_subnet.app_subnets[*].id

  tags = {
    Name = "app-lb"
  }
}

resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.app_lb.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.frontend_tg.arn
  }
}

resource "aws_lb_listener_rule" "backend_rule" {
  listener_arn = aws_lb_listener.http.arn
  priority     = 100

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.backend_tg.arn
  }

  condition {
    path_pattern {
      values = ["/api/v1/*"]
    }
  }

}

resource "aws_ecs_task_definition" "go_learning_backend" {
  family                   = "go-learning-backend"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  execution_role_arn       = aws_iam_role.ecs_execution_role.arn

  container_definitions = jsonencode([
    {
      name  = "backend"
      image = "registry.gitlab.com/eristow_dev/go_learning/backend:latest"

      repositoryCredentials = {
        credentialsParameter = data.aws_secretsmanager_secret.gitlab_deploy_token.arn
      }

      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]

      environment = [
        {
          name  = "DATABASE_URL"
          value = local.db_connection_string
        }
      ]

      logConfiguration = {
        logDriver = "awslogs"

        options = {
          "awslogs-group"         = "/ecs/go-learning-backend"
          "awslogs-region"        = var.aws_region
          "awslogs-stream-prefix" = "ecs"
          "awslogs-create-group"  = "true"
        }
      }
    }
  ])
}

resource "aws_lb_target_group" "backend_tg" {
  name        = "backend-tg"
  target_type = "ip"
  port        = 8080
  # TODO: switch to HTTPS using AWS Certificate Manager
  protocol = "HTTP"
  vpc_id   = aws_vpc.app_vpc.id

  health_check {
    path                = "/api/v1/health"
    interval            = 30
    timeout             = 5
    healthy_threshold   = 3
    unhealthy_threshold = 3
    matcher             = "200-299"
  }
}

resource "aws_ecs_service" "go_learning_backend" {
  name            = "go-learning-backend"
  cluster         = aws_ecs_cluster.go_learning_cluster.id
  task_definition = aws_ecs_task_definition.go_learning_backend.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = aws_subnet.app_subnets[*].id
    security_groups  = [aws_security_group.app_sg.id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.backend_tg.arn
    container_name   = "backend"
    container_port   = 8080
  }

  service_registries {
    registry_arn   = aws_service_discovery_service.backend.arn
    container_name = "backend"
  }

  depends_on = [aws_db_instance.app_db]
}


resource "aws_ecs_task_definition" "go_learning_frontend" {
  family                   = "go-learning-frontend"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  execution_role_arn       = aws_iam_role.ecs_execution_role.arn

  container_definitions = jsonencode([
    {
      name  = "frontend"
      image = "registry.gitlab.com/eristow_dev/go_learning/frontend:latest"
      repositoryCredentials = {
        credentialsParameter = data.aws_secretsmanager_secret.gitlab_deploy_token.arn
      }

      portMappings = [
        {
          containerPort = 3000
          hostPort      = 3000
        }
      ]

      environment = [
        {
          name = "PUBLIC_BACKEND_URL"
          # TODO: switch to HTTPS using AWS Certificate Manager
          value = "http://backend.internal.app:8080/api/v1"
        },
        {
          name  = "PROTOCOL_HEADER"
          value = "x-forwarded-proto"
        },
        {
          name  = "HOST_HEADER"
          value = "x-forwarded-host"
        },
        {
          name  = "ORIGIN"
          value = "http://${aws_lb.app_lb.dns_name}"
        }
      ]

      logConfiguration = {
        logDriver = "awslogs"

        options = {
          "awslogs-group"         = "/ecs/go-learning-frontend"
          "awslogs-region"        = var.aws_region
          "awslogs-stream-prefix" = "ecs"
          "awslogs-create-group"  = "true"
        }
      }
    }
  ])
}

resource "aws_lb_target_group" "frontend_tg" {
  name        = "frontend-tg"
  target_type = "ip"
  port        = 3000
  # TODO: switch to HTTPS using AWS Certificate Manager
  protocol = "HTTP"
  vpc_id   = aws_vpc.app_vpc.id

  health_check {
    path                = "/health"
    interval            = 30
    timeout             = 5
    healthy_threshold   = 3
    unhealthy_threshold = 3
    matcher             = "200-299"
  }
}

resource "aws_ecs_service" "go_learning_frontend" {
  name            = "go-learning-frontend"
  cluster         = aws_ecs_cluster.go_learning_cluster.id
  task_definition = aws_ecs_task_definition.go_learning_frontend.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = aws_subnet.app_subnets[*].id
    security_groups  = [aws_security_group.app_sg.id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.frontend_tg.arn
    container_name   = "frontend"
    container_port   = 3000
  }

  depends_on = [aws_db_instance.app_db, aws_ecs_service.go_learning_backend]
}
