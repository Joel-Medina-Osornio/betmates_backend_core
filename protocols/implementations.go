package protocols

import (
	"net/http"

	"github.com/Joel-Medina-Osornio/betmates_backend_core/errors"
)

type DefaultHTTPErrorHandler struct {
	errorMapping map[errors.ErrorCode]int
}

func NewDefaultHTTPErrorHandler() *DefaultHTTPErrorHandler {
	return &DefaultHTTPErrorHandler{
		errorMapping: getDefaultErrorMapping(),
	}
}

func NewCustomHTTPErrorHandler(mapping map[errors.ErrorCode]int) *DefaultHTTPErrorHandler {
	return &DefaultHTTPErrorHandler{
		errorMapping: mapping,
	}
}

func (h *DefaultHTTPErrorHandler) HandleError(err errors.LayerError) ProtocolResponse {
	return ProtocolResponse{
		Error:   err.Error(),
		Code:    string(err.Code()),
		Type:    string(err.Type()),
		Details: err.Details(),
	}
}

func (h *DefaultHTTPErrorHandler) HandleHTTPError(err errors.LayerError) HTTPErrorResponse {
	response := h.HandleError(err)

	// Obtener código HTTP del mapeo
	httpStatus, exists := h.errorMapping[err.Code()]
	if !exists {
		httpStatus = h.getFallbackHTTPStatus(err.Type())
	}

	return HTTPErrorResponse{
		ProtocolResponse: response,
		HTTPStatus:       httpStatus,
	}
}

func (h *DefaultHTTPErrorHandler) getFallbackHTTPStatus(errType errors.ErrorType) int {
	switch errType {
	case errors.ValidationError:
		return http.StatusBadRequest
	case errors.AuthenticationError:
		return http.StatusUnauthorized
	case errors.AuthorizationError:
		return http.StatusForbidden
	case errors.NotFoundError:
		return http.StatusNotFound
	case errors.ConflictError:
		return http.StatusConflict
	case errors.BusinessRuleError:
		return http.StatusUnprocessableEntity
	case errors.InfrastructureError:
		return http.StatusFailedDependency
	case errors.InternalError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func getDefaultErrorMapping() map[errors.ErrorCode]int {
	return map[errors.ErrorCode]int{
		// Validation Errors (400)
		errors.ErrInvalidEmail:    http.StatusBadRequest,
		errors.ErrInvalidPassword: http.StatusBadRequest,
		errors.ErrMissingRequired: http.StatusBadRequest,
		errors.ErrInvalidFormat:   http.StatusBadRequest,

		// Authentication Errors (401)
		errors.ErrInvalidToken:       http.StatusUnauthorized,
		errors.ErrExpiredToken:       http.StatusUnauthorized,
		errors.ErrInvalidCredentials: http.StatusUnauthorized,

		// Authorization Errors (403)
		errors.ErrInsufficientPermissions: http.StatusForbidden,
		errors.ErrAccessDenied:            http.StatusForbidden,

		// Not Found Errors (404)
		errors.ErrUserNotFound:     http.StatusNotFound,
		errors.ErrResourceNotFound: http.StatusNotFound,

		// Conflict Errors (409)
		errors.ErrUserAlreadyExists: http.StatusConflict,
		errors.ErrEmailAlreadyTaken: http.StatusConflict,

		// Business Rule Errors (422)
		errors.ErrInvalidBusinessRule: http.StatusUnprocessableEntity,
		errors.ErrInvalidState:        http.StatusUnprocessableEntity,

		// Infrastructure Errors (424)
		errors.ErrDatabaseConnection:  http.StatusFailedDependency,
		errors.ErrExternalService:     http.StatusFailedDependency,
		errors.ErrRepositoryOperation: http.StatusFailedDependency,
	}
}

type DefaultGRPCErrorHandler struct{}

func NewDefaultGRPCErrorHandler() *DefaultGRPCErrorHandler {
	return &DefaultGRPCErrorHandler{}
}

func (h *DefaultGRPCErrorHandler) HandleError(err errors.LayerError) ProtocolResponse {
	return ProtocolResponse{
		Error:   err.Error(),
		Code:    string(err.Code()),
		Type:    string(err.Type()),
		Details: err.Details(),
	}
}

func (h *DefaultGRPCErrorHandler) HandleGRPCError(err errors.LayerError) GRPCErrorResponse {
	response := h.HandleError(err)

	// Mapeo básico de tipos de error a códigos gRPC
	grpcCode := h.mapToGRPCCode(err.Type())

	return GRPCErrorResponse{
		ProtocolResponse: response,
		GRPCCode:         grpcCode,
		GRPCMessage:      err.Error(),
	}
}

func (h *DefaultGRPCErrorHandler) mapToGRPCCode(errType errors.ErrorType) int {
	// Códigos gRPC básicos (puedes importar google.golang.org/grpc/codes para códigos completos)
	switch errType {
	case errors.ValidationError:
		return 3 // InvalidArgument
	case errors.AuthenticationError:
		return 16 // Unauthenticated
	case errors.AuthorizationError:
		return 7 // PermissionDenied
	case errors.NotFoundError:
		return 5 // NotFound
	case errors.ConflictError:
		return 6 // AlreadyExists
	case errors.BusinessRuleError:
		return 3 // InvalidArgument
	case errors.InfrastructureError:
		return 14 // Unavailable
	case errors.InternalError:
		return 13 // Internal
	default:
		return 13 // Internal
	}
}
