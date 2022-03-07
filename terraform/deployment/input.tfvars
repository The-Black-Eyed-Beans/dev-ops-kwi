vpc_cidr = "12.22.0.0/16"

vpc_sg = {
    sg1 = {
        name = "Gateway security group"
        description = "gateway"
        from_port = 80
        to_port = 80
        protocol = "tcp"
    },
    sg2 = {
        name = "User Microservice security group"
        description = "user"
        from_port = 8070
        to_port = 8070
        protocol = "tcp"
    },
    sg3 = {
        name = "Underwriter Microservice security group"
        description = "underwriter"
        from_port = 8071
        to_port = 8071
        protocol = "tcp"
    },
    sg4 = {
        name = "Transaction Microservice security group"
        description = "transaction"
        from_port = 8073
        to_port = 8073
        protocol = "tcp"
    },
    sg5 = {
        name = "Bank Microservice security group"
        description = "bank"
        from_port = 8083
        to_port = 8083
        protocol = "tcp"
    }
}

public = {
    subnet1 = {
        cidr_block = "12.22.1.0/24"
        tag_name = "aline-kwi-public-1"
        az = "us-east-1a"
    },
    subnet2 = {
        cidr_block = "12.22.2.0/24"
        tag_name = "aline-kwi-public-2"
        az = "us-east-1b"
    }
}

private = {
    subnet3 = {
        cidr_block = "12.22.3.0/24"
        tag_name = "aline-kwi-private-1"
        az = "us-east-1a"
    },
    subnet4 = {
        cidr_block = "12.22.4.0/24"
        tag_name = "aline-kwi-private-2"
        az = "us-east-1b"
    }
}

alb = "aline-kwi-alb"
