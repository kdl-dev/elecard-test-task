package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kdl-dev/elecard-test-task/pkg/input"
	"github.com/kdl-dev/elecard-test-task/pkg/models"
)

const GetTasks = "GetTasks"
const CheckResults = "CheckResults"
const AutoExecution = "AutoExec"

type Client interface {
	GetTasks(url string, params interface{}) (*input.GetTasksResponse, error)
	CalculateRectangles(resp *input.GetTasksResponse) (*[]input.TestResults, error)
	CheckResults(url string, params interface{}) (*input.ResultTestsResponse, error)
}

func Resolve(cl Client, url string, method string, params interface{}) error {
	switch method {
	case AutoExecution:
		return MethodAutoExecution(cl, url, params)
	case GetTasks:
		return MethodGetTasks(cl, url, params)
	case CheckResults:
		return MethodCheckResults(cl, url, params)
	default:
		return fmt.Errorf("unknown method \"%s\"", method)
	}
}

func MethodGetTasks(cl Client, url string, params interface{}) error {
	tasks, err := cl.GetTasks(url, nil)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", *tasks.Tests)

	return nil
}

func MethodCheckResults(cl Client, url string, params interface{}) error {
	_, ok := params.(string)
	if !ok {
		return errors.New("parameters must be of string type")
	}

	testResults, err := ParseParams(params.(string))
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	resp, err := cl.CheckResults(url, testResults)
	if err != nil {
		return err
	}

	for index, isDone := range resp.Result {
		fmt.Printf("Test #%d: %v\n", index+1, isDone)
	}

	return nil
}

func MethodAutoExecution(cl Client, url string, params interface{}) error {
	tasks, err := cl.GetTasks(url, nil)
	if err != nil {
		return err
	}

	testResults, err := cl.CalculateRectangles(tasks)
	if err != nil {
		return err
	}

	for index, testResult := range *testResults {
		fmt.Printf("Rectangle #%d :\n", index+1)
		fmt.Printf("\tleft bottom: %v;\n", testResult.Left_bottom)
		fmt.Printf("\tright top: %v;\n\n", testResult.Right_top)
	}

	resp, err := cl.CheckResults(url, testResults)
	if err != nil {
		return err
	}

	allIsDone := true
	for index, isDone := range resp.Result {
		if allIsDone && !isDone {
			allIsDone = !allIsDone
		}
		fmt.Printf("Test #%d: %v\n", index+1, isDone)
	}

	if allIsDone {
		fmt.Printf("\nEverything is OK!\n")
	}

	return nil
}

func ParseParams(params string) (*[]input.TestResults, error) {
	paramsArr := strings.Split(params, ",")
	if len(paramsArr)%4 != 0 {
		return nil, errors.New("not enough parameters")
	}

	var results []input.TestResults = make([]input.TestResults, 0, len(paramsArr)/2)

	for i := 0; i < len(paramsArr); i += 4 {

		if _, err := strconv.ParseFloat(paramsArr[i], 64); err != nil {
			return nil, err
		}

		if _, err := strconv.ParseFloat(paramsArr[i+1], 64); err != nil {
			return nil, err
		}

		if _, err := strconv.ParseFloat(paramsArr[i+2], 64); err != nil {
			return nil, err
		}

		if _, err := strconv.ParseFloat(paramsArr[i+3], 64); err != nil {
			return nil, err
		}

		results = append(results, input.TestResults{
			Left_bottom: models.Coordinates{
				X: json.Number(paramsArr[i]),
				Y: json.Number(paramsArr[i+1]),
			},
			Right_top: models.Coordinates{
				X: json.Number(paramsArr[i+2]),
				Y: json.Number(paramsArr[i+3]),
			},
		})
	}

	return &results, nil
}
