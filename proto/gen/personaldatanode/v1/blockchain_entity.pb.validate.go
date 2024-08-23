// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: personaldatanode/v1/blockchain_entity.proto

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

// Validate checks the field values on BlockchainEntity with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *BlockchainEntity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BlockchainEntity with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BlockchainEntityMultiError, or nil if none found.
func (m *BlockchainEntity) ValidateAll() error {
	return m.validate(true)
}

func (m *BlockchainEntity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for T

	// no validation rules for Id

	// no validation rules for Nnid

	if len(errors) > 0 {
		return BlockchainEntityMultiError(errors)
	}

	return nil
}

// BlockchainEntityMultiError is an error wrapping multiple validation errors
// returned by BlockchainEntity.ValidateAll() if the designated constraints
// aren't met.
type BlockchainEntityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BlockchainEntityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BlockchainEntityMultiError) AllErrors() []error { return m }

// BlockchainEntityValidationError is the validation error returned by
// BlockchainEntity.Validate if the designated constraints aren't met.
type BlockchainEntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BlockchainEntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BlockchainEntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BlockchainEntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BlockchainEntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BlockchainEntityValidationError) ErrorName() string { return "BlockchainEntityValidationError" }

// Error satisfies the builtin error interface
func (e BlockchainEntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBlockchainEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BlockchainEntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BlockchainEntityValidationError{}
