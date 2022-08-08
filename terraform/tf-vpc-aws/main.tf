provider "aws" {
   region = "us-east-1"
   profile = "developer"
}

resource "aws_subnet" "subnet-1" {
   # Referencing the VPC that we created above
   # <resource_type>.<resource_name>.<resource_property>
   vpc_id = aws_vpc.main-vpc.id
   cidr_block = "10.0.1.0/24"
   tags = {
      Name = "prod-subnet"
   }
}

resource "aws_vpc" "main-vpc" {
   # It is going to be the network that's going
   # to be used for that VPC
   cidr_block = "10.0.0.0/16"
   tags = {
      Name = "production"
   }
}