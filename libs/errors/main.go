package errors

import (
	"fmt"
	"runtime"
	"strings"
)

type errorString string

func (err errorString) Error() string {
	return string(err)
}

func New(message string) error {
	return errorString(message)
}

type errorWrapper interface {
	Error() string
	GetOriginalError() error
}

// WrappedError holds an error wrapped with a context message
type WrappedError struct {
	originalError error
	path          string
	messages      []string
}

func (err WrappedError) Error() string {
	if len(err.messages) > 0 {
		retVal := fmt.Sprintf("%s: ", err.path)

		for _, message := range err.messages {
			retVal += message + "; "
		}

		return fmt.Sprintf("%s => %v", retVal, err.originalError)
	}

	return fmt.Sprintf("%s => %v", err.path, err.originalError)
}

// GetOriginalError returns the original error
func (err WrappedError) GetOriginalError() error {
	if err.originalError != nil {
		if originalError, ok := (err.originalError).(errorWrapper); ok {
			return originalError.GetOriginalError()
		}
	}

	return err.originalError
}

func Is(a, b error) bool {
	return GetOriginalError(a) == GetOriginalError((b))
}

func Equals(a, b error) bool {
	if a == nil || b == nil {
		return Is(a, b)
	}

	return GetOriginalError(a).Error() == GetOriginalError(b).Error()
}

func caller() string {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	funcRef := runtime.FuncForPC(pc[0])

	pathArr := strings.Split(funcRef.Name(), "/")

	return pathArr[len(pathArr)-1]
}

// Wrap wraps an error with a context message
func Wrap(err error, messages ...string) error {
	if err == nil {
		return nil
	}

	return &WrappedError{
		originalError: err,
		path:          caller(),
		messages:      messages,
	}
}

func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return &WrappedError{
		originalError: err,
		path:          caller(),
		messages:      []string{fmt.Sprintf(format, args...)},
	}
}

// GetOriginalError returns the original error if the provided error is a WrappedError.
// Returns the provided error otherwise
func GetOriginalError(err error) error {
	wrappedErr, ok := err.(errorWrapper)
	if ok {
		return wrappedErr.GetOriginalError()
	}

	return err
}

// NullArgumentError represents a error that is used to sinalize that a provided argument is null
type NullArgumentError struct {
	ArgumentName string
}

func (e *NullArgumentError) Error() string {
	return fmt.Sprintf("Parameter %s can't be null", e.ArgumentName)
}
