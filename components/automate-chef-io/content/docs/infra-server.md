+++
title = "Install Chef Infra Server With Automate"
description = "Use Automate to install Chef Infra Server"
date = 2020-02-11T14:24:00-08:00
weight = 20
draft = false
bref = ""
toc = true
[menu]
  [menu.docs]
    parent = "get_started"
    weight = 60
+++

## Overview
Use Chef Automate to install Chef Infra Server.

## System requirements and prerequisites
### Hardware
See the [hardware requirements for Chef Infra Server](https://docs.chef.io/install_server_pre.html) for guidance on memory and number of CPUs for a standalone Chef Infra Server installation.

For a single-host installation that contains Chef Infra Server and Chef Automate:

* up to 200 managed nodes: 8GB RAM, 2 vCPUs
* between 200 and 500 managed nodes: 30GB RAM, 4 vCPUs
* between 500 and 5000 managed nodes: 61GB RAM, 8vCPUs
* 80 GB of disk space plus 2 MB of disk space per managed node, available to `/hab`

### Operating System
* a Linux kernel of version 3.2 or greater
* `systemd` as the init system
* `useradd`
* `curl` or `wget`
* The shell that starts Automate should have a max open files setting of at least 65535
* Run the installation and bootstrapping procedures as the superuser or use `sudo` at the start of each command.

### Download the `chef-automate` command line tool
Download and unzip Chef Automate:

```shell
curl https://packages.chef.io/files/current/latest/chef-automate-cli/chef-automate_linux_amd64.zip | gunzip - > chef-automate && chmod +x chef-automate
```

## Install Chef Automate and Chef Infra Server on the same host
You can use either a command-line option or a configuration file to install Chef Automate
and Chef Infra Server on the same host.

### Via command line
To install Chef Automate and Chef Infra Server on the same host, run the
command:

```shell
sudo chef-automate deploy --product automate --product chef-server
```

### Via configuration

1. Generate a skeleton configuration file by running
    ```shell
      sudo chef-automate init-config
    ```

1. Add a stanza to the configuration file to deploy Automate and Chef Infra Server
    ```toml
      [deployment.v1.svc]
      products=["automate", "chef-server"]
    ```

1. Make any other configuration changes desired.

1. Run `chef-automate deploy` with your configuration file:

    ```shell
      sudo chef-automate deploy config.toml
    ```

## Install standalone Chef Infra Server using Chef Automate
You can use either a command-line option or a configuration file to install Chef Infra Server via Chef Automate.

### Via command line
1. Generate a skeleton configuration file by running
    ```shell
       sudo chef-automate init-config
    ```

1. Add a stanza to the configuration file to disable Automate data collection:
    ```toml
       [erchef.v1.sys.data_collector]
       enabled = false
    ```

1. Use the configuration file to deploy Chef Infra Server:
    ```shell
       sudo chef-automate deploy --product chef-server <configuration_file>
    ```

### Via configuration
1. Generate a skeleton configuration file by running
    ```shell
      sudo chef-automate init-config
    ```

1. Add a stanza to the configuration file to deploy Chef Infra Server
    ```toml
       [deployment.v1.svc]
       products=["chef-server"]

       # Disable Automate data collection as Automate will not be deployed
       [erchef.v1.sys.data_collector]
       enabled = false
    ```

1. Run `chef-automate deploy` with your configuration file:

    ```shell
      sudo chef-automate deploy config.toml
    ```

## Further reading
See the [Chef Infra Server documentation](https://docs.chef.io/server_overview.html) for
instructions and guidance on using and managing your Chef Infra Server.