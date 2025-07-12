package errors

type LayerType string

type ErrorType string

type ErrorCode string

// Only type definitions are kept. No concrete error codes or log levels.

type LayerError interface {
	error
	Layer() LayerType
	Code() ErrorCode
	Type() ErrorType
	Details() map[string]interface{}
}

type baseError struct {
	layer   LayerType
	code    ErrorCode
	errType ErrorType
	message string
	details map[string]interface{}
}

func (e *baseError) Error() string {
	return e.message
}

func (e *baseError) Layer() LayerType {
	return e.layer
}

func (e *baseError) Code() ErrorCode {
	return e.code
}

func (e *baseError) Type() ErrorType {
	return e.errType
}

func (e *baseError) Details() map[string]interface{} {
	return e.details
}

func (l LayerType) String() string {
	return string(l)
}

func (e ErrorType) String() string {
	return string(e)
}

func (c ErrorCode) String() string {
	return string(c)
}
