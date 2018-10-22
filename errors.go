package jrpc2errors

import "errors"

// ErrorCode type for error codes
type ErrorCode int

const (
	// ParseError Parse error - Invalid JSON was received by the server.
	// An error occurred on the server while parsing the JSON text.
	ParseError ErrorCode = -32700
	// InvalidRequestError Invalid Request - The JSON sent is not a valid Request object.
	InvalidRequestError ErrorCode = -32600
	// MethodNotFoundError Method not found - The method does not exist / is not available.
	MethodNotFoundError ErrorCode = -32601
	// InvalidParamsError Invalid params - Invalid method parameter(s).
	InvalidParamsError ErrorCode = -32602
	// InternalError Internal error - Internal JSON-RPC error.
	InternalError ErrorCode = -32603
	// ServerError Server error - Reserved for implementation-defined server-errors.
	ServerError ErrorCode = -32000
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
