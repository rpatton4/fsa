// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsatypes

import (
	"fmt"
	"github.com/rpatton4/fsa/pkg/fsaerrors"
)

const exactDomesticStateCodeLength = 2
const minInternationalStateProvinceLength = 1
const maxInternationalStateProvinceLength = 30

type StateProvince string

// NewDomesticStateProvince validates the input string in terms of what COD will accept, which is StateProvince == 2 characters
func NewDomesticStateProvince(s string) (StateProvince, *fsaerrors.Error) {
	if len(s) != exactDomesticStateCodeLength {
		return "", &fsaerrors.Error{
			Code:    fsaerrors.PostalAddressStateProvinceLengthInvalid,
			Message: fmt.Sprintf("state province length is invalid, value is '%s', length: %d, length must be: %d", s, len(s), exactDomesticStateCodeLength),
		}
	}
	return StateProvince(s), nil
}

// NewInternationalStateProvince validates the input string in terms of what COD will accept, which is 1 <= StateProvince <= 40 characters
func NewInternationalStateProvince(s string) (StateProvince, *fsaerrors.Error) {
	l := len(s)
	if l < minInternationalStateProvinceLength {
		return "", &fsaerrors.Error{
			Code:    fsaerrors.PostalAddressStateProvinceLengthInvalid,
			Message: fmt.Sprintf("state province is too short, value is '%s', length: %d, min length: %d", s, l, minInternationalStateProvinceLength),
		}
	}
	if l > maxInternationalStateProvinceLength {
		return "", &fsaerrors.Error{
			Code:    fsaerrors.PostalAddressStateProvinceLengthInvalid,
			Message: fmt.Sprintf("state province is too long, value is '%s', length: %d, min length: %d", s, l, maxInternationalStateProvinceLength),
		}
	}
	return StateProvince(s), nil
}
