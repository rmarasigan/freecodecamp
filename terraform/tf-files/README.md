# Terraform Files

### File Hierarchy
```
project-name
│
└───.terraform
│   │
│   └───providers
│     │
│     └───registry.terraform.io
│        │
│        └───hashicorp
│           │
│           └───aws
│              │
│              └───4.25.0
│                 │
│                 └───linux_amd64
│                   │
│                   └───terraform-provider-aws_v4.25.0_x5
│   
└───.terraform.lock.hcl
│   
└───main.tf
|
└───terraform.tfstate
│   
└───terraform.tfstate.backup
```

## main.tf
It contains the main set of configurations for your module. You can also create other configuration files and organize them however makes sense for your project.

## .terraform
This folder gets created whenever we initialize any plugins. When we did a `terraform init`, it is going to create that directory and then install all the necessary plugins for our code to run in this folder. Since we have only one provider, all of the code for the AWS provider is going to get installed into `terraform-provider-aws_v4.25.0_x5`. This directory contains the modules and plugins used to provision your infrastructure. These files are specific to a specific instance of Terraform when provisioning infrastructure, not the configuration of the infrastructure defined in `.tf` files.

If we try to delete the **`.terraform`** folder and run `terraform apply` it will throw an error.
```bash
dev@dev:~$ terraform apply
╷
│ Error: Required plugins are not installed
│ 
│ The installed provider plugins are not consistent with the packages selected in the dependency lock file:
│   - registry.terraform.io/hashicorp/aws: there is no package for registry.terraform.io/hashicorp/aws 4.25.0 cached in .terraform/providers
│ 
│ Terraform uses external plugins to integrate with a variety of different infrastructure services. To download the plugins required for this configuration, run:
│   terraform init
╵
```

So what we have to do is we have to initialize it again by using the `terraform init` command.

## terraform.tfstate
It represents all of the state for Terraform. Anytime we create a resource within AWS or any cloud provider, we need a way for Terraform to keep track of everything that's created. That way, if we go to modify a parameter, like an extra tag, or maybe change the instance type from like `t2.micro` to another size, it needs to be able to check what is the current status of that resource, what are its configuration, and compare it to what is in the code. The way it does that is it just creates a simple text file and so we've got the **`terraform.tfstate`** and it looks like it is in a JSON format.

We can store all of the resources that we create with this state file. So this file is very important. For some reason, if you ever go into this and start deleting things, you will break Terraform. So Terraform will have a mismatched state from what's being deployed into AWS. So never mess around with the `tfstate` file. Anytime we create an instance or modify an instance within AWS, we're going to update our state file so that we can keep track of all of that information.

## terraform.tfstate.backup
By default, a backup of your state file is written to **`terraform.tfstate.backup`** in case the state file is lost or corrupted to simplify recovery. The state file is used by Terraform to keep track of resources and metadata information about your infrastructure. By default, the state of your environment is stored locally in your Terraform workspace directory in a file called `terraform.tfstate` along with a backup file called `terraform.tfstate.backup`. The state file will not exist until you have completed at least one `terraform apply`.

## .terraform.lock.hcl
The lock file is always named `.terraform.lock.hcl`, and this name is intended to signify that it is a lock file for various items that Terraform caches in the `.terraform` subdirectory of your working directory. Terraform automatically creates or updates the dependency lock file each time you run the `terraform init` command.


## Reference
* [Dependency Lock](https://www.terraform.io/language/files/dependency-lock#lock-file-location)
* [Build and Use a Local Module](https://learn.hashicorp.com/tutorials/terraform/module-create)
* [What is terraform.tfstate.backup file in terraform?](https://www.devopsschool.com/blog/what-is-terraform-tfstate-backup-file-in-terraform/)