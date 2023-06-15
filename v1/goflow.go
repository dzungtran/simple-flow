package v1

import (
	"fmt"
	"time"

	"github.com/dzungtran/simple-flow/core/sdk"
)

type FlowService struct {
	WorkerConcurrency int
	RetryCount        int
	// Flows               map[string]runtime.FlowDefinitionHandler
	RequestReadTimeout  time.Duration
	RequestWriteTimeout time.Duration
	OpenTraceUrl        string
	Logger              sdk.Logger
	EnableMonitoring    bool
	DebugEnabled        bool
}

type Request struct {
	Body      []byte
	RequestId string
	Query     map[string][]string
	Header    map[string][]string
}

const (
	DefaultWorkerConcurrency = 2
	DefaultRetryCount        = 2
)

func (fs *FlowService) Execute(flowName string, req *Request) error {
	if flowName == "" {
		return fmt.Errorf("flowName must be provided to execute flow")
	}

	fs.ConfigureDefault()

	// request := &runtimePkg.Request{
	// 	Header:    req.Header,
	// 	RequestID: req.RequestId,
	// 	Body:      req.Body,
	// 	Query:     req.Query,
	// }

	// err := fs.runtime.Execute(flowName, request)
	// if err != nil {
	// 	return fmt.Errorf("failed to execute request, %v", err)
	// }

	return nil
}

func (fs *FlowService) Register(
	flowName string,
	// handler runtime.FlowDefinitionHandler,
) error {
	if flowName == "" {
		return fmt.Errorf("flow-name must not be empty")
	}
	// if handler == nil {
	// 	return fmt.Errorf("handler must not be nil")
	// }

	// if fs.Flows == nil {
	// 	fs.Flows = make(map[string]runtime.FlowDefinitionHandler)
	// }

	// if fs.Flows[flowName] != nil {
	// 	return fmt.Errorf("flow-name must be unique for each flow")
	// }

	// fs.Flows[flowName] = handler

	return nil
}

func (fs *FlowService) Start() error {
	// if len(fs.Flows) == 0 {
	// 	return fmt.Errorf("must register atleast one flow")
	// }

	fs.ConfigureDefault()
	errorChan := make(chan error)
	defer close(errorChan)
	if err := fs.initRuntime(); err != nil {
		return err
	}
	err := <-errorChan
	return err
}

func (fs *FlowService) StartWorker() error {
	fs.ConfigureDefault()

	errorChan := make(chan error)
	defer close(errorChan)
	if err := fs.initRuntime(); err != nil {
		return err
	}
	err := <-errorChan
	return fmt.Errorf("worker has stopped, error: %v", err)
}

func (fs *FlowService) ConfigureDefault() {
	if fs.RetryCount == 0 {
		fs.RetryCount = DefaultRetryCount
	}
}

func (fs *FlowService) initRuntime() error {
	// err := fs.runtime.Init()
	// if err != nil {
	// 	return err
	// }
	return nil
}
