package results

import "fmt"

const (
	CodeOk                      = "000000"
	CodeBadRequest              = "000001"
	CodeUnknownBusiness         = "000002"
	CodeUnknownType             = "000003"
	CodeDuplicateEntry          = "000004"
	CodeNotFound                = "000005"
	CodeAuthToken               = "000006"
	CodeInvalidParamType        = "000007"
	CodeTrainIsRunning          = "000008"
	CodeOperationConflict       = "000009"
	CodeDataOutOfDate           = "000010"
	CodeTestIsRunning           = "000011"
	CodeFilesystemError         = "000012"
	CodeAlreadyExistingDirError = "000013"
	CodeUnknownTestsetVersion   = "000014"
	CodeHanabiError             = "000015"
	CodeDuplicateTest           = "000016"
	CodePermissionDenied        = "000017"
	CodeUnknown                 = "999999"
)

const (
	ReasonDecodeRequest   = "FailedToDecodeRequest"
	ReasonUnknownBusiness = "UnknownBusiness"
	ReasonUnknownType     = "UnknownType"
)

type Error interface {
	Code() string
	Message() string
	Reason() string
	Error() string
	IsNotFound() bool
	IsConflict() bool
	IsDuplicate() bool
	Wrap(format string, args ...interface{}) Error
}

type simpleError struct {
	code    string
	message string
	reason  string
	details string
	err     error
}

func NewError(code, reason, message string, err error) Error {
	return &simpleError{
		code:    code,
		reason:  reason,
		message: message,
		err:     err,
	}
}

func (e *simpleError) Code() string {
	return e.code
}

func (e *simpleError) Message() string {
	return e.message
}

func (e *simpleError) Reason() string {
	return e.reason
}

func (e *simpleError) Error() string {
	message := `code: ` + e.code + ` message: ` + e.message
	if e.details != "" {
		message = message + " details: " + e.details
	}
	if e.err != nil {
		return message + " error: " + e.err.Error()
	}
	return message
}

func (e *simpleError) IsNotFound() bool {
	if e == nil {
		return false
	}
	return e.code == CodeNotFound
}

func (e *simpleError) IsConflict() bool {
	if e == nil {
		return false
	}
	return e.code == CodeOperationConflict
}

func (e *simpleError) IsDuplicate() bool {
	if e == nil {
		return false
	}
	return e.code == CodeDuplicateEntry
}

func (e *simpleError) Wrap(format string, args ...interface{}) Error {
	details := fmt.Sprintf(format, args...)
	if e.details == "" {
		e.details = details
	} else {
		e.details = details + " " + e.details
	}
	return e
}

func NewDuplicateError(message string, err error) Error {
	return &simpleError{
		code:    CodeDuplicateEntry,
		message: message,
		err:     err,
	}
}

func NewUnknownError(message string, err error) Error {
	return &simpleError{
		code:    CodeUnknown,
		message: message,
		err:     err,
	}
}

func NewNotFoundError(message string) Error {
	return &simpleError{
		code:    CodeNotFound,
		message: message,
	}
}

func NewAuthTokenError(message string) Error {
	return &simpleError{
		code:    CodeAuthToken,
		message: message,
	}
}

func FromError(err error) Error {
	switch v := err.(type) {
	case nil:
		return nil
	case Error:
		return v
	default:
		return NewUnknownError("未知错误", err)
	}
}

func NewInvalidParamTypeError(message string) Error {
	return &simpleError{
		code:    CodeInvalidParamType,
		message: message,
	}
}

func NewOperationConflictError(message string) Error {
	return &simpleError{
		code:    CodeOperationConflict,
		message: message,
	}
}

func NewFilesystemError(message string, err error) Error {
	return &simpleError{
		code:    CodeFilesystemError,
		message: message,
		err:     err,
	}
}

func NewAlreadyExistingDirError(message string) Error {
	return &simpleError{
		code:    CodeAlreadyExistingDirError,
		message: message,
	}
}

func NewUnknownTestsetVersion(version int64) Error {
	return &simpleError{
		code:    CodeUnknownTestsetVersion,
		message: fmt.Sprintf("未知测试集版本: %d", version),
	}
}

func NewPermissionDenied(message string) Error {
	return &simpleError{
		code:    CodePermissionDenied,
		message: message,
	}
}
