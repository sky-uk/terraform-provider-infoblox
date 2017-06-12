# Terraform Infoblox Provider

This document is in place for developer documentation.  User documentation is located [HERE](https://www.terraform.io/docs/providers/infoblox/) on Terraform's website.

A Terraform provider for the Infoblox.  The Infoblox provider is used to interact with resources supported by the Infoblox Server.
The provider needs to be configured with the proper credentials before it can be used.

## Introductory Documentation

Both [README.md](../../../README.md) and [BUILDING.md](../../../BUILDING.md) should be read first!

## Base API Dependency ~ [go-infoblox](https://github.com/sky-uk/go-infoblox)

This provider utilizes [go-infoblox](https://github.com/sky-uk/go-infoblox) Go Library for communicating Infoblox REST API.
Because of the dependency this provider is compatible with Infoblox systems that are supported by go-infoblox. If you want to contributed additional functionality into go-infoblox API bindings
please feel free to send the pull requests.


## Resources Implemented
| Feature                 | Create | Read  | Update  | Delete |
|-------------------------|--------|-------|---------|--------|
| DNS Zones               |   N    |   N   |    N    |   N    |
| DNS Records             |   N    |   N   |    N    |   N    |


### Limitations

This is currently a proof of concept and only has a very limited number of
supported resources.  These resources also have a very limited number
of attributes.

This section is a work in progress and additional contributions are more than welcome.