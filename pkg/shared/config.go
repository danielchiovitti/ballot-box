package shared

import "github.com/spf13/viper"

func NewConfig(viper *viper.Viper) ConfigInterface {
	return &Config{
		TimeOut: viper.GetInt(string(TIMEOUT)),
	}
}

type Config struct {
	TimeOut int
}

func (c *Config) GetTimeOut() int {
	return c.TimeOut
}
