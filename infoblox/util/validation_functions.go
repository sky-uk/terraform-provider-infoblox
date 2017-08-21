package util

import (
	"fmt"
	"strings"
)

// ValidateZoneFormat - Checks the Zone Format is valid
func ValidateZoneFormat(v interface{}, k string) (ws []string, errors []error) {
	zoneFormat := v.(string)
	if zoneFormat != "FORWARD" && zoneFormat != "IPV4" && zoneFormat != "IPV6" {
		errors = append(errors, fmt.Errorf("%q must be one of FORWARD, IPV4 or IPV6", k))
	}
	return
}

//ValidateTSIGAlgorithm - Check for valid Encription Algorithms
func ValidateTSIGAlgorithm(v interface{}, k string) (ws []string, errors []error) {
	tsigAlgorithm := v.(string)
	if tsigAlgorithm != "HMAC-MD5" && tsigAlgorithm != "HMAC-SHA256" {
		errors = append(errors, fmt.Errorf("%q must be one of HMAC-MD5 or HMAC-SHA256", k))
	}
	return
}

// ValidateUnsignedInteger - Checks the value we are passing is actually >0
func ValidateUnsignedInteger(v interface{}, k string) (ws []string, errors []error) {
	ttl := v.(int)
	if ttl < 0 {
		errors = append(errors, fmt.Errorf("%q can't be negative", k))
	}
	return
}

// CheckLeadingTrailingSpaces - Checks strings for any leading/trailing space
func CheckLeadingTrailingSpaces(v interface{}, k string) (ws []string, errors []error) {
	stringToCheck := v.(string)
	trimedString := strings.Trim(stringToCheck, " ")
	if trimedString != stringToCheck {
		errors = append(errors, fmt.Errorf("%q must not contain trailing or leading white space", k))
	}
	return
}
