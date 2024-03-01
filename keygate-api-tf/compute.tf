resource "aws_instance" "instance_1" {
        ami = "ami-0e731c8a588258d0d"
        instance_type = var.instance_type
        security_groups = [aws_security_group.instances.name]
        user_data       = <<-EOF
                #!/bin/bash
                echo "What if Harambe never died ...? (EC2 instance 1)" > index.html
                python3 -m http.server 8080 &
                EOF

        tags = {
                Name = "${var.environment_name}-ec2-1"
        }
}

resource "aws_instance" "instance_2" {
        ami = "ami-0e731c8a588258d0d"
        instance_type = var.instance_type
        security_groups = [aws_security_group.instances.name]
        user_data       = <<-EOF
                #!/bin/bash
                echo "Would the world be any different if Harambe never died? (EC2 instance 2)" > index.html
                python3 -m http.server 8080 &
                EOF

        tags = {
        Name = "${var.environment_name}-ec2-1"
        }
} 