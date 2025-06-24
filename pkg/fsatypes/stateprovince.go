// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsatypes

import (
	"fmt"
)

const exactDomesticStateCodeLength = 2
const minInternationalStateProvinceLength = 1
const maxInternationalStateProvinceLength = 30

type StateProvince string

// NewDomesticStateProvince validates the input string in terms of what COD will accept, which is StateProvince == 2 characters
func NewDomesticStateProvince(s string) (StateProvince, error) {
	if len(s) != exactDomesticStateCodeLength {
		return "", fmt.Errorf("state province length is invalid, value is '%s', length: %d, length must be: %d", s, len(s), minAddressLineLength)
	}
	return StateProvince(s), nil
}

// NewInternationalStateProvince validates the input string in terms of what COD will accept, which is 1 <= StateProvince <= 40 characters
func NewInternationalStateProvince(s string) (StateProvince, error) {
	l := len(s)
	if l < minInternationalStateProvinceLength {
		return "", fmt.Errorf("state province is too short, value is '%s', length: %d, min length: %d", s, l, minInternationalStateProvinceLength)
	}
	if l > maxInternationalStateProvinceLength {
		return "", fmt.Errorf("state province is too long, value is '%s', length: %d, max length: %d", s, l, maxInternationalStateProvinceLength)
	}
	return StateProvince(s), nil
}
