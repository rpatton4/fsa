// Robert Patton ("DEVELOPER") CONFIDENTIAL
// Unpublished Copyright (c) 2025 Robert Patton, All Rights Reserved.
//
// NOTICE: All information contained herein is, and remains the property of DEVELOPER. The
// intellectual and technical concepts contained herein are proprietary to DEVELOPER and may be
// covered by U.S. and Foreign Patents, patents in process, and are protected by trade secret or
// copyright law. Dissemination of this information or reproduction of this material is strictly
// forbidden unless prior written permission is obtained from DEVELOPER. Access to the source code
// contained herein is hereby forbidden to anyone except current DEVELOPER employees, managers or
// contractors who have executed Confidentiality and Non-disclosure agreements explicitly covering
// such access.
//
// The copyright notice above does not evidence any actual or intended publication or disclosure of
// this source code, which includes information that is confidential and/or proprietary, and is a
// trade secret, of DEVELOPER. ANY REPRODUCTION, MODIFICATION, DISTRIBUTION, PUBLIC  PERFORMANCE,
// OR PUBLIC DISPLAY OF OR THROUGH USE OF THIS SOURCE CODE WITHOUT THE EXPRESS WRITTEN CONSENT OF
// COMPANY IS STRICTLY PROHIBITED, AND IN VIOLATION OF APPLICABLE LAWS AND INTERNATIONAL TREATIES.
// THE RECEIPT OR POSSESSION OF THIS SOURCE CODE AND/OR RELATED INFORMATION DOES NOT CONVEY OR
// IMPLY ANY RIGHTS TO REPRODUCE, DISCLOSE OR DISTRIBUTE ITS CONTENTS, OR TO MANUFACTURE, USE, OR
// SELL ANYTHING THAT IT MAY DESCRIBE, IN WHOLE OR IN PART.

package fsamodels

import (
	"FSAModule/pkg"
	"fmt"
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
	AddressLine1      string
	AddressLine2      string
	AddressLine3      string
	City              string
	StateProvinceCode string
	PostalCode        string
	CountryCode       string
}

// Validate is a method to check whether the required fields of an address are valid in terms of
// the format (length etc.) expected by COD
func (p PostalAddress) Validate() (bool, []pkg.FSAError) {
	var errors []pkg.FSAError = make([]pkg.FSAError, 1)
	var valid bool = true

	if len(p.AddressLine1) < MinAddressLineLength {
		valid = false
		errors = append(errors, pkg.FSAError{
			Code:    pkg.PostalAddressLineInvalid,
			Message: fmt.Sprintf("AddressLine1 is too short, value is '%s', length: %d, min length: %d", p.AddressLine1, len(p.AddressLine1), MinAddressLineLength),
		})
	}
	if len(p.AddressLine1) > MaxAddressLineLength {
		valid = false
		errors = append(errors, pkg.FSAError{
			Code:    pkg.PostalAddressLineInvalid,
			Message: fmt.Sprintf("AddressLine1 is too long, value is '%s', length: %d, max length: %d", p.AddressLine1, len(p.AddressLine1), MaxAddressLineLength),
		})
	}

	if p.AddressLine2 != "" && len(p.AddressLine2) > MaxAddressLineLength {
		valid = false
		errors = append(errors, pkg.FSAError{
			Code:    pkg.PostalAddressLineInvalid,
			Message: fmt.Sprintf("AddressLine2 is too long, value is '%s', length: %d, max length: %d", p.AddressLine2, len(p.AddressLine2), MaxAddressLineLength),
		})
	}

	if p.AddressLine3 != "" && len(p.AddressLine3) > MaxAddressLineLength {
		valid = false
		errors = append(errors, pkg.FSAError{
			Code:    pkg.PostalAddressLineInvalid,
			Message: fmt.Sprintf("AddressLine3 is too long, value is '%s', length: %d, max length: %d", p.AddressLine3, len(p.AddressLine3), MaxAddressLineLength),
		})
	}

	if len(p.City) < MinCityLength {
		valid = false
		errors = append(errors, pkg.FSAError{
			Code:    pkg.PostalAddressCityLengthInvalid,
			Message: fmt.Sprintf("City is too short, value is '%s', length: %d, min length: %d", p.City, len(p.City), MinCityLength),
		})
	}
	if len(p.City) > MaxCityLength {
		valid = false
		errors = append(errors, pkg.FSAError{
			Code:    pkg.PostalAddressCityLengthInvalid,
			Message: fmt.Sprintf("City is too long, value is '%s', length: %d, mac length: %d", p.City, len(p.City), MaxCityLength),
		})
	}

	if len(p.PostalCode) < MinPostalCodeLength {
		valid = false
		errors = append(errors, pkg.FSAError{
			Code:    pkg.PostalAddressPostalCodeInvalid,
			Message: fmt.Sprintf("PostalCode is too short, value is '%s', length: %d, min length: %d", p.PostalCode, len(p.PostalCode), MinPostalCodeLength),
		})
	}
	if len(p.PostalCode) > MaxPostalCodeLength {
		valid = false
		errors = append(errors, pkg.FSAError{
			Code:    pkg.PostalAddressPostalCodeInvalid,
			Message: fmt.Sprintf("PostalCode is too long, value is '%s', length: %d, max length: %d", p.PostalCode, len(p.PostalCode), MaxPostalCodeLength),
		})
	}

	return valid, errors
}
