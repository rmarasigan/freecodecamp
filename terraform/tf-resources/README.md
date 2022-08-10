# Terraform Resources

**Resources** are the most important element in the Terraform language. Each resource block describes one or more infrastructure objects, such as virtual networks, compute instances, or higher-level components such as DNS records.

* **Resource Blocks** documents the syntax for declaring resources.
* **Resource Behavior** explains in more detail how Terraform handles resource declarations when applying a configuration.

With the decorative nature of Terraform, it is kind of an all or nothing situation when you want to deploy your Terraform config or destroy it. So you can either spin up everything that's been configured in your Terrafom filr or destroy all of it. However, Terraform does understand that there's going to be instances where you want to just delete an individual resource or you're just trying to roll out with a deployment to do staged deployments where only certain resources are to play are deployed one day and then the next day, another set of resources that are deployed. So we can target individual resources within our configs to either apply or destroy them by passing in the `-target` flag.

```bash
dev@dev:~$ terraform destroy -target aws_instance.web-server-instance
aws_vpc.prod-vpc: Refreshing state... [id=vpc-0d234d6543ab21ae0]
aws_subnet.subnet-1: Refreshing state... [id=subnet-01234d5678d87e6bb]
aws_security_group.allow_web: Refreshing state... [id=sg-08c7654bc3b21a9c8]
aws_network_interface.web-server-nic: Refreshing state... [id=eni-0b0b45fdb4e76cc98]
aws_instance.web-server-instance: Refreshing state... [id=i-0cf1833b18a2e022c]

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  - destroy

Terraform will perform the following actions:

   # aws_instance.web-server-instance will be destroyed
   - resource "aws_instance" "web-server-instance" {
      - ami                                  = "ami-0729e439b6769d6ab" -> null
      - arn                                  = "arn:aws:ec2:us-east-1:xxxxxxxxxxxx:instance/i-0a2f34a56f7ebea8b" -> null
      - associate_public_ip_address          = true -> null
      - availability_zone                    = "us-east-1a" -> null
      - cpu_core_count                       = 1 -> null
      - cpu_threads_per_core                 = 1 -> null
      - disable_api_stop                     = false -> null
      - disable_api_termination              = false -> null
      - ebs_optimized                        = false -> null
      - get_password_data                    = false -> null
      - hibernation                          = false -> null
      - id                                   = "i-0a2f34a56f7ebea8b" -> null
      - instance_initiated_shutdown_behavior = "stop" -> null
      - instance_state                       = "running" -> null
      - instance_type                        = "t2.micro" -> null
      - ipv6_address_count                   = 0 -> null
      - ipv6_addresses                       = [] -> null
      - key_name                             = "terraform-key" -> null
      - monitoring                           = false -> null
      - primary_network_interface_id         = "eni-0b0b45fdb4e76cc98" -> null
      - private_dns                          = "ip-10-0-1-50.ec2.internal" -> null
      - private_ip                           = "10.0.1.50" -> null
      - public_ip                            = "5.210.68.161" -> null
      - secondary_private_ips                = [] -> null
      - security_groups                      = [] -> null
      - source_dest_check                    = true -> null
      - subnet_id                            = "subnet-01234d5678d87e6bb" -> null
      - tags                                 = {
          - "Name" = "web-server"
        } -> null
      - tags_all                             = {
          - "Name" = "web-server"
        } -> null
      - tenancy                              = "default" -> null
      - user_data                            = "2842f5e731dd3841e732fd56c7bc84fa045626b0" -> null
      - user_data_replace_on_change          = false -> null
      - vpc_security_group_ids               = [
          - "sg-08c7654bc3b21a9c8",
        ] -> null

      - capacity_reservation_specification {
          - capacity_reservation_preference = "open" -> null
        }

      - credit_specification {
          - cpu_credits = "standard" -> null
        }

      - enclave_options {
          - enabled = false -> null
        }

      - maintenance_options {
          - auto_recovery = "default" -> null
        }

      - metadata_options {
          - http_endpoint               = "enabled" -> null
          - http_put_response_hop_limit = 1 -> null
          - http_tokens                 = "optional" -> null
          - instance_metadata_tags      = "disabled" -> null
        }

      - network_interface {
          - delete_on_termination = false -> null
          - device_index          = 0 -> null
          - network_card_index    = 0 -> null
          - network_interface_id  = "eni-0b0b45fdb4e76cc98" -> null
        }

      - private_dns_name_options {
          - enable_resource_name_dns_a_record    = false -> null
          - enable_resource_name_dns_aaaa_record = false -> null
          - hostname_type                        = "ip-name" -> null
        }

      - root_block_device {
          - delete_on_termination = true -> null
          - device_name           = "/dev/sda1" -> null
          - encrypted             = false -> null
          - iops                  = 100 -> null
          - tags                  = {} -> null
          - throughput            = 0 -> null
          - volume_id             = "vol-093efa95fb6c55620" -> null
          - volume_size           = 8 -> null
          - volume_type           = "gp2" -> null
        }
    }

Plan: 0 to add, 0 to change, 1 to destroy.
╷
│ Warning: Resource targeting is in effect
│ 
│ You are creating a plan with the -target option, which means that the result of this plan may not represent all of the changes requested by the current configuration.
│ 
│ The -target option is not for routine use, and is provided only for exceptional situations such as recovering from errors or mistakes, or when Terraform specifically suggests to use it as part of an
│ error message.
╵

Do you really want to destroy all resources?
  Terraform will destroy all your managed infrastructure, as shown above.
  There is no undo. Only 'yes' will be accepted to confirm.

  Enter a value: yes

aws_instance.web-server-instance: Destroying... [id=i-0a2f34a56f7ebea8b]
aws_instance.web-server-instance: Still destroying... [id=i-0a2f34a56f7ebea8b, 10s elapsed]
aws_instance.web-server-instance: Still destroying... [id=i-0a2f34a56f7ebea8b, 20s elapsed]
aws_instance.web-server-instance: Still destroying... [id=i-0a2f34a56f7ebea8b, 30s elapsed]
aws_instance.web-server-instance: Destruction complete after 31s
╷
│ Warning: Applied changes may be incomplete
│ 
│ The plan was created with the -target option in effect, so some changes requested in the configuration may have been ignored and the output values may not be fully updated. Run the following command to
│ verify that no other changes are pending:
│     terraform plan
│ 	
│ Note that the -target option is not suitable for routine use, and is provided only for exceptional situations such as recovering from errors or mistakes, or when Terraform specifically suggests to use
│ it as part of an error message.
╵

Destroy complete! Resources: 1 destroyed.
```

So it successfully destroyed just that one resource. If we want to do the exact opposite thing, and deploy the `web-server-instance` resource, we can use the `-target` flag once again to accomplish that.

```bash
dev@dev:~$ terraform destroy -target aws_instance.web-server-instance
aws_vpc.prod-vpc: Refreshing state... [id=vpc-0d234d6543ab21ae0]
aws_subnet.subnet-1: Refreshing state... [id=subnet-01234d5678d87e6bb]
aws_security_group.allow_web: Refreshing state... [id=sg-08c7654bc3b21a9c8]
aws_network_interface.web-server-nic: Refreshing state... [id=eni-0b0b45fdb4e76cc98]

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

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
          + network_interface_id  = "eni-0b0b45fdb4e76cc98"
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

Plan: 1 to add, 0 to change, 0 to destroy.
╷
│ Warning: Resource targeting is in effect
│ 
│ You are creating a plan with the -target option, which means that the result of this plan may not represent all of the changes requested by the current configuration.
│ 
│ The -target option is not for routine use, and is provided only for exceptional situations such as recovering from errors or mistakes, or when Terraform specifically suggests to use it as part of an
│ error message.
╵

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

aws_instance.web-server-instance: Creating...
aws_instance.web-server-instance: Still creating... [10s elapsed]
aws_instance.web-server-instance: Still creating... [20s elapsed]
aws_instance.web-server-instance: Still creating... [30s elapsed]
aws_instance.web-server-instance: Still creating... [40s elapsed]
aws_instance.web-server-instance: Creation complete after 45s [id=i-123456a7dd8c90bf1]
╷
│ Warning: Applied changes may be incomplete
│ 
│ The plan was created with the -target option in effect, so some changes requested in the configuration may have been ignored and the output values may not be fully updated. Run the following command to
│ verify that no other changes are pending:
│     terraform plan
│ 	
│ Note that the -target option is not suitable for routine use, and is provided only for exceptional situations such as recovering from errors or mistakes, or when Terraform specifically suggests to use
│ it as part of an error message.
╵
```

In this case, we just deployed the `web-server-instance` that we just deleted. We can see that it is going to add one resource, which is exactly what we want.

## Reference
* [Resources](https://www.terraform.io/language/resources)