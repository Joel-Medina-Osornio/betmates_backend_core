package errors

func NewInfrastructureError(code ErrorCode, message string, details ...map[string]interface{}) LayerError {
	var d map[string]interface{}
	if len(details) > 0 {
		d = details[0]
	}
	return &baseError{
		layer:   InfrastructureLayer,
		code:    code,
		errType: InfrastructureError,
		message: message,
		details: d,
	}
}

func NewApplicationError(code ErrorCode, errType ErrorType, message string, details ...map[string]interface{}) LayerError {
	var d map[string]interface{}
	if len(details) > 0 {
		d = details[0]
	}
	return &baseError{
		layer:   ApplicationLayer,
		code:    code,
		errType: errType,
		message: message,
		details: d,
	}
}

func NewDomainError(code ErrorCode, errType ErrorType, message string, details ...map[string]interface{}) LayerError {
	var d map[string]interface{}
	if len(details) > 0 {
		d = details[0]
	}
	return &baseError{
		layer:   DomainLayer,
		code:    code,
		errType: errType,
		message: message,
		details: d,
	}
}

func NewValidationError(code ErrorCode, message string, details ...map[string]interface{}) LayerError {
	return NewApplicationError(code, ValidationError, message, details...)
}

func NewAuthenticationError(code ErrorCode, message string, details ...map[string]interface{}) LayerError {
	return NewApplicationError(code, AuthenticationError, message, details...)
}

func NewAuthorizationError(code ErrorCode, message string, details ...map[string]interface{}) LayerError {
	return NewApplicationError(code, AuthorizationError, message, details...)
}

func NewNotFoundError(code ErrorCode, message string, details ...map[string]interface{}) LayerError {
	return NewDomainError(code, NotFoundError, message, details...)
}

func NewConflictError(code ErrorCode, message string, details ...map[string]interface{}) LayerError {
	return NewDomainError(code, ConflictError, message, details...)
}

func NewBusinessRuleError(code ErrorCode, message string, details ...map[string]interface{}) LayerError {
	return NewDomainError(code, BusinessRuleError, message, details...)
}
