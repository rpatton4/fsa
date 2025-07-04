// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsaservices

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"github.com/rpatton4/fsa/internal/isirparser"
	"github.com/rpatton4/fsa/pkg/fsaconstants"
	"github.com/rpatton4/fsa/pkg/fsaerrors"
	"github.com/rpatton4/fsa/pkg/fsamodels"
	"io"
	"strings"
)

func ParseISIRStream(stream io.Reader) ([]fsamodels.ISIRecord, *fsaerrors.Error) {
	cid, err := uuid.NewV7()
	if err != nil {
		return nil, &fsaerrors.Error{
			Code:    fsaerrors.LibrarySystemErrorCorrelationIDGeneration,
			Message: fmt.Sprintf("failed to generate UUID for ISIR stream parsing %s", err.Error()),
			Func:    "ParseISIRStream()",
		}
	}
	return ParseISIRStreamWithCorrelationId(stream, cid)
}

func ParseISIRStreamWithCorrelationId(stream io.Reader, cid uuid.UUID) ([]fsamodels.ISIRecord, *fsaerrors.Error) {
	records := make([]fsamodels.ISIRecord, 0)
	linesParsed, linesSkipped := 0, 0
	fScanner := bufio.NewScanner(stream)
	var parsers = make(map[fsaconstants.AwardYear]isirparser.ISIRParser)

	for cur := 1; fScanner.Scan(); cur++ {
		line := fScanner.Text()

		// Determine whether the line is empty, meaning invalid, and skip if so
		if strings.TrimSpace(line) == "" {
			continue
		}

		ay, err := isirparser.DetermineAYFromISIRLine(line, cid)
		if err != nil {
			linesSkipped++
			continue
		}

		p, ok := parsers[ay]

		// Reuse parsers we've already created
		if !ok {
			np, err := isirparser.CreateISIRParser(ay, cid)
			if err != nil {
				linesSkipped++
				continue
			}
			parsers[ay] = np
			p = np
		}

		rec, err := p.ParseISIR(line, cid)
		if err != nil {
			return records, err
		}
		records = append(records, rec)
		linesParsed++
	}
	return records, nil
}
