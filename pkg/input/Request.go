package input

import "github.com/kdl-dev/elecard-test-task/pkg/models"

type TestResults models.Rectangle

type Request struct {
	Key    string       `json:"key"`
	Method string       `json:"method"`
	Params *interface{} `json:"params"`
}
