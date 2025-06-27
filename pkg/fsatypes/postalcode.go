// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsatypes

import (
	"fmt"
	"github.com/rpatton4/fsa/pkg/fsaerrors"
)

const minPostalCodeLength = 1
const maxPostalCodeLength = 40

type PostalCode string

// NewPostalCode validates the input string in terms of what COD will accept, which is 1 <= PostalCode <= 10 characters
func NewPostalCode(s string) (PostalCode, *fsaerrors.Error) {
	l := len(s)
	if l < minAddressLineLength {
		return "", &fsaerrors.Error{
			Code:    fsaerrors.PostalAddressPostalCodeInvalid,
			Message: fmt.Sprintf("postal code is too short, value is '%s', length: %d, min length: %d", s, l, minPostalCodeLength),
		}
	}
	if l > maxAddressLineLength {
		return "", &fsaerrors.Error{
			Code:    fsaerrors.PostalAddressPostalCodeInvalid,
			Message: fmt.Sprintf("postal code is too long, value is '%s', length: %d, max length: %d", s, l, maxPostalCodeLength),
		}
	}
	return PostalCode(s), nil
}
