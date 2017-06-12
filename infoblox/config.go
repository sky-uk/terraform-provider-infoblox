package infoblox

import (
	"github.com/sky-uk/skyinfoblox"
	"log"
)

// Config is a struct for containing the provider parameters.
type Config struct {
	Debug       bool
	Insecure    bool
	IBXUsername string
	IBXPassword string
	IBXServer   string
}

// Client returns a new client for accessing VMWare vSphere.
func (c *Config) Client() (*skyinfoblox.InfobloxClient, error) {
	log.Printf("[INFO] VMWare NSX Client configured for URL: %s", c.IBXServer)
	nsxclient := skyinfoblox.NewInfobloxClient("https://"+c.IBXServer, c.IBXUsername, c.IBXPassword, c.Insecure, c.Debug)
	return nsxclient, nil
}
