AzureRM-SRV Terraform Provider
==================

General Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.11.x (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/abhinavdahiya/terraform-provider-azurerm-srv`

```sh
$ mkdir -p $GOPATH/src/github.com/abhinavdahiya; cd $GOPATH/src/github.com/abhinavdahiya
$ git clone git@github.com:abhinavdahiya/terraform-provider-azurerm-srv
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-azurerm
$ GOBIN=~/.terraform.d/plugins go install
```

Why this special Provider
-------------------------

This code is almost all from upstream [azurerm provider][terraform-provider-azurerm], with the only exception that the
`record` field is a `TypeList` compared to `TypeSet`.

The `TypeList` allows users to create the list of records in the SRV record dynamically, which is not possible with upstream unless
we wait for `terraform 0.12` to become generally available with `foreach` semantics.

This allows users to create SRV record like:

```tf
provider "azurerm" {}
provider "azurerm-srv" {}

resource "azurerm_resource_group" "main" {
  name     = "example"
  location = "westus2"
}

resource "azurerm_virtual_network" "main" {
  name                = "example"
  resource_group_name = "${azurerm_resource_group.main.name}"
  location            = "westus2"
  address_space       = ["10.0.0.0/16"]
}

resource "azurerm_dns_zone" "private" {
  name                           = "thisistest.com"
  resource_group_name            = "${azurerm_resource_group.main.name}"
  zone_type                      = "Private"
  resolution_virtual_network_ids = ["${azurerm_virtual_network.main.id}"]
}

resource "null_resource" "etcd_mappings" {
  count = 4

  triggers {
    priority = 10
    weight   = 10
    port     = 2380
    target   = "etcd-${count.index}.${azurerm_dns_zone.private.name}"
  }
}

resource "azurerm-srv_dns_srv_record" "etcd_cluster" {
  name                = "_etcd-server-ssl._tcp"
  zone_name           = "${azurerm_dns_zone.private.name}"
  resource_group_name = "${azurerm_resource_group.main.name}"
  ttl                 = 60

  record = ["${null_resource.etcd_mappings.*.triggers}"]
}
```

[terraform-provider-azurerm]: https://github.com/terraform-providers/terraform-provider-azurerm