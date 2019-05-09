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
