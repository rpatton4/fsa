// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsaservices

import (
	"bufio"
	"github.com/google/uuid"
	"github.com/rpatton4/fsa/internal/isirparser"
	"github.com/rpatton4/fsa/pkg/fsaconstants"
	"github.com/rpatton4/fsa/pkg/fsaerrors"
	"github.com/rpatton4/fsa/pkg/fsamodels"
	"io"
	"log/slog"
	"strings"
)

func ParseISIRStream(stream io.Reader) ([]fsamodels.ISIRecord, *fsaerrors.Error) {
	cid, err := uuid.NewV7()
	if err != nil {
		slog.Error("Error generating UUID for ISIR stream parsing", "errorMessage", err.Error(), "func", "ParseISIRStreamWithCorrelationId()")
		return nil, &fsaerrors.Error{
			Code:    fsaerrors.LibrarySystemErrorCorrelationIDGeneration,
			Message: "ParseISIRStream() Failed to generate UUID for ISIR stream parsing",
		}
	}
	slog.Debug("Generated correlation ID for ISIR stream parsing", "correlationId", cid.String(), "func", "ParseISIRStream()")
	return ParseISIRStreamWithCorrelationId(stream, cid)
}

func ParseISIRStreamWithCorrelationId(stream io.Reader, cid uuid.UUID) ([]fsamodels.ISIRecord, *fsaerrors.Error) {
	slog.Info("Parsing ISIR stream", "correlationId", cid.String(), "func", "ParseISIRStreamWithCorrelationId()")
	records := make([]fsamodels.ISIRecord, 0)
	linesParsed, linesSkipped := 0, 0
	fScanner := bufio.NewScanner(stream)
	var parsers = make(map[fsaconstants.AwardYear]isirparser.ISIRParser)

	for cur := 1; fScanner.Scan(); cur++ {
		line := fScanner.Text()

		// Determine whether the line is empty, meaning invalid, and skip if so
		if strings.TrimSpace(line) == "" {
			slog.Debug("Skipping empty line in ISIR stream", "correlationId", cid.String(), "func", "ParseISIRStreamWithCorrelationId()")
			continue
		}

		ay, err := isirparser.DetermineAYFromISIRLine(line, cid)
		if err != nil {
			slog.Error("Error determining AY from ISIR line, skipping ISIR line", "errorMessage", err.Error(), "correlationId", cid.String(), "func", "ParseISIRStreamWithCorrelationId()")
			linesSkipped++
			continue
		}
		slog.Debug("Determined AY from ISIR line", "current line number", cur, "ay", ay)

		p, ok := parsers[ay]

		// Reuse parsers we've already created
		if !ok {
			np, err := isirparser.CreateISIRParser(ay, cid)
			if err != nil {
				slog.Error("Error creating parser for ISIR line, skipping line", "errorMessage", err.Error(), "correlationId", cid.String(), "func", "ParseISIRStreamWithCorrelationId()")
				linesSkipped++
				continue
			}
			parsers[ay] = np
			p = np
			slog.Debug("Created new ISIR parser for AY", "ay", ay)
		} else {
			slog.Debug("Reusing existing ISIR parser for AY", "ay", ay)
		}

		rec, err := p.ParseISIR(line, cid)
		if err != nil {
			slog.Error("Error parsing line from ISIR file, stopping stream processing", "errorMessage", err.Error())
			return records, err
		}
		records = append(records, rec)
		linesParsed++
	}
	slog.Info("Parsed ISIR stream", "lines_parsed", linesParsed, "lines_skipped", linesSkipped)
	return records, nil
}
