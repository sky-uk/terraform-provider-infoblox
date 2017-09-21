# skyinfoblox - Go library for the Infoblox appliance

This is the GoLang API wrapper for Infoblox. This is currently used for building terraform provider for the same appliance.
This package is based on the Infoblox WAPI library version v2.6.1.
Wapi library documentation can be accessed here:

https://h1infoblox.devops.int.ovp.bskyb.com/wapidoc/index.html


## Run Unit tests
```
make test

```

## Building the cli binary
```
make all

```

This will give you skyinfoblox-cli file which you can use to interact with InfoBlox API.

## Using the library
Note: starting from release 0.1.0, a new API has been provided. The new API supports CRUD operations on all object types.

## WAPI version used
The Server WAPI version can be configured at client creation time (it defaults anyhow to v2.6.1). Support for older versions schemes is driven by the Infoblox WAPI server itself (query your server schema to find out the list of supported versions).

### Getting an API client object
In order to get an API client object first set a list of connnection parameters and pass it to the Connect() function:

```
	params := Params{
		WapiVersion: "v2.6.1", // this is anyhow the default...
		URL:         server,
		User:        username,
		Password:    password,
		IgnoreSSL:   true,
		Debug:       true,
	}

    client := Connect(params)

```

### Creating an object
You can create any object you like setting object profile in a map:

```
	adminRole := make(map[string]interface{})
	adminRole["name"] = "a role name"
	adminRole["comment"] = "An initial comment"

	refObj, err := client.Create("adminrole", adminRole)
```

Or, for a limited selection of objects we directly support, you can use our provided structs:

```
		disable := true
		superUser := false

		adminGroup := model.IBXAdminGroup{
			AccessMethod:   []string{"API"},
			Comment:        "API Access only",
			Disable:        &disable,
			EmailAddresses: []string{"test@example-test.com"},
			Name:           "test",
			Roles:          []string{"test-role"},
			SuperUser:      &superUser,
		}

        refObj, err := client.Create("admingroup", adminGroup)

```

### Deleting an object

```
    refObj, err = client.Delete(refObj)
```

The Delete() function returns the deleted object reference or an error otherwise

#### Reading an object

```
    obj := make(map[string]interface{})
    err = client.Read(objRef, []string{<list of attrs you want back>}, &obj)
```

### Updating an object

```
    updatedRefObj, err := client.Update(refObj, newObjectProfileAsMap)
```

### Creating and reading an object
This function can come handy if you like to get anyhow the created
object back instead of only its reference.
It returns the created object as map[string]interface{}
You can create any object you like setting object profile in a map:

```
	adminRole := make(map[string]interface{})
	adminRole["name"] = "a role name"
	adminRole["comment"] = "An initial comment"

	myObj, err := client.CreateAndRead("adminrole", adminRole)
```

### Updating and reading an object
Again this can be handy to get the updated object back

```
    updatedObj, err := client.UpdateAndRead(refObj, newObjectProfileAsMap)
```


## CLI Usage
If called without parameters, the cli will print its needed input parameters and the 
provided commands

```
$ ./skyinfoblox-cli
  -debug
    	Debug output. Default:false
  -password string
    	Authentication password (Env: IBX_PASSWORD)
  -port int
    	Infoblox API server port. Default:443 (default 443)
  -server string
    	Infoblox API server hostname or address. (Env: IBX_SERVER)
  -username string
    	Authentication username (Env: IBX_USERNAME)
  -wapiVersion string
    	WAPI version (defaults to v2.6.1) 
  Commands:
    create-and-read-object
    create-object
    delete-object
    read-all-objects
    read-object
    update-and-read-object
    update-object

```

### Creating an object with the CLI

To create an object, define first its profile in a json-encoded file, then issue the command:

```
./skyinfoblox-cli <all needed input params> create-object -type <object type> -profile <json-file with encoded object profile>
```
The command should print the created object reference or the error returned by the Infoblox REST API (note: passing the -debug parameter will result in the detailed request and response being dumped on screen).

### Deleting 
To delete an object:

```
./skyinfoblox-cli <all needed input params> delete-object -ref <the object reference>
```

### Reading all objects of a given type
To get the list of all objects of a given type:

```
./skyinfoblox-cli <all needed input params> read-all-objects -type <object type>
```

You should get back the list of objects (default fields)

### Reading an object
To get the profile of an object:

```
./skyinfoblox-cli <all needed input params> read-object -ref <object reference> -return-params <comma-separated list of attributes to get back>
```

### To update an object
To update an object:

```
./skyinfoblox-cli <all needed input params> update-object -ref <object reference> -profile <json-encoded file with updated object profile>
```
