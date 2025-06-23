// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

// The isirparser package contains behavior for reading an Institutional Student
// Information Record (ISIR) from the standard fixed format provided by the
// Department of Education and populating an isirecord struct with the data
//
// The fixed format does change over time, and the changes coincide with new
// award years (AY) so this package is broken into functionality by AY.
//
// The bulk of the logic is intended to be entirely internal to this package,
// while the ISIRParser interface and implementations are the external entry
// points available outside of the package.
package isirparser

import (
	"errors"
	"github.com/rpatton4/fsa/pkg/fsaconstants"
	"github.com/rpatton4/fsa/pkg/isirmodels"
	"log/slog"
)

type ISIRParser interface {
	ParseISIR(record string) (isirmodels.ISIRecord, error)
}

// Factory method to create a parser which understands the format for the given award year
func CreateISIRParser(y fsaconstants.AwardYear) (ISIRParser, error) {
	switch y {
	case fsaconstants.AwardYear2526:
		return &ISIRParser2526{}, nil
	default:
		return nil, errors.New("no ISIR Parser available for AY " + y.String())
	}
}

// Reads a line containing an ISIR record and determines which Award Year the ISIR is for
func DetermineAYFromISIRLine(l string) (fsaconstants.AwardYear, error) {
	v, err := getAwardYearValue(l)
	if err != nil {
		return fsaconstants.AwardYearUnknown, err
	}

	switch v {
	case "5":
		return fsaconstants.AwardYear2425, nil
	case "6":
		return fsaconstants.AwardYear2526, nil
	default:
		msg := "unable to determine the Award Year from the ISIR"
		slog.Error(msg, "AY value", v)
		return fsaconstants.AwardYearUnknown, errors.New(msg)
	}
}

// Retrieves the value for the award year from the given ISIR line.  Contains any logic to locate that field
// based on formats for different ISIR years
func getAwardYearValue(l string) (string, error) {
	if len(l) < 1 {
		msg := "unable to determine AY value from an empty record"
		slog.Error(msg)
		return "", errors.New(msg)
	}
	// As of June 2025, the AY field is always the first character of the line
	return string(l[0]), nil
}
