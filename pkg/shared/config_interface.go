package shared

type ConfigInterface interface {
	GetTimeOut() int
	GetRateMaxReq() int
	GetRateWindow() int
	GetRateGlobalMaxReq() int
	GetRateGlobalWindow() int
}
