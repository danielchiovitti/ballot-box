package shared

type ErrorCode int
type ErrorMessage string

const (
	INTERNAL_ERROR_CODE                ErrorCode = 0
	MAX_REQ_LIMIT_EXCEEDED_CODE        ErrorCode = 100
	MAX_GLOBAL_REQ_LIMIT_EXCEEDED_CODE ErrorCode = 101
	MANDATORY_HEADERS_MISSING_CODE     ErrorCode = 1000
)

const (
	INTERNAL_ERROR_MESSAGE                ErrorMessage = "Internal Error"
	MAX_REQ_LIMIT_EXCEEDED                ErrorMessage = "Max Request limit Exceeded"
	MANDATORY_HEADERS_MISSING_MESSAGE     ErrorMessage = "Mandatory headers are missing"
	MAX_GLOBAL_REQ_LIMIT_EXCEEDED_MESSAGE ErrorMessage = "Max Global Request limit Exceeded"
)
