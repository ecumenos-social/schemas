// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: personaldatanode/v1/blockchain_peer_service.proto

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

// Validate checks the field values on
// BlockchainPeerServiceCreateTransactionRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BlockchainPeerServiceCreateTransactionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// BlockchainPeerServiceCreateTransactionRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// BlockchainPeerServiceCreateTransactionRequestMultiError, or nil if none found.
func (m *BlockchainPeerServiceCreateTransactionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BlockchainPeerServiceCreateTransactionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BlockchainPeerServiceCreateTransactionRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BlockchainPeerServiceCreateTransactionRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BlockchainPeerServiceCreateTransactionRequestValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BlockchainPeerServiceCreateTransactionRequestMultiError(errors)
	}

	return nil
}

// BlockchainPeerServiceCreateTransactionRequestMultiError is an error wrapping
// multiple validation errors returned by
// BlockchainPeerServiceCreateTransactionRequest.ValidateAll() if the
// designated constraints aren't met.
type BlockchainPeerServiceCreateTransactionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BlockchainPeerServiceCreateTransactionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BlockchainPeerServiceCreateTransactionRequestMultiError) AllErrors() []error { return m }

// BlockchainPeerServiceCreateTransactionRequestValidationError is the
// validation error returned by
// BlockchainPeerServiceCreateTransactionRequest.Validate if the designated
// constraints aren't met.
type BlockchainPeerServiceCreateTransactionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BlockchainPeerServiceCreateTransactionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BlockchainPeerServiceCreateTransactionRequestValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e BlockchainPeerServiceCreateTransactionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BlockchainPeerServiceCreateTransactionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BlockchainPeerServiceCreateTransactionRequestValidationError) ErrorName() string {
	return "BlockchainPeerServiceCreateTransactionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e BlockchainPeerServiceCreateTransactionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBlockchainPeerServiceCreateTransactionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BlockchainPeerServiceCreateTransactionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BlockchainPeerServiceCreateTransactionRequestValidationError{}

// Validate checks the field values on
// BlockchainPeerServiceCreateTransactionResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BlockchainPeerServiceCreateTransactionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// BlockchainPeerServiceCreateTransactionResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// BlockchainPeerServiceCreateTransactionResponseMultiError, or nil if none found.
func (m *BlockchainPeerServiceCreateTransactionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *BlockchainPeerServiceCreateTransactionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return BlockchainPeerServiceCreateTransactionResponseMultiError(errors)
	}

	return nil
}

// BlockchainPeerServiceCreateTransactionResponseMultiError is an error
// wrapping multiple validation errors returned by
// BlockchainPeerServiceCreateTransactionResponse.ValidateAll() if the
// designated constraints aren't met.
type BlockchainPeerServiceCreateTransactionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BlockchainPeerServiceCreateTransactionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BlockchainPeerServiceCreateTransactionResponseMultiError) AllErrors() []error { return m }

// BlockchainPeerServiceCreateTransactionResponseValidationError is the
// validation error returned by
// BlockchainPeerServiceCreateTransactionResponse.Validate if the designated
// constraints aren't met.
type BlockchainPeerServiceCreateTransactionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BlockchainPeerServiceCreateTransactionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BlockchainPeerServiceCreateTransactionResponseValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e BlockchainPeerServiceCreateTransactionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BlockchainPeerServiceCreateTransactionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BlockchainPeerServiceCreateTransactionResponseValidationError) ErrorName() string {
	return "BlockchainPeerServiceCreateTransactionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e BlockchainPeerServiceCreateTransactionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBlockchainPeerServiceCreateTransactionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BlockchainPeerServiceCreateTransactionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BlockchainPeerServiceCreateTransactionResponseValidationError{}

// Validate checks the field values on BlockchainPeerServiceValidateRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *BlockchainPeerServiceValidateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BlockchainPeerServiceValidateRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// BlockchainPeerServiceValidateRequestMultiError, or nil if none found.
func (m *BlockchainPeerServiceValidateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BlockchainPeerServiceValidateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BlockchainPeerServiceValidateRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BlockchainPeerServiceValidateRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BlockchainPeerServiceValidateRequestValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BlockchainPeerServiceValidateRequestMultiError(errors)
	}

	return nil
}

// BlockchainPeerServiceValidateRequestMultiError is an error wrapping multiple
// validation errors returned by
// BlockchainPeerServiceValidateRequest.ValidateAll() if the designated
// constraints aren't met.
type BlockchainPeerServiceValidateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BlockchainPeerServiceValidateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BlockchainPeerServiceValidateRequestMultiError) AllErrors() []error { return m }

// BlockchainPeerServiceValidateRequestValidationError is the validation error
// returned by BlockchainPeerServiceValidateRequest.Validate if the designated
// constraints aren't met.
type BlockchainPeerServiceValidateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BlockchainPeerServiceValidateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BlockchainPeerServiceValidateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BlockchainPeerServiceValidateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BlockchainPeerServiceValidateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BlockchainPeerServiceValidateRequestValidationError) ErrorName() string {
	return "BlockchainPeerServiceValidateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e BlockchainPeerServiceValidateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBlockchainPeerServiceValidateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BlockchainPeerServiceValidateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BlockchainPeerServiceValidateRequestValidationError{}

// Validate checks the field values on BlockchainPeerServiceValidateResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *BlockchainPeerServiceValidateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BlockchainPeerServiceValidateResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// BlockchainPeerServiceValidateResponseMultiError, or nil if none found.
func (m *BlockchainPeerServiceValidateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *BlockchainPeerServiceValidateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Valid

	if len(errors) > 0 {
		return BlockchainPeerServiceValidateResponseMultiError(errors)
	}

	return nil
}

// BlockchainPeerServiceValidateResponseMultiError is an error wrapping
// multiple validation errors returned by
// BlockchainPeerServiceValidateResponse.ValidateAll() if the designated
// constraints aren't met.
type BlockchainPeerServiceValidateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BlockchainPeerServiceValidateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BlockchainPeerServiceValidateResponseMultiError) AllErrors() []error { return m }

// BlockchainPeerServiceValidateResponseValidationError is the validation error
// returned by BlockchainPeerServiceValidateResponse.Validate if the
// designated constraints aren't met.
type BlockchainPeerServiceValidateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BlockchainPeerServiceValidateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BlockchainPeerServiceValidateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BlockchainPeerServiceValidateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BlockchainPeerServiceValidateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BlockchainPeerServiceValidateResponseValidationError) ErrorName() string {
	return "BlockchainPeerServiceValidateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e BlockchainPeerServiceValidateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBlockchainPeerServiceValidateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BlockchainPeerServiceValidateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BlockchainPeerServiceValidateResponseValidationError{}

// Validate checks the field values on BlockchainPeerServiceGetBlocksRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *BlockchainPeerServiceGetBlocksRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BlockchainPeerServiceGetBlocksRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// BlockchainPeerServiceGetBlocksRequestMultiError, or nil if none found.
func (m *BlockchainPeerServiceGetBlocksRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BlockchainPeerServiceGetBlocksRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.From != nil {
		// no validation rules for From
	}

	if m.To != nil {
		// no validation rules for To
	}

	if m.Pagination != nil {

		if all {
			switch v := interface{}(m.GetPagination()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BlockchainPeerServiceGetBlocksRequestValidationError{
						field:  "Pagination",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BlockchainPeerServiceGetBlocksRequestValidationError{
						field:  "Pagination",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BlockchainPeerServiceGetBlocksRequestValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return BlockchainPeerServiceGetBlocksRequestMultiError(errors)
	}

	return nil
}

// BlockchainPeerServiceGetBlocksRequestMultiError is an error wrapping
// multiple validation errors returned by
// BlockchainPeerServiceGetBlocksRequest.ValidateAll() if the designated
// constraints aren't met.
type BlockchainPeerServiceGetBlocksRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BlockchainPeerServiceGetBlocksRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BlockchainPeerServiceGetBlocksRequestMultiError) AllErrors() []error { return m }

// BlockchainPeerServiceGetBlocksRequestValidationError is the validation error
// returned by BlockchainPeerServiceGetBlocksRequest.Validate if the
// designated constraints aren't met.
type BlockchainPeerServiceGetBlocksRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BlockchainPeerServiceGetBlocksRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BlockchainPeerServiceGetBlocksRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BlockchainPeerServiceGetBlocksRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BlockchainPeerServiceGetBlocksRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BlockchainPeerServiceGetBlocksRequestValidationError) ErrorName() string {
	return "BlockchainPeerServiceGetBlocksRequestValidationError"
}

// Error satisfies the builtin error interface
func (e BlockchainPeerServiceGetBlocksRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBlockchainPeerServiceGetBlocksRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BlockchainPeerServiceGetBlocksRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BlockchainPeerServiceGetBlocksRequestValidationError{}

// Validate checks the field values on BlockchainPeerServiceGetBlocksResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *BlockchainPeerServiceGetBlocksResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// BlockchainPeerServiceGetBlocksResponse with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// BlockchainPeerServiceGetBlocksResponseMultiError, or nil if none found.
func (m *BlockchainPeerServiceGetBlocksResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *BlockchainPeerServiceGetBlocksResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BlockchainPeerServiceGetBlocksResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BlockchainPeerServiceGetBlocksResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BlockchainPeerServiceGetBlocksResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return BlockchainPeerServiceGetBlocksResponseMultiError(errors)
	}

	return nil
}

// BlockchainPeerServiceGetBlocksResponseMultiError is an error wrapping
// multiple validation errors returned by
// BlockchainPeerServiceGetBlocksResponse.ValidateAll() if the designated
// constraints aren't met.
type BlockchainPeerServiceGetBlocksResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BlockchainPeerServiceGetBlocksResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BlockchainPeerServiceGetBlocksResponseMultiError) AllErrors() []error { return m }

// BlockchainPeerServiceGetBlocksResponseValidationError is the validation
// error returned by BlockchainPeerServiceGetBlocksResponse.Validate if the
// designated constraints aren't met.
type BlockchainPeerServiceGetBlocksResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BlockchainPeerServiceGetBlocksResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BlockchainPeerServiceGetBlocksResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BlockchainPeerServiceGetBlocksResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BlockchainPeerServiceGetBlocksResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BlockchainPeerServiceGetBlocksResponseValidationError) ErrorName() string {
	return "BlockchainPeerServiceGetBlocksResponseValidationError"
}

// Error satisfies the builtin error interface
func (e BlockchainPeerServiceGetBlocksResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBlockchainPeerServiceGetBlocksResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BlockchainPeerServiceGetBlocksResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BlockchainPeerServiceGetBlocksResponseValidationError{}

// Validate checks the field values on BlockchainPeerServiceGetEntitiesRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *BlockchainPeerServiceGetEntitiesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// BlockchainPeerServiceGetEntitiesRequest with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// BlockchainPeerServiceGetEntitiesRequestMultiError, or nil if none found.
func (m *BlockchainPeerServiceGetEntitiesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BlockchainPeerServiceGetEntitiesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Pagination != nil {

		if all {
			switch v := interface{}(m.GetPagination()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BlockchainPeerServiceGetEntitiesRequestValidationError{
						field:  "Pagination",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BlockchainPeerServiceGetEntitiesRequestValidationError{
						field:  "Pagination",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BlockchainPeerServiceGetEntitiesRequestValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return BlockchainPeerServiceGetEntitiesRequestMultiError(errors)
	}

	return nil
}

// BlockchainPeerServiceGetEntitiesRequestMultiError is an error wrapping
// multiple validation errors returned by
// BlockchainPeerServiceGetEntitiesRequest.ValidateAll() if the designated
// constraints aren't met.
type BlockchainPeerServiceGetEntitiesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BlockchainPeerServiceGetEntitiesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BlockchainPeerServiceGetEntitiesRequestMultiError) AllErrors() []error { return m }

// BlockchainPeerServiceGetEntitiesRequestValidationError is the validation
// error returned by BlockchainPeerServiceGetEntitiesRequest.Validate if the
// designated constraints aren't met.
type BlockchainPeerServiceGetEntitiesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BlockchainPeerServiceGetEntitiesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BlockchainPeerServiceGetEntitiesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BlockchainPeerServiceGetEntitiesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BlockchainPeerServiceGetEntitiesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BlockchainPeerServiceGetEntitiesRequestValidationError) ErrorName() string {
	return "BlockchainPeerServiceGetEntitiesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e BlockchainPeerServiceGetEntitiesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBlockchainPeerServiceGetEntitiesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BlockchainPeerServiceGetEntitiesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BlockchainPeerServiceGetEntitiesRequestValidationError{}

// Validate checks the field values on BlockchainPeerServiceGetEntitiesResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *BlockchainPeerServiceGetEntitiesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// BlockchainPeerServiceGetEntitiesResponse with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// BlockchainPeerServiceGetEntitiesResponseMultiError, or nil if none found.
func (m *BlockchainPeerServiceGetEntitiesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *BlockchainPeerServiceGetEntitiesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BlockchainPeerServiceGetEntitiesResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BlockchainPeerServiceGetEntitiesResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BlockchainPeerServiceGetEntitiesResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return BlockchainPeerServiceGetEntitiesResponseMultiError(errors)
	}

	return nil
}

// BlockchainPeerServiceGetEntitiesResponseMultiError is an error wrapping
// multiple validation errors returned by
// BlockchainPeerServiceGetEntitiesResponse.ValidateAll() if the designated
// constraints aren't met.
type BlockchainPeerServiceGetEntitiesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BlockchainPeerServiceGetEntitiesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BlockchainPeerServiceGetEntitiesResponseMultiError) AllErrors() []error { return m }

// BlockchainPeerServiceGetEntitiesResponseValidationError is the validation
// error returned by BlockchainPeerServiceGetEntitiesResponse.Validate if the
// designated constraints aren't met.
type BlockchainPeerServiceGetEntitiesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BlockchainPeerServiceGetEntitiesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BlockchainPeerServiceGetEntitiesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BlockchainPeerServiceGetEntitiesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BlockchainPeerServiceGetEntitiesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BlockchainPeerServiceGetEntitiesResponseValidationError) ErrorName() string {
	return "BlockchainPeerServiceGetEntitiesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e BlockchainPeerServiceGetEntitiesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBlockchainPeerServiceGetEntitiesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BlockchainPeerServiceGetEntitiesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BlockchainPeerServiceGetEntitiesResponseValidationError{}
