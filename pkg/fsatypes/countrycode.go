// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsatypes

import (
	"fmt"
	"github.com/rpatton4/fsa/pkg/fsaerrors"
)

const exactCountryCodeLength = 2

type CountryCode string

// NewCountryCode validates the input string in terms of what COD will accept, which is CountryCode == 2 characters
func NewCountryCode(s string) (CountryCode, *fsaerrors.Error) {
	if len(s) != exactCountryCodeLength {
		return "", &fsaerrors.Error{
			Code:    fsaerrors.PostalAddressCountryCodeInvalid,
			Message: fmt.Sprintf("country code length is invalid, value is '%s', length: %d, length must be: %d", s, len(s), minAddressLineLength),
		}
	}
	return CountryCode(s), nil
}
