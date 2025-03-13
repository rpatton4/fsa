package fsaservices

import (
	"FSAModule/internal/isirconv"
	"FSAModule/pkg/isirmodels"
	"bufio"
	"fmt"
	"io"
	"strings"
)

// TODO Look for award year info (from file names for ex.) to determine which internal implementation of services to use

func ParseISIRStream(stream io.Reader) ([]isirmodels.ISIRecord, error) {
	records := make([]isirmodels.ISIRecord, 0)
	fScanner := bufio.NewScanner(stream)

	for fScanner.Scan() {
		line := fScanner.Text()

		// Determine whether the line is empty, meaning invalid, and skip if so
		if strings.TrimSpace(line) == "" {
			continue
		}
		rec, err := isirconv.ParseISIR(line)
		if err != nil {
			fmt.Println("Error parsing line from ISIR file", err)
			return records, err
		}
		records = append(records, rec)

	}
	return records, nil
}
