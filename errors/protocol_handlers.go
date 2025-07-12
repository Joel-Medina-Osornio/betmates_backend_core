package errors

// ProtocolHandler define la interfaz base para manejar errores en diferentes protocolos
type ProtocolHandler interface {
	// HandleError procesa un error de capa y retorna información para la respuesta del protocolo
	HandleError(err LayerError) ProtocolResponse
}

// ProtocolResponse representa la respuesta de error estandarizada para cualquier protocolo
type ProtocolResponse struct {
	Error   string                 `json:"error"`
	Code    string                 `json:"code,omitempty"`
	Type    string                 `json:"type,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// HTTPErrorResponse extiende ProtocolResponse con información HTTP específica
type HTTPErrorResponse struct {
	ProtocolResponse
	HTTPStatus int `json:"-"`
}

// HTTPErrorHandler extiende ProtocolHandler con funcionalidad HTTP específica
type HTTPErrorHandler interface {
	ProtocolHandler
	// HandleHTTPError procesa un error y retorna una respuesta HTTP
	HandleHTTPError(err LayerError) HTTPErrorResponse
}

// SOAPErrorResponse extiende ProtocolResponse con información SOAP específica
type SOAPErrorResponse struct {
	ProtocolResponse
	SOAPFaultCode   string `json:"-"`
	SOAPFaultString string `json:"-"`
	SOAPFaultActor  string `json:"-"`
}

// SOAPErrorHandler extiende ProtocolHandler con funcionalidad SOAP específica
type SOAPErrorHandler interface {
	ProtocolHandler
	// HandleSOAPError procesa un error y retorna una respuesta SOAP
	HandleSOAPError(err LayerError) SOAPErrorResponse
}

// GRPCErrorResponse extiende ProtocolResponse con información gRPC específica
type GRPCErrorResponse struct {
	ProtocolResponse
	GRPCCode    int    `json:"-"`
	GRPCMessage string `json:"-"`
}

// GRPCErrorHandler extiende ProtocolHandler con funcionalidad gRPC específica
type GRPCErrorHandler interface {
	ProtocolHandler
	// HandleGRPCError procesa un error y retorna una respuesta gRPC
	HandleGRPCError(err LayerError) GRPCErrorResponse
}

// GraphQLErrorResponse extiende ProtocolResponse con información GraphQL específica
type GraphQLErrorResponse struct {
	ProtocolResponse
	GraphQLErrorCode  string                 `json:"-"`
	GraphQLPath       []string               `json:"-"`
	GraphQLExtensions map[string]interface{} `json:"-"`
}

// GraphQLErrorHandler extiende ProtocolHandler con funcionalidad GraphQL específica
type GraphQLErrorHandler interface {
	ProtocolHandler
	// HandleGraphQLError procesa un error y retorna una respuesta GraphQL
	HandleGraphQLError(err LayerError) GraphQLErrorResponse
}
