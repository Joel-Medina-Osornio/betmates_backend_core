package errors

// ErrorCode representa un código de error específico
type ErrorCode string

// ErrorType define el tipo de error para categorización
type ErrorType string

// LayerType define la capa donde ocurrió el error
type LayerType string

const (
	// Tipos de capa
	InfrastructureLayer LayerType = "infrastructure"
	ApplicationLayer    LayerType = "application"
	DomainLayer         LayerType = "domain"

	// Tipos de error
	ValidationError     ErrorType = "validation"
	AuthenticationError ErrorType = "authentication"
	AuthorizationError  ErrorType = "authorization"
	NotFoundError       ErrorType = "not_found"
	ConflictError       ErrorType = "conflict"
	BusinessRuleError   ErrorType = "business_rule"
	InfrastructureError ErrorType = "infrastructure"
	InternalError       ErrorType = "internal"
)

// Códigos de error comunes (cada aplicación puede definir los suyos)
const (
	// Validation Errors
	ErrInvalidEmail    ErrorCode = "INVALID_EMAIL"
	ErrInvalidPassword ErrorCode = "INVALID_PASSWORD"
	ErrMissingRequired ErrorCode = "MISSING_REQUIRED"
	ErrInvalidFormat   ErrorCode = "INVALID_FORMAT"

	// Authentication Errors
	ErrInvalidToken       ErrorCode = "INVALID_TOKEN"
	ErrExpiredToken       ErrorCode = "EXPIRED_TOKEN"
	ErrInvalidCredentials ErrorCode = "INVALID_CREDENTIALS"

	// Authorization Errors
	ErrInsufficientPermissions ErrorCode = "INSUFFICIENT_PERMISSIONS"
	ErrAccessDenied            ErrorCode = "ACCESS_DENIED"

	// Not Found Errors
	ErrUserNotFound     ErrorCode = "USER_NOT_FOUND"
	ErrResourceNotFound ErrorCode = "RESOURCE_NOT_FOUND"

	// Conflict Errors
	ErrUserAlreadyExists ErrorCode = "USER_ALREADY_EXISTS"
	ErrEmailAlreadyTaken ErrorCode = "EMAIL_ALREADY_TAKEN"

	// Business Rule Errors
	ErrInvalidBusinessRule ErrorCode = "INVALID_BUSINESS_RULE"
	ErrInvalidState        ErrorCode = "INVALID_STATE"

	// Infrastructure Errors
	ErrDatabaseConnection  ErrorCode = "DATABASE_CONNECTION"
	ErrExternalService     ErrorCode = "EXTERNAL_SERVICE"
	ErrRepositoryOperation ErrorCode = "REPOSITORY_OPERATION"
)
