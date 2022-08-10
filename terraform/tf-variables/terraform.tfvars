# String
# subnet_prefix = "10.0.200.0/24"

# Objects
# subnet_prefix = ["10.0.200.0/24", "10.0.2.0/24"]

# List of objects
subnet_prefix = [
  {
    cidr_block = "10.0.1.0/24"
    name       = "prod_subnet"
  },
  {
    cidr_block = "10.0.2.0/24"
    name       = "dev_subnet"
  }
]