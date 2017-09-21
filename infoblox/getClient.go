package infoblox

import (
	"github.com/sky-uk/skyinfoblox"
	"os"
)

//GetClient - get a skyinfoblox client..
func GetClient() *skyinfoblox.Client {
	params := skyinfoblox.Params{
		WapiVersion: "v2.6.1", // this is anyhow the default...
		URL:         os.Getenv("INFOBLOX_SERVER"),
		User:        os.Getenv("INFOBLOX_USERNAME"),
		Password:    os.Getenv("INFOBLOX_PASSWORD"),
		IgnoreSSL:   true,
		Debug:       true,
	}
	client := skyinfoblox.Connect(params)

	return client
}
