// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/placeNoteUserAccountService.proto

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

// Validate checks the field values on AuthGoogleAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AuthGoogleAccountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthGoogleAccountRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthGoogleAccountRequestMultiError, or nil if none found.
func (m *AuthGoogleAccountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthGoogleAccountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetAuthCode()) < 1 {
		err := AuthGoogleAccountRequestValidationError{
			field:  "AuthCode",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AuthGoogleAccountRequestMultiError(errors)
	}

	return nil
}

// AuthGoogleAccountRequestMultiError is an error wrapping multiple validation
// errors returned by AuthGoogleAccountRequest.ValidateAll() if the designated
// constraints aren't met.
type AuthGoogleAccountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthGoogleAccountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthGoogleAccountRequestMultiError) AllErrors() []error { return m }

// AuthGoogleAccountRequestValidationError is the validation error returned by
// AuthGoogleAccountRequest.Validate if the designated constraints aren't met.
type AuthGoogleAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthGoogleAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthGoogleAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthGoogleAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthGoogleAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthGoogleAccountRequestValidationError) ErrorName() string {
	return "AuthGoogleAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AuthGoogleAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthGoogleAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthGoogleAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthGoogleAccountRequestValidationError{}

// Validate checks the field values on AuthGoogleAccountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AuthGoogleAccountResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthGoogleAccountResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthGoogleAccountResponseMultiError, or nil if none found.
func (m *AuthGoogleAccountResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthGoogleAccountResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return AuthGoogleAccountResponseMultiError(errors)
	}

	return nil
}

// AuthGoogleAccountResponseMultiError is an error wrapping multiple validation
// errors returned by AuthGoogleAccountResponse.ValidateAll() if the
// designated constraints aren't met.
type AuthGoogleAccountResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthGoogleAccountResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthGoogleAccountResponseMultiError) AllErrors() []error { return m }

// AuthGoogleAccountResponseValidationError is the validation error returned by
// AuthGoogleAccountResponse.Validate if the designated constraints aren't met.
type AuthGoogleAccountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthGoogleAccountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthGoogleAccountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthGoogleAccountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthGoogleAccountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthGoogleAccountResponseValidationError) ErrorName() string {
	return "AuthGoogleAccountResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthGoogleAccountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthGoogleAccountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthGoogleAccountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthGoogleAccountResponseValidationError{}

// Validate checks the field values on RegisterUserAccountRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterUserAccountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterUserAccountRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterUserAccountRequestMultiError, or nil if none found.
func (m *RegisterUserAccountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterUserAccountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetAuthToken()) < 1 {
		err := RegisterUserAccountRequestValidationError{
			field:  "AuthToken",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _RegisterUserAccountRequest_AuthMethod_NotInLookup[m.GetAuthMethod()]; ok {
		err := RegisterUserAccountRequestValidationError{
			field:  "AuthMethod",
			reason: "value must not be in list [UNKNOWN]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := AuthMethod_name[int32(m.GetAuthMethod())]; !ok {
		err := RegisterUserAccountRequestValidationError{
			field:  "AuthMethod",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetUserSettingId()) < 1 {
		err := RegisterUserAccountRequestValidationError{
			field:  "UserSettingId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := RegisterUserAccountRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RegisterUserAccountRequestMultiError(errors)
	}

	return nil
}

// RegisterUserAccountRequestMultiError is an error wrapping multiple
// validation errors returned by RegisterUserAccountRequest.ValidateAll() if
// the designated constraints aren't met.
type RegisterUserAccountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterUserAccountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterUserAccountRequestMultiError) AllErrors() []error { return m }

// RegisterUserAccountRequestValidationError is the validation error returned
// by RegisterUserAccountRequest.Validate if the designated constraints aren't met.
type RegisterUserAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterUserAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterUserAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterUserAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterUserAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterUserAccountRequestValidationError) ErrorName() string {
	return "RegisterUserAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterUserAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterUserAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterUserAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterUserAccountRequestValidationError{}

var _RegisterUserAccountRequest_AuthMethod_NotInLookup = map[AuthMethod]struct{}{
	0: {},
}

// Validate checks the field values on UserAccountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserAccountResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserAccountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserAccountResponseMultiError, or nil if none found.
func (m *UserAccountResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UserAccountResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for AuthMethod

	// no validation rules for UserSettingId

	// no validation rules for Name

	if len(errors) > 0 {
		return UserAccountResponseMultiError(errors)
	}

	return nil
}

// UserAccountResponseMultiError is an error wrapping multiple validation
// errors returned by UserAccountResponse.ValidateAll() if the designated
// constraints aren't met.
type UserAccountResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserAccountResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserAccountResponseMultiError) AllErrors() []error { return m }

// UserAccountResponseValidationError is the validation error returned by
// UserAccountResponse.Validate if the designated constraints aren't met.
type UserAccountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserAccountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserAccountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserAccountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserAccountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserAccountResponseValidationError) ErrorName() string {
	return "UserAccountResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UserAccountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserAccountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserAccountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserAccountResponseValidationError{}