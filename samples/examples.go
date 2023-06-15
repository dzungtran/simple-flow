package main

import (
	"fmt"
	"github.com/dzungtran/simple-flow/samples/condition"
	"github.com/dzungtran/simple-flow/samples/loop"
	"github.com/dzungtran/simple-flow/samples/parallel"
	"github.com/dzungtran/simple-flow/samples/serial"
	"github.com/dzungtran/simple-flow/samples/single"

	goflow "github.com/dzungtran/simple-flow/v1"
)

func main() {
	fs := &goflow.FlowService{
		Port:              8080,
		RedisURL:          "localhost:6379",
		OpenTraceUrl:      "localhost:5775",
		WorkerConcurrency: 5,
		EnableMonitoring:  true,
		DebugEnabled:      true,
	}
	fs.Register("single", single.DefineWorkflow)
	fs.Register("serial", serial.DefineWorkflow)
	fs.Register("parallel", parallel.DefineWorkflow)
	fs.Register("condition", condition.DefineWorkflow)
	fs.Register("loop", loop.DefineWorkflow)
	fmt.Println(fs.Start())
}
