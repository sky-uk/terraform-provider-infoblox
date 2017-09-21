package infoblox

import (
	"fmt"
	"log"
)

// TestAccCheckDestroy - checks that no object of a given type and with a given
// key/value pair exists
func TestAccCheckDestroy(objType, key, value string) error {
	client := GetClient()
	recs, err := client.ReadAll(objType)
	if err != nil {
		return err
	}
	for _, rec := range recs {
		if rec[key] == value {
			return fmt.Errorf("Object with key %s = %s still exists", key, value)
		}
	}
	return nil
}

// TestAccCheckExists - checks that object of a given type and with a given
// key/value pair exists
func TestAccCheckExists(objType, key, value string) error {
	client := GetClient()
	recs, err := client.ReadAll(objType)
	if err != nil {
		return err
	}
	for _, rec := range recs {
		if rec[key] == value {
			log.Printf("Resource with %s equal to %s exists\n", key, value)
			return nil
		}
	}
	return fmt.Errorf("Object with key %s = %s does not exists", key, value)
}
