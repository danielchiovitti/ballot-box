package shared

import "github.com/spf13/viper"

func NewConfig(viper *viper.Viper) ConfigInterface {
	return &Config{
		TimeOut:          viper.GetInt(string(TIMEOUT)),
		RateMaxReq:       viper.GetInt(string(RATE_MAX_REQ)),
		RateWindow:       viper.GetInt(string(RATE_WINDOW)),
		RateGlobalMaxReq: viper.GetInt(string(RATE_GLOBAL_MAX_REQ)),
		RateGlobalWindow: viper.GetInt(string(RATE_GLOBAL_WINDOW)),
	}
}

type Config struct {
	TimeOut          int
	RateMaxReq       int
	RateWindow       int
	RateGlobalMaxReq int
	RateGlobalWindow int
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
