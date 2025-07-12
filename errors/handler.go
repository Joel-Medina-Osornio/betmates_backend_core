package errors

// ErrorHandler define la interfaz para manejar errores de capa
// Cada aplicación implementa esta interfaz según sus necesidades
type ErrorHandler interface {
	// HandleError procesa un error de capa y retorna información para la respuesta
	HandleError(err LayerError) ErrorResponse
}

// ErrorResponse representa la respuesta de error estandarizada
type ErrorResponse struct {
	Error   string                 `json:"error"`
	Code    string                 `json:"code,omitempty"`
	Type    string                 `json:"type,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}
