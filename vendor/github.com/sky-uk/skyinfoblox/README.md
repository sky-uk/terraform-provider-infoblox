# skyinfoblox - Go library for the Infoblox appliance

This is the GoLang API wrapper for Infoblox. This is currently used for building terraform provider for the same.
This package is based on the Infoblox WAPI library version 2.3.1.
Wapi library documentation can be accessed here:

https://h1infoblox.devops.int.ovp.bskyb.com/wapidoc/index.html



## Building the cli binary
```
make all

```

## Run Unit tests
```
make test

```

This will give you skyinfoblox-cli file which you can use to interact with InfoBlox API.

```
$ ./skyinfoblox-cli
  -debug
    	Debug output. Default:false
  -password string
    	Authentication password (Env: IBX_PASSWORD)
  -port int
    	Infoblox API server port. Default:443 (default 443)
  -server string
    	Infoblox API server hostname or address (default "localhost")
  -username string
    	Authentication username (Env: IBX_USERNAME)
  Commands:
    admin-group-create
    admin-group-delete
    admin-group-show
    admin-group-show-all
    admin-group-update
    admin-role-create
    admin-role-delete
    admin-role-show
    admin-role-show-all
    admin-role-update
    adminuser-create
    adminuser-delete
    adminuser-show
    adminuser-update
    network-create
    network-delete
    network-show
    network-update
    networks-show-all
    permission-create
    permission-delete
    permission-show
    permission-show-all
    permission-update
    range-create
    range-delete
    range-show
    range-update
    record
    record-ns-create
    record-ns-delete
    record-ns-update
    record-txt-create
    records-show-all
    zone-create
    zone-delete
    zone-show
    zone-show-all
    zone-update
    zoneforward-create
    zoneforward-delete
    zoneforward-show
    zoneforward-show-all
    zoneforward-update

```

```
./skyinfoblox-cli -server=https://serverhostnameOrIP  -username=admin -password=password records-list -a
|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Name                                  | IPv4        | Ref                                                                                                                                                   |
|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| yorg.test.np.ovp.sky.com              | 10.10.10.10 | record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLnNreS5vdnAubnAsdGVzdC55b3JnLDEwLjEwLjEwLjEw:yorg.test.np.ovp.sky.com/default                                |
|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| h1ins01.devops.int.ovp.bskyb.com      | 10.77.58.10 | record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLm92cCxpbnQuZGV2b3BzLmgxaW5zMDEsMTAuNzcuNTguMTA:h1ins01.devops.int.ovp.bskyb.com/default             |
|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| s1ins01.devops.int.ovp.bskyb.com      | 10.93.58.10 | record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLm92cCxpbnQuZGV2b3BzLnMxaW5zMDEsMTAuOTMuNTguMTA:s1ins01.devops.int.ovp.bskyb.com/default             |
|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| h1ifbr02-v01.devops.int.ovp.bskyb.com | 10.77.58.20 | record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLm92cCxpbnQuZGV2b3BzLmgxaWZicjAyLXYwMSwxMC43Ny41OC4yMA:h1ifbr02-v01.devops.int.ovp.bskyb.com/default |
|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| craigtest.test-ovp.bskyb.com          | 10.10.10.1  | record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlndGVzdCwxMC4xMC4xMC4x:craigtest.test-ovp.bskyb.com/default                        |
|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| craig2test.test-ovp.bskyb.com         | 10.10.10.2  | record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnMnRlc3QsMTAuMTAuMTAuMg:craig2test.test-ovp.bskyb.com/default                     |
|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| craig3test.test-ovp.bskyb.com         | 10.10.1.80  | record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnM3Rlc3QsMTAuMTAuMS44MA:craig3test.test-ovp.bskyb.com/default                     |
|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|


```

## Development

during your development, you can run the cli with following command.
```
go run cli/*.go -server=https://infobloxserver.com  -username=admin -password=pass records-list

```
