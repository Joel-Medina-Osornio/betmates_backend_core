package protocols

import "github.com/Joel-Medina-Osornio/betmates_backend_core/errors"

type ProtocolHandler interface {
	HandleError(err errors.LayerError) ProtocolResponse
}

type ProtocolResponse struct {
	Error   string                 `json:"error"`
	Code    string                 `json:"code,omitempty"`
	Type    string                 `json:"type,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}

type HTTPErrorResponse struct {
	ProtocolResponse
	HTTPStatus int `json:"-"`
}

type HTTPErrorHandler interface {
	ProtocolHandler
	HandleHTTPError(err errors.LayerError) HTTPErrorResponse
}

type SOAPErrorResponse struct {
	ProtocolResponse
	SOAPFaultCode   string `json:"-"`
	SOAPFaultString string `json:"-"`
	SOAPFaultActor  string `json:"-"`
}

type SOAPErrorHandler interface {
	ProtocolHandler
	HandleSOAPError(err errors.LayerError) SOAPErrorResponse
}

type GRPCErrorResponse struct {
	ProtocolResponse
	GRPCCode    int    `json:"-"`
	GRPCMessage string `json:"-"`
}

type GRPCErrorHandler interface {
	ProtocolHandler
	HandleGRPCError(err errors.LayerError) GRPCErrorResponse
}

type GraphQLErrorResponse struct {
	ProtocolResponse
	GraphQLErrorCode  string                 `json:"-"`
	GraphQLPath       []string               `json:"-"`
	GraphQLExtensions map[string]interface{} `json:"-"`
}

type GraphQLErrorHandler interface {
	ProtocolHandler
	HandleGraphQLError(err errors.LayerError) GraphQLErrorResponse
}
