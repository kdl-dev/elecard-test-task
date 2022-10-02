package client

import (
	"testing"
)

func TestGetBigFloat(t *testing.T) {
	value1 := "-352265234985208769785987385345.14325798579835834545123154235"
	value2 := "-352265234985208769785987385345,14325798579835834545123154235"
	value3 := "Hi Elecard!"

	bigFloatValue, _, err := getBigFloat(value1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if bigFloatValue.Text('f', -1) != value1 {
		t.Errorf("unexpected error: %s != %s", bigFloatValue.Text('f', -1), value1)
	}

	_, _, err = getBigFloat(value2)
	if err == nil {
		t.Errorf("expected error, but none (invalid format for float number :%s)", value2)
	}

	_, _, err = getBigFloat(value3)
	if err == nil {
		t.Errorf("expected error, but none (invalid format for float number :%s)", value3)
	}

}
