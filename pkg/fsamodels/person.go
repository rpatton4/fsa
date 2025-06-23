// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsamodels

import (
	"github.com/rpatton4/fsa/pkg/fsatypes"
	"time"
)

// Person A representation of a person (versus student, parent etc.) with properties commonly
// used by FSA.  This was primarily influenced by the CommonRecord's XML schema definition of
// person as defined in CommonRecord5pt0c.xsd
type Person struct {
	// Fields are listed in the order in which they appear in the CommonRecord XML schema.
	// Index: SSN, BirthDate, and LastName are part of the CommonRecord's Index complex element

	// The SSN must match the following format: ^\\d{7}$|^\\d{8}$|^\\d{9}$|^\\d{3}[-]\\d{2}[-]\\d{4}$ which
	// works out to meaning it can be either 7, 8, or 9 digits, or the common current format of 123-45-6789
	SSN fsatypes.SSN
	// The time portion of the birthdate is ignored
	DateOfBirth time.Time
	LastName    fsatypes.Name

	// PersonIdentifiers: The following 2 identifiers (DL # and School ID) are used with the previously
	// declared SSN to construct the Person's PersonIdentifier complex type

	DriversLicenseState string
	// Free text identifier assigned by the state which is present on a driver's license
	DriversLicenseId fsatypes.StateID
	// CommonRecord calls this the SchoolAssignedPersonId, but it would more typically be known as the
	// Student ID.
	SchoolAssignedPersonId fsatypes.StudentID

	// Name: The following fields are used to construct the Person's Name complex type along with the
	// previously declared LastName

	FirstName  fsatypes.Name
	MiddleName fsatypes.Name
	NameSuffix fsatypes.Suffix

	// One contact record is allowed. While CommonRecord allows multiple contacts, the majority of use
	// cases in a Financial Aid context can be satisfied with a single contact record due to having fields
	// for two addresses and multiple phone numbers and email addresses.
	Contact ContactInfo
}
