// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

// Package fsatypes defines custom types and related builders for fields. These types
// are often used in the fsamodels package to represent data in a structured way that adheres to specific
// validation rules according to specifications from COD (Common Origination and Disbursement) and other
// FSA (Federal Student Aid) systems.
// All types will have a constructor function which validates the input string and returns the type if valid,
// these should always be used instead of directly instantiating the type to ensure validation is performed.
package fsatypes

import (
	"errors"
	"regexp"
)

var rs string = "^\\d{7}$|^\\d{8}$|^\\d{9}$|^\\d{3}[-]\\d{2}[-]\\d{4}$"
var re *regexp.Regexp

type SSN string

// NewSSN validates the input string in terms of what COD will accept.
// The SSN must match the following format: ^\\d{7}$|^\\d{8}$|^\\d{9}$|^\\d{3}[-]\\d{2}[-]\\d{4}$ which
// works out to meaning it can be either 7, 8, or 9 digits, or the common current format of 123-45-6789
func NewSSN(s string) (SSN, error) {
	if re == nil {
		re = regexp.MustCompile(rs)
	}

	if !re.MatchString(s) {
		return "", errors.New("invalid SSN")
	}
	return SSN(s), nil
}
