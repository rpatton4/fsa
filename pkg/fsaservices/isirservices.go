package fsaservices

import (
	"bufio"
	"github.com/rpatton4/fsa/internal/isirparser"
	"github.com/rpatton4/fsa/pkg/isirmodels"
	"io"
	"log/slog"
	"strings"
)

func ParseISIRStream(stream io.Reader) ([]isirmodels.ISIRecord, error) {
	slog.Info("Parsing ISIR stream")
	records := make([]isirmodels.ISIRecord, 0)
	linesParsed, linesSkipped := 0, 0
	fScanner := bufio.NewScanner(stream)

	for fScanner.Scan() {
		line := fScanner.Text()

		// Determine whether the line is empty, meaning invalid, and skip if so
		if strings.TrimSpace(line) == "" {
			slog.Debug("Skipping empty line in ISIR stream")
			continue
		}

		ay, err := isirparser.DetermineAYFromISIRLine(line)
		if err != nil {
			slog.Error("Error determining AY from ISIR line, skipping ISIR line", "error", err.Error())
			linesSkipped++
			continue
		}

		p, err := isirparser.CreateISIRParser(ay)
		if err != nil {
			slog.Error("Error creating parser for ISIR line, skipping line", "error", err.Error())
			linesSkipped++
			continue
		}

		rec, err := p.ParseISIR(line)
		if err != nil {
			slog.Error("Error parsing line from ISIR file, stopping stream processing", "error", err.Error())
			return records, err
		}
		records = append(records, rec)
		linesParsed++
	}
	slog.Info("Parsed ISIR stream", "lines_parsed", linesParsed, "lines_skipped", linesSkipped)
	return records, nil
}
