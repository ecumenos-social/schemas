// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: networkwarden/v1/network_warden.proto

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

// Validate checks the field values on NetworkWarden with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NetworkWarden) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NetworkWarden with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NetworkWardenMultiError, or
// nil if none found.
func (m *NetworkWarden) ValidateAll() error {
	return m.validate(true)
}

func (m *NetworkWarden) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for CreatedAt

	// no validation rules for IdGenNode

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Label

	// no validation rules for Address

	// no validation rules for PdnCapacity

	// no validation rules for NnCapacity

	if all {
		switch v := interface{}(m.GetLocation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NetworkWardenValidationError{
					field:  "Location",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NetworkWardenValidationError{
					field:  "Location",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NetworkWardenValidationError{
				field:  "Location",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IsOpen

	// no validation rules for Url

	// no validation rules for Alive

	// no validation rules for Version

	if all {
		switch v := interface{}(m.GetRateLimit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NetworkWardenValidationError{
					field:  "RateLimit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NetworkWardenValidationError{
					field:  "RateLimit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRateLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NetworkWardenValidationError{
				field:  "RateLimit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.LastModifiedAt != nil {
		// no validation rules for LastModifiedAt
	}

	if m.LastPingedAt != nil {
		// no validation rules for LastPingedAt
	}

	if len(errors) > 0 {
		return NetworkWardenMultiError(errors)
	}

	return nil
}

// NetworkWardenMultiError is an error wrapping multiple validation errors
// returned by NetworkWarden.ValidateAll() if the designated constraints
// aren't met.
type NetworkWardenMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NetworkWardenMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NetworkWardenMultiError) AllErrors() []error { return m }

// NetworkWardenValidationError is the validation error returned by
// NetworkWarden.Validate if the designated constraints aren't met.
type NetworkWardenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NetworkWardenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NetworkWardenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NetworkWardenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NetworkWardenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NetworkWardenValidationError) ErrorName() string { return "NetworkWardenValidationError" }

// Error satisfies the builtin error interface
func (e NetworkWardenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNetworkWarden.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NetworkWardenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NetworkWardenValidationError{}
