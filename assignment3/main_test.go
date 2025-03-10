package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	// test request not error
	_, err := GetData()
	if err != nil {
		t.Errorf("Error Get Data function %v", err)
	}

	// test format data
	input := "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone."
	expected := []string{"fatback", "t-bone", "t-bone", "pastrami", "t-bone", "pork", "meatloaf", "jowl", "enim", "bresaola", "t-bone"}
	actual := FormatData(input)
	if ok := reflect.DeepEqual(expected, actual); !ok {
		t.Errorf("Error Format Data Input: \"%s\", Expected: %+v, Actual: %+v", input, expected, actual)
	}

	// test convert to map
	inp := []string{"fatback", "t-bone", "t-bone", "pastrami", "t-bone", "pork", "meatloaf", "jowl", "enim", "bresaola", "t-bone"}
	exp := map[string]int{"bresaola": 1, "enim": 1, "fatback": 1, "jowl": 1, "meatloaf": 1, "pastrami": 1, "t-bone": 4, "pork": 1}
	act := ConvertToMap(inp)
	if ok := reflect.DeepEqual(exp, act); !ok {
		t.Errorf("Error Format Data Input: \"%s\", Expected: %+v, Actual: %+v", inp, exp, act)
	}
}
