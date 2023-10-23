// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: systemcc/lscc/lscc.proto

package lscc

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

// Validate checks the field values on GetChaincodeDataRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetChaincodeDataRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetChaincodeDataRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetChaincodeDataRequestMultiError, or nil if none found.
func (m *GetChaincodeDataRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetChaincodeDataRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Channel

	// no validation rules for Chaincode

	if len(errors) > 0 {
		return GetChaincodeDataRequestMultiError(errors)
	}

	return nil
}

// GetChaincodeDataRequestMultiError is an error wrapping multiple validation
// errors returned by GetChaincodeDataRequest.ValidateAll() if the designated
// constraints aren't met.
type GetChaincodeDataRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetChaincodeDataRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetChaincodeDataRequestMultiError) AllErrors() []error { return m }

// GetChaincodeDataRequestValidationError is the validation error returned by
// GetChaincodeDataRequest.Validate if the designated constraints aren't met.
type GetChaincodeDataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetChaincodeDataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetChaincodeDataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetChaincodeDataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetChaincodeDataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetChaincodeDataRequestValidationError) ErrorName() string {
	return "GetChaincodeDataRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetChaincodeDataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetChaincodeDataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetChaincodeDataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetChaincodeDataRequestValidationError{}

// Validate checks the field values on GetChaincodesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetChaincodesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetChaincodesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetChaincodesRequestMultiError, or nil if none found.
func (m *GetChaincodesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetChaincodesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Channel

	if len(errors) > 0 {
		return GetChaincodesRequestMultiError(errors)
	}

	return nil
}

// GetChaincodesRequestMultiError is an error wrapping multiple validation
// errors returned by GetChaincodesRequest.ValidateAll() if the designated
// constraints aren't met.
type GetChaincodesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetChaincodesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetChaincodesRequestMultiError) AllErrors() []error { return m }

// GetChaincodesRequestValidationError is the validation error returned by
// GetChaincodesRequest.Validate if the designated constraints aren't met.
type GetChaincodesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetChaincodesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetChaincodesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetChaincodesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetChaincodesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetChaincodesRequestValidationError) ErrorName() string {
	return "GetChaincodesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetChaincodesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetChaincodesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetChaincodesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetChaincodesRequestValidationError{}

// Validate checks the field values on GetDeploymentSpecRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeploymentSpecRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeploymentSpecRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDeploymentSpecRequestMultiError, or nil if none found.
func (m *GetDeploymentSpecRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeploymentSpecRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Channel

	// no validation rules for Chaincode

	if len(errors) > 0 {
		return GetDeploymentSpecRequestMultiError(errors)
	}

	return nil
}

// GetDeploymentSpecRequestMultiError is an error wrapping multiple validation
// errors returned by GetDeploymentSpecRequest.ValidateAll() if the designated
// constraints aren't met.
type GetDeploymentSpecRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeploymentSpecRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeploymentSpecRequestMultiError) AllErrors() []error { return m }

// GetDeploymentSpecRequestValidationError is the validation error returned by
// GetDeploymentSpecRequest.Validate if the designated constraints aren't met.
type GetDeploymentSpecRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeploymentSpecRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeploymentSpecRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeploymentSpecRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeploymentSpecRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeploymentSpecRequestValidationError) ErrorName() string {
	return "GetDeploymentSpecRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeploymentSpecRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeploymentSpecRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeploymentSpecRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeploymentSpecRequestValidationError{}

// Validate checks the field values on DeployRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeployRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeployRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeployRequestMultiError, or
// nil if none found.
func (m *DeployRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeployRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Channel

	if all {
		switch v := interface{}(m.GetDeploymentSpec()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeployRequestValidationError{
					field:  "DeploymentSpec",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeployRequestValidationError{
					field:  "DeploymentSpec",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDeploymentSpec()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeployRequestValidationError{
				field:  "DeploymentSpec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPolicy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeployRequestValidationError{
					field:  "Policy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeployRequestValidationError{
					field:  "Policy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPolicy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeployRequestValidationError{
				field:  "Policy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ESCC

	// no validation rules for VSCC

	if all {
		switch v := interface{}(m.GetCollectionConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeployRequestValidationError{
					field:  "CollectionConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeployRequestValidationError{
					field:  "CollectionConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCollectionConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeployRequestValidationError{
				field:  "CollectionConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Transient

	if len(errors) > 0 {
		return DeployRequestMultiError(errors)
	}

	return nil
}

// DeployRequestMultiError is an error wrapping multiple validation errors
// returned by DeployRequest.ValidateAll() if the designated constraints
// aren't met.
type DeployRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeployRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeployRequestMultiError) AllErrors() []error { return m }

// DeployRequestValidationError is the validation error returned by
// DeployRequest.Validate if the designated constraints aren't met.
type DeployRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeployRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeployRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeployRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeployRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeployRequestValidationError) ErrorName() string { return "DeployRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeployRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeployRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeployRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeployRequestValidationError{}
