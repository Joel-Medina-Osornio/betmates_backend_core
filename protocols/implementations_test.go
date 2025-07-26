package protocols

import (
	"net/http"
	"testing"

	"github.com/Joel-Medina-Osornio/betmates_backend_core/errors"
)

func TestDefaultHTTPErrorHandler_HandleHTTPError(t *testing.T) {
	handler := NewDefaultHTTPErrorHandler()

	tests := []struct {
		name       string
		err        errors.LayerError
		wantStatus int
	}{
		{
			name:       "Validation Error should return 400",
			err:        errors.NewValidationError(errors.ErrInvalidEmail, "Invalid email"),
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Authentication Error should return 401",
			err:        errors.NewAuthenticationError(errors.ErrInvalidToken, "Invalid token"),
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "Authorization Error should return 403",
			err:        errors.NewApplicationError(errors.ErrInsufficientPermissions, errors.AuthorizationError, "Insufficient permissions"),
			wantStatus: http.StatusForbidden,
		},
		{
			name:       "Not Found Error should return 404",
			err:        errors.NewNotFoundError(errors.ErrUserNotFound, "User not found"),
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Conflict Error should return 409",
			err:        errors.NewConflictError(errors.ErrUserAlreadyExists, "User already exists"),
			wantStatus: http.StatusConflict,
		},
		{
			name:       "Business Rule Error should return 422",
			err:        errors.NewBusinessRuleError(errors.ErrInvalidBusinessRule, "Invalid business rule"),
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name:       "Infrastructure Error should return 424",
			err:        errors.NewInfrastructureError(errors.ErrDatabaseConnection, "Database connection failed"),
			wantStatus: http.StatusFailedDependency,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := handler.HandleHTTPError(tt.err)
			if response.HTTPStatus != tt.wantStatus {
				t.Errorf("HandleHTTPError() HTTPStatus = %v, want %v", response.HTTPStatus, tt.wantStatus)
			}
			if response.Error != tt.err.Error() {
				t.Errorf("HandleHTTPError() Error = %v, want %v", response.Error, tt.err.Error())
			}
			if response.Code != string(tt.err.Code()) {
				t.Errorf("HandleHTTPError() Code = %v, want %v", response.Code, tt.err.Code())
			}
			if response.Type != string(tt.err.Type()) {
				t.Errorf("HandleHTTPError() Type = %v, want %v", response.Type, tt.err.Type())
			}
		})
	}
}

func TestCustomHTTPErrorHandler_HandleHTTPError(t *testing.T) {
	customMapping := map[errors.ErrorCode]int{
		errors.ErrInvalidEmail:       http.StatusUnprocessableEntity, // 422 en lugar de 400
		errors.ErrUserNotFound:       http.StatusGone,                // 410 en lugar de 404
		errors.ErrDatabaseConnection: http.StatusServiceUnavailable,  // 503 en lugar de 424
	}

	handler := NewCustomHTTPErrorHandler(customMapping)

	tests := []struct {
		name       string
		err        errors.LayerError
		wantStatus int
	}{
		{
			name:       "Custom mapping for InvalidEmail should return 422",
			err:        errors.NewValidationError(errors.ErrInvalidEmail, "Invalid email"),
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name:       "Custom mapping for UserNotFound should return 410",
			err:        errors.NewNotFoundError(errors.ErrUserNotFound, "User not found"),
			wantStatus: http.StatusGone,
		},
		{
			name:       "Custom mapping for DatabaseConnection should return 503",
			err:        errors.NewInfrastructureError(errors.ErrDatabaseConnection, "Database connection failed"),
			wantStatus: http.StatusServiceUnavailable,
		},
		{
			name:       "Non-custom error should use fallback",
			err:        errors.NewAuthenticationError(errors.ErrInvalidToken, "Invalid token"),
			wantStatus: http.StatusUnauthorized, // Fallback to default
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := handler.HandleHTTPError(tt.err)
			if response.HTTPStatus != tt.wantStatus {
				t.Errorf("HandleHTTPError() HTTPStatus = %v, want %v", response.HTTPStatus, tt.wantStatus)
			}
		})
	}
}

func TestDefaultGRPCErrorHandler_HandleGRPCError(t *testing.T) {
	handler := NewDefaultGRPCErrorHandler()

	tests := []struct {
		name     string
		err      errors.LayerError
		wantCode int
		wantMsg  string
	}{
		{
			name:     "Validation Error should return InvalidArgument (3)",
			err:      errors.NewValidationError(errors.ErrInvalidEmail, "Invalid email"),
			wantCode: 3, // InvalidArgument
			wantMsg:  "Invalid email",
		},
		{
			name:     "Authentication Error should return Unauthenticated (16)",
			err:      errors.NewAuthenticationError(errors.ErrInvalidToken, "Invalid token"),
			wantCode: 16, // Unauthenticated
			wantMsg:  "Invalid token",
		},
		{
			name:     "Authorization Error should return PermissionDenied (7)",
			err:      errors.NewApplicationError(errors.ErrInsufficientPermissions, errors.AuthorizationError, "Insufficient permissions"),
			wantCode: 7, // PermissionDenied
			wantMsg:  "Insufficient permissions",
		},
		{
			name:     "Not Found Error should return NotFound (5)",
			err:      errors.NewNotFoundError(errors.ErrUserNotFound, "User not found"),
			wantCode: 5, // NotFound
			wantMsg:  "User not found",
		},
		{
			name:     "Conflict Error should return AlreadyExists (6)",
			err:      errors.NewConflictError(errors.ErrUserAlreadyExists, "User already exists"),
			wantCode: 6, // AlreadyExists
			wantMsg:  "User already exists",
		},
		{
			name:     "Infrastructure Error should return Unavailable (14)",
			err:      errors.NewInfrastructureError(errors.ErrDatabaseConnection, "Database connection failed"),
			wantCode: 14, // Unavailable
			wantMsg:  "Database connection failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := handler.HandleGRPCError(tt.err)
			if response.GRPCCode != tt.wantCode {
				t.Errorf("HandleGRPCError() GRPCCode = %v, want %v", response.GRPCCode, tt.wantCode)
			}
			if response.GRPCMessage != tt.wantMsg {
				t.Errorf("HandleGRPCError() GRPCMessage = %v, want %v", response.GRPCMessage, tt.wantMsg)
			}
			if response.Error != tt.err.Error() {
				t.Errorf("HandleGRPCError() Error = %v, want %v", response.Error, tt.err.Error())
			}
		})
	}
}

func TestDefaultHTTPErrorHandler_HandleError(t *testing.T) {
	handler := NewDefaultHTTPErrorHandler()
	err := errors.NewValidationError(errors.ErrInvalidEmail, "Invalid email", map[string]interface{}{
		"field": "email",
		"value": "invalid-email",
	})

	response := handler.HandleError(err)

	if response.Error != err.Error() {
		t.Errorf("HandleError() Error = %v, want %v", response.Error, err.Error())
	}
	if response.Code != string(err.Code()) {
		t.Errorf("HandleError() Code = %v, want %v", response.Code, err.Code())
	}
	if response.Type != string(err.Type()) {
		t.Errorf("HandleError() Type = %v, want %v", response.Type, err.Type())
	}
	if response.Details["field"] != "email" {
		t.Errorf("HandleError() Details[field] = %v, want %v", response.Details["field"], "email")
	}
}

func TestDefaultGRPCErrorHandler_HandleError(t *testing.T) {
	handler := NewDefaultGRPCErrorHandler()
	err := errors.NewValidationError(errors.ErrInvalidEmail, "Invalid email", map[string]interface{}{
		"field": "email",
		"value": "invalid-email",
	})

	response := handler.HandleError(err)

	if response.Error != err.Error() {
		t.Errorf("HandleError() Error = %v, want %v", response.Error, err.Error())
	}
	if response.Code != string(err.Code()) {
		t.Errorf("HandleError() Code = %v, want %v", response.Code, err.Code())
	}
	if response.Type != string(err.Type()) {
		t.Errorf("HandleError() Type = %v, want %v", response.Type, err.Type())
	}
	if response.Details["field"] != "email" {
		t.Errorf("HandleError() Details[field] = %v, want %v", response.Details["field"], "email")
	}
}
