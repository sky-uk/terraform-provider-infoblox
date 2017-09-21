package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"log"
	"reflect"
)

// CreateResource - Creates a new resource provided its resource schema
func CreateResource(name string, resource *schema.Resource, d *schema.ResourceData, m interface{}) error {

	obj := make(map[string]interface{})
	attrs := GetAttrs(resource)
	for _, attr := range attrs {
		key := attr.Name
		log.Println("Found attribute: ", key)
		if v, ok := d.GetOk(key); ok {
			attr.Value = v
			obj[key] = GetValue(attr)
		}
	}

	params := m.(map[string]interface{})
	client := params["ibxClient"].(*skyinfoblox.Client)

	log.Printf("Going to create an %s object: %+v", name, obj)
	ref, err := client.Create(name, obj)
	if err != nil {
		d.SetId("")
		return err
	}
	d.SetId(ref)
	return ReadResource(resource, d, m)
}

// CreateAndReadResource - Creates a new resource provided its resource schema and read it back
func CreateAndReadResource(name string, resource *schema.Resource, d *schema.ResourceData, m interface{}) error {

	obj := make(map[string]interface{})
	attrs := GetAttrs(resource)
	for _, attr := range attrs {
		key := attr.Name
		log.Println("Found attribute: ", key)
		if v, ok := d.GetOk(key); ok {
			attr.Value = v
			obj[key] = GetValue(attr)
		}
	}

	params := m.(map[string]interface{})
	client := params["ibxClient"].(*skyinfoblox.Client)

	log.Printf("Going to create an %s object: %+v", name, obj)
	createdObj, err := client.CreateAndRead(name, obj)
	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(createdObj["_ref"].(string))
	delete(createdObj, "_ref")
	for key := range createdObj {
		if isScalar(createdObj[key]) == true {
			log.Printf("Setting key %s to %+v\n", key, createdObj[key])
			d.Set(key, createdObj[key])
		}
	}
	return nil
}
func isScalar(field interface{}) bool {
	t := reflect.TypeOf(field)
	if t == nil {
		return false
	}
	k := t.Kind()
	switch k {
	case reflect.Slice:
		return false
	case reflect.Map:
		return false
	}
	return true
}

// ReadResource - Reads a resource provided its resource schema
func ReadResource(resource *schema.Resource, d *schema.ResourceData, m interface{}) error {

	params := m.(map[string]interface{})
	client := params["ibxClient"].(*skyinfoblox.Client)

	ref := d.Id()
	obj := make(map[string]interface{})

	attrs := GetAttrs(resource)
	keys := []string{}
	for _, attr := range attrs {
		keys = append(keys, attr.Name)
	}
	err := client.Read(ref, keys, &obj)
	if err != nil {
		d.SetId("")
		return err
	}

	delete(obj, "_ref")
	for key := range obj {
		if isScalar(obj[key]) == true {
			log.Printf("Setting key %s to %+v\n", key, obj[key])
			d.Set(key, obj[key])
		}
	}

	return nil
}

// DeleteResource - Deletes a resource
func DeleteResource(d *schema.ResourceData, m interface{}) error {

	params := m.(map[string]interface{})
	client := params["ibxClient"].(*skyinfoblox.Client)

	ref := d.Id()
	ref, err := client.Delete(ref)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

// UpdateResource - Updates a resource provided its schema
func UpdateResource(resource *schema.Resource, d *schema.ResourceData, m interface{}) error {

	needsUpdate := false

	params := m.(map[string]interface{})
	client := params["ibxClient"].(*skyinfoblox.Client)

	ref := d.Id()
	obj := make(map[string]interface{})

	attrs := GetAttrs(resource)
	for _, attr := range attrs {
		key := attr.Name
		if d.HasChange(key) {
			attr.Value = d.Get(key)
			obj[key] = GetValue(attr)
			log.Printf("Updating field %s, value: %+v\n", key, obj[key])
			needsUpdate = true
		}
	}

	log.Printf("UPDATE: going to update reference %s with obj: \n%+v\n", ref, obj)

	if needsUpdate {
		newRef, err := client.Update(ref, obj)
		if err != nil {
			return err
		}
		d.SetId(newRef)
	}

	return ReadResource(resource, d, m)
}

// UpdateAndReadResource - Updates a resource provided its schema
func UpdateAndReadResource(resource *schema.Resource, d *schema.ResourceData, m interface{}) error {

	needsUpdate := false

	params := m.(map[string]interface{})
	client := params["ibxClient"].(*skyinfoblox.Client)

	ref := d.Id()
	obj := make(map[string]interface{})

	attrs := GetAttrs(resource)
	for _, attr := range attrs {
		key := attr.Name
		if d.HasChange(key) {
			attr.Value = d.Get(key)
			obj[key] = GetValue(attr)
			log.Printf("Updating field %s, value: %+v\n", key, obj[key])
			needsUpdate = true
		}
	}

	log.Printf("UPDATE: going to update reference %s with obj: \n%+v\n", ref, obj)

	if needsUpdate {
		newObject, err := client.UpdateAndRead(ref, obj)
		if err != nil {
			return err
		}
		d.SetId(newObject["_ref"].(string))
		delete(newObject, "_ref")
		for key := range newObject {
			if isScalar(newObject[key]) == true {
				log.Printf("Updating key %s to %+v\n", key, newObject[key])
				d.Set(key, newObject[key])
			}
		}
	}

	return nil
}
