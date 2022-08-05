# Setting up authentication and configuration
provider "aws" {
   # A Region is a location where AWS has a data center
   # https://docs.aws.amazon.com/directoryservice/latest/admin-guide/regions.html
   region = "us-east-1"
   # If no named profile is specified, the `default` profile is used.
   # Use the `profile` parameter to specify a named profile.
   profile = "developer"
}

/************************************************************
   resource <provider>_<resource_type> "name" {
      config options...
      key1 = "value1"
      key2 = "value2"
   }
************************************************************/
resource "aws_instance" "web_server" {
   ami = "ami-0729e439b6769d6ab"
   instance_type =  "t2.micro"
   tags = {
      name = "ubuntu"
   }
}