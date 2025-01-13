package shared

import "github.com/spf13/viper"

func NewConfig(viper *viper.Viper) ConfigInterface {
	return &Config{
		TimeOut:             viper.GetInt(string(TIMEOUT)),
		RateMaxReq:          viper.GetInt(string(RATE_MAX_REQ)),
		RateWindow:          viper.GetInt(string(RATE_WINDOW)),
		RateGlobalMaxReq:    viper.GetInt(string(RATE_GLOBAL_MAX_REQ)),
		RateGlobalWindow:    viper.GetInt(string(RATE_GLOBAL_WINDOW)),
		BloomPrecision:      viper.GetFloat64(string(BLOOM_PRECISION)),
		BloomName:           viper.GetString(string(BLOOM_NAME)),
		BloomInitial:        viper.GetInt(string(BLOOM_INITIAL)),
		OltpStreamName:      viper.GetString(string(OLTP_STREAM_NAME)),
		OltpStreamGroupName: viper.GetString(string(OLTP_STREAM_GROUP_NAME)),
		OlapStreamName:      viper.GetString(string(OLAP_STREAM_NAME)),
		OlapStreamGroupName: viper.GetString(string(OLAP_STREAM_GROUP_NAME)),
	}
}

type Config struct {
	TimeOut             int
	RateMaxReq          int
	RateWindow          int
	RateGlobalMaxReq    int
	RateGlobalWindow    int
	BloomPrecision      float64
	BloomName           string
	BloomInitial        int
	OltpStreamName      string
	OltpStreamGroupName string
	OlapStreamName      string
	OlapStreamGroupName string
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
