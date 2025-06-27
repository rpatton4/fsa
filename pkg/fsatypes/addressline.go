// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsatypes

import (
	"fmt"
	"github.com/rpatton4/fsa/pkg/fsaerrors"
)

const minAddressLineLength = 1
const maxAddressLineLength = 40

type AddressLine string

// NewAddressLine validates the input string in terms of what COD will accept, which is 1 <= AddressLine <= 40 characters
func NewAddressLine(s string) (AddressLine, *fsaerrors.Error) {
	l := len(s)
	if l < minAddressLineLength {
		return "", &fsaerrors.Error{
			Code:    fsaerrors.PostalAddressLineInvalid,
			Message: fmt.Sprintf("address line is too short, value is '%s', length: %d, min length: %d", s, l, minAddressLineLength),
		}
	}
	if l > maxAddressLineLength {
		return "", &fsaerrors.Error{
			Code:    fsaerrors.PostalAddressLineInvalid,
			Message: fmt.Sprintf("address line is too long, value is '%s', length: %d, max length: %d", s, l, maxAddressLineLength),
		}
	}
	return AddressLine(s), nil
}
