package errors

// AppError represents a custom application error with an HTTP status code, internal code, and message.
type AppError struct {
	HTTPCode int
	Code     string
	Message  string
}

// Error implements the error interface for AppError.
func (e AppError) Error() string {
	return e.Message
}

// New creates a new AppError with a custom message.
func (e AppError) New(msg string) AppError {
	return AppError{
		HTTPCode: e.HTTPCode,
		Code:     e.Code,
		Message:  msg,
	}
}

// Predefined errors with standard HTTP status codes and messages.
var (
	ErrBadRequest          = AppError{HTTPCode: 400, Code: "BAD_REQUEST", Message: "The request could not be understood by the server due to malformed syntax."}
	ErrUnauthorized        = AppError{HTTPCode: 401, Code: "UNAUTHORIZED", Message: "The request requires user authentication."}
	ErrForbidden           = AppError{HTTPCode: 403, Code: "FORBIDDEN", Message: "The server understood the request, but refuses to authorize it."}
	ErrNotFound            = AppError{HTTPCode: 404, Code: "NOT_FOUND", Message: "The server has not found anything matching the Request-URI."}
	ErrMethodNotAllowed    = AppError{HTTPCode: 405, Code: "METHOD_NOT_ALLOWED", Message: "The method specified in the Request-Line is not allowed for the resource identified by the Request-URI."}
	ErrUnprocessableEntity = AppError{HTTPCode: 422, Code: "UNPROCESSABLE_ENTITY", Message: "The request was well-formed but could not be processed due to semantic errors."}
	ErrInternalServer      = AppError{HTTPCode: 500, Code: "INTERNAL_SERVER_ERROR", Message: "The server encountered an unexpected condition which prevented it from fulfilling the request."}
)
