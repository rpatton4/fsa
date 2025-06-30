// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/rpatton4/fsa/pkg/fsaservices"
	"log/slog"
	"os"
	"strings"
)

var cid uuid.UUID

// This is the main entry point for the FSA module.
// The module is designed to be used as a library, so this main function
// is primarily for testing and demonstration purposes.
// In a real application, you would import the FSA module and use its
func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 3 || args[0] == "--help" {
		displayHelp()
		return
	}

	c, err := uuid.NewV7()
	if err != nil {
		slog.Error("Error generating UUID ", "errorMessage", err.Error(), "func", "main()")
		return
	}
	cid = c
	executeCommand(args)
}

func displayHelp() {
	fmt.Println("Usage: fsa [command] <file path> <--verbose>")
	fmt.Println("Commands:")
	fmt.Println("    --help\t\tDisplay this help message")
	fmt.Println("    --read-isirs\tRead ISIRs from a file and print them to the console in JSON format")
	fmt.Println("Verbose Mode:")
	fmt.Println("    --verbose\t\tSet logging to debug level. This must be the last argument.")
}

func executeCommand(args []string) {
	cmd := args[0]
	switch cmd {
	case "--read-isirs":
		readIsirs(args[1:])
		return
	}
}

// This function would read ISIRs from the specified file
// and print them to the console in JSON format.
func readIsirs(args []string) {
	if len(args) == 2 {
		if args[1] != "--verbose" {
			fmt.Println("Error: --verbose must be the last argument")
			return
		}
		setVerboseLogging()
	}

	if len(args) < 1 {
		fmt.Println("Error: No file path provided for reading ISIRs")
		return
	}

	f := args[0]

	slog.Debug("Reading ISIRs from file", "filePath", f)

	c, err := os.ReadFile(f)
	if err != nil {
		slog.Error("Error reading file: ", "errorMessage", err)
		return
	}

	records, ferr := fsaservices.ParseISIRStreamWithCorrelationId(strings.NewReader(string(c)), cid)
	if ferr != nil {
		slog.Error("error parsing ISIR stream: ", "errorMessage", ferr)
	}
	slog.Info("Parsed ISIR records", "count", len(records))

	for _, r := range records {
		fmt.Print(r.JsonString(cid))
	}
}

func setVerboseLogging() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}
