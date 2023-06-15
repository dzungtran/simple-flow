package sdk

import (
	"net/url"
)

// Context execution context and execution state
type Context struct {
	requestId string     // the request id
	node      string     // the execution position
	Query     url.Values // provides request Query
	State     string     // state of the request
	Name      string     // name of the faas-flow

	NodeInput map[string][]byte // stores inputs form each node
}

const (
	// StateSuccess denotes success state
	StateSuccess = "success"
	// StateFailure denotes failure state
	StateFailure = "failure"
	// StateOngoing denotes ongoing state
	StateOngoing = "ongoing"
)

// CreateContext create request context (used by template)
func CreateContext(id string, node string, name string) *Context {

	context := &Context{}
	context.requestId = id
	context.node = node
	context.Name = name
	context.State = StateOngoing
	context.NodeInput = make(map[string][]byte)

	return context
}

// GetRequestId returns the request id
func (context *Context) GetRequestId() string {
	return context.requestId
}

// GetPhase return the node no
func (context *Context) GetNode() string {
	return context.node
}
