package validation

import "fmt"

type AppError interface {
	ErrorCode() string
	Message() string
	Details() []error
	Target() string
	error
}

const (
	ErrCodeRequireField            = "ErrCodeRequiredField"
	ErrCodeNotFound                = "ErrCodeNotFound"
	ErrCodeInvalidInput            = "ErrCodeInvalidInput"
	ErrCodeValidation              = "ErrCodeValidation"
	ErrCodeRuleRequired            = "ErrCodeRuleRequired"
	ErrCodeRuleDuplicate           = "ErrCodeRuleDuplicate"
	ErrCodeRuleInvalidValue        = "ErrCodeRuleInvalidValue"
	ErrCodeRuleInvalidType         = "ErrCodeRuleInvalidType"
	ErrCodeInvalidLength           = "ErrCodeInvalidLength"
	ErrCodeDuplicate               = "ErrCodeDuplicate"
	ErrCodeFallback                = "ErrCodeFallback"
	ErrCodeAuthFailed              = "ErrCodeAuthFailed"
	ErrCodeForbidden               = "ErrCodeForbidden"
	ErrCodeRuleViolation           = "ErrCodeRuleViolation"
	ErrCodeInsufficientFunds       = "ErrCodeInsufficientFunds"
	ErrCodeInvalidRequestParameter = "ErrCodeInvalidRequestParameter"
)

func NewAppError(code string, msg string, target string) AppError {
	return baseAppError{code: code, msg: msg, target: target}
}

func NewNotFoundError(msg string) AppError {
	return baseAppError{code: ErrCodeNotFound, msg: msg, target: ""}
}

func NewRequiredFieldError(field string) AppError {
	return baseAppError{code: ErrCodeRequireField, msg: fmt.Sprintf("'%s' is a required field", field), target: field}
}

func NewInsufficientFundsError(target string) AppError {
	return baseAppError{code: ErrCodeInsufficientFunds, msg: "Insufficient funds, the requested amount exceeds the available balance", target: target}
}

type baseAppError struct {
	code   string
	msg    string
	target string
}

func (b baseAppError) ErrorCode() string {
	return b.code
}

func (b baseAppError) Message() string {
	return b.msg
}

func (b baseAppError) Details() []error {
	return nil
}

func (b baseAppError) Target() string {
	return b.target
}

func (b baseAppError) Error() string {
	return b.msg
}

var _ AppError = (*baseAppError)(nil)
