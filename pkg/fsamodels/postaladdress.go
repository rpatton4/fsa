// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsamodels

import (
	"fmt"
	"github.com/rpatton4/fsa/pkg"
	"github.com/rpatton4/fsa/pkg/fsatypes"
)

// MinAddressLineLength is the minimum length in characters for any address line as mandated for CommonRecord
const MinAddressLineLength = 1

// MaxAddressLineLength is the maximum length in characters for any address line as mandated for CommonRecord
const MaxAddressLineLength = 40

// MinCityLength is the minimum length in characters for any city name as mandated for CommonRecord
const MinCityLength = 2

// MaxCityLength is the maximum length in characters for any city name as mandated for CommonRecord
const MaxCityLength = 30

// MinPostalCodeLength is the minimum length in characters for any postal code as mandated for CommonRecord
const MinPostalCodeLength = 1

// MaxPostalCodeLength is the maximum length in characters for any postal code as mandated for CommonRecord
const MaxPostalCodeLength = 10

// PostalAddress A representation of a mailing address. This is a simplified version of the CommonRecord XML
// schema structures for address, combining US domestic and international address info
type PostalAddress struct {
	AddressLine1      fsatypes.AddressLine
	AddressLine2      fsatypes.AddressLine
	AddressLine3      fsatypes.AddressLine
	City              fsatypes.City
	StateProvinceCode string
	PostalCode        string
	CountryCode       string
}

// Validate is a method to check whether the required fields of an address are valid in terms of
// the format (length etc.) expected by COD
func (p PostalAddress) Validate() (bool, []pkg.FSAError) {
	var errors = make([]pkg.FSAError, 1)
	var valid = true

	valid, errors = validateAddressLineLength(p.AddressLine1, valid, errors)

	if p.AddressLine2 != "" {
		valid, errors = validateAddressLineLength(p.AddressLine2, valid, errors)
	}

	if p.AddressLine3 != "" {
		valid, errors = validateAddressLineLength(p.AddressLine3, valid, errors)
	}

	valid, errors = validateCityLength(p.City, valid, errors)

	valid, errors = validatePostalCodeLength(p.PostalCode, valid, errors)

	return valid, errors
}

func validatePostalCodeLength(pc string, valid bool, errors []pkg.FSAError) (bool, []pkg.FSAError) {
	if len(pc) < MinPostalCodeLength {
		valid = false
		errors = append(errors, pkg.FSAError{
			Code:    pkg.PostalAddressPostalCodeInvalid,
			Message: fmt.Sprintf("PostalCode is too short, value is '%s', length: %d, min length: %d", pc, len(pc), MinPostalCodeLength),
		})
	} else if len(pc) > MaxPostalCodeLength {
		valid = false
		errors = append(errors, pkg.FSAError{
			Code:    pkg.PostalAddressPostalCodeInvalid,
			Message: fmt.Sprintf("PostalCode is too long, value is '%s', length: %d, max length: %d", pc, len(pc), MaxPostalCodeLength),
		})
	}
	return valid, errors
}

func validateCityLength(c string, valid bool, errors []pkg.FSAError) (bool, []pkg.FSAError) {
	if len(c) < MinCityLength {
		valid = false
		errors = append(errors, pkg.FSAError{
			Code:    pkg.PostalAddressCityLengthInvalid,
			Message: fmt.Sprintf("City is too short, value is '%s', length: %d, min length: %d", c, len(c), MinCityLength),
		})
	} else if len(c) > MaxCityLength {
		valid = false
		errors = append(errors, pkg.FSAError{
			Code:    pkg.PostalAddressCityLengthInvalid,
			Message: fmt.Sprintf("City is too long, value is '%s', length: %d, mac length: %d", c, len(c), MaxCityLength),
		})
	}

	return valid, errors
}

func validateAddressLineLength(al string, valid bool, errors []pkg.FSAError) (bool, []pkg.FSAError) {
	if len(al) < MinAddressLineLength {
		valid = false
		errors = append(errors, pkg.FSAError{
			Code:    pkg.PostalAddressLineInvalid,
			Message: fmt.Sprintf("AddressLine is too short, value is '%s', length: %d, min length: %d", al, len(al), MinAddressLineLength),
		})
	} else if len(al) > MaxAddressLineLength {
		valid = false
		errors = append(errors, pkg.FSAError{
			Code:    pkg.PostalAddressLineInvalid,
			Message: fmt.Sprintf("AddressLine is too long, value is '%s', length: %d, max length: %d", al, len(al), MaxAddressLineLength),
		})
	}
	return valid, errors
}
