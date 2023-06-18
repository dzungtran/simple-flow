package runtime

import (
	"fmt"

	"github.com/dzungtran/simple-flow/sdk"
	"github.com/dzungtran/simple-flow/sdk/executor"
)

type FlowRuntime struct {
	Flows  map[string]FlowDefinitionHandler
	Logger sdk.Logger
}

type Worker struct {
	ID          string   `json:"id"`
	Flows       []string `json:"flows"`
	Concurrency int      `json:"concurrency"`
}

type Task struct {
	FlowName    string              `json:"flow_name"`
	RequestID   string              `json:"request_id"`
	Body        string              `json:"body"`
	Header      map[string][]string `json:"header"`
	RawQuery    string              `json:"raw_query"`
	Query       map[string][]string `json:"query"`
	RequestType string              `json:"request_type"`
}

const (
	PartialRequest = "PARTIAL"
	NewRequest     = "NEW"
	// PauseRequest   = "PAUSE"
	// ResumeRequest  = "RESUME"
	// StopRequest    = "STOP"
)

func (fRuntime *FlowRuntime) Init() error {

	return nil
}

func (fRuntime *FlowRuntime) CreateExecutor(req *Request) (executor.Executor, error) {
	flowHandler, ok := fRuntime.Flows[req.FlowName]
	if !ok {
		return nil, fmt.Errorf("could not find handler for flow %s", req.FlowName)
	}
	ex := &FlowExecutor{
		Handler: flowHandler,
		Runtime: fRuntime,
	}
	err := ex.Init(req)
	return ex, err
}

func (fRuntime *FlowRuntime) Execute(flowName string, request *Request) error {

	// data, _ := json.Marshal(&Task{
	// 	FlowName:    flowName,
	// 	RequestID:   request.RequestID,
	// 	Body:        string(request.Body),
	// 	Header:      request.Header,
	// 	RawQuery:    request.RawQuery,
	// 	Query:       request.Query,
	// 	RequestType: NewRequest,
	// })

	return nil
}

func (fRuntime *FlowRuntime) EnqueuePartialRequest(pr *Request) error {
	// data, _ := json.Marshal(&Task{
	// 	FlowName:    pr.FlowName,
	// 	RequestID:   pr.RequestID,
	// 	Body:        string(pr.Body),
	// 	Header:      pr.Header,
	// 	RawQuery:    pr.RawQuery,
	// 	Query:       pr.Query,
	// 	RequestType: PartialRequest,
	// })
	return nil
}

func (fRuntime *FlowRuntime) handleRequest(request *Request, requestType string) error {
	var err error
	switch requestType {
	case PartialRequest:
		err = fRuntime.handlePartialRequest(request)
	case NewRequest:
		err = fRuntime.handleNewRequest(request)
	default:
		return fmt.Errorf("invalid request %v received with type %s", request, requestType)
	}
	return err
}

func (fRuntime *FlowRuntime) handleNewRequest(request *Request) error {
	flowExecutor, err := fRuntime.CreateExecutor(request)
	if err != nil {
		return fmt.Errorf("failed to execute request " + request.RequestID + ", error: " + err.Error())
	}

	response := &Response{}
	response.RequestID = request.RequestID
	response.Header = make(map[string][]string)

	err = ExecuteFlowHandler(response, request, flowExecutor)
	if err != nil {
		return fmt.Errorf("request failed to be processed. error: " + err.Error())
	}

	return nil
}

func (fRuntime *FlowRuntime) handlePartialRequest(request *Request) error {
	flowExecutor, err := fRuntime.CreateExecutor(request)
	if err != nil {
		fRuntime.Logger.Log(fmt.Sprintf("[request `%s`] failed to execute request, error: %v", request.RequestID, err))
		return fmt.Errorf("[goflow] failed to execute request " + request.RequestID + ", error: " + err.Error())
	}
	response := &Response{}
	response.RequestID = request.RequestID
	response.Header = make(map[string][]string)

	err = PartialExecuteFlowHandler(response, request, flowExecutor)
	if err != nil {
		fRuntime.Logger.Log(fmt.Sprintf("[request `%s`] failed to be processed. error: %v", request.RequestID, err.Error()))
		return fmt.Errorf("[goflow] request failed to be processed. error: " + err.Error())
	}
	return nil
}
