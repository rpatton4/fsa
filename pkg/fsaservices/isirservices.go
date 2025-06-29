// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsaservices

import (
	"bufio"
	"github.com/rpatton4/fsa/internal/isirparser"
	"github.com/rpatton4/fsa/pkg/fsaconstants"
	"github.com/rpatton4/fsa/pkg/fsaerrors"
	"github.com/rpatton4/fsa/pkg/fsamodels"
	"io"
	"log/slog"
	"strings"
)

func ParseISIRStream(stream io.Reader) ([]fsamodels.ISIRecord, *fsaerrors.Error) {
	slog.Debug("ParseISIRStream(stream) starting")
	records := make([]fsamodels.ISIRecord, 0)
	linesParsed, linesSkipped := 0, 0
	fScanner := bufio.NewScanner(stream)
	var parsers = make(map[fsaconstants.AwardYear]isirparser.ISIRParser)

	for cur := 1; fScanner.Scan(); cur++ {
		line := fScanner.Text()

		// Determine whether the line is empty, meaning invalid, and skip if so
		if strings.TrimSpace(line) == "" {
			slog.Debug("Skipping empty line in ISIR stream")
			continue
		}

		ay, err := isirparser.DetermineAYFromISIRLine(line)
		if err != nil {
			slog.Error("Error determining AY from ISIR line, skipping ISIR line", "errorMessage", err.Error())
			linesSkipped++
			continue
		}
		slog.Debug("Determined AY from ISIR line", "current line number", cur, "ay", ay)

		p, ok := parsers[ay]

		// Reuse parsers we've already created
		if !ok {
			np, err := isirparser.CreateISIRParser(ay)
			if err != nil {
				slog.Error("Error creating parser for ISIR line, skipping line", "errorMessage", err.Error())
				linesSkipped++
				continue
			}
			parsers[ay] = np
			p = np
			slog.Debug("Created new ISIR parser for AY", "ay", ay)
		} else {
			slog.Debug("Reusing existing ISIR parser for AY", "ay", ay)
		}

		rec, err := p.ParseISIR(line)
		if err != nil {
			slog.Error("Error parsing line from ISIR file, stopping stream processing", "errorMessage", err.Error())
			return records, err
		}
		records = append(records, rec)
		linesParsed++
	}
	slog.Debug("ParseISIRStream(stream) finished", "lines_parsed", linesParsed, "lines_skipped", linesSkipped, "records produced", len(records))
	return records, nil
}
