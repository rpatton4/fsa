// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsamodels

import "testing"

func Test_PostalAddressValidation(t *testing.T) {
	// Test creating a valid postal address
	validAddress, err := NewPostalAddress("123 Main St", "Apt 4B", "Under the Stairs", "Springfield", "IL", "62701", "US")
	if err != nil {
		t.Errorf("Expected no error for valid postal address, got: %v", err)
	}
	if validAddress.AddressLine1 != "123 Main St" {
		t.Errorf("Expected AddressLine1 to be '123 Main St', got: %s", validAddress.AddressLine1)
	}

	// Test invalid address line 1
	_, err = NewPostalAddress("1234567890123456789012345678901234567890 Main St", "", "", "Springfield", "IL", "12345678901", "US")
	if err == nil || len(err.UpstreamErrors) == 0 {
		t.Error("Expected error for invalid Address Line 1 length, but got none")
	}

	// Test invalid address line 2
	_, err = NewPostalAddress("abc", "1234567890123456789012345678901234567890 Main St", "", "Springfield", "IL", "12345678901", "US")
	if err == nil || len(err.UpstreamErrors) == 0 {
		t.Error("Expected error for invalid Address Line 2 length, but got none")
	}

	// Test invalid address line 3
	_, err = NewPostalAddress("abc", "def", "1234567890123456789012345678901234567890 Main St", "Springfield", "IL", "12345678901", "US")
	if err == nil || len(err.UpstreamErrors) == 0 {
		t.Error("Expected error for invalid Address Line 3 length, but got none")
	}

	// Test for invalid city
	_, err = NewPostalAddress("123 Main St", "", "", "1234567890123456789012345678901234567890", "IL", "12345678901", "US")
	if err == nil || len(err.UpstreamErrors) == 0 {
		t.Error("Expected error for invalid city length, but got none")
	}

	// Test for invalid domestic state code
	_, err = NewPostalAddress("123 Main St", "", "", "Springfield", "I", "12345678901", "US")
	if err == nil || len(err.UpstreamErrors) == 0 {
		t.Error("Expected error for invalid domestic state length, but got none")
	}

	// Test for invalid international state code
	_, err = NewPostalAddress("123 Main St", "", "", "Springfield", "1234567890123456789012345678901234567890", "12345678901", "UK")
	if err == nil || len(err.UpstreamErrors) == 0 {
		t.Error("Expected error for invalid international state length, but got none")
	}

	// Test invalid postal code
	_, err = NewPostalAddress("123 Main St", "", "", "Springfield", "IL", "12345678901234567890123456789012345678901234567890123456789012345678901234567890", "US")
	if err == nil || len(err.UpstreamErrors) == 0 {
		t.Error("Expected error for invalid postal code length, but got none")
	}
}
