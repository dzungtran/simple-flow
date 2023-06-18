package runtime

import (
	"errors"
	"fmt"
	"log"

	"github.com/dzungtran/simple-flow/sdk/executor"
)

func ExecuteFlowHandler(response *Response, request *Request, ex executor.Executor) error {
	log.Printf("Executing flow %s\n", request.FlowName)

	var stateOption executor.ExecutionStateOption
	flowExecutor := executor.CreateFlowExecutor(ex, nil)
	resp, err := flowExecutor.Execute(stateOption)
	if err != nil {
		return fmt.Errorf("failed to execute request. %s", err.Error())
	}

	return nil
}

func PartialExecuteFlowHandler(response *Response, request *Request, ex executor.Executor) error {

	var stateOption executor.ExecutionStateOption

	if request.RequestID == "" {
		return errors.New("request ID must be set in partial request")
	}
	partialState, err := executor.DecodePartialReq(request.Body)
	if err != nil {
		return errors.New("failed to decode partial state")
	}
	stateOption = executor.PartialRequest(partialState)

	// Create a flow executor with provided executor
	flowExecutor := executor.CreateFlowExecutor(ex, nil)
	resp, err := flowExecutor.Execute(stateOption)
	if err != nil {
		return fmt.Errorf("failed to execute request. %s", err.Error())
	}

	response.Body = resp

	return nil
}
