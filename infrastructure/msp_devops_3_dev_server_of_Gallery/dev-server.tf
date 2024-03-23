//This Terraform template prepares development environment for Petclinic Microservices Application.
//User needs to select appropriate key name when launching the template.

provider "aws" {
  region = var.region
  //  If you have entered your credentials in AWS CLI before, you do not need to use these arguments.
}

resource "aws_launch_template" "GalleryServerLT" {
  image_id = var.ami
  instance_type = var.instance_type
  key_name = var.mykey
  vpc_security_group_ids = [aws_security_group.dev-server-sg.id]
  user_data = filebase64("gallery-userdata.sh")
}

resource "aws_instance" "GalleryServer" {
  launch_template {
    id = aws_launch_template.GalleryServerLT.id
    version = aws_launch_template.GalleryServerLT.latest_version
  }
  tags = {
    Name = var.devservertag
  }
}

resource "aws_security_group" "dev-server-sg" {
  name = var.devops_server_secgr
  tags = {
    Name = var.devops_server_secgr
  }
  dynamic "ingress" {
    for_each = var.dev-server-ports
    iterator = port
    content {
      from_port = port.value
      to_port = port.value
      protocol = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }

  egress {
    from_port =0
    protocol = "-1"
    to_port =0
    cidr_blocks = ["0.0.0.0/0"]
  }

}

output "GalleryServerDNSName" {
  value = aws_instance.GalleryServer.public_dns
}