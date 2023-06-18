package runtime

import (
	"github.com/dzungtran/simple-flow/sdk"
	"github.com/dzungtran/simple-flow/workflow"
)

type FlowExecutor struct {
	gateway     string
	flowName    string // the name of the function
	reqID       string // the request id
	CallbackURL string // the callback url

	IsLoggingEnabled bool
	partialState     []byte

	// Logger  sdk.Logger
	Handler FlowDefinitionHandler
	Runtime *FlowRuntime
}

type FlowDefinitionHandler func(flow *workflow.Workflow, context *workflow.Context) error

func (fe *FlowExecutor) HandleNextNode(partial *Request) error {
	// var err error
	// request := &runtime.Request{}
	// request.Body, err = partial.Encode()
	// if err != nil {
	// 	return fmt.Errorf("failed to encode partial state, error %v", err)
	// }
	// request.RequestID = fe.reqID
	// request.FlowName = fe.flowName
	// request.Header = make(map[string][]string)

	// err = fe.Runtime.EnqueuePartialRequest(request)
	// if err != nil {
	// 	return fmt.Errorf("failed to enqueue request, error %v", err)
	// }
	return nil
}

// func (fe *FlowExecutor) GetExecutionOption(_ sdk.Operation) map[string]interface{} {
// 	options := make(map[string]interface{})
// 	return options
// }

func (fe *FlowExecutor) HandleExecutionCompletion(data []byte) error {
	return nil
}

func (fe *FlowExecutor) GetFlowName() string {
	return fe.flowName
}

func (fe *FlowExecutor) GetFlowDefinition(pipeline *sdk.Pipeline, context *sdk.Context) error {
	workflow := workflow.GetWorkflow(pipeline)
	return fe.Handler(workflow, context)
}

func (fe *FlowExecutor) Init(request interface{}) error {
	return nil
}
