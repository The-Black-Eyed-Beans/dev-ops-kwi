resource "aws_ecs_cluster" "ecs" {
    name = "aline-kwi-ecs"

    setting {
        name = "containerInsights"
        value = "enabled"
    }

    tags = {
        Name = "aline-kwi-ecs"
    }
}

