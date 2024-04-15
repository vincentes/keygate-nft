resource "tls_private_key" "example" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "aws_key_pair" "generated_key" {
  key_name   = "keygate-ssh"
  public_key = tls_private_key.example.public_key_openssh
}

resource "aws_instance" "instance_1" {
  ami           = "ami-051f8a213df8bc089"
  instance_type = var.instance_type
  security_groups = [aws_security_group.instances.name]
  key_name = aws_key_pair.generated_key.key_name

  tags = {
    Name = "${var.environment_name}-ec2-2"
  }
}

resource "aws_instance" "instance_2" {
  ami           = "ami-051f8a213df8bc089"
  instance_type = var.instance_type
  security_groups = [aws_security_group.instances.name]
  key_name = aws_key_pair.generated_key.key_name

  tags = {
    Name = "${var.environment_name}-ec2-2"
  }
}

output "private_key" {
  value     = tls_private_key.example.private_key_pem
  sensitive = true
}