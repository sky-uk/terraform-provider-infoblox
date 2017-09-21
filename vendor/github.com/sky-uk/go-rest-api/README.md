# go-rest-api

A fairly generic HTTP API.

Supports both HTTP and TLS (https).
Encoding schemes supported: json/xml.


## Importing

```
import(
    "rest"
    "net/http"
)
```

## Usage

### Get a REST Client object

```
    client := rest.Client{
        URL: url,       // mandatory
        User: user, 
        Password: password, 
        IgnoreSSL: ignoreSSL, 
        Debug: debug, 
        Headers: headers,
        Timeout: 30     // in seconds
    } 

    // for a simple http client...
    client := rest.Client{URL: url}

```

### Perform a request

```
    // Prepare a request...
    api := rest.NewBaseAPI(
        http.MethodGet,         // request method
        "/",                    // request path
        nil,                    // request payload object
        new(string),            // (pointer to) response payload object
        nil,                    // (pointer to) error object
    )

    // Perform the request...
    err := client.Do(api)
    if err != nil {
        // handle errors....
    }
```

### Getting the response object

```

    // example (json payload)
    type JSONFoo struct {
	    Fields map[string]string `json:"fields"`
    }

    // Prepare a request...
    api := rest.NewRestAPI(
        http.MethodGet,         // request method
        "/",                    // request path
        nil,                    // request payload object
        new(JSONFoo),           // (pointer to) response payload object
        nil,                    // (pointer to) error object
    )

    //...Perform request...
    // Get response object...
    respObj := *api.ResponseObject().(*JSONFoo)
```

More examples for the supported encodings in the client_test.go module.

### Getting response status code

```
    status := api.StatusCode()
```

### Getting the raw response as a byte stream

```
    raw := api.RawResponse()
```
