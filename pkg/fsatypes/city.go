// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsatypes

import (
	"fmt"
)

const minCityLength = 2
const maxCityLength = 30

type City string

// NewCity validates the input string in terms of what COD will accept, which is 2 <= City <= 30 characters
func NewCity(s string) (City, error) {
	l := len(s)
	if l < minCityLength {
		return "", fmt.Errorf("city is too short, value is '%s', length: %d, min length: %d", s, l, minCityLength)
	}
	if l > maxCityLength {
		return "", fmt.Errorf("city is too long, value is '%s', length: %d, max length: %d", s, l, maxCityLength)
	}
	return City(s), nil
}
