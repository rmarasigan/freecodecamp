# Setting up authentication and configuration
provider "aws" {
  # A Region is a location where AWS has a data center
  # https://docs.aws.amazon.com/directoryservice/latest/admin-guide/regions.html
  region = "us-east-1"
  # If no named profile is specified, the `default` profile is used.
  # Use the `profile` parameter to specify a named profile.
  profile = "developer"
}

# Virtual Private Cloud (VPC)
resource "aws_vpc" "prod-vpc" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "production"
  }
}

# Internet Gateway
resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.prod-vpc.id
}

# Route Table
resource "aws_route_table" "prod-route-table" {
  vpc_id = aws_vpc.prod-vpc.id

  # We are going to send all traffic, all IPv4 traffic wherever this
  # route points
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw.id
  }

  # The traffic from the subnet that we're going to create can get out
  # to the internet
  route {
    ipv6_cidr_block = "::/0"
    gateway_id      = aws_internet_gateway.gw.id
  }

  tags = {
    Name = "Prod"
  }
}

# Subnet
# It is where our web server is going to reside on
resource "aws_subnet" "subnet-1" {
  vpc_id     = aws_vpc.prod-vpc.id
  cidr_block = "10.0.1.0/24"

  # Availability Zone is, within each Region, AWS actually have
  # a couple of data centers, so if one of their data centers goes
  # down in a region, the whole region doesn't go down. They have
  # redundancy. So you can actually deploy resources to a specific
  # availability zone. 
  availability_zone = "us-east-1a"

  tags = {
    Name = "prod-subnet"
  }
}

# Route Table Association
resource "aws_route_table_association" "prod-route-table-association" {
  subnet_id      = aws_subnet.subnet-1.id
  route_table_id = aws_route_table.prod-route-table.id
}

# Security Group
resource "aws_security_group" "allow_web" {
  name        = "allow_web_traffic"
  description = "Allow Web inbound traffic"
  vpc_id      = aws_vpc.prod-vpc.id

  # We are going to allow HTTPS TCP traffic on port 443
  ingress {
    description = "HTTPS"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"

    # We can actually put our own IP address of of our
    # computer or something so that only certain devices
    # can actually reach this. Sometimes it is best not to
    # open it up to the internet.
    # The default IP (0.0.0.0/0) means that any IP address can access it.
    cidr_blocks = ["0.0.0.0/0"]
  }

  # We are going to allow HTTP TCP traffic on port 80
  ingress {
    description = "HTTP"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # We are going to allow SSH TCP traffic on port 22
  ingress {
    description = "SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # It allows all ports in the egress direction
  egress {
    from_port = 0
    to_port   = 0

    # This means any protocol
    protocol = "-1"

    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "allow_web"
  }
}


# Network Interface
resource "aws_network_interface" "web-server-nic" {
  subnet_id       = aws_subnet.subnet-1.id
  private_ips     = ["10.0.1.50"]
  security_groups = [aws_security_group.allow_web.id]
}

# Elastic IP
# Relies on the deployment of the Internet Gateway which
# is created above this. If you try to do this through the
# console, and you try to create an Elastic IP and assign it
# to a device, on a subnet that are in a VPC that doesn't have
# an internet gateway, it will throw an error. In order to
# have a public address you need to have an Internet Gateway
# So deploying an Elastic IP requires the Internet Gateway to
# be deployed first before the actual Elastic IP gets deployed.
resource "aws_eip" "web_server_eip" {
  vpc                       = true
  network_interface         = aws_network_interface.web-server-nic.id
  associate_with_private_ip = "10.0.1.50"
  depends_on                = [aws_internet_gateway.gw]
}

# Ubuntu Server
resource "aws_instance" "web-server-instance" {
  ami           = "ami-0729e439b6769d6ab"
  instance_type = "t2.micro"

  # If you don't hard code your availability zone,
  # Amazon will pick a random availability zone to
  # deploy that
  availability_zone = "us-east-1a"
  key_name          = "terraform-key"

  # We have tell them that which interface you want
  # it to be 
  network_interface {
    device_index         = 0
    network_interface_id = aws_network_interface.web-server-nic.id
  }

  # On deployment of this server, to actually run a
  # few commands we use `user_data`. The EOF at the
  # top starts all the configs and the EOF at the
  # bottom basically tells Terraform that we're done.
  user_data = <<-EOF
                #!/bin/bash
                sudo apt update -y
                sudo apt install apache2 -y
                sudo systemctl start apache2
                sudo bash -c 'echo Hello Developer this is your very first web server > /var/www/html/index.html'
                EOF

  tags = {
    Name = "web-server"
  }
}
