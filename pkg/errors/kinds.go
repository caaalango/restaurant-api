package errs

import "net/http"

type kind uint8

const (
	UnauthorizedCredential kind = iota + 1
	InvalidCredential
	RepositoryError
	SupplierError
	ValidationError
	JSONError
	InvalidInput
	NotFound
	CacheMalfunction
	ParsingError
	UnexpectedError
	InternalError
	Unprocessable
)

func (k kind) String() string {
	switch k {
	case UnauthorizedCredential:
		return "UNAUTHORIZED_CREDENTIAL"
	case InvalidCredential:
		return "INVALID_CREDENTIAL"
	case RepositoryError:
		return "REPOSITORY_ERROR"
	case SupplierError:
		return "SUPPLIER_ERROR"
	case ValidationError:
		return "VALIDATION_ERROR"
	case JSONError:
		return "JSON_ERROR"
	case InvalidInput:
		return "INVALID_INPUT"
	case NotFound:
		return "NOT_FOUND"
	case CacheMalfunction:
		return "CACHE_ERROR"
	case ParsingError:
		return "PARSING_ERROR"
	case UnexpectedError:
		return "UNEXPECTED_ERROR"
	case InternalError:
		return "INTERNAL_ERROR"
	case Unprocessable:
		return "UNPROCESSABLE_ERROR"
	default:
		return "UNEXPECTED_ERROR"
	}
}

func (k kind) HttpStatus() int {
	switch k {
	case InvalidCredential:
		return http.StatusBadRequest
	case UnauthorizedCredential:
		return http.StatusUnauthorized
	case RepositoryError:
		return http.StatusInternalServerError
	case SupplierError:
		return http.StatusBadGateway
	case ValidationError:
		return http.StatusBadRequest
	case JSONError:
		return http.StatusInternalServerError
	case InvalidInput:
		return http.StatusBadRequest
	case NotFound:
		return http.StatusNotFound
	case CacheMalfunction:
		return http.StatusInternalServerError
	case ParsingError:
		return http.StatusInternalServerError
	case UnexpectedError:
		return http.StatusInternalServerError
	case InternalError:
		return http.StatusInternalServerError
	case Unprocessable:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
