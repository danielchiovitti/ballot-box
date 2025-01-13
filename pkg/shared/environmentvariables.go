package shared

type EnvironmentVariables string

const (
	TIMEOUT EnvironmentVariables = "TIMEOUT"
	PORT    EnvironmentVariables = "PORT"

	REDIS_ADDRESS   EnvironmentVariables = "REDIS_ADDRESS"
	REDIS_DATABASE  EnvironmentVariables = "REDIS_DATABASE"
	REDIS_PROTOCOL  EnvironmentVariables = "REDIS_PROTOCOL"
	REDIS_PASSWORD  EnvironmentVariables = "REDIS_PASSWORD"
	REDIS_PORT      EnvironmentVariables = "REDIS_PORT"
	REDIS_POOL_SIZE EnvironmentVariables = "REDIS_POOL_SIZE"

	REDIS_BLOOM_ADDRESS   EnvironmentVariables = "REDIS_BLOOM_ADDRESS"
	REDIS_BLOOM_DATABASE  EnvironmentVariables = "REDIS_BLOOM_DATABASE"
	REDIS_BLOOM_PROTOCOL  EnvironmentVariables = "REDIS_BLOOM_PROTOCOL"
	REDIS_BLOOM_PASSWORD  EnvironmentVariables = "REDIS_BLOOM_PASSWORD"
	REDIS_BLOOM_PORT      EnvironmentVariables = "REDIS_BLOOM_PORT"
	REDIS_BLOOM_POOL_SIZE EnvironmentVariables = "REDIS_BLOOM_POOL_SIZE"

	RATE_MAX_REQ EnvironmentVariables = "RATE_MAX_REQ"
	RATE_WINDOW  EnvironmentVariables = "RATE_WINDOW"

	RATE_GLOBAL_MAX_REQ EnvironmentVariables = "RATE_GLOBAL_MAX_REQ"
	RATE_GLOBAL_WINDOW  EnvironmentVariables = "RATE_GLOBAL_WINDOW"

	BLOOM_PRECISION EnvironmentVariables = "BLOOM_PRECISION"
	BLOOM_NAME      EnvironmentVariables = "BLOOM_NAME"
	BLOOM_INITIAL   EnvironmentVariables = "BLOOM_INITIAL"

	OLTP_STREAM_NAME       EnvironmentVariables = "OLTP_STREAM_NAME"
	OLTP_STREAM_GROUP_NAME EnvironmentVariables = "OLTP_STREAM_GROUP_NAME"

	OLAP_STREAM_NAME       EnvironmentVariables = "OLAP_STREAM_NAME"
	OLAP_STREAM_GROUP_NAME EnvironmentVariables = "OLAP_STREAM_GROUP_NAME"
)
