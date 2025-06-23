// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsatypes

import (
	"fmt"
)

const minNameLength = 0
const maxNameLength = 35
const minSuffixLength = 1
const maxSuffixLength = 10

type Name string

// NewName validates the input string in terms of what COD will accept, which is 0 <= Name <= 35 characters
func NewName(s string) (Name, error) {
	l := len(s)
	if l < minNameLength {
		return "", fmt.Errorf("name is too short, value is '%s', length: %d, min length: %d", s, l, minNameLength)
	}
	if l > maxNameLength {
		return "", fmt.Errorf("name is too long, value is '%s', length: %d, max length: %d", s, l, maxNameLength)
	}
	return Name(s), nil
}

type Suffix string

// NewSuffix validates the input string in terms of what COD will accept, which is 1 <= Suffix <= 10 characters
func NewSuffix(s string) (Suffix, error) {
	l := len(s)
	if l < minSuffixLength {
		return "", fmt.Errorf("suffix is too short, value is '%s', length: %d, min length: %d", s, l, minNameLength)
	}
	if l > maxSuffixLength {
		return "", fmt.Errorf("suffix is too long, value is '%s', length: %d, max length: %d", s, l, maxNameLength)
	}
	return Suffix(s), nil
}
