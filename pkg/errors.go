// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package pkg

// ErrorCode An error code used or emitted by the FSA module
type ErrorCode int

// Enumeration of all possible error codes used within or emitted by the FSA Module.
// The integer values may change, use the constant names for safety
const (
	PostalAddressLineInvalid = iota + 1
	PostalAddressCityLengthInvalid
	PostalAddressPostalCodeInvalid
)

// FSAError The implementation of an error, specific to the FSA module.  This simply extends the concept of
// errors.Error with a code for the purposes of comparisons and internationalization of messages.
// In many cases the code will be slightly less detailed than the included message, for example
// the message may indicate that a field value is too long while the code only indicates that the
// length is invalid
type FSAError struct {
	Code    ErrorCode
	Message string
}
