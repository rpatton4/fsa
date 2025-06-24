// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsatypes

import (
	"fmt"
)

const minStudentIdLength = 0
const maxStudentIdLength = 20

type StudentID string

// NewStudentID validates the input string in terms of what COD will accept, which is 0 <= StudentID <= 20 characters
func NewStudentID(s string) (StudentID, error) {
	l := len(s)
	if l < minStudentIdLength {
		return "", fmt.Errorf("student ID is too short, value is '%s', length: %d, min length: %d", s, l, minStudentIdLength)
	}
	if l > maxStudentIdLength {
		return "", fmt.Errorf("student ID is too long, value is '%s', length: %d, max length: %d", s, l, maxStudentIdLength)
	}
	return StudentID(s), nil
}
