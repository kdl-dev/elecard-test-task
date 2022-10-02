package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"strings"

	"github.com/kdl-dev/elecard-test-task/pkg/input"
	"github.com/kdl-dev/elecard-test-task/pkg/models"
)

type CLI struct {
	authKey string
}

func NewCLIClient(authKey string) *CLI {
	return &CLI{authKey: authKey}
}

func (cli *CLI) GetTasks(url string, params interface{}) (*input.GetTasksResponse, error) {

	body := input.Request{
		Key:    cli.authKey,
		Method: GetTasks,
		Params: &params,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	if resp.Status != "200 OK" {
		return nil, fmt.Errorf("http status code is not 200 OK (%s)", resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := new(input.GetTasksResponse)

	err = json.Unmarshal(respBody, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cli *CLI) CheckResults(url string, params interface{}) (*input.ResultTestsResponse, error) {

	body := input.Request{
		Key:    cli.authKey,
		Method: CheckResults,
		Params: &params,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	if resp.Status != "200 OK" {
		return nil, fmt.Errorf("http status code is not 200 OK (%s)", resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var results input.ResultTestsResponse

	if err = json.Unmarshal(respBody, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

func (cli *CLI) CalculateRectangles(resp *input.GetTasksResponse) (*[]input.TestResults, error) {

	var result = make([]input.TestResults, len(*resp.Tests))

	for index, test := range *resp.Tests {

		arrSquaresRT := calculateArrSquaresRT(&test)
		arrSquaresLB := calculateArrSquaresLB(&test)

		lb, err := calculateGlobalLeftBottom(*arrSquaresLB)
		if err != nil {
			return nil, err
		}

		rt, err := calculateGlobalRightTop(*arrSquaresRT)
		if err != nil {
			return nil, err
		}

		result[index] = input.TestResults{
			Left_bottom: *lb,
			Right_top:   *rt,
		}
	}
	return &result, nil
}

func calculateArrSquaresLB(test *input.TestTasks) *[]models.Coordinates {
	squares_left_bottom := make([]models.Coordinates, 0, len(*test))

	for _, circle := range *test {
		bigFloatX, precisionX := getBigFloat(circle.X.String())
		bigFloatY, precisionY := getBigFloat(circle.Y.String())
		bigFloatR, precisionR := getBigFloat(circle.R.String())

		leftBottomX := bigFloatX.Sub(bigFloatX, bigFloatR)
		leftBottomY := bigFloatY.Sub(bigFloatY, bigFloatR)

		squares_left_bottom = append(squares_left_bottom, models.Coordinates{
			X: json.Number(leftBottomX.Text('f', int(math.Max(float64(precisionX), float64(precisionR))))),
			Y: json.Number(leftBottomY.Text('f', int(math.Max(float64(precisionY), float64(precisionR))))),
		})
	}

	return &squares_left_bottom
}

func calculateArrSquaresRT(test *input.TestTasks) *[]models.Coordinates {
	squares_right_top := make([]models.Coordinates, 0, len(*test))

	for _, circle := range *test {
		bigFloatX, precisionX := getBigFloat(circle.X.String())
		bigFloatY, precisionY := getBigFloat(circle.Y.String())
		bigFloatR, precisionR := getBigFloat(circle.R.String())

		rightTopX := bigFloatX.Add(bigFloatX, bigFloatR)
		rightTopY := bigFloatY.Add(bigFloatY, bigFloatR)

		squares_right_top = append(squares_right_top, models.Coordinates{
			X: json.Number(rightTopX.Text('f', int(math.Max(float64(precisionX), float64(precisionR))))),
			Y: json.Number(rightTopY.Text('f', int(math.Max(float64(precisionY), float64(precisionR))))),
		})
	}

	return &squares_right_top
}

func calculateGlobalLeftBottom(points []models.Coordinates) (*models.Coordinates, error) {

	var bfPoint_X *big.Float
	var bfPoint_Y *big.Float
	min_x, _ := getBigFloat(points[0].X.String())
	min_y, _ := getBigFloat(points[0].Y.String())

	for _, point := range points {
		bfPoint_X, _ = getBigFloat(point.X.String())
		if bfPoint_X.Cmp(min_x) == -1 {
			min_x = bfPoint_X
		}

		bfPoint_Y, _ = getBigFloat(point.Y.String())
		if bfPoint_Y.Cmp(min_y) == -1 {
			min_y = bfPoint_Y
		}
	}

	return &models.Coordinates{
		X: json.Number(min_x.Text('f', -1)),
		Y: json.Number(min_y.Text('f', -1)),
	}, nil
}

func calculateGlobalRightTop(points []models.Coordinates) (*models.Coordinates, error) {

	var bfPoint_X *big.Float
	var bfPoint_Y *big.Float

	max_x, _ := getBigFloat(points[0].X.String())
	max_y, _ := getBigFloat(points[0].Y.String())

	for _, point := range points {
		bfPoint_X, _ = getBigFloat(point.X.String())
		if bfPoint_X.Cmp(max_x) == 1 {
			max_x = bfPoint_X
		}

		bfPoint_Y, _ = getBigFloat(point.Y.String())
		if bfPoint_Y.Cmp(max_y) == 1 {
			max_y = bfPoint_Y
		}
	}

	return &models.Coordinates{
		X: json.Number(max_x.Text('f', -1)),
		Y: json.Number(max_y.Text('f', -1)),
	}, nil
}

func getBigFloat(value string) (*big.Float, int) {

	var integerDigitCount int
	var fractionalDigitCount int
	if !strings.Contains(value, ".") {
		integerDigitCount = len(value)
		fractionalDigitCount = 0
	} else {
		integerDigitCount = strings.Index(value, ".")
		fractionalDigitCount = len(value) - (integerDigitCount + 1) // exclude the decimal from the count
	}

	precision := uint(math.Ceil(float64(fractionalDigitCount)*math.Log2(10.0)) + float64(integerDigitCount)*math.Log2(10.0))
	bf := big.NewFloat(0.0)
	bf.SetMode(bf.Mode()).SetPrec(precision + 10).SetString(value)

	return bf, fractionalDigitCount
}
