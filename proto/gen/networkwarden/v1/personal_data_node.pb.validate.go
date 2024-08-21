// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: networkwarden/v1/personal_data_node.proto

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

// Validate checks the field values on PersonalDataNode with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PersonalDataNode) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PersonalDataNode with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PersonalDataNodeMultiError, or nil if none found.
func (m *PersonalDataNode) ValidateAll() error {
	return m.validate(true)
}

func (m *PersonalDataNode) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for CreatedAt

	// no validation rules for NwId

	// no validation rules for Address

	// no validation rules for Label

	// no validation rules for Name

	// no validation rules for Description

	if all {
		switch v := interface{}(m.GetLocation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PersonalDataNodeValidationError{
					field:  "Location",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PersonalDataNodeValidationError{
					field:  "Location",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PersonalDataNodeValidationError{
				field:  "Location",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AccountsCapacity

	// no validation rules for Alive

	// no validation rules for LastPingedAt

	// no validation rules for IsOpen

	// no validation rules for IsInviteCodeRequired

	// no validation rules for OwnerHolderId

	// no validation rules for Url

	// no validation rules for Version

	if all {
		switch v := interface{}(m.GetRateLimit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PersonalDataNodeValidationError{
					field:  "RateLimit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PersonalDataNodeValidationError{
					field:  "RateLimit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRateLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PersonalDataNodeValidationError{
				field:  "RateLimit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCrawlRateLimit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PersonalDataNodeValidationError{
					field:  "CrawlRateLimit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PersonalDataNodeValidationError{
					field:  "CrawlRateLimit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCrawlRateLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PersonalDataNodeValidationError{
				field:  "CrawlRateLimit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IdGenNode

	// no validation rules for Status

	if m.LastModifiedAt != nil {
		// no validation rules for LastModifiedAt
	}

	if len(errors) > 0 {
		return PersonalDataNodeMultiError(errors)
	}

	return nil
}

// PersonalDataNodeMultiError is an error wrapping multiple validation errors
// returned by PersonalDataNode.ValidateAll() if the designated constraints
// aren't met.
type PersonalDataNodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PersonalDataNodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PersonalDataNodeMultiError) AllErrors() []error { return m }

// PersonalDataNodeValidationError is the validation error returned by
// PersonalDataNode.Validate if the designated constraints aren't met.
type PersonalDataNodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PersonalDataNodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PersonalDataNodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PersonalDataNodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PersonalDataNodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PersonalDataNodeValidationError) ErrorName() string { return "PersonalDataNodeValidationError" }

// Error satisfies the builtin error interface
func (e PersonalDataNodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPersonalDataNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PersonalDataNodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PersonalDataNodeValidationError{}
