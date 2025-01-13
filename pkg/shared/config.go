package shared

import "github.com/spf13/viper"

func NewConfig(viper *viper.Viper) ConfigInterface {
	return &Config{
		TimeOut:                 viper.GetInt(string(TIMEOUT)),
		RateMaxReq:              viper.GetInt(string(RATE_MAX_REQ)),
		RateWindow:              viper.GetInt(string(RATE_WINDOW)),
		RateGlobalMaxReq:        viper.GetInt(string(RATE_GLOBAL_MAX_REQ)),
		RateGlobalWindow:        viper.GetInt(string(RATE_GLOBAL_WINDOW)),
		BloomPrecision:          viper.GetFloat64(string(BLOOM_PRECISION)),
		BloomName:               viper.GetString(string(BLOOM_NAME)),
		BloomInitial:            viper.GetInt(string(BLOOM_INITIAL)),
		OltpStreamName:          viper.GetString(string(OLTP_STREAM_NAME)),
		OltpStreamGroupName:     viper.GetString(string(OLTP_STREAM_GROUP_NAME)),
		OlapStreamName:          viper.GetString(string(OLAP_STREAM_NAME)),
		OlapStreamGroupName:     viper.GetString(string(OLAP_STREAM_GROUP_NAME)),
		MongoDbHost:             viper.GetString(string(MONGODB_HOST)),
		MongoDbPort:             viper.GetInt(string(MONGODB_PORT)),
		MongoDbDatabaseName:     viper.GetString(string(MONGODB_DATABASE_NAME)),
		MongoDbUser:             viper.GetString(string(MONGODB_USER)),
		MongoDbPassword:         viper.GetString(string(MONGODB_PASSWORD)),
		MongoDbMinPoolSize:      viper.GetInt(string(MONGODB_MIN_POOL_SIZE)),
		MongoDbMaxPoolSize:      viper.GetInt(string(MONGODB_MAX_POOL_SIZE)),
		MongoDbMaxIdleTimeout:   viper.GetInt(string(MONGODB_MAX_IDLE_TIMEOUT)),
		MongoDbWaitQueueTimeout: viper.GetInt(string(MONGODB_WAIT_QUEUE_TIMEOUT)),
		MongoDbAuthSource:       viper.GetString(string(MONGODB_AUTH_SOURCE)),
		OltpConsumersQty:        viper.GetInt(string(OLTP_CONSUMERS_QTY)),
		OlapConsumersQty:        viper.GetInt(string(OLAP_CONSUMERS_QTY)),
	}
}

type Config struct {
	TimeOut                 int
	RateMaxReq              int
	RateWindow              int
	RateGlobalMaxReq        int
	RateGlobalWindow        int
	BloomPrecision          float64
	BloomName               string
	BloomInitial            int
	OltpStreamName          string
	OltpStreamGroupName     string
	OlapStreamName          string
	OlapStreamGroupName     string
	MongoDbHost             string
	MongoDbPort             int
	MongoDbDatabaseName     string
	MongoDbUser             string
	MongoDbPassword         string
	MongoDbMinPoolSize      int
	MongoDbMaxPoolSize      int
	MongoDbMaxIdleTimeout   int
	MongoDbWaitQueueTimeout int
	MongoDbAuthSource       string
	OltpConsumersQty        int
	OlapConsumersQty        int
}

func (c *Config) GetTimeOut() int {
	return c.TimeOut
}

func (c *Config) GetRateMaxReq() int {
	return c.RateMaxReq
}

func (c *Config) GetRateWindow() int {
	return c.RateWindow
}

func (c *Config) GetRateGlobalMaxReq() int {
	return c.RateGlobalMaxReq
}

func (c *Config) GetRateGlobalWindow() int {
	return c.RateGlobalWindow
}

func (c *Config) GetBloomPrecision() float64 {
	return c.BloomPrecision
}

func (c *Config) GetBloomName() string {
	return c.BloomName
}

func (c *Config) GetBloomInitial() int {
	return c.BloomInitial
}

func (c *Config) GetOltpStreamName() string {
	return c.OltpStreamName
}

func (c *Config) GetOltpStreamGroupName() string {
	return c.OltpStreamGroupName
}

func (c *Config) GetOlapStreamName() string {
	return c.OlapStreamName
}

func (c *Config) GetOlapStreamGroupName() string {
	return c.OlapStreamGroupName
}

func (c *Config) GetMongoDbHost() string {
	return c.MongoDbHost
}

func (c *Config) GetMongoDbPort() int {
	return c.MongoDbPort
}

func (c *Config) GetMongoDbDatabaseName() string {
	return c.MongoDbDatabaseName
}

func (c *Config) GetMongoDbUser() string {
	return c.MongoDbUser
}

func (c *Config) GetMongoDbPassword() string {
	return c.MongoDbPassword
}

func (c *Config) GetMongoDbMinPoolSize() int {
	return c.MongoDbMinPoolSize
}

func (c *Config) GetMongoDbMaxPoolSize() int {
	return c.MongoDbMaxPoolSize
}

func (c *Config) GetMongoDbMaxIdleTimeout() int {
	return c.MongoDbMaxIdleTimeout
}

func (c *Config) GetMongoDbWaitQueueTimeout() int {
	return c.MongoDbWaitQueueTimeout
}

func (c *Config) GetMongoDbAuthSource() string {
	return c.MongoDbAuthSource
}

func (c *Config) GetOltpConsumersQty() int {
	return c.OltpConsumersQty
}

func (c *Config) GetOlapConsumersQty() int {
	return c.OlapConsumersQty
}
