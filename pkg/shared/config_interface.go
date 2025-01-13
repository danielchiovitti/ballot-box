package shared

type ConfigInterface interface {
	GetTimeOut() int
	GetRateMaxReq() int
	GetRateWindow() int
	GetRateGlobalMaxReq() int
	GetRateGlobalWindow() int
	GetBloomPrecision() float64
	GetBloomName() string
	GetBloomInitial() int
	GetOltpStreamName() string
	GetOltpStreamGroupName() string
	GetOlapStreamName() string
	GetOlapStreamGroupName() string
}
