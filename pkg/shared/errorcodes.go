package shared

type ErrorCode int
type ErrorMessage string

const (
	MANDATORY_HEADERS_MISSING_CODE ErrorCode = 1000
)

const (
	MANDATORY_HEADERS_MISSING_MESSAGE ErrorMessage = "Mandatory headers are missing"
)
