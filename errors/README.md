# Sistema de Errores por Capas - Librería Compartida

Esta librería proporciona un sistema de manejo de errores desacoplado que permite a cada aplicación definir su propio mapeo de errores a códigos HTTP, manteniendo la separación de responsabilidades entre capas.

## Características

- ✅ **Desacoplado**: La librería base no conoce HTTP
- ✅ **Flexible**: Cada aplicación define su propio mapeo
- ✅ **Reutilizable**: Se puede usar en cualquier proyecto
- ✅ **Trazable**: Cada error tiene capa, código y detalles
- ✅ **Extensible**: Fácil agregar nuevos tipos de errores

## Arquitectura

```
┌─────────────────────────────────────────────────────────────┐
│                    Librería Compartida                      │
│  ┌─────────────────┐  ┌─────────────────┐  ┌──────────────┐ │
│  │   Error Types   │  │  Layer Error    │  │   Factory    │ │
│  │                 │  │   Interface     │  │              │ │
│  └─────────────────┘  └─────────────────┘  └──────────────┘ │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    Aplicación Específica                    │
│  ┌─────────────────┐  ┌─────────────────┐  ┌──────────────┐ │
│  │ Error Mapping   │  │ Error Handler   │  │ Middleware   │ │
│  │ (HTTP Codes)    │  │ Implementation  │  │ Integration  │ │
│  └─────────────────┘  └─────────────────┘  └──────────────┘ │
└─────────────────────────────────────────────────────────────┘
```

## Uso Básico

### 1. Crear Errores

```go
import sharedErrors "betting-app/shared/errors"

// Error de infraestructura
err := sharedErrors.NewInfrastructureError(
    sharedErrors.ErrDatabaseConnection,
    "Database connection failed",
    map[string]interface{}{
        "database": "postgres",
        "host": "localhost",
    },
)

// Error de aplicación
err := sharedErrors.NewValidationError(
    sharedErrors.ErrInvalidEmail,
    "Invalid email format",
    map[string]interface{}{
        "email": "invalid-email",
    },
)

// Error de dominio
err := sharedErrors.NewBusinessRuleError(
    sharedErrors.ErrInvalidBusinessRule,
    "User cannot place bet with insufficient funds",
    map[string]interface{}{
        "user_id": "123",
        "required_amount": 100,
        "available_amount": 50,
    },
)
```

### 2. Definir Mapeo de Errores (Aplicación Específica)

```go
// auth-service/internal/infrastructure/http/errors/error_mapping.go
var AuthServiceErrorMapping = map[sharedErrors.ErrorCode]int{
    // Validation Errors (400)
    sharedErrors.ErrInvalidEmail:    http.StatusBadRequest,
    sharedErrors.ErrMissingRequired: http.StatusBadRequest,
    
    // Authentication Errors (401)
    sharedErrors.ErrInvalidToken:    http.StatusUnauthorized,
    sharedErrors.ErrExpiredToken:    http.StatusUnauthorized,
    
    // Not Found Errors (404)
    sharedErrors.ErrUserNotFound:    http.StatusNotFound,
    
    // Conflict Errors (409)
    sharedErrors.ErrUserAlreadyExists: http.StatusConflict,
    
    // Infrastructure Errors (424)
    sharedErrors.ErrDatabaseConnection: http.StatusFailedDependency,
}

// Códigos específicos del auth-service
const (
    ErrInvalidProviderToken sharedErrors.ErrorCode = "INVALID_PROVIDER_TOKEN"
    ErrProviderUnavailable  sharedErrors.ErrorCode = "PROVIDER_UNAVAILABLE"
)

func init() {
    AuthServiceErrorMapping[ErrInvalidProviderToken] = http.StatusUnauthorized
    AuthServiceErrorMapping[ErrProviderUnavailable] = http.StatusServiceUnavailable
}
```

### 3. Implementar Error Handler

```go
// auth-service/internal/infrastructure/http/errors/http_error_handler.go
type AuthServiceErrorHandler struct{}

func (h *AuthServiceErrorHandler) HandleHTTPError(err sharedErrors.LayerError) sharedErrors.HTTPErrorResponse {
    response := sharedErrors.ErrorResponse{
        Error:   err.Error(),
        Code:    string(err.Code()),
        Type:    string(err.Type()),
        Details: err.Details(),
    }
    
    // Obtener código HTTP del mapeo
    httpStatus, exists := AuthServiceErrorMapping[err.Code()]
    if !exists {
        httpStatus = h.getFallbackHTTPStatus(err.Type())
    }
    
    return sharedErrors.HTTPErrorResponse{
        ErrorResponse: response,
        HTTPStatus:    httpStatus,
    }
}
```

### 4. Integrar en Middleware

```go
// auth-service/internal/infrastructure/http/middlewares/error_handler.go
func ErrorHandlerMiddleware(errorHandler sharedErrors.HTTPErrorHandler) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        
        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err
            
            if layerErr, ok := err.(sharedErrors.LayerError); ok {
                httpResponse := errorHandler.HandleHTTPError(layerErr)
                c.JSON(httpResponse.HTTPStatus, httpResponse.ErrorResponse)
                return
            }
            
            // Error no controlado
            c.JSON(http.StatusInternalServerError, sharedErrors.ErrorResponse{
                Error: "Internal server error",
                Code:  "INTERNAL_ERROR",
                Type:  "internal",
            })
        }
    }
}
```

### 5. Configurar en Router

```go
// auth-service/internal/infrastructure/http/configuration/router.go
func SetupRouter(authController *http.AuthController, healthController *http.HealthController, jwtSecret string) *gin.Engine {
    router := gin.Default()
    
    // Crear el error handler específico del auth-service
    errorHandler := errors.NewAuthServiceErrorHandler()
    
    router.Use(
        middlewares.CORS(),
        middlewares.Logger(),
        middlewares.Recovery(),
        middlewares.ErrorHandlerMiddleware(errorHandler), // ← Agregar aquí
    )
    
    // ... resto de la configuración
    return router
}
```

## Ejemplos de Uso por Capa

### Infraestructura (Repositorios)

```go
func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
    user, exists := r.users[id]
    if !exists {
        return nil, sharedErrors.NewInfrastructureError(
            sharedErrors.ErrUserNotFound,
            "User not found in database",
            map[string]interface{}{
                "user_id": id.String(),
            },
        )
    }
    return user, nil
}
```

### Aplicación (Servicios)

```go
func (s *AuthService) validateProviderToken(provider entities.Provider, token string) (*dto.ProviderUserInfo, error) {
    if token == "" {
        return nil, sharedErrors.NewValidationError(
            sharedErrors.ErrMissingRequired,
            "Authentication token is required",
        )
    }
    
    switch provider {
    case entities.Google:
        return s.validateGoogleToken(token)
    case entities.Apple:
        return s.validateAppleToken(token)
    default:
        return nil, sharedErrors.NewValidationError(
            sharedErrors.ErrInvalidBusinessRule,
            "Unsupported authentication provider",
            map[string]interface{}{
                "provider": string(provider),
            },
        )
    }
}
```

### Dominio (Entidades)

```go
func (u *User) PlaceBet(amount float64) error {
    if amount <= 0 {
        return sharedErrors.NewBusinessRuleError(
            sharedErrors.ErrInvalidBusinessRule,
            "Bet amount must be positive",
            map[string]interface{}{
                "amount": amount,
            },
        )
    }
    
    if u.Balance < amount {
        return sharedErrors.NewBusinessRuleError(
            sharedErrors.ErrInvalidBusinessRule,
            "Insufficient funds for bet",
            map[string]interface{}{
                "required": amount,
                "available": u.Balance,
            },
        )
    }
    
    u.Balance -= amount
    return nil
}
```

## Ventajas del Diseño

### 1. **Separación de Responsabilidades**
- La librería base no conoce HTTP
- Cada aplicación define su propio mapeo
- Las capas internas no tienen dependencias externas

### 2. **Flexibilidad**
- Diferentes servicios pueden mapear el mismo error a diferentes códigos HTTP
- Fácil agregar nuevos códigos de error específicos del dominio
- Configuración independiente por aplicación

### 3. **Reutilización**
- La librería se puede usar en cualquier proyecto
- Los códigos de error comunes están predefinidos
- Fácil extensión para nuevos tipos de error

### 4. **Trazabilidad**
- Cada error tiene información completa de capa, código y detalles
- Logs estructurados para debugging
- Respuestas consistentes para el cliente

## Estructura del Proyecto

```
backend-services/
├── shared/
│   └── errors/
│       ├── error_types.go          # Tipos y códigos de error base
│       ├── layer_error.go          # Interfaz base de errores
│       ├── factory.go              # Factory methods
│       ├── handler.go              # Interfaces de handlers
│       ├── protocol_handlers.go    # Interfaces para diferentes protocolos
│       ├── examples/               # Ejemplos de implementación
│       │   ├── protocol_examples.go
│       │   └── usage_examples.go
│       └── README.md
├── auth-service/
│   └── internal/
│       └── infrastructure/
│           └── http/
│               └── errors/
│                   ├── error_mapping.go      # Mapeo específico del auth-service
│                   └── http_error_handler.go # Handler HTTP específico
├── betting-service/
│   └── internal/
│       └── infrastructure/
│           └── grpc/
│               └── errors/
│                   ├── error_mapping.go      # Mapeo específico del betting-service
│                   └── grpc_error_handler.go # Handler gRPC específico
└── wallet-service/
    └── internal/
        └── infrastructure/
            └── graphql/
                └── errors/
                    ├── error_mapping.go         # Mapeo específico del wallet-service
                    └── graphql_error_handler.go # Handler GraphQL específico
```

## Separación de Responsabilidades

### **Librería Compartida (`shared/errors/`)**
- ✅ **Agnóstica de protocolo**: No conoce HTTP, gRPC, GraphQL, etc.
- ✅ **Reutilizable**: Se puede usar en cualquier proyecto
- ✅ **Extensible**: Fácil agregar nuevos tipos de errores
- ✅ **Ejemplos genéricos**: Muestra cómo implementar diferentes protocolos

### **Servicios Específicos**
- ✅ **Mapeo específico**: Cada servicio define su propio mapeo de errores
- ✅ **Handler específico**: Cada servicio implementa su handler de protocolo
- ✅ **Independiente**: Cada servicio puede usar diferentes protocolos

## Migración Gradual

1. **Paso 1**: Crear la librería compartida
2. **Paso 2**: Implementar el mapeo específico del servicio
3. **Paso 3**: Actualizar repositorios gradualmente
4. **Paso 4**: Actualizar servicios y casos de uso
5. **Paso 5**: Actualizar controladores
6. **Paso 6**: Agregar el middleware de error handling

Este diseño te permite tener un sistema de errores robusto, flexible y reutilizable que se adapta a las necesidades específicas de cada servicio mientras mantiene la consistencia en toda la aplicación. 