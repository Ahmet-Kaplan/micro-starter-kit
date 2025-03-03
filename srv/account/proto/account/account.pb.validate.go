// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: srv/account/proto/account/account.proto

package accountsrv

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _account_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate is disabled for User. This method will always return nil.
func (m *User) Validate() error {
	return nil
}

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

// Validate checks the field values on UserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserRequest) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetId(); wrapper != nil {

		if err := m._validateUuid(wrapper.GetValue()); err != nil {
			return UserRequestValidationError{
				field:  "Id",
				reason: "value must be a valid UUID",
				cause:  err,
			}
		}

	}

	if wrapper := m.GetUsername(); wrapper != nil {

		if l := utf8.RuneCountInString(wrapper.GetValue()); l < 4 || l > 16 {
			return UserRequestValidationError{
				field:  "Username",
				reason: "value length must be between 4 and 16 runes, inclusive",
			}
		}

		if len(wrapper.GetValue()) > 256 {
			return UserRequestValidationError{
				field:  "Username",
				reason: "value length must be at most 256 bytes",
			}
		}

		if !_UserRequest_Username_Pattern.MatchString(wrapper.GetValue()) {
			return UserRequestValidationError{
				field:  "Username",
				reason: "value does not match regex pattern \"^[a-z0-9_-]{3,15}$\"",
			}
		}

	}

	if wrapper := m.GetFirstName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return UserRequestValidationError{
				field:  "FirstName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetLastName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return UserRequestValidationError{
				field:  "LastName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetEmail(); wrapper != nil {

		if err := m._validateEmail(wrapper.GetValue()); err != nil {
			return UserRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
		}

	}

	return nil
}

// UserRequestValidationError is the validation error returned by
// UserRequest.Validate if the designated constraints aren't met.
type UserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRequestValidationError) ErrorName() string { return "UserRequestValidationError" }

// Error satisfies the builtin error interface
func (e UserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRequestValidationError{}

var _UserRequest_Username_Pattern = regexp.MustCompile("^[a-z0-9_-]{3,15}$")

// Validate is disabled for UserResponse. This method will always return nil.
func (m *UserResponse) Validate() error {
	return nil
}

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

// Validate is disabled for UserExistResponse. This method will always return nil.
func (m *UserExistResponse) Validate() error {
	return nil
}

// UserExistResponseValidationError is the validation error returned by
// UserExistResponse.Validate if the designated constraints aren't met.
type UserExistResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserExistResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserExistResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserExistResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserExistResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserExistResponseValidationError) ErrorName() string {
	return "UserExistResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UserExistResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserExistResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserExistResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserExistResponseValidationError{}

// Validate checks the field values on UserListQuery with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserListQuery) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetLimit(); wrapper != nil {

		if val := wrapper.GetValue(); val < 1 || val > 100 {
			return UserListQueryValidationError{
				field:  "Limit",
				reason: "value must be inside range [1, 100]",
			}
		}

	}

	if wrapper := m.GetPage(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			return UserListQueryValidationError{
				field:  "Page",
				reason: "value must be greater than or equal to 1",
			}
		}

	}

	if v, ok := interface{}(m.GetSort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserListQueryValidationError{
				field:  "Sort",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetUsername(); wrapper != nil {

		if l := utf8.RuneCountInString(wrapper.GetValue()); l < 4 || l > 16 {
			return UserListQueryValidationError{
				field:  "Username",
				reason: "value length must be between 4 and 16 runes, inclusive",
			}
		}

		if len(wrapper.GetValue()) > 256 {
			return UserListQueryValidationError{
				field:  "Username",
				reason: "value length must be at most 256 bytes",
			}
		}

		if !_UserListQuery_Username_Pattern.MatchString(wrapper.GetValue()) {
			return UserListQueryValidationError{
				field:  "Username",
				reason: "value does not match regex pattern \"^[a-z0-9_-]{3,15}$\"",
			}
		}

	}

	if wrapper := m.GetFirstName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return UserListQueryValidationError{
				field:  "FirstName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetLastName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return UserListQueryValidationError{
				field:  "LastName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetEmail(); wrapper != nil {

		if err := m._validateEmail(wrapper.GetValue()); err != nil {
			return UserListQueryValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
		}

	}

	return nil
}

// UserListQueryValidationError is the validation error returned by
// UserListQuery.Validate if the designated constraints aren't met.
type UserListQueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserListQueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserListQueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserListQueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserListQueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserListQueryValidationError) ErrorName() string { return "UserListQueryValidationError" }

// Error satisfies the builtin error interface
func (e UserListQueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserListQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserListQueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserListQueryValidationError{}

var _UserListQuery_Username_Pattern = regexp.MustCompile("^[a-z0-9_-]{3,15}$")

// Validate is disabled for UserListResponse. This method will always return nil.
func (m *UserListResponse) Validate() error {
	return nil
}

// UserListResponseValidationError is the validation error returned by
// UserListResponse.Validate if the designated constraints aren't met.
type UserListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserListResponseValidationError) ErrorName() string { return "UserListResponseValidationError" }

// Error satisfies the builtin error interface
func (e UserListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserListResponseValidationError{}

// Validate is disabled for Profile. This method will always return nil.
func (m *Profile) Validate() error {
	return nil
}

// ProfileValidationError is the validation error returned by Profile.Validate
// if the designated constraints aren't met.
type ProfileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProfileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProfileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProfileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProfileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProfileValidationError) ErrorName() string { return "ProfileValidationError" }

// Error satisfies the builtin error interface
func (e ProfileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProfile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProfileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProfileValidationError{}

// Validate checks the field values on ProfileRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ProfileRequest) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetId(); wrapper != nil {

		if err := m._validateUuid(wrapper.GetValue()); err != nil {
			return ProfileRequestValidationError{
				field:  "Id",
				reason: "value must be a valid UUID",
				cause:  err,
			}
		}

	}

	if wrapper := m.GetUserId(); wrapper != nil {

		if err := m._validateUuid(wrapper.GetValue()); err != nil {
			return ProfileRequestValidationError{
				field:  "UserId",
				reason: "value must be a valid UUID",
				cause:  err,
			}
		}

	}

	if v, ok := interface{}(m.GetTz()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProfileRequestValidationError{
				field:  "Tz",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetAvatar(); wrapper != nil {

		if _, err := url.Parse(wrapper.GetValue()); err != nil {
			return ProfileRequestValidationError{
				field:  "Avatar",
				reason: "value must be a valid URI",
				cause:  err,
			}
		}

	}

	if wrapper := m.GetGender(); wrapper != nil {

		if _, ok := _ProfileRequest_Gender_InLookup[wrapper.GetValue()]; !ok {
			return ProfileRequestValidationError{
				field:  "Gender",
				reason: "value must be in list [M F]",
			}
		}

	}

	if t := m.GetBirthday(); t != nil {
		ts, err := ptypes.Timestamp(t)
		if err != nil {
			return ProfileRequestValidationError{
				field:  "Birthday",
				reason: "value is not a valid timestamp",
				cause:  err,
			}
		}

		now := time.Now()

		if ts.Sub(now) >= 0 {
			return ProfileRequestValidationError{
				field:  "Birthday",
				reason: "value must be less than now",
			}
		}

	}

	return nil
}

// ProfileRequestValidationError is the validation error returned by
// ProfileRequest.Validate if the designated constraints aren't met.
type ProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProfileRequestValidationError) ErrorName() string { return "ProfileRequestValidationError" }

// Error satisfies the builtin error interface
func (e ProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProfileRequestValidationError{}

var _ProfileRequest_Gender_InLookup = map[string]struct{}{
	"M": {},
	"F": {},
}

// Validate is disabled for ProfileResponse. This method will always return nil.
func (m *ProfileResponse) Validate() error {
	return nil
}

// ProfileResponseValidationError is the validation error returned by
// ProfileResponse.Validate if the designated constraints aren't met.
type ProfileResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProfileResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProfileResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProfileResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProfileResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProfileResponseValidationError) ErrorName() string { return "ProfileResponseValidationError" }

// Error satisfies the builtin error interface
func (e ProfileResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProfileResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProfileResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProfileResponseValidationError{}

// Validate checks the field values on ProfileListQuery with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ProfileListQuery) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetLimit(); wrapper != nil {

		if val := wrapper.GetValue(); val < 1 || val > 100 {
			return ProfileListQueryValidationError{
				field:  "Limit",
				reason: "value must be inside range [1, 100]",
			}
		}

	}

	if wrapper := m.GetPage(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			return ProfileListQueryValidationError{
				field:  "Page",
				reason: "value must be greater than or equal to 1",
			}
		}

	}

	if v, ok := interface{}(m.GetSort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProfileListQueryValidationError{
				field:  "Sort",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetUserId(); wrapper != nil {

		if err := m._validateUuid(wrapper.GetValue()); err != nil {
			return ProfileListQueryValidationError{
				field:  "UserId",
				reason: "value must be a valid UUID",
				cause:  err,
			}
		}

	}

	if wrapper := m.GetGender(); wrapper != nil {

		if _, ok := _ProfileListQuery_Gender_InLookup[wrapper.GetValue()]; !ok {
			return ProfileListQueryValidationError{
				field:  "Gender",
				reason: "value must be in list [M F]",
			}
		}

	}

	return nil
}

// ProfileListQueryValidationError is the validation error returned by
// ProfileListQuery.Validate if the designated constraints aren't met.
type ProfileListQueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProfileListQueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProfileListQueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProfileListQueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProfileListQueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProfileListQueryValidationError) ErrorName() string { return "ProfileListQueryValidationError" }

// Error satisfies the builtin error interface
func (e ProfileListQueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProfileListQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProfileListQueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProfileListQueryValidationError{}

var _ProfileListQuery_Gender_InLookup = map[string]struct{}{
	"M": {},
	"F": {},
}

// Validate is disabled for ProfileListResponse. This method will always return nil.
func (m *ProfileListResponse) Validate() error {
	return nil
}

// ProfileListResponseValidationError is the validation error returned by
// ProfileListResponse.Validate if the designated constraints aren't met.
type ProfileListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProfileListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProfileListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProfileListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProfileListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProfileListResponseValidationError) ErrorName() string {
	return "ProfileListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProfileListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProfileListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProfileListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProfileListResponseValidationError{}
