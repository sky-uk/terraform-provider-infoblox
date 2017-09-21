package util

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

// AccessControlSchema - returns the schema for an access control
func AccessControlSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Determines whether dynamic DNS updates are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks, and TSIG keys for the hosts.",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"_struct": {
					Type:         schema.TypeString,
					Description:  "Specifies the type of struct we're passing",
					Optional:     true,
					ValidateFunc: ValidateAcType,
				},
				"address": {
					Type:         schema.TypeString,
					Description:  "The address this rule applies to or ANY",
					Optional:     true,
					ValidateFunc: CheckLeadingTrailingSpaces,
				},
				"permission": {
					Type:         schema.TypeString,
					Description:  "The permission to use for this address",
					Optional:     true,
					ValidateFunc: ValidateAddressAcPermission,
				},
				"tsig_key": {
					Type:         schema.TypeString,
					Description:  "A generated TSIG key",
					Optional:     true,
					ValidateFunc: CheckLeadingTrailingSpaces,
				},
				"tsig_key_alg": {
					Type:         schema.TypeString,
					Description:  "The TSIG key algorithm",
					Optional:     true,
					ValidateFunc: ValidateTSIGAlgorithm,
				},
				"tsig_key_name": {
					Type:         schema.TypeString,
					Description:  "The name of the TSIG key",
					Optional:     true,
					ValidateFunc: CheckLeadingTrailingSpaces,
				},
				"use_tsig_key_name": {
					Type:        schema.TypeBool,
					Description: "Use flag for: tsig_key_name",
					Optional:    true,
				},
			},
		},
	}
}

// ValidateAcType - validates if the access control type is correct
func ValidateAcType(v interface{}, k string) (ws []string, errors []error) {
	acType := v.(string)
	if acType != "addressac" && acType != "tsigac" {
		errors = append(errors, fmt.Errorf("%q must be one of addressac or tsigac", k))
	}
	return
}

// ValidateAddressAcPermission - validates if the permission type is correct
func ValidateAddressAcPermission(v interface{}, k string) (ws []string, errors []error) {
	permission := v.(string)
	if permission != "ALLOW" && permission != "DENY" {
		errors = append(errors, fmt.Errorf("%q must be one of ALLOW or DENY", k))
	}
	return
}
