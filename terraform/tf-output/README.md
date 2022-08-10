# Terraform Output

Output values make information about your infrastructure available on the command line, and it can expose information for other Terraform configurations to use. Output values are similar to return values in programming languages.

So you know, any of the properties that we can see by doing a `terraform state show` on anyone of these resources, we can tell Terraform to automatically print those out when the resources are created. We can do that using the `terraform output` command. One of the main use cases is extracting the public IP that is going to get assigned to an Elastic IP or a EC2 instance that gets created.

Syntax:
```
output "output_name" {
   value = aws_instance.resource.property
}
```

So to get the value, what we have to do is we have to use this `terraform state show`, pass the resource name and grab the property. 
:
```
output "server_public_ip" {
   value = aws_eip.web_server_eip.public_ip
}
```

That will get the property `public_ip` and it will print out into the console when we run `terraform apply`.

```bash
dev@dev:~$ terraform apply

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # aws_eip.web_server_eip will be created
  + resource "aws_eip" "web_server_eip" {
      + allocation_id             = (known after apply)
      + associate_with_private_ip = "10.0.1.50"
      + association_id            = (known after apply)
      + carrier_ip                = (known after apply)
      + customer_owned_ip         = (known after apply)
      + domain                    = (known after apply)
      + id                        = (known after apply)
      + instance                  = (known after apply)
      + network_border_group      = (known after apply)
      + network_interface         = (known after apply)
      + private_dns               = (known after apply)
      + private_ip                = (known after apply)
      + public_dns                = (known after apply)
      + public_ip                 = (known after apply)
      + public_ipv4_pool          = (known after apply)
      + tags_all                  = (known after apply)
      + vpc                       = true
    }

  # aws_instance.web-server-instance will be created
  + resource "aws_instance" "web-server-instance" {
      + ami                                  = "ami-0729e439b6769d6ab"
      + arn                                  = (known after apply)
      + associate_public_ip_address          = (known after apply)
      + availability_zone                    = "us-east-1a"
      + cpu_core_count                       = (known after apply)
      + cpu_threads_per_core                 = (known after apply)
      + disable_api_stop                     = (known after apply)
      + disable_api_termination              = (known after apply)
      + ebs_optimized                        = (known after apply)
      + get_password_data                    = false
      + host_id                              = (known after apply)
      + id                                   = (known after apply)
      + instance_initiated_shutdown_behavior = (known after apply)
      + instance_state                       = (known after apply)
      + instance_type                        = "t2.micro"
      + ipv6_address_count                   = (known after apply)
      + ipv6_addresses                       = (known after apply)
      + key_name                             = "terraform-key"
      + monitoring                           = (known after apply)
      + outpost_arn                          = (known after apply)
      + password_data                        = (known after apply)
      + placement_group                      = (known after apply)
      + placement_partition_number           = (known after apply)
      + primary_network_interface_id         = (known after apply)
      + private_dns                          = (known after apply)
      + private_ip                           = (known after apply)
      + public_dns                           = (known after apply)
      + public_ip                            = (known after apply)
      + secondary_private_ips                = (known after apply)
      + security_groups                      = (known after apply)
      + subnet_id                            = (known after apply)
      + tags                                 = {
          + "Name" = "web-server"
        }
      + tags_all                             = {
          + "Name" = "web-server"
        }
      + tenancy                              = (known after apply)
      + user_data                            = "2842f5e731dd3841e732fd56c7bc84fa045626b0"
      + user_data_base64                     = (known after apply)
      + user_data_replace_on_change          = false
      + vpc_security_group_ids               = (known after apply)

      + capacity_reservation_specification {
          + capacity_reservation_preference = (known after apply)

          + capacity_reservation_target {
              + capacity_reservation_id                 = (known after apply)
              + capacity_reservation_resource_group_arn = (known after apply)
            }
        }

      + ebs_block_device {
          + delete_on_termination = (known after apply)
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + snapshot_id           = (known after apply)
          + tags                  = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = (known after apply)
          + volume_type           = (known after apply)
        }

      + enclave_options {
          + enabled = (known after apply)
        }

      + ephemeral_block_device {
          + device_name  = (known after apply)
          + no_device    = (known after apply)
          + virtual_name = (known after apply)
        }

      + maintenance_options {
          + auto_recovery = (known after apply)
        }

      + metadata_options {
          + http_endpoint               = (known after apply)
          + http_put_response_hop_limit = (known after apply)
          + http_tokens                 = (known after apply)
          + instance_metadata_tags      = (known after apply)
        }

      + network_interface {
          + delete_on_termination = false
          + device_index          = 0
          + network_card_index    = 0
          + network_interface_id  = (known after apply)
        }

      + private_dns_name_options {
          + enable_resource_name_dns_a_record    = (known after apply)
          + enable_resource_name_dns_aaaa_record = (known after apply)
          + hostname_type                        = (known after apply)
        }

      + root_block_device {
          + delete_on_termination = (known after apply)
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + tags                  = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = (known after apply)
          + volume_type           = (known after apply)
        }
    }

  # aws_internet_gateway.gw will be created
  + resource "aws_internet_gateway" "gw" {
      + arn      = (known after apply)
      + id       = (known after apply)
      + owner_id = (known after apply)
      + tags_all = (known after apply)
      + vpc_id   = (known after apply)
    }

  # aws_network_interface.web-server-nic will be created
  + resource "aws_network_interface" "web-server-nic" {
      + arn                       = (known after apply)
      + id                        = (known after apply)
      + interface_type            = (known after apply)
      + ipv4_prefix_count         = (known after apply)
      + ipv4_prefixes             = (known after apply)
      + ipv6_address_count        = (known after apply)
      + ipv6_address_list         = (known after apply)
      + ipv6_address_list_enabled = false
      + ipv6_addresses            = (known after apply)
      + ipv6_prefix_count         = (known after apply)
      + ipv6_prefixes             = (known after apply)
      + mac_address               = (known after apply)
      + outpost_arn               = (known after apply)
      + owner_id                  = (known after apply)
      + private_dns_name          = (known after apply)
      + private_ip                = (known after apply)
      + private_ip_list           = (known after apply)
      + private_ip_list_enabled   = false
      + private_ips               = [
          + "10.0.1.50",
        ]
      + private_ips_count         = (known after apply)
      + security_groups           = (known after apply)
      + source_dest_check         = true
      + subnet_id                 = (known after apply)
      + tags_all                  = (known after apply)

      + attachment {
          + attachment_id = (known after apply)
          + device_index  = (known after apply)
          + instance      = (known after apply)
        }
    }

  # aws_route_table.prod-route-table will be created
  + resource "aws_route_table" "prod-route-table" {
      + arn              = (known after apply)
      + id               = (known after apply)
      + owner_id         = (known after apply)
      + propagating_vgws = (known after apply)
      + route            = [
          + {
              + carrier_gateway_id         = ""
              + cidr_block                 = ""
              + core_network_arn           = ""
              + destination_prefix_list_id = ""
              + egress_only_gateway_id     = ""
              + gateway_id                 = (known after apply)
              + instance_id                = ""
              + ipv6_cidr_block            = "::/0"
              + local_gateway_id           = ""
              + nat_gateway_id             = ""
              + network_interface_id       = ""
              + transit_gateway_id         = ""
              + vpc_endpoint_id            = ""
              + vpc_peering_connection_id  = ""
            },
          + {
              + carrier_gateway_id         = ""
              + cidr_block                 = "0.0.0.0/0"
              + core_network_arn           = ""
              + destination_prefix_list_id = ""
              + egress_only_gateway_id     = ""
              + gateway_id                 = (known after apply)
              + instance_id                = ""
              + ipv6_cidr_block            = ""
              + local_gateway_id           = ""
              + nat_gateway_id             = ""
              + network_interface_id       = ""
              + transit_gateway_id         = ""
              + vpc_endpoint_id            = ""
              + vpc_peering_connection_id  = ""
            },
        ]
      + tags             = {
          + "Name" = "Prod"
        }
      + tags_all         = {
          + "Name" = "Prod"
        }
      + vpc_id           = (known after apply)
    }

  # aws_route_table_association.prod-route-table-association will be created
  + resource "aws_route_table_association" "prod-route-table-association" {
      + id             = (known after apply)
      + route_table_id = (known after apply)
      + subnet_id      = (known after apply)
    }

  # aws_security_group.allow_web will be created
  + resource "aws_security_group" "allow_web" {
      + arn                    = (known after apply)
      + description            = "Allow Web inbound traffic"
      + egress                 = [
          + {
              + cidr_blocks      = [
                  + "0.0.0.0/0",
                ]
              + description      = ""
              + from_port        = 0
              + ipv6_cidr_blocks = []
              + prefix_list_ids  = []
              + protocol         = "-1"
              + security_groups  = []
              + self             = false
              + to_port          = 0
            },
        ]
      + id                     = (known after apply)
      + ingress                = [
          + {
              + cidr_blocks      = [
                  + "0.0.0.0/0",
                ]
              + description      = "HTTP"
              + from_port        = 80
              + ipv6_cidr_blocks = []
              + prefix_list_ids  = []
              + protocol         = "tcp"
              + security_groups  = []
              + self             = false
              + to_port          = 80
            },
          + {
              + cidr_blocks      = [
                  + "0.0.0.0/0",
                ]
              + description      = "HTTPS"
              + from_port        = 443
              + ipv6_cidr_blocks = []
              + prefix_list_ids  = []
              + protocol         = "tcp"
              + security_groups  = []
              + self             = false
              + to_port          = 443
            },
          + {
              + cidr_blocks      = [
                  + "0.0.0.0/0",
                ]
              + description      = "SSH"
              + from_port        = 22
              + ipv6_cidr_blocks = []
              + prefix_list_ids  = []
              + protocol         = "tcp"
              + security_groups  = []
              + self             = false
              + to_port          = 22
            },
        ]
      + name                   = "allow_web_traffic"
      + name_prefix            = (known after apply)
      + owner_id               = (known after apply)
      + revoke_rules_on_delete = false
      + tags                   = {
          + "Name" = "allow_web"
        }
      + tags_all               = {
          + "Name" = "allow_web"
        }
      + vpc_id                 = (known after apply)
    }

  # aws_subnet.subnet-1 will be created
  + resource "aws_subnet" "subnet-1" {
      + arn                                            = (known after apply)
      + assign_ipv6_address_on_creation                = false
      + availability_zone                              = "us-east-1a"
      + availability_zone_id                           = (known after apply)
      + cidr_block                                     = "10.0.1.0/24"
      + enable_dns64                                   = false
      + enable_resource_name_dns_a_record_on_launch    = false
      + enable_resource_name_dns_aaaa_record_on_launch = false
      + id                                             = (known after apply)
      + ipv6_cidr_block_association_id                 = (known after apply)
      + ipv6_native                                    = false
      + map_public_ip_on_launch                        = false
      + owner_id                                       = (known after apply)
      + private_dns_hostname_type_on_launch            = (known after apply)
      + tags                                           = {
          + "Name" = "prod-subnet"
        }
      + tags_all                                       = {
          + "Name" = "prod-subnet"
        }
      + vpc_id                                         = (known after apply)
    }

  # aws_vpc.prod-vpc will be created
  + resource "aws_vpc" "prod-vpc" {
      + arn                                  = (known after apply)
      + cidr_block                           = "10.0.0.0/16"
      + default_network_acl_id               = (known after apply)
      + default_route_table_id               = (known after apply)
      + default_security_group_id            = (known after apply)
      + dhcp_options_id                      = (known after apply)
      + enable_classiclink                   = (known after apply)
      + enable_classiclink_dns_support       = (known after apply)
      + enable_dns_hostnames                 = (known after apply)
      + enable_dns_support                   = true
      + id                                   = (known after apply)
      + instance_tenancy                     = "default"
      + ipv6_association_id                  = (known after apply)
      + ipv6_cidr_block                      = (known after apply)
      + ipv6_cidr_block_network_border_group = (known after apply)
      + main_route_table_id                  = (known after apply)
      + owner_id                             = (known after apply)
      + tags                                 = {
          + "Name" = "production"
        }
      + tags_all                             = {
          + "Name" = "production"
        }
    }

Plan: 9 to add, 0 to change, 0 to destroy.

Changes to Outputs:
  + server_public_ip = (known after apply)

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

aws_vpc.prod-vpc: Creating...
aws_vpc.prod-vpc: Creation complete after 4s [id=vpc-0d234d6543ab21ae0]
aws_internet_gateway.gw: Creating...
aws_subnet.subnet-1: Creating...
aws_security_group.allow_web: Creating...
aws_subnet.subnet-1: Creation complete after 1s [id=subnet-01234d5678d87e6bb]
aws_internet_gateway.gw: Creation complete after 4s [id=igw-0bac1ebb234fb0105]
aws_route_table.prod-route-table: Creating...
aws_security_group.allow_web: Creation complete after 5s [id=sg-08c7654bc3b21a9c8]
aws_network_interface.web-server-nic: Creating...
aws_network_interface.web-server-nic: Creation complete after 1s [id=eni-0b0b45fdb4e76cc98]
aws_eip.web_server_eip: Creating...
aws_instance.web-server-instance: Creating...
aws_route_table.prod-route-table: Creation complete after 3s [id=rtb-0282e480b60d991bf]
aws_route_table_association.prod-route-table-association: Creating...
aws_eip.web_server_eip: Creation complete after 2s [id=eipalloc-010d7e670c3625d6a]
aws_route_table_association.prod-route-table-association: Creation complete after 1s [id=rtbassoc-03da971722e7de0e5]
aws_instance.web-server-instance: Still creating... [10s elapsed]
aws_instance.web-server-instance: Still creating... [20s elapsed]
aws_instance.web-server-instance: Still creating... [30s elapsed]
aws_instance.web-server-instance: Still creating... [40s elapsed]
aws_instance.web-server-instance: Creation complete after 46s [id=i-0a2f34a56f7ebea8b]

Apply complete! Resources: 9 added, 0 changed, 0 destroyed.

Outputs:

server_public_ip = 5.210.68.161
```

The output we created, is going to print out the server public IP so that we don't have to go into the `terraform state` and verify it ourselves. This can come in handy because we can print out all the details that we are concerned with without having to go in and check the states ourselves.

However, if you decide to add another output into your config but you don't want to do a `terraform apply` because technically a `terraform apply` could potentially make changes to your network and if you are in a production network, you don't want to accidentally deploy or delete something just to see what the output is. You can always use `terraform refresh`.

This should just refresh all of your states and it will run the outputs again so that you can verify them without actually deploying anything or doing an actual apply.

```bash
dev@dev:~$ terraform refresh
aws_vpc.prod-vpc: Refreshing state... [id=vpc-0d234d6543ab21ae0]
aws_internet_gateway.gw: Refreshing state... [id=igw-0bac1ebb234fb0105]
aws_subnet.subnet-1: Refreshing state... [id=subnet-01234d5678d87e6bb]
aws_security_group.allow_web: Refreshing state... [id=sg-08c7654bc3b21a9c8]
aws_route_table.prod-route-table: Refreshing state... [id=rtb-0282e480b60d991bf]
aws_network_interface.web-server-nic: Refreshing state... [id=eni-0b0b45fdb4e76cc98]
aws_route_table_association.prod-route-table-association: Refreshing state... [id=rtbassoc-03da971722e7de0e5]
aws_eip.web_server_eip: Refreshing state... [id=eipalloc-010d7e670c3625d6a]
aws_instance.web-server-instance: Refreshing state... [id=i-0a2f34a56f7ebea8b]

Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

Outputs:

server_id = i-0a2f34a56f7ebea8b
server_private_ip = 10.0.1.50
server_public_ip = 5.210.68.161
```