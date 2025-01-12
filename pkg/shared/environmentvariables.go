package shared

type EnvironmentVariables string

const (
	TIMEOUT         EnvironmentVariables = "TIMEOUT"
	PORT            EnvironmentVariables = "PORT"
	REDIS_ADDRESS   EnvironmentVariables = "REDIS_ADDRESS"
	REDIS_DATABASE  EnvironmentVariables = "REDIS_DATABASE"
	REDIS_PROTOCOL  EnvironmentVariables = "REDIS_PROTOCOL"
	REDIS_PASSWORD  EnvironmentVariables = "REDIS_PASSWORD"
	REDIS_PORT      EnvironmentVariables = "REDIS_PORT"
	REDIS_POOL_SIZE EnvironmentVariables = "REDIS_POOL_SIZE"
	RATE_MAX_REQ    EnvironmentVariables = "RATE_MAX_REQ"
	RATE_WINDOW     EnvironmentVariables = "RATE_WINDOW"
)
