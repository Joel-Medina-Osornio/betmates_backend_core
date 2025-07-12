# Betting App Core Library

Este módulo contiene utilidades y paquetes comunes para todos los servicios del ecosistema Betting App. Actualmente incluye:

- `errors`: Manejo de errores desacoplado por capas (infraestructura, dominio, aplicación).

## Ejemplo de uso básico (paquete errors)

```go
import "github.com/betting-app/core/errors"

err := errors.NewValidationError(
    errors.ErrInvalidEmail,
    "Invalid email format",
    map[string]interface{}{
        "email": "invalid-email",
    },
)
```

## Tareas por hacer (TODO)

- [ ] Agregar paquete de logging común
- [ ] Agregar utilidades de configuración
- [ ] Documentar ejemplos de integración multi-protocolo
- [ ] Mejorar cobertura de tests
- [ ] Publicar como módulo versionado

---

Licencia MIT. Ver archivo LICENSE. 