package exceptions

type Kind uint8

const (
	Other           Kind = iota // Unclassified error. This value is not printed in the error message.
	IO                          // External I/O error such as network failure.
	Unanticipated               // Unanticipated error.
	Internal                    // Internal error or inconsistency.
	Database                    // Error from database.
	Invalid                     // Invalid operation for this type of item.
	Exist                       // Item already exists.
	NotExist                    // Item does not exist.
	Private                     // Information withheld.
	Validation                  // Input validation error.
	InvalidRequest              // Invalid Request
	Unauthenticated             //
	Unauthorized                //
)

func (k Kind) String() string {
	switch k {
	case Other:
		return "other_error"
	case Invalid:
		return "invalid_operation"
	case IO:
		return "I/O_error"
	case Exist:
		return "item_already_exists"
	case NotExist:
		return "item_does_not_exist"
	case Private:
		return "information_withheld"
	case Internal:
		return "internal_error"
	case Database:
		return "database_error"
	case Validation:
		return "input_validation_error"
	case Unanticipated:
		return "unanticipated_error"
	case InvalidRequest:
		return "invalid_request_error"
	case Unauthenticated:
		return "unauthenticated_request"
	case Unauthorized:
		return "unauthorized_request"
	}
	return "unknown_error_kind"
}

// --

type Error struct {
	// Kind is the class of error, such as permission failure,
	// or "Other" if its class is unknown or irrelevant.
	Kind Kind
	// Code is a human-readable, short representation of the error,
	// for example "invalid_date_format"
	Code string
	// Param represents the parameter related to the error for example "release_date"
	Param string
	// The underlying error that triggered this one, if any, for example
	// "parsing time \"1984a-03-02T00:00:00Z\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"a-03-02T00:00:00Z\" as \"-\""
	Inner error
}

func (e Error) Unwrap() error {
	return e.Inner
}

func (e *Error) Error() string {
	return e.Inner.Error()
}

// --

type ErrorResponse struct {
	Error ServiceError `json:"error"`
}

type ServiceError struct {
	Kind    string `json:"kind,omitempty"`
	Code    string `json:"code,omitempty"`
	Param   string `json:"param,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewErrorResponse(error *Error) ErrorResponse {
	const msg string = "Internal Server Error"
	switch error.Kind {
	case Internal, Database:
		return ErrorResponse{
			Error: ServiceError{
				Kind:    Internal.String(),
				Message: msg,
			},
		}
	default:
		return ErrorResponse{
			Error: ServiceError{
				Kind:    error.Kind.String(),
				Code:    string(error.Code),
				Param:   string(error.Param),
				Message: error.Error(),
			},
		}
	}
}
