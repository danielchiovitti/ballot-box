package shared

import "github.com/spf13/viper"

func NewConfig(viper *viper.Viper) ConfigInterface {
	return &Config{
		TimeOut:    viper.GetInt(string(TIMEOUT)),
		RateMaxReq: viper.GetInt(string(RATE_MAX_REQ)),
		RateWindow: viper.GetInt(string(RATE_WINDOW)),
	}
}

type Config struct {
	TimeOut    int
	RateMaxReq int
	RateWindow int
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
