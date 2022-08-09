# Terraform Practice Project

### Pre-requisite
* AWS Account
* Terraform CLI
* AWS EC2 Key Pair
* Terraform Account
* AWS CLI configured

### AWS EC2 Key Pair
To set up the keys, we need to go to the EC2 section (Services → Compute → EC2). There's a section under *Network & Security* called **Key Pairs**. Do **Create key pair** and we'll just give it a name. So there are different **file formats** that you can use. Go ahead and download the **pem** format.

---
**NOTE**

**Key pair type**
* **RSA** is the most widely implemented and supported public-key algorithm for SSH. Compared to ED25519, it is slower and even considered not safe if it is generated with a key smaller than 2048-bit length.
* **ED25519** public key is compact. It only contains 68 characters. ED25519 signatures are elliptic-curve signatures, carefully engineered at several levels of design and implementation to achieve very high speeds without compromising security.

**Private key file format**

**PEM** (Privacy Enhanced Mail) is a base64 container format for encoding keys and certificates. PEM certificates are frequently used for web servers as they can easily be translated into readable data using a simple text editor. 

Example:
```
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAxZls7NgRS5rtDH+mCgEeqSO12BiPAY2dhZZbTkPEoDUVqG46
PHGIBBERISHFKS3xE4eMMY+uWPriviY0W12EFqYBUJDq6GmGSQPMSI8Kep7OgVGK
d2t3/Lq+I8ZqlFwAgPwS7n18mlgTV/YyA0HCCIa+kZGKqbsuylfsODUAJGQqOtUM
FKX4+O45RrJwcQGI2M+pbD4+Ow28iEsNP3flQKvamNuezm2VallhGiO/E39UITrB
27FWM1pLQo5etyE378h5q8BKMGKt1ZPxPhjiNjBefGZ1Dscw7CA5
-----END RSA PRIVATE KEY-----
```

**PPK** (Putty Private Key) files are used by PuTTY, a free [**SSH**](https://techterms.com/definition/ssh) and [**Telnet**](https://techterms.com/definition/telnet) client. It is a Windows SSH client, it does not support .pem format. Hence you have to convert .ppk format using PuTTyGen.

Example:
```
PuTTY-User-Key-File-2: ssh-rsa
Encryption: none
Comment: imported-openssh-key
Public-Lines: 3
AAAAB3NzaC1yc2EAAAADAQABAAABAQDFmWzs2BFLmu0Mf6YKAR6pI7XYGI8BjZ2F
llGIBBERISHobjo8c6DjH6iBW0UpLfETh4wxj65Y+uK+JjRbXYQWpgFQkOroaYZJ
1DxehJUvZpuZqLSUGvZrUzbYmx25oVE5+x6UBSA0NUiTpq6jWFPH
Private-Lines: 3
AAABAFLQsak2CVGhldYWLbWlsH78+3qdGFpzSuCxGoVbu63VccVMpVWutHBcVVgZ
s4GIBBERISHcVhru01cUVf8w22pOmuB0TC+qAmDZmaWTZbXslun2Dc1tGjWu4n/F
S0KOXrchN+/IeavASjBirdWT8T4Y4jYwXnxmdQ7HMOwgOQ==
Private-MAC: 5fGIBBERISHda60d2e8b049e655e4a8a78bb18dc
```

**System platform compatibility**

**.pem file format**
* Mac
* Linux
* Windows PowerShell

**.ppk file format**
* Windows PuTTY/Cygwin
---

So we'll create a key pair and it is going to automatically download this for you. Just remember where it is getting saved to.

![tf-ec2-server-key](../../assets/img/tf-ec2-server-key.png)

Now that we have our key pair set, this key is going to allow us to actually connect to our server once we've deployed it.

![tf-ec2-key-pair](../../assets/img/tf-ec2-key-pair.png)

### Steps
1. Create a VPC
2. Create an Internet Gateway
   > This is to send traffic out to the internet because we do want to be able to assign a public IP address to the server so that anybody in the world can reach it
3. Create a custom Route Table
4. Create a Subnet
   > Anytime you create a subnet you need to assign it to a route table, or then they'll get assigned to the default route table for your VPC
5. Associate subnet with Route Table
6. Create Security Group to allow ports 22, 80, 443
   > A **Security Group** is just responsible for determining what kind of traffic is allowed to get to any one of your EC2 instances. Since this is a Web Server, we're going to allow port 80 and port 443 for HTTP and HTTPS traffic. We're also going to allow port 22 so that we can SSH to it and connect to it and make changes to it.
7. Create a network interface with an IP in the subnet that was created in *Step 4*
8. Assign an Elastic IP to the network interface created in *Step 7*
   > An **Elastic IP** in Amazon is just a public IP address that is routable on the internet
9. Create Ubuntu Server and install or enable Apache2

## `terraform init`

```bash
dev@dev:~$ terraform init

Initializing the backend...

Initializing provider plugins...
- Finding latest version of hashicorp/aws...
- Installing hashicorp/aws v4.25.0...
- Installed hashicorp/aws v4.25.0 (signed by HashiCorp)

Terraform has created a lock file .terraform.lock.hcl to record the provider
selections it made above. Include this file in your version control repository
so that Terraform can guarantee to make the same selections by default when
you run "terraform init" in the future.

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
```

## `terraform apply`

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
```


## Connecting to the Web Server
Since it has the public IP, we can actually SSH to it, and the way to do that, it is going to vary depending on what operating system you're on.

### Windows
So if you're on Windows, the first thing that you need to do is download a software called [PuTTy](https://www.putty.org/) and download the installer. So get the 64-bit if you are using 64-bit, or get the 32-bit if you are using the 32 bit.

![tf-windows-putty](../../assets/img/tf-windows-putty.png)

You can see there's PuTTY and PuTTYgen. The first thing that we need to do is to open up the PuTTygen. The reason we need to do this is because we have a `.pem` file and PuTTy uses a `.ppk` file format. So we have to convert the `.pem` file to the right format to be able to use it on a Windows machine. Hit on the **Load** button and choose the **`All Files (*.*)`** so that you can get your `.pem` file and hit **Open**.

![tf-windows-putty-load](../../assets/img/tf-windows-putty-load.png)


Then all we need to do is to select the **Save private key** and it is going to give you a warning. Go ahead and just hit **Yes**.

![tf-windows-putty-private-key](../../assets/img/tf-windows-putty-private-key.png)

The next step is to open up the **PuTTy** application and go back to the AWS EC2 Instances and select the instance and click on **Connect**.

![tf-vpc-ec2-connect](../../assets/img/tf-vpc-ec2-connect.png)

It is going to give you all the information on how to connect to the instance. On the **SSH client** tab, you'll see that the user for the instance is `ubuntu`, and it is going to give you the IP address. On the PuTTY Configuration application paste the user and specify the IP address (`ubuntu@xxx.xxx.xxx.xxx`).

Now, we have to log in using the `.ppk` file and to tell PuTTY to use that, go into the **SSH** on the left panel and hit the **Auth** button, and then there is a section saying `Private key file for authentication`. So hit **Browse** and find your `.ppk` file and hit **Open** and **Open** again. It should pop up with a security alert so just hit **Yes**.

![tf-putty-security-alert](../../assets/img/tf-putty-security-alert.png)

At this point, you should be logged in using PuTTy. You know we're on a Linux machine, you can run any command you want. So that is how we connect to a Windows machine.

![tf-windows-connect](../../assets/img/tf-windows-connect.png)

### Linux / Mac
The first thing that you need to do is go to your web browser and go to the **AWS EC2 Instances** section. Select the `web-server` and then hit **Connect**.

![tf-vpc-ec2-linux](../../assets/img/tf-vpc-ec2-linux.png)


This command line means that we are going to set the permissions to the correct values and only the `User` or `Owner` has the read permission.
```bash
dev@dev:~$ chmod 400 terraform-key.pem
```

So here's all what you have to do.
```bash
dev@dev:~$ chmod 400 terraform-key.pem
dev@dev:~$ ssh -i "terraform-key.pem" ubuntu@5.210.68.161
The authenticity of host '34.240.36.100 (34.240.36.100)' can't be established.
ECDSA key fingerprint is SHA256:1F2DC7a8Dd3f5908f9adc+47E70fD49dA6659fefB71.
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
Warning: Permanently added '34.240.36.100' (ECDSA) to the list of known hosts.
Welcome to Ubuntu 18.04.6 LTS (GNU/Linux 5.4.0-1078-aws x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

 System information disabled due to load higher than 1.0

41 updates can be applied immediately.
36 of these updates are standard security updates.
To see these additional updates run: apt list --upgradable



The programs included with the Ubuntu system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Ubuntu comes with ABSOLUTELY NO WARRANTY, to the extent permitted by
applicable law.

To run a command as administrator (user "root"), use "sudo <command>".
See "man sudo_root" for details.

ubuntu@ip-10-0-1-50:~$
```

## `terraform destroy`

```bash
dev@dev:~$ terraform destroy
aws_vpc.prod-vpc: Refreshing state... [id=vpc-0d234d6543ab21ae0]
aws_internet_gateway.gw: Refreshing state... [id=igw-0bac1ebb234fb0105]
aws_subnet.subnet-1: Refreshing state... [id=subnet-01234d5678d87e6bb]
aws_security_group.allow_web: Refreshing state... [id=sg-08c7654bc3b21a9c8]
aws_route_table.prod-route-table: Refreshing state... [id=rtb-0282e480b60d991bf]
aws_network_interface.web-server-nic: Refreshing state... [id=eni-0b0b45fdb4e76cc98]
aws_route_table_association.prod-route-table-association: Refreshing state... [id=rtbassoc-03da971722e7de0e5]
aws_eip.web_server_eip: Refreshing state... [id=eipalloc-010d7e670c3625d6a]
aws_instance.web-server-instance: Refreshing state... [id=i-0a2f34a56f7ebea8b]

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  - destroy

Terraform will perform the following actions:

  # aws_eip.web_server_eip will be destroyed
  - resource "aws_eip" "web_server_eip" {
      - allocation_id             = "eipalloc-010d7e670c3625d6a" -> null
      - associate_with_private_ip = "10.0.1.50" -> null
      - association_id            = "eipassoc-02f18f6e37d46ddb2" -> null
      - domain                    = "vpc" -> null
      - id                        = "eipalloc-010d7e670c3625d6a" -> null
      - instance                  = "i-0a2f34a56f7ebea8b" -> null
      - network_border_group      = "us-east-1" -> null
      - network_interface         = "eni-0b0b45fdb4e76cc98" -> null
      - private_dns               = "ip-10-0-1-50.ec2.internal" -> null
      - private_ip                = "10.0.1.50" -> null
      - public_dns                = "ec2-5-210-68-161.compute-1.amazonaws.com" -> null
      - public_ip                 = "5.210.68.161" -> null
      - public_ipv4_pool          = "amazon" -> null
      - tags                      = {} -> null
      - tags_all                  = {} -> null
      - vpc                       = true -> null
    }

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

  # aws_internet_gateway.gw will be destroyed
  - resource "aws_internet_gateway" "gw" {
      - arn      = "arn:aws:ec2:us-east-1:xxxxxxxxxxxx:internet-gateway/igw-0bac1ebb234fb0105" -> null
      - id       = "igw-0bac1ebb234fb0105" -> null
      - owner_id = "xxxxxxxxxxxx" -> null
      - tags     = {} -> null
      - tags_all = {} -> null
      - vpc_id   = "vpc-0d234d6543ab21ae0" -> null
    }

  # aws_network_interface.web-server-nic will be destroyed
  - resource "aws_network_interface" "web-server-nic" {
      - arn                       = "arn:aws:ec2:us-east-1:xxxxxxxxxxxx:network-interface/eni-0b0b45fdb4e76cc98" -> null
      - id                        = "eni-0b0b45fdb4e76cc98" -> null
      - interface_type            = "interface" -> null
      - ipv4_prefix_count         = 0 -> null
      - ipv4_prefixes             = [] -> null
      - ipv6_address_count        = 0 -> null
      - ipv6_address_list         = [] -> null
      - ipv6_address_list_enabled = false -> null
      - ipv6_addresses            = [] -> null
      - ipv6_prefix_count         = 0 -> null
      - ipv6_prefixes             = [] -> null
      - mac_address               = "12:34:5f:67:8c:4d" -> null
      - owner_id                  = "xxxxxxxxxxxx" -> null
      - private_ip                = "10.0.1.50" -> null
      - private_ip_list           = [
          - "10.0.1.50",
        ] -> null
      - private_ip_list_enabled   = false -> null
      - private_ips               = [
          - "10.0.1.50",
        ] -> null
      - private_ips_count         = 0 -> null
      - security_groups           = [
          - "sg-08c7654bc3b21a9c8",
        ] -> null
      - source_dest_check         = true -> null
      - subnet_id                 = "subnet-01234d5678d87e6bb" -> null
      - tags                      = {} -> null
      - tags_all                  = {} -> null

      - attachment {
          - attachment_id = "eni-attach-0b4e6cc9cccdaa4ae" -> null
          - device_index  = 0 -> null
          - instance      = "i-0a2f34a56f7ebea8b" -> null
        }
    }

  # aws_route_table.prod-route-table will be destroyed
  - resource "aws_route_table" "prod-route-table" {
      - arn              = "arn:aws:ec2:us-east-1:xxxxxxxxxxxx:route-table/rtb-0282e480b60d991bf" -> null
      - id               = "rtb-0282e480b60d991bf" -> null
      - owner_id         = "xxxxxxxxxxxx" -> null
      - propagating_vgws = [] -> null
      - route            = [
          - {
              - carrier_gateway_id         = ""
              - cidr_block                 = ""
              - core_network_arn           = ""
              - destination_prefix_list_id = ""
              - egress_only_gateway_id     = ""
              - gateway_id                 = "igw-0bac1ebb234fb0105"
              - instance_id                = ""
              - ipv6_cidr_block            = "::/0"
              - local_gateway_id           = ""
              - nat_gateway_id             = ""
              - network_interface_id       = ""
              - transit_gateway_id         = ""
              - vpc_endpoint_id            = ""
              - vpc_peering_connection_id  = ""
            },
          - {
              - carrier_gateway_id         = ""
              - cidr_block                 = "0.0.0.0/0"
              - core_network_arn           = ""
              - destination_prefix_list_id = ""
              - egress_only_gateway_id     = ""
              - gateway_id                 = "igw-0bac1ebb234fb0105"
              - instance_id                = ""
              - ipv6_cidr_block            = ""
              - local_gateway_id           = ""
              - nat_gateway_id             = ""
              - network_interface_id       = ""
              - transit_gateway_id         = ""
              - vpc_endpoint_id            = ""
              - vpc_peering_connection_id  = ""
            },
        ] -> null
      - tags             = {
          - "Name" = "Prod"
        } -> null
      - tags_all         = {
          - "Name" = "Prod"
        } -> null
      - vpc_id           = "vpc-0d234d6543ab21ae0" -> null
    }

  # aws_route_table_association.prod-route-table-association will be destroyed
  - resource "aws_route_table_association" "prod-route-table-association" {
      - id             = "rtbassoc-03da971722e7de0e5" -> null
      - route_table_id = "rtb-0282e480b60d991bf" -> null
      - subnet_id      = "subnet-01234d5678d87e6bb" -> null
    }

  # aws_security_group.allow_web will be destroyed
  - resource "aws_security_group" "allow_web" {
      - arn                    = "arn:aws:ec2:us-east-1:xxxxxxxxxxxx:security-group/sg-08c7654bc3b21a9c8" -> null
      - description            = "Allow Web inbound traffic" -> null
      - egress                 = [
          - {
              - cidr_blocks      = [
                  - "0.0.0.0/0",
                ]
              - description      = ""
              - from_port        = 0
              - ipv6_cidr_blocks = []
              - prefix_list_ids  = []
              - protocol         = "-1"
              - security_groups  = []
              - self             = false
              - to_port          = 0
            },
        ] -> null
      - id                     = "sg-08c7654bc3b21a9c8" -> null
      - ingress                = [
          - {
              - cidr_blocks      = [
                  - "0.0.0.0/0",
                ]
              - description      = "HTTP"
              - from_port        = 80
              - ipv6_cidr_blocks = []
              - prefix_list_ids  = []
              - protocol         = "tcp"
              - security_groups  = []
              - self             = false
              - to_port          = 80
            },
          - {
              - cidr_blocks      = [
                  - "0.0.0.0/0",
                ]
              - description      = "HTTPS"
              - from_port        = 443
              - ipv6_cidr_blocks = []
              - prefix_list_ids  = []
              - protocol         = "tcp"
              - security_groups  = []
              - self             = false
              - to_port          = 443
            },
          - {
              - cidr_blocks      = [
                  - "0.0.0.0/0",
                ]
              - description      = "SSH"
              - from_port        = 22
              - ipv6_cidr_blocks = []
              - prefix_list_ids  = []
              - protocol         = "tcp"
              - security_groups  = []
              - self             = false
              - to_port          = 22
            },
        ] -> null
      - name                   = "allow_web_traffic" -> null
      - owner_id               = "xxxxxxxxxxxx" -> null
      - revoke_rules_on_delete = false -> null
      - tags                   = {
          - "Name" = "allow_web"
        } -> null
      - tags_all               = {
          - "Name" = "allow_web"
        } -> null
      - vpc_id                 = "vpc-0d234d6543ab21ae0" -> null
    }

  # aws_subnet.subnet-1 will be destroyed
  - resource "aws_subnet" "subnet-1" {
      - arn                                            = "arn:aws:ec2:us-east-1:xxxxxxxxxxxx:subnet/subnet-01234d5678d87e6bb" -> null
      - assign_ipv6_address_on_creation                = false -> null
      - availability_zone                              = "us-east-1a" -> null
      - availability_zone_id                           = "use1-az2" -> null
      - cidr_block                                     = "10.0.1.0/24" -> null
      - enable_dns64                                   = false -> null
      - enable_resource_name_dns_a_record_on_launch    = false -> null
      - enable_resource_name_dns_aaaa_record_on_launch = false -> null
      - id                                             = "subnet-01234d5678d87e6bb" -> null
      - ipv6_native                                    = false -> null
      - map_customer_owned_ip_on_launch                = false -> null
      - map_public_ip_on_launch                        = false -> null
      - owner_id                                       = "xxxxxxxxxxxx" -> null
      - private_dns_hostname_type_on_launch            = "ip-name" -> null
      - tags                                           = {
          - "Name" = "prod-subnet"
        } -> null
      - tags_all                                       = {
          - "Name" = "prod-subnet"
        } -> null
      - vpc_id                                         = "vpc-0d234d6543ab21ae0" -> null
    }

  # aws_vpc.prod-vpc will be destroyed
  - resource "aws_vpc" "prod-vpc" {
      - arn                              = "arn:aws:ec2:us-east-1:xxxxxxxxxxxx:vpc/vpc-0d234d6543ab21ae0" -> null
      - assign_generated_ipv6_cidr_block = false -> null
      - cidr_block                       = "10.0.0.0/16" -> null
      - default_network_acl_id           = "acl-0e2eec1651a2f5e54" -> null
      - default_route_table_id           = "rtb-0c0f2386563f79db7" -> null
      - default_security_group_id        = "sg-0c5a6d7d7d05a90ab" -> null
      - dhcp_options_id                  = "dopt-0a38da625cd8f4d69" -> null
      - enable_classiclink               = false -> null
      - enable_classiclink_dns_support   = false -> null
      - enable_dns_hostnames             = false -> null
      - enable_dns_support               = true -> null
      - id                               = "vpc-0d234d6543ab21ae0" -> null
      - instance_tenancy                 = "default" -> null
      - ipv6_netmask_length              = 0 -> null
      - main_route_table_id              = "rtb-0c0f2386563f79db7" -> null
      - owner_id                         = "xxxxxxxxxxxx" -> null
      - tags                             = {
          - "Name" = "production"
        } -> null
      - tags_all                         = {
          - "Name" = "production"
        } -> null
    }

Plan: 0 to add, 0 to change, 9 to destroy.

Do you really want to destroy all resources?
  Terraform will destroy all your managed infrastructure, as shown above.
  There is no undo. Only 'yes' will be accepted to confirm.

  Enter a value: yes

aws_route_table_association.prod-route-table-association: Destroying... [id=rtbassoc-03da971722e7de0e5]
aws_eip.web_server_eip: Destroying... [id=eipalloc-010d7e670c3625d6a]
aws_instance.web-server-instance: Destroying... [id=i-0a2f34a56f7ebea8b]
aws_route_table_association.prod-route-table-association: Destruction complete after 1s
aws_route_table.prod-route-table: Destroying... [id=rtb-0282e480b60d991bf]
aws_route_table.prod-route-table: Destruction complete after 2s
aws_eip.web_server_eip: Destruction complete after 4s
aws_internet_gateway.gw: Destroying... [id=igw-0bac1ebb234fb0105]
aws_internet_gateway.gw: Destruction complete after 1s
aws_instance.web-server-instance: Still destroying... [id=i-0a2f34a56f7ebea8b, 10s elapsed]
aws_instance.web-server-instance: Still destroying... [id=i-0a2f34a56f7ebea8b, 20s elapsed]
aws_instance.web-server-instance: Still destroying... [id=i-0a2f34a56f7ebea8b, 30s elapsed]
aws_instance.web-server-instance: Still destroying... [id=i-0a2f34a56f7ebea8b, 40s elapsed]
aws_instance.web-server-instance: Destruction complete after 43s
aws_network_interface.web-server-nic: Destroying... [id=eni-0b0b45fdb4e76cc98]
aws_network_interface.web-server-nic: Destruction complete after 0s
aws_subnet.subnet-1: Destroying... [id=subnet-01234d5678d87e6bb]
aws_security_group.allow_web: Destroying... [id=sg-08c7654bc3b21a9c8]
aws_subnet.subnet-1: Destruction complete after 2s
aws_security_group.allow_web: Destruction complete after 2s
aws_vpc.prod-vpc: Destroying... [id=vpc-0d234d6543ab21ae0]
aws_vpc.prod-vpc: Destruction complete after 1s

Destroy complete! Resources: 9 destroyed.
```

## Reference
* [`aws_vpc`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc)
* [`aws_eip`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eip)
* [`aws_subnet`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet)
* [`aws_route_table`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table)
* [`aws_security_group`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group)
* [`aws_internet_gateway`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/internet_gateway)
* [`aws_network_interface`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_interface)
* [`aws_route_table_association`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table_association)
* [Upgrade Your SSH Key to Ed25519](https://medium.com/risan/upgrade-your-ssh-key-to-ed25519-c6e8d60d3c54)