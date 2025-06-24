// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsamodels

import (
	"github.com/rpatton4/fsa/pkg"
	"github.com/rpatton4/fsa/pkg/fsatypes"
)

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
	var pa PostalAddress
	wrapperError := pkg.FSAError{
		Code:    pkg.PostalAddressValidationFailed,
		Message: "postal address failed validation, see upstream errors",
	}

	if line1 != "" {
		l1, err := fsatypes.NewAddressLine(line1)
		if err != nil {
			wrapperError.UpstreamErrors = append(wrapperError.UpstreamErrors, err)
		}
		pa.AddressLine1 = l1
	}

	if line2 != "" {
		l2, err := fsatypes.NewAddressLine(line2)
		if err != nil {
			wrapperError.UpstreamErrors = append(wrapperError.UpstreamErrors, err)
		}
		pa.AddressLine2 = l2
	}

	if line3 != "" {
		l3, err := fsatypes.NewAddressLine(line3)
		if err != nil {
			wrapperError.UpstreamErrors = append(wrapperError.UpstreamErrors, err)
		}
		pa.AddressLine2 = l3
	}

	if city != "" {
		c, err := fsatypes.NewCity(city)
		if err != nil {
			wrapperError.UpstreamErrors = append(wrapperError.UpstreamErrors, err)
		}
		pa.City = c
	}

	// Default to US / Domestic
	if countryCode == "" || countryCode == "US" {
		sp, err := fsatypes.NewDomesticStateProvince(stateProvinceCode)
		if err != nil {
			wrapperError.UpstreamErrors = append(wrapperError.UpstreamErrors, err)
		}
		pa.StateProvinceCode = sp
	} else {
		sp, err := fsatypes.NewInternationalStateProvince(stateProvinceCode)
		if err != nil {
			wrapperError.UpstreamErrors = append(wrapperError.UpstreamErrors, err)
		}
		pa.StateProvinceCode = sp
	}

	if postalCode != "" {
		pc, err := fsatypes.NewPostalCode(postalCode)
		if err != nil {
			wrapperError.UpstreamErrors = append(wrapperError.UpstreamErrors, err)
		}
		pa.PostalCode = pc
	}

	if countryCode != "" {
		cc, err := fsatypes.NewCountryCode(countryCode)
		if err != nil {
			wrapperError.UpstreamErrors = append(wrapperError.UpstreamErrors, err)
		}
		pa.CountryCode = cc
	}

	if len(wrapperError.UpstreamErrors) > 0 {
		return nil, &wrapperError
	}
	return &pa, nil
}
