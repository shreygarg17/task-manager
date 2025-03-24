package validator

import (
	"errors"
	"strings"
)

var (
	ErrEmptyTitle       = errors.New("task title cannot be empty")
	ErrTitleTooLong     = errors.New("task title is too long")
	ErrDescriptionTooLong = errors.New("task description is too long")
)

const (
	MaxTitleLength       = 100
	MaxDescriptionLength = 500
)

// ValidateTitle checks that title is not empty and within max length
func ValidateTitle(title string) error {
	title = strings.TrimSpace(title)

	if title == "" {
		return ErrEmptyTitle
	}
	if len(title) > MaxTitleLength {
		return ErrTitleTooLong
	}
	return nil
}

// ValidateDescription checks that description is within allowed length
func ValidateDescription(desc string) error {
	desc = strings.TrimSpace(desc)

	if len(desc) > MaxDescriptionLength {
		return ErrDescriptionTooLong
	}
	return nil
}

// ValidateTaskFields performs multiple validations
func ValidateTaskFields(title, desc string) error {
	if err := ValidateTitle(title); err != nil {
		return err
	}
	if err := ValidateDescription(desc); err != nil {
		return err
	}
	return nil
}
