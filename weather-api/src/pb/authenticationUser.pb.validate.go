// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: src/proto/authenticationUser.proto

package pb

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

// Validate checks the field values on VerifyGoogleAuthCodeRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyGoogleAuthCodeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyGoogleAuthCodeRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyGoogleAuthCodeRequestMultiError, or nil if none found.
func (m *VerifyGoogleAuthCodeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyGoogleAuthCodeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetGoogleAuthCode()) < 1 {
		err := VerifyGoogleAuthCodeRequestValidationError{
			field:  "GoogleAuthCode",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return VerifyGoogleAuthCodeRequestMultiError(errors)
	}

	return nil
}

// VerifyGoogleAuthCodeRequestMultiError is an error wrapping multiple
// validation errors returned by VerifyGoogleAuthCodeRequest.ValidateAll() if
// the designated constraints aren't met.
type VerifyGoogleAuthCodeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyGoogleAuthCodeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyGoogleAuthCodeRequestMultiError) AllErrors() []error { return m }

// VerifyGoogleAuthCodeRequestValidationError is the validation error returned
// by VerifyGoogleAuthCodeRequest.Validate if the designated constraints
// aren't met.
type VerifyGoogleAuthCodeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyGoogleAuthCodeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyGoogleAuthCodeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyGoogleAuthCodeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyGoogleAuthCodeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyGoogleAuthCodeRequestValidationError) ErrorName() string {
	return "VerifyGoogleAuthCodeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyGoogleAuthCodeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyGoogleAuthCodeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyGoogleAuthCodeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyGoogleAuthCodeRequestValidationError{}

// Validate checks the field values on UserResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserResponseMultiError, or
// nil if none found.
func (m *UserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for AuthToken

	if len(errors) > 0 {
		return UserResponseMultiError(errors)
	}

	return nil
}

// UserResponseMultiError is an error wrapping multiple validation errors
// returned by UserResponse.ValidateAll() if the designated constraints aren't met.
type UserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserResponseMultiError) AllErrors() []error { return m }

// UserResponseValidationError is the validation error returned by
// UserResponse.Validate if the designated constraints aren't met.
type UserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserResponseValidationError) ErrorName() string { return "UserResponseValidationError" }

// Error satisfies the builtin error interface
func (e UserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserResponseValidationError{}

// Validate checks the field values on VerifyAuthTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyAuthTokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyAuthTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyAuthTokenRequestMultiError, or nil if none found.
func (m *VerifyAuthTokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyAuthTokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return VerifyAuthTokenRequestMultiError(errors)
	}

	return nil
}

// VerifyAuthTokenRequestMultiError is an error wrapping multiple validation
// errors returned by VerifyAuthTokenRequest.ValidateAll() if the designated
// constraints aren't met.
type VerifyAuthTokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyAuthTokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyAuthTokenRequestMultiError) AllErrors() []error { return m }

// VerifyAuthTokenRequestValidationError is the validation error returned by
// VerifyAuthTokenRequest.Validate if the designated constraints aren't met.
type VerifyAuthTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyAuthTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyAuthTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyAuthTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyAuthTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyAuthTokenRequestValidationError) ErrorName() string {
	return "VerifyAuthTokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyAuthTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyAuthTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyAuthTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyAuthTokenRequestValidationError{}

// Validate checks the field values on AuthTokenResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AuthTokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthTokenResponseMultiError, or nil if none found.
func (m *AuthTokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthTokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if len(errors) > 0 {
		return AuthTokenResponseMultiError(errors)
	}

	return nil
}

// AuthTokenResponseMultiError is an error wrapping multiple validation errors
// returned by AuthTokenResponse.ValidateAll() if the designated constraints
// aren't met.
type AuthTokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthTokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthTokenResponseMultiError) AllErrors() []error { return m }

// AuthTokenResponseValidationError is the validation error returned by
// AuthTokenResponse.Validate if the designated constraints aren't met.
type AuthTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthTokenResponseValidationError) ErrorName() string {
	return "AuthTokenResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthTokenResponseValidationError{}
