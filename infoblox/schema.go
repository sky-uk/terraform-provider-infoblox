package infoblox

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

//ResourceAttr - attribute metadata
type ResourceAttr struct {
	Name  string
	Type  schema.ValueType
	Value interface{}
}

//GetAttrs - returns the list of attributes names and types
func GetAttrs(resource *schema.Resource) []ResourceAttr {
	attrs := make([]ResourceAttr, 0)

	s := resource.Schema

	str := spew.Sdump(s)
	log.Println("Schema:\n", str)

	for key := range s {
		var attr ResourceAttr
		attr.Name = key
		attr.Type = s[key].Type
		attrs = append(attrs, attr)
	}
	str = spew.Sdump(attrs)
	log.Println("Attributes:\n", str)
	return attrs
}

// GetValue - returns the value of an attribute
// id does some transformations for specific types
func GetValue(attr ResourceAttr) interface{} {
	switch attr.Type {
	case schema.TypeSet:
		v := attr.Value.(*schema.Set)
		return v.List()
	}
	return attr.Value
}
