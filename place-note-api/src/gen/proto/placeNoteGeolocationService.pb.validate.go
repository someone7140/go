// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/placeNoteGeolocationService.proto

package placeNote

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetLatLonFromAddressRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetLatLonFromAddressRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLatLonFromAddressRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetLatLonFromAddressRequestMultiError, or nil if none found.
func (m *GetLatLonFromAddressRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLatLonFromAddressRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetAddress()) < 1 {
		err := GetLatLonFromAddressRequestValidationError{
			field:  "Address",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetLatLonFromAddressRequestMultiError(errors)
	}

	return nil
}

// GetLatLonFromAddressRequestMultiError is an error wrapping multiple
// validation errors returned by GetLatLonFromAddressRequest.ValidateAll() if
// the designated constraints aren't met.
type GetLatLonFromAddressRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLatLonFromAddressRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLatLonFromAddressRequestMultiError) AllErrors() []error { return m }

// GetLatLonFromAddressRequestValidationError is the validation error returned
// by GetLatLonFromAddressRequest.Validate if the designated constraints
// aren't met.
type GetLatLonFromAddressRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLatLonFromAddressRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLatLonFromAddressRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLatLonFromAddressRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLatLonFromAddressRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLatLonFromAddressRequestValidationError) ErrorName() string {
	return "GetLatLonFromAddressRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetLatLonFromAddressRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLatLonFromAddressRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLatLonFromAddressRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLatLonFromAddressRequestValidationError{}

// Validate checks the field values on GetLatLonFromAddressResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetLatLonFromAddressResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLatLonFromAddressResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetLatLonFromAddressResponseMultiError, or nil if none found.
func (m *GetLatLonFromAddressResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLatLonFromAddressResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetLatLon()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetLatLonFromAddressResponseValidationError{
					field:  "LatLon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetLatLonFromAddressResponseValidationError{
					field:  "LatLon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLatLon()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetLatLonFromAddressResponseValidationError{
				field:  "LatLon",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetLatLonFromAddressResponseMultiError(errors)
	}

	return nil
}

// GetLatLonFromAddressResponseMultiError is an error wrapping multiple
// validation errors returned by GetLatLonFromAddressResponse.ValidateAll() if
// the designated constraints aren't met.
type GetLatLonFromAddressResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLatLonFromAddressResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLatLonFromAddressResponseMultiError) AllErrors() []error { return m }

// GetLatLonFromAddressResponseValidationError is the validation error returned
// by GetLatLonFromAddressResponse.Validate if the designated constraints
// aren't met.
type GetLatLonFromAddressResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLatLonFromAddressResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLatLonFromAddressResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLatLonFromAddressResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLatLonFromAddressResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLatLonFromAddressResponseValidationError) ErrorName() string {
	return "GetLatLonFromAddressResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetLatLonFromAddressResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLatLonFromAddressResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLatLonFromAddressResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLatLonFromAddressResponseValidationError{}
