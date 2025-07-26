package errors

// Layer types
const (
	InfrastructureLayer LayerType = "infrastructure"
	ApplicationLayer    LayerType = "application"
	DomainLayer         LayerType = "domain"
)

// Error types
const (
	ValidationError     ErrorType = "validation"
	AuthenticationError ErrorType = "authentication"
	AuthorizationError  ErrorType = "authorization"
	NotFoundError       ErrorType = "not_found"
	ConflictError       ErrorType = "conflict"
	BusinessRuleError   ErrorType = "business_rule"
	InfrastructureError ErrorType = "infrastructure"
	InternalError       ErrorType = "internal"
)

// Error codes
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
