
# terraform-provider-catalystcenter

terraform-provider-catalystcenter is a Terraform Provider for [Cisco Catalyst Center](https://developer.cisco.com/docs/dna-center/)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.13.x
- [Go](https://golang.org/doc/install) 1.20 (to build the provider plugin)

## Introduction

The terraform-provider-catalystcenter provides a Terraform provider for managing and automating your Cisco Catalyst Center environment. It consists of a set of resources and data-sources for performing tasks related to Catalyst Center.

This collection has been tested and supports Cisco Catalyst Center 2.3.5.3.

Other versions of this collection have support for previous Cisco Catalyst Center versions. The recommended versions are listed below on the [Compatibility matrix](#compatibility-matrix).

## Compatibility matrix
The following table shows the supported versions.

| Cisco Catalyst Center version | Terraform "catalystcenter" provider version |
|--------------------------|----------------------------------------|
| 2.3.7.6                  | 1.0.0-beta                             |
| 2.3.7.9                  | 1.1.1-beta                             |

If your SDK, Terraform provider is older please consider updating it first.

## Using the provider

There are two ways to get and use the provider.
1. Downloading & installing it from registry.terraform.io
2. Building it from source

### From build (For test)

Clone this repository to: `$GOPATH/src/github.com/cisco-en-programmability/terraform-provider-catalystcenter`

```sh
$ mkdir -p $GOPATH/src/github.com/cisco-en-programmability/
$ cd $GOPATH/src/github.com/cisco-en-programmability/
$ git clone https://github.com/cisco-en-programmability/terraform-provider-catalystcenter.git
```

Enter the provider directory and build the provider

! **NOTE**:
It is important to check the architecture of your operating system in the file [MakeFile](./Makefile)

```sh
$ cd $GOPATH/src/github.com/cisco-en-programmability/terraform-provider-catalystcenter
$ make developtest
```

If the Makefile values (HOSTNAME, NAMESPACE, NAME, VERSION) were not changed, then the following code could used without changes.
Otherwise change the values accordingly.


To use this provider, copy and paste this code into your Terraform configuration. Then, run terraform init.

```hcl
terraform {
  required_providers {
    catalystcenter = {
      source = "cisco-en-programmability/catalystcenter"
      version = "1.1.1-beta"
    }
  }
}

# Configure provider with your Cisco Catalyst Center SDK credentials
provider "catalystcenter" {
  # Cisco Catalyst Center user name
  username = "admin"
  # it can be set using the environment variable CATALYST_BASE_URL

  # Cisco Catalyst Center password
  password = "admin123"
  # it can be set using the environment variable CATALYST_USERNAME

  # Cisco Catalyst Center base URL, FQDN or IP
  base_url = "https://172.168.196.2"
  # it can be set using the environment variable CATALYST_PASSWORD

  # Boolean to enable debugging
  debug = "false"
  # it can be set using the environment variable CATALYST_DEBUG

  # Boolean to enable or disable SSL certificate verification
  ssl_verify = "false"
  # it can be set using the environment variable CATALYST_SSL_VERIFY
}
```

There are several examples of the use of the provider within the folder [samples](./examples/samples)

## Example for_each
```hcl
locals {
 interfaces = {
   "1" = { description = "desc1", interface_uuid = "c6820b57-ecde-4b6d-98db-06ba10486809" },
   "2" = { description = "desc2", interface_uuid = "c6820b57-ecde-4b6d-98db-06ba10486801" },
   "3" = { description = "desc3", interface_uuid = "c6820b57-ecde-4b6d-98db-06ba10486802" }
 }
}

resource "catalystcenter_interface_update" "example" {
  provider = catalystcenter
  for_each = local.interfaces
  parameters {
    description    = each.value.description
    interface_uuid = each.value.interface_uuid
    vlan_id        = each.key
  }
}
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed
on your machine (version 1.15+ is _required_). You'll also need to correctly setup a
[GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-catalystcenter
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

_Note:_ Acceptance tests create real resources.

```sh
$ make testacc
```

## Documentation

In the [docs directory](./docs/) you can find the documentation source for this 1.0.19-beta version.

You can find the documentation online for the previously released versions at [Terraform Registry - Cisco Catalyst Center provider](https://registry.terraform.io/providers/cisco-en-programmability/catalystcenter/latest/docs).


# Automatic Version Management for Resources and DataSources
In this Terraform provider, version management is handled dynamically for **all resources** and **datasources**. The version of a resource or datasource is automatically selected based on the available versions in the environment, allowing users to either use a default version or specify a particular version when needed. Below is an explanation of how this automatic version selection works for both resources and datasources.

## General Version Selection Behavior
For all resources and datasources, the behavior follows these general rules:

1. **If multiple versions of a resource or datasource exist (e.g., `resource_v1`, `resource_v2`)**:
   * By **default**, the **oldest** available version will be selected. That is, if both `resource_v1` and `resource_v2` are available, the provider will automatically select `v1` unless explicitly specified otherwise.
   * If a user wants to use the latest version (e.g., `v2`), they need to specify it explicitly in the configuration.

2. **If only one version of a resource or datasource exists (e.g., only `resource_v1`)**:
   * **The resource or datasource will automatically be used in that version**, without needing to specify the version suffix in the configuration.

3. **If only the latest version is available (e.g., only `resource_v2` and not `resource_v1`)**:
   * The resource or datasource will automatically refer to the **latest available version**, i.e., `resource_v2`.


## Example Usage with Resources
### Case 1: Both `resource_v1` and `resource_v2` are available
If both `resource_v1` and `resource_v2` are available in your environment, the provider will select `resource_v1` by default. If you want to use `resource_v2`, you can specify it explicitly:

```hcl
# Using version 1 (default)
resource "resource" "default" {
  # This resource will refer to `resource_v1` by default
}

# Explicitly using version 2
resource "resource_v2" "default" {
  # This resource will refer to `resource_v2`
}
```
### Case 2: Only resource\_v1 exists
If only resource\_v1 is available in your environment, there is no need to specify the version, as the provider will automatically use that version:

```hcl
resource "resource" "default" {

# This resource will refer to `resource_v1` since it's the only version available
}
```
### Case 3: Only resource\_v2 exists
If only resource\_v2 is available in your environment, the provider will automatically refer to that version, since resource\_v1 is not available:

```hcl
resource "resource" "default" {

# This resource will refer to `resource_v2` since `resource_v1` is not available
}
```
## Example Usage with DataSources
The same version selection logic applies to datasources. Below is how version management works for datasources:

### Case 1: Both datasource\_v1 and datasource\_v2 are available
```hcl
# Using version 1 (default)
data "datasource" "default" {

# This datasource will refer to `datasource_v1` by default
}

# Explicitly using version 2
data "datasource_v2" "default" {

# This datasource will refer to `datasource_v2`
}
```
### Case 2: Only datasource\_v1 exists
```hcl
data "datasource" "default" {

# This datasource will refer to `datasource_v1` since it's the only version available
}
```
### Case 3: Only datasource\_v2 exists
```hcl
data "datasource" "default" {

# This datasource will refer to `datasource_v2` since `datasource_v1` is not available
}
```

# Contributing

Ongoing development efforts and contributions to this provider are tracked as issues in this repository.

We welcome community contributions to this project. If you find problems, need an enhancement or need a new data-source or resource, please open an issue or create a PR against the [Terraform Provider for Cisco Catalyst Center repository](https://github.com/cisco-en-programmability/terraform-provider-catalystcenter/issues).

# Change log

All notable changes to this project will be documented in the [CHANGELOG](./CHANGELOG.md) file.

The development team may make additional changes as the library evolves with the Cisco Catalyst Center.

**NOTE**: Consider reviewing the Changelog to review the new features of the 1.0.19-beta version.

## License

This library is distributed under the license found in the [LICENSE](./LICENSE) file.
