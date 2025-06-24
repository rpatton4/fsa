// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsamodels

import (
	"github.com/rpatton4/fsa/pkg"
	"github.com/rpatton4/fsa/pkg/fsatypes"
)

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
	StateProvinceCode fsatypes.StateProvince
	PostalCode        fsatypes.PostalCode
	CountryCode       fsatypes.CountryCode
}

func NewPostalAddress(line1 string, line2 string, line3 string, city string, stateProvinceCode string, postalCode string, countryCode string) (*PostalAddress, *pkg.FSAError) {
	wrapperError := pkg.FSAError{
		Code:    pkg.PostalAddressValidationFailed,
		Message: "postal address failed validation, see upstream errors",
	}
	l1, err := fsatypes.NewAddressLine(line1)

	if len(wrapperError.UpstreamErrors) > 0 {
		return nil, &wrapperError
	}
}
