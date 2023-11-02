// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: medical.proto

package v1

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

// Validate checks the field values on UserGetRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserGetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserGetRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserGetRequestMultiError,
// or nil if none found.
func (m *UserGetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UserGetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return UserGetRequestMultiError(errors)
	}

	return nil
}

// UserGetRequestMultiError is an error wrapping multiple validation errors
// returned by UserGetRequest.ValidateAll() if the designated constraints
// aren't met.
type UserGetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserGetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserGetRequestMultiError) AllErrors() []error { return m }

// UserGetRequestValidationError is the validation error returned by
// UserGetRequest.Validate if the designated constraints aren't met.
type UserGetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserGetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserGetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserGetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserGetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserGetRequestValidationError) ErrorName() string { return "UserGetRequestValidationError" }

// Error satisfies the builtin error interface
func (e UserGetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserGetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserGetRequestValidationError{}

// Validate checks the field values on UserGetReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserGetReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserGetReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserGetReplyMultiError, or
// nil if none found.
func (m *UserGetReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UserGetReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserGetReplyValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserGetReplyValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserGetReplyValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserGetReplyMultiError(errors)
	}

	return nil
}

// UserGetReplyMultiError is an error wrapping multiple validation errors
// returned by UserGetReply.ValidateAll() if the designated constraints aren't met.
type UserGetReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserGetReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserGetReplyMultiError) AllErrors() []error { return m }

// UserGetReplyValidationError is the validation error returned by
// UserGetReply.Validate if the designated constraints aren't met.
type UserGetReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserGetReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserGetReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserGetReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserGetReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserGetReplyValidationError) ErrorName() string { return "UserGetReplyValidationError" }

// Error satisfies the builtin error interface
func (e UserGetReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserGetReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserGetReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserGetReplyValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Surname

	// no validation rules for Email

	// no validation rules for Role

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on ProtocolGetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProtocolGetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtocolGetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProtocolGetRequestMultiError, or nil if none found.
func (m *ProtocolGetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtocolGetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if all {
		switch v := interface{}(m.GetAnamnesis()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProtocolGetRequestValidationError{
					field:  "Anamnesis",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProtocolGetRequestValidationError{
					field:  "Anamnesis",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAnamnesis()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProtocolGetRequestValidationError{
				field:  "Anamnesis",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProtocolGetRequestMultiError(errors)
	}

	return nil
}

// ProtocolGetRequestMultiError is an error wrapping multiple validation errors
// returned by ProtocolGetRequest.ValidateAll() if the designated constraints
// aren't met.
type ProtocolGetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtocolGetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtocolGetRequestMultiError) AllErrors() []error { return m }

// ProtocolGetRequestValidationError is the validation error returned by
// ProtocolGetRequest.Validate if the designated constraints aren't met.
type ProtocolGetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolGetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolGetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolGetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolGetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolGetRequestValidationError) ErrorName() string {
	return "ProtocolGetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolGetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolGetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolGetRequestValidationError{}

// Validate checks the field values on ProtocolGetReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ProtocolGetReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtocolGetReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProtocolGetReplyMultiError, or nil if none found.
func (m *ProtocolGetReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtocolGetReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProtocol()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProtocolGetReplyValidationError{
					field:  "Protocol",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProtocolGetReplyValidationError{
					field:  "Protocol",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProtocol()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProtocolGetReplyValidationError{
				field:  "Protocol",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProtocolGetReplyMultiError(errors)
	}

	return nil
}

// ProtocolGetReplyMultiError is an error wrapping multiple validation errors
// returned by ProtocolGetReply.ValidateAll() if the designated constraints
// aren't met.
type ProtocolGetReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtocolGetReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtocolGetReplyMultiError) AllErrors() []error { return m }

// ProtocolGetReplyValidationError is the validation error returned by
// ProtocolGetReply.Validate if the designated constraints aren't met.
type ProtocolGetReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolGetReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolGetReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolGetReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolGetReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolGetReplyValidationError) ErrorName() string { return "ProtocolGetReplyValidationError" }

// Error satisfies the builtin error interface
func (e ProtocolGetReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolGetReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolGetReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolGetReplyValidationError{}

// Validate checks the field values on Protocol with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Protocol) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Protocol with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProtocolMultiError, or nil
// if none found.
func (m *Protocol) ValidateAll() error {
	return m.validate(true)
}

func (m *Protocol) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Description

	for idx, item := range m.GetActions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProtocolValidationError{
						field:  fmt.Sprintf("Actions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProtocolValidationError{
						field:  fmt.Sprintf("Actions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProtocolValidationError{
					field:  fmt.Sprintf("Actions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ProtocolMultiError(errors)
	}

	return nil
}

// ProtocolMultiError is an error wrapping multiple validation errors returned
// by Protocol.ValidateAll() if the designated constraints aren't met.
type ProtocolMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtocolMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtocolMultiError) AllErrors() []error { return m }

// ProtocolValidationError is the validation error returned by
// Protocol.Validate if the designated constraints aren't met.
type ProtocolValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolValidationError) ErrorName() string { return "ProtocolValidationError" }

// Error satisfies the builtin error interface
func (e ProtocolValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocol.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolValidationError{}

// Validate checks the field values on Actions with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Actions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Actions with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ActionsMultiError, or nil if none found.
func (m *Actions) ValidateAll() error {
	return m.validate(true)
}

func (m *Actions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Drug

	// no validation rules for Dosage

	// no validation rules for Time

	// no validation rules for Duration

	// no validation rules for Description

	if len(errors) > 0 {
		return ActionsMultiError(errors)
	}

	return nil
}

// ActionsMultiError is an error wrapping multiple validation errors returned
// by Actions.ValidateAll() if the designated constraints aren't met.
type ActionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActionsMultiError) AllErrors() []error { return m }

// ActionsValidationError is the validation error returned by Actions.Validate
// if the designated constraints aren't met.
type ActionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActionsValidationError) ErrorName() string { return "ActionsValidationError" }

// Error satisfies the builtin error interface
func (e ActionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActionsValidationError{}

// Validate checks the field values on Anamnesis with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Anamnesis) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Anamnesis with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AnamnesisMultiError, or nil
// if none found.
func (m *Anamnesis) ValidateAll() error {
	return m.validate(true)
}

func (m *Anamnesis) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Description

	if len(errors) > 0 {
		return AnamnesisMultiError(errors)
	}

	return nil
}

// AnamnesisMultiError is an error wrapping multiple validation errors returned
// by Anamnesis.ValidateAll() if the designated constraints aren't met.
type AnamnesisMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AnamnesisMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AnamnesisMultiError) AllErrors() []error { return m }

// AnamnesisValidationError is the validation error returned by
// Anamnesis.Validate if the designated constraints aren't met.
type AnamnesisValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AnamnesisValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AnamnesisValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AnamnesisValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AnamnesisValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AnamnesisValidationError) ErrorName() string { return "AnamnesisValidationError" }

// Error satisfies the builtin error interface
func (e AnamnesisValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAnamnesis.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AnamnesisValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AnamnesisValidationError{}
