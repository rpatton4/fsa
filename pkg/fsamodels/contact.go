// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsamodels

// ContactInfo A representation of the contact info for a person with properties commonly
// used by FSA.  This was primarily influenced by the CommonRecord's XML schema definition of
// contact, though simplified for typical use cases.
type ContactInfo struct {
	PermanentAddress PostalAddress
	PhoneNumber      string
	EmailAddress     string
}
