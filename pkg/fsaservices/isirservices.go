package fsaservices

import (
	"bufio"
	"github.com/rpatton4/fsa/internal/isirlayout"
	"github.com/rpatton4/fsa/pkg/isirmodels"
	"io"
	"log/slog"
	"strings"
)

// TODO Look for award year info (from file names for ex.) to determine which internal implementation of services to use

func ParseISIRStream(stream io.Reader) ([]isirmodels.ISIRecord, error) {
	slog.Info("Parsing ISIR stream")
	records := make([]isirmodels.ISIRecord, 0)
	fScanner := bufio.NewScanner(stream)

	for fScanner.Scan() {
		line := fScanner.Text()

		// Determine whether the line is empty, meaning invalid, and skip if so
		if strings.TrimSpace(line) == "" {
			slog.Debug("Skipping empty line in ISIR stream")
			continue
		}
		rec, err := isirlayout.ParseISIR(line)
		if err != nil {
			slog.Error("Error parsing line from ISIR file", "error", err.Error())
			return records, err
		}
		records = append(records, rec)

	}
	return records, nil
}
