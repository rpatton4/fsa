// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsatypes

import (
	"fmt"
)

const minStateIDLength = 0
const maxStateIDLength = 20

type StateID string

// NewStateID validates the input string in terms of what COD will accept, which is 0 <= NewStateID <= 20 characters
func NewStateID(s string) (StateID, error) {
	l := len(s)
	if l < minStateIDLength {
		return "", fmt.Errorf("state ID is too short, value is '%s', length: %d, min length: %d", s, l, minStateIDLength)
	}
	if l > maxStateIDLength {
		return "", fmt.Errorf("state ID is too long, value is '%s', length: %d, max length: %d", s, l, maxStateIDLength)
	}
	return StateID(s), nil
}
