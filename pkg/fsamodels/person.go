// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsamodels

import (
	"time"
)

// Person A representation of a person (versus student, parent etc.) with properties commonly
// used by FSA.  This was primarily influenced by the CommonRecord's XML schema definition of
// person.
type Person struct {
	FirstName   string
	MiddleName  string
	LastName    string
	NameSuffix  string
	DateOfBirth time.Time
	SSN         string
	ITIN        string
}
