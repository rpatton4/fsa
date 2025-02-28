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
