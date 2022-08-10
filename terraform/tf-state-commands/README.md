# Terraform State Commands

If you don't know what commands to run, always just start with `terraform` and you'll see that the command line will kind of help you with the different available commands.

```bash
dev@dev:~$ terraform
Usage: terraform [global options] <subcommand> [args]

The available commands for execution are listed below.
The primary workflow commands are given first, followed by
less common or more advanced commands.

Main commands:
  init          Prepare your working directory for other commands
  validate      Check whether the configuration is valid
  plan          Show changes required by the current configuration
  apply         Create or update infrastructure
  destroy       Destroy previously-created infrastructure

All other commands:
  console       Try Terraform expressions at an interactive command prompt
  fmt           Reformat your configuration in the standard style
  force-unlock  Release a stuck lock on the current workspace
  get           Install or upgrade remote Terraform modules
  graph         Generate a Graphviz graph of the steps in an operation
  import        Associate existing infrastructure with a Terraform resource
  login         Obtain and save credentials for a remote host
  logout        Remove locally-stored credentials for a remote host
  output        Show output values from your root module
  providers     Show the providers required for this configuration
  refresh       Update the state to match remote systems
  show          Show the current state or a saved plan
  state         Advanced state management
  taint         Mark a resource instance as not fully functional
  test          Experimental support for module integration testing
  untaint       Remove the 'tainted' state from a resource instance
  version       Show the current Terraform version
  workspace     Workspace management

Global options (use these before the subcommand, if any):
  -chdir=DIR    Switch to a different working directory before executing the
                given subcommand.
  -help         Show this help output, or the help for a specified subcommand.
  -version      An alias for the "version" subcommand.
```

## `terraform state`
It is used for advanced state management and can be used in many cases. This command is a nested subcommand, meaning it has further subcommands.

```bash
dev@dev:~$ terraform state
Usage: terraform [global options] state <subcommand> [options] [args]

  This command has subcommands for advanced state management.

  These subcommands can be used to slice and dice the Terraform state.
  This is sometimes necessary in advanced cases. For your safety, all
  state management commands that modify the state create a timestamped
  backup of the state prior to making modifications.

  The structure and output of the commands is specifically tailored to work
  well with the common Unix utilities such as grep, awk, etc. We recommend
  using those tools to perform more advanced state tasks.

Subcommands:
    list                List resources in the state
    mv                  Move an item in the state
    pull                Pull current state and output to stdout
    push                Update remote state from a local state file
    replace-provider    Replace provider in the state
    rm                  Remove instances from the state
    show                Show a resource in the state
```

### `terraform state list`
It is used to list resources within a Terraform state. The command will list all resources in the state file matching the given addresses (if any). If no addresses are given, all resources are listed.

```bash
dev@dev:~$ terraform state list
aws_eip.web_server_eip
aws_instance.web-server-instance
aws_internet_gateway.gw
aws_network_interface.web-server-nic
aws_route_table.prod-route-table
aws_route_table_association.prod-route-table-association
aws_security_group.allow_web
aws_subnet.subnet-1
aws_vpc.prod-vpc
```

To show the detail of a specific resource, you can use `terraform state show resource_name`. So you can see it is going to give us a detailed output regarding that state that include a lot of information that would normally only be stored in the AWS Console. This is super helpful because there's gonna be a lot of times when you deploy something, and you want to know what the IP address is, especially the public IP.

Normally you'd have to do it through the AWS Console. But you know, we have this command so that we can just verify it from the command line really quickly without having to log in.

```bash
dev@dev:~$ terraform state show aws_eip.web_server_eip
# aws_eip.web_server_eip:
resource "aws_eip" "web_server_eip" {
    allocation_id             = "eipalloc-012d3e456c7895d6a"
    associate_with_private_ip = "10.0.1.50"
    association_id            = "eipassoc-02f13f4e56d78ddb2"
    domain                    = "vpc"
    id                        = "eipalloc-012d3e456c7895d6a"
    instance                  = "i-0a2f34a56f7ebea8b"
    network_border_group      = "us-east-1"
    network_interface         = "eni-0b0b45fdb4e76cc98"
    private_dns               = "ip-10-0-1-50.ec2.internal"
    private_ip                = "10.0.1.50"
    public_dns                = "ec2-5-210-68-161.compute-1.amazonaws.com"
    public_ip                 = "5.210.68.161"
    public_ipv4_pool          = "amazon"
    tags_all                  = {}
    vpc                       = true
```

## Reference
* [State Command](https://www.terraform.io/cli/commands/state)
* [Command: state list](https://www.terraform.io/cli/commands/state/list)