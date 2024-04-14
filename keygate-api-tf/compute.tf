resource "aws_instance" "instance_1" {
  ami           = "ami-051f8a213df8bc089"
  instance_type = var.instance_type
  security_groups = [aws_security_group.instances.name]

  user_data = <<-EOF
#!/bin/bash
# Install dependencies
sudo yum install -y golang git

# Create a directory for the API
mkdir -p /home/ec2-user/api

# Set up the environment
echo "export PATH=$PATH:/usr/local/go/bin" >> /home/ec2-user/.bashrc
source /home/ec2-user/.bashrc

# Set up SSH key for accessing the private repository
mkdir -p /home/ec2-user/.ssh
echo "${var.github_private_key}" > /home/ec2-user/.ssh/id_rsa
chmod 600 /home/ec2-user/.ssh/id_rsa

# Clone the private repository
git clone git@github.com:vincentes/keygate.git /home/ec2-user/api

# Build and run the Golang API
cd /home/ec2-user/api
go build -o api
nohup ./api &
EOF

  tags = {
    Name = "${var.environment_name}-ec2-1"
  }
}

resource "aws_instance" "instance_2" {
  ami           = "ami-051f8a213df8bc089"
  instance_type = var.instance_type
  security_groups = [aws_security_group.instances.name]

  user_data = <<-EOF
#!/bin/bash
# Install dependencies
sudo yum install -y golang git

# Create a directory for the API
mkdir -p /home/ec2-user/api

# Set up the environment
echo "export PATH=$PATH:/usr/local/go/bin" >> /home/ec2-user/.bashrc
source /home/ec2-user/.bashrc

# Set up SSH key for accessing the private repository
mkdir -p /home/ec2-user/.ssh
echo "${var.github_private_key}" > /home/ec2-user/.ssh/id_rsa
chmod 600 /home/ec2-user/.ssh/id_rsa

# Clone the private repository
git clone git@github.com:vincentes/keygate.git /home/ec2-user/api

# Build and run the Golang API
cd /home/ec2-user/api
go build -o api
nohup ./api &
EOF

  tags = {
    Name = "${var.environment_name}-ec2-2"
  }
}