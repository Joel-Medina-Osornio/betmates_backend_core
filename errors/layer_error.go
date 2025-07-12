package errors

// LayerError es la interfaz base para todos los errores de capa
type LayerError interface {
	error
	Layer() LayerType
	Code() ErrorCode
	Type() ErrorType
	Details() map[string]interface{}
}

// baseError implementa la interfaz LayerError
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
