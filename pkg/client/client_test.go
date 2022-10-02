package client

import (
	"encoding/json"
	"testing"
)

func TestParseParams(t *testing.T) {
	params1 := "0,0,0,0"
	params2 := "0,0,0,0,1,1"
	params3 := "Hello Elecard!"

	testResults, err := ParseParams(params1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	for _, result := range *testResults {
		if result.Left_bottom.X != json.Number(params1[0]) ||
			result.Left_bottom.Y != json.Number(params1[2]) ||
			result.Right_top.X != json.Number(params1[4]) ||
			result.Right_top.Y != json.Number(params1[6]) {
			t.Errorf("parsing is incorrect, expected %s, in fact %v", params1, result)
		}
	}

	_, err = ParseParams(params2)
	if err == nil {
		t.Errorf("expected error, incorect format: %s, but none", params2)
	}

	_, err = ParseParams(params3)
	if err == nil {
		t.Errorf("expected error, incorect format: %s, but none", params3)
	}
}
