package jrpc2errors

import "errors"

// ErrorCode type for error codes
type ErrorCode int

const (
	// ParseError Parse error - Invalid JSON was received by the server.
	// An error occurred on the server while parsing the JSON text.
	ParseError ErrorCode = -32700

	// ParseErrorMessage Parse error - Invalid JSON was received by the server.
	// An error occurred on the server while parsing the JSON text.
	ParseErrorMessage = "Parse error"

	// InvalidRequestError Invalid Request - The JSON sent is not a valid Request object.
	InvalidRequestError ErrorCode = -32600

	// InvalidRequestMessage Invalid Request - The JSON sent is not a valid Request object.
	InvalidRequestMessage = "Invalid request"

	// MethodNotFoundError Method not found - The method does not exist / is not available.
	MethodNotFoundError ErrorCode = -32601

	// MethodNotFoundMessage Method not found - The method does not exist / is not available.
	MethodNotFoundMessage = "Method not found"

	// InvalidParamsError Invalid params - Invalid method parameter(s).
	InvalidParamsError ErrorCode = -32602

	// InvalidParamsMessage Invalid params - Invalid method parameter(s).
	InvalidParamsMessage = "Invalid params"

	// InternalError Internal error - Internal JSON-RPC error.
	InternalError ErrorCode = -32603

	// InternalErrorMessage Internal error - Internal JSON-RPC error.
	InternalErrorMessage = "Internal error"

	// ServerError Server error - Reserved for implementation-defined server-errors.
	ServerError ErrorCode = -32000

	// ServerErrorMessage Server error - Reserved for implementation-defined server-errors.
	ServerErrorMessage = "Server error"

	// DataNotFoundError ...
	DataNotFoundError ErrorCode = -32001

	// DataNotFoundTitle ...
	DataNotFoundTitle = "Data not found"

	// PermissionDeniedError ...
	PermissionDeniedError ErrorCode = -32602

	// PermissionDeniedMessage ...
	PermissionDeniedMessage = "Permission denied"

	// AlreadyExecutedError ...
	AlreadyExecutedError ErrorCode = -32003

	// AlreadyExecutedMessage ...
	AlreadyExecutedMessage = "Already executed"

	// RegistrationIsRequiredError ...
	RegistrationIsRequiredError ErrorCode = -32004

	// RegistrationIsRequiredMessage ...
	RegistrationIsRequiredMessage = "Registration is required"

	// CheckIsRequiredError ...
	CheckIsRequiredError ErrorCode = -32005

	// CheckIsRequiredMessage ...
	CheckIsRequiredMessage = "Check is required"

	// ExternalServiceError ...
	ExternalServiceError ErrorCode = -32006

	// ExternalServiceMessage ...
	ExternalServiceMessage = "Error occurred when calling external service, try again later"
)

// ErrNullResult null result error
var ErrNullResult = errors.New("result is null")

// Error basic error struct for API answer
type Error struct {
	// A Number that indicates the error type that occurred.
	Code ErrorCode `json:"code"` /* required */

	// A String providing a short description of the error.
	// The message SHOULD be limited to a concise single sentence.
	Message string `json:"message"` /* required */

	// A Primitive or Structured value that contains additional information about the error.
	Data interface{} `json:"data"` /* optional */
}

// Error returns error message in string format
func (e *Error) Error() string {
	return e.Message
}
