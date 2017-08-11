Terraform Provider
==================

Terraform provider for Infoblox appliance.

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/sky-uk/terraform-provider-infoblox`

```sh
$ mkdir -p $GOPATH/src/github.com/sky-uk; cd $GOPATH/src/github.com/sky-uk
$ git clone git@github.com:sky-uk/terraform-provider-infoblox
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/sky-uk/terraform-provider-infoblox
$ make build
```

Using the provider
----------------------
## Fill in for each provider


Template examples
------------------

 - Network
  
   This is an essential object to have DHCP scopes in Infoblox, by default when you create a network it makes it available for dhcp, this is controlled by the parameter "disable" which disables 
    the network for DHCP as the name suggests.
   Here we can see an example of how to create a network and the required parameters 
   ```
   resource "infoblox_network" "mynet" {
        comment = "My awesome network",
        network = "172.17.10.0/24",
        option {
               name = "routers",
               num = 3,
               useoption = true,
               value =  "172.17.10.1",
               vendorclass =  "DHCP"
        }
        option {
             name =  "domain-name-servers",
             num= 6,
             useoption =  true
             value = "8.8.8.8",
             vendorclass =  "DHCP"
        }
   } 
    ```
   For each DHCP option you need, you have to setup a option block as per the example above. 
   The only required field is the network field, so the bare minimum required to create a network is 
    ```
    resource "infoblox_network" "mynet" {
              network = "172.17.10.0/24"
    }
    ```
    

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-infoblox
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

