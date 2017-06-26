# ZoneAuth Resource
The zone auth resource is used for managing Infoblox DNS zones. 

## Attributes
|Attribute                        |Create|Read|Update|Delete|
|---------------------------------|------|----|------|------|
|fqdn                             |   Y  |  Y |   N* |   Y  |
|view                             |   Y  |  Y |   N* |   Y  |
|comment                          |   Y  |  Y |   Y  |   Y  |
|zoneformat                       |   Y  |  Y |   N* |   Y  |
|prefix                           |   Y  |  Y |   Y  |   Y  |
|disable                          |   Y  |  Y |   Y  |   Y  |
|dnsintegrityenable               |   Y  |  Y |   Y  |   Y  |
|dnsintegritymember               |   Y  |  Y |   Y  |   Y  |
|locked                           |   Y  |  Y |   Y  |   Y  |
|lockedby**                       |   N  |  Y |   N  |   N  |
|networkview**                    |   N  |  Y |   N  |   N  |
|soaserialnumber***               |   N  |  Y |   N  |   N  |
|soattl                           |   Y  |  Y |   Y  |   Y  |
|soanegativettl                   |   Y  |  Y |   Y  |   Y  |
|soarefresh                       |   Y  |  Y |   Y  |   Y  |
|soaretry                         |   Y  |  Y |   Y  |   Y  |
|allowupdate                      |   Y  |  Y |   Y  |   Y  |
|allowupdate => type              |   Y  |  Y |   Y  |   Y  |
|allowupdate => address           |   Y  |  Y |   Y  |   Y  |
|allowupdate => permission        |   Y  |  Y |   Y  |   Y  |
|allowupdate => tsigkey           |   Y  |  Y |   Y  |   Y  |
|allowupdate => tsigkeyalgorithm  |   Y  |  Y |   Y  |   Y  |
|allowupdate => tsigkeyname       |   Y  |  Y |   Y  |   Y  |
|allowupdate => usetsigkeyname    |   Y  |  Y |   Y  |   Y  |

*Infoblox API doesn't permit updates to these attributes. Changing the attribute will force delete/create!!!  
**The Infoblox API doesn't permit changing or setting the lockedby nor networkview attributes.  
***The attribute soaserialnumber is read-only from Terraform's point of view as Infoblox auto-increments the serial.

## Sample Reverse Zone Template

```
resource "infoblox_zone_auth" "some_reverse_zone" {  
  fqdn = "192.168.234.0/24"  
  comment = "This is a comment about my zone"  
  zoneformat = "IPV4"  
  view = "default"  
  prefix = "128/6"  
  disable = false  
  dnsintegrityenable = false  
  dnsintegritymember = "s1ins01.devops.int.ovp.bskyb.com"  
  locked = false  
  soattl = 120  
  soanegativettl = 12  
  soarefresh = 1200  
  soaretry = 150  
  allowupdate = [  
  {  
    type = "addressac"  
    address = "192.168.100.11"  
    permission = "DENY"  
  },  
  {  
    type = "tsigac"  
    tsigkey = "0jnu3SdsMvzzlmTDPYTceA=="  
    tsigkeyalgorithm = "HMAC-SHA256"  
    tsigkeyname = "test.key"  
    usetsigkeyname = true  
  },  
  {  
    type = "addressac"  
    address = "192.168.101.15"  
    permission = "ALLOW"  
  },  
  {  
    type = "addressac"  
    address = "192.168.111.10"  
    permission = "ALLOW"  
  },  
  ]  
}    
```

## Sample Forward Zone Template

```
resource "infoblox_zone_auth" "some_forward_zone" {  
  fqdn = "paas-testing.com"  
  comment = "Some forward zone comment"  
  zoneformat = "FORWARD"  
  view = "default"  
  disable = false  
  dnsintegrityenable = false  
  dnsintegritymember = "s1ins01.devops.int.ovp.bskyb.com"  
  locked = false   
  soattl = 120  
  soanegativettl = 12  
  soarefresh = 1200  
  soaretry = 150  
  allowupdate = [  
  {  
    type = "addressac"  
    address = "192.168.100.11"  
    permission = "DENY"  
  },  
  {  
    type = "tsigac"  
    tsigkey = "0jnu3SdsMvzzlmTDPYTceA=="  
    tsigkeyalgorithm = "HMAC-SHA256"  
    tsigkeyname = "test.key"  
    usetsigkeyname = true  
  },  
  {  
    type = "addressac"  
    address = "192.168.101.15"  
    permission = "ALLOW"  
  },  
  {  
    type = "addressac"  
    address = "192.168.111.10"  
    permission = "ALLOW"  
  },  
  ]  
}      
```
