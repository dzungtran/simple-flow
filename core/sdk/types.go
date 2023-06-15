package sdk

// Logger logs the flow logs
type Logger interface {
	// Configure configure a logger with flowname and requestID
	Configure(flowName string, requestId string)
	// Init initialize a logger
	Init() error
	// Log logs a flow log
	Log(str string)
}
