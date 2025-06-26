// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsaservices

import "fmt"

// ErrorCode An error code used or emitted by the FSA module
type ErrorCode int

// Enumeration of all possible error codes used within or emitted by the FSA Module.
// The integer values may change, use the constant names for safety
const (
	PostalAddressValidationFailed = iota + 1
	PostalAddressLineInvalid
	PostalAddressCityLengthInvalid
	PostalAddressStateProvinceLengthInvalid
	PostalAddressPostalCodeInvalid
	PostalAddressCountryCodeInvalid
	ISIRParseError
	LibraryConfigurationErrorISIRAYUnrecognized
	AYDeterminationErrorEmptyISIRInputLine
)

// FSAError The implementation of an error, specific to the FSA module.  This simply extends the concept of
// errors.Error with a code for the purposes of comparisons and internationalization of messages.
// In many cases the code will be slightly less detailed than the included message, for example
// the message may indicate that a field value is too long while the code only indicates that the
// length is invalid
type FSAError struct {
	Code           ErrorCode
	Message        string
	UpstreamErrors []*FSAError
}

func (e *FSAError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, upstream errors: %d", e.Code, e.Message, len(e.UpstreamErrors))
}
