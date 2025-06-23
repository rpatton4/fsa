// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsatypes

import (
	"fmt"
)

const minAddressLineLength = 1
const maxAddressLineLength = 40

type AddressLine string

// NewAddressLine validates the input string in terms of what COD will accept, which is 1 <= AddressLine <= 40 characters
func NewAddressLine(s string) (AddressLine, error) {
	l := len(s)
	if l < minAddressLineLength {
		return "", fmt.Errorf("address line is too short, value is '%s', length: %d, min length: %d", s, l, minNameLength)
	}
	if l > maxAddressLineLength {
		return "", fmt.Errorf("address line is too long, value is '%s', length: %d, max length: %d", s, l, maxNameLength)
	}
	return AddressLine(s), nil
}
