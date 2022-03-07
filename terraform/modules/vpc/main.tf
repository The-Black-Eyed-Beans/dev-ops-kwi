resource "aws_vpc" "vpc" {
    cidr_block = var.vpc_cidr
    
    tags = {
        Name = "aline-kwi-vpc"
    }
}

resource "aws_internet_gateway" "gateway" {
    vpc_id = aws_vpc.vpc.id
    depends_on = [aws_vpc.vpc]

    tags = {
        Name = "aline-kwi-ig"
    }
}

resource "aws_security_group" "sg" {
    for_each = var.vpc_sg
    name = each.value["name"]
    description = each.value["description"]
    vpc_id = aws_vpc.vpc.id

    ingress {
        from_port = each.value["from_port"]
        to_port = each.value["to_port"]
        protocol = each.value["protocol"]
        cidr_blocks = ["0.0.0.0/0"]
    }

    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }

    tags = {
        Name = "aline-kwi-sg"
    }
    depends_on = [aws_vpc.vpc]
}

resource "aws_subnet" "public" {
    for_each = var.public
    vpc_id = aws_vpc.vpc.id
    cidr_block = each.value["cidr_block"]
    map_public_ip_on_launch = true
    availability_zone = each.value["az"]

    tags = {
        Name = each.value["tag_name"] == "" ? "aline-kwi-subnet" : each.value["tag_name"]
        Tier = "Public"
    }

    depends_on = [aws_vpc.vpc]
}

resource "aws_subnet" "private" {
    for_each = var.private
    vpc_id = aws_vpc.vpc.id
    cidr_block = each.value["cidr_block"]
    map_public_ip_on_launch = false
    availability_zone = each.value["az"]

    tags = {
        Name = each.value["tag_name"] == "" ? "aline-kwi-subnet" : each.value["tag_name"]
        Tier = "Private"
    }

    depends_on = [aws_vpc.vpc]
}

resource "aws_eip" "eip" {
    vpc = true
}

resource "aws_nat_gateway" "nat_gateway" {
    allocation_id = aws_eip.eip.id
    subnet_id = aws_subnet.public[element(keys(aws_subnet.public), 0)]["id"]

    tags = {
        Name = "aline-kwi-nat-gateway"
    }

    depends_on = [aws_eip.eip, aws_subnet.public]
}


resource "aws_route_table" "public_route_table" {
    vpc_id = aws_vpc.vpc.id

    route {
        cidr_block = "0.0.0.0/0"
        gateway_id = aws_internet_gateway.gateway.id
    }

    tags = {
        Name = "aline-kwi-public-RT"
    }

    depends_on = [aws_vpc.vpc, aws_internet_gateway.gateway]
}

resource "aws_route_table" "private_route_table" {
    vpc_id = aws_vpc.vpc.id

    route {
        cidr_block = "0.0.0.0/0"
        nat_gateway_id = aws_nat_gateway.nat_gateway.id
    }

    tags = {
        Name = "aline-kwi-private-RT"
    }

    depends_on = [aws_vpc.vpc, aws_nat_gateway.nat_gateway]
}

resource "aws_route_table_association" "public_association" {
    for_each = aws_subnet.public
    subnet_id = each.value["id"]
    route_table_id = aws_route_table.public_route_table.id
    depends_on = [aws_route_table.public_route_table, aws_subnet.public]
}

resource "aws_route_table_association" "private_association" {
    for_each = aws_subnet.private
    subnet_id = each.value["id"]
    route_table_id = aws_route_table.private_route_table.id
    depends_on = [aws_route_table.private_route_table, aws_subnet.private]
}

resource "aws_lb" "alb" {
    name = var.alb
    internal = false
    load_balancer_type = "application"
    security_groups = [for security_group in aws_security_group.sg : security_group.id]
    subnets = [for subnet in aws_subnet.public : subnet.id]
    depends_on = [aws_subnet.public, aws_security_group.sg]
}