package model

type DatabaseOptions struct {
	Host               string
	Port               int
	DatabaseName       string
	User               string
	Password           string
	MinPoolSize        int
	MaxPoolSize        int
	MaxIdleTimeMS      int
	ConnectTimeoutMS   int
	WaitQueueTimeoutMS int
	AuthSource         string
}

type DatabaseOptionsFunc func(opt *DatabaseOptions)
