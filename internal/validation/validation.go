// Package validation provides functionality for validating various types of
// data, including API keys.
//
// This package includes several validation functions to ensure that input data
// adheres to defined formats and constraints. Specifically, it provides the
// following validation functions:
//
//   - ValidateTelegramToken: Validates the format of a Telegram bot token,
//     ensuring it matches the required pattern of `8-10 digits:followed by a 35-character string`.
//
//   - Validate: A generic function that checks if a provided value exists within a list of
//     valid options, applicable to any comparable type.
//
// The purpose of this package is to enhance the robustness and reliability of
// the application by enforcing input validation rules across various
// components.
package validation

import (
	"fmt"
)

// Validate checks if the provided value is in the list of valid options.
func Validate[T comparable](value T, validOptions []T) error {
	for _, validOption := range validOptions {
		if value == validOption {
			return nil
		}
	}
	return fmt.Errorf("invalid value: %v", value)
}
