package input

import "github.com/kdl-dev/elecard-test-task/pkg/models"

type TestTasks []models.Circle

type GetTasksResponse struct {
	Tests *[]TestTasks `json:"result"`
}

type ResultTestsResponse struct {
	Result []bool `json:"result"`
}
