package tegrastats

import (
	"reflect"
	"runtime"
	"slices"
	"strings"
	"testing"
)

type TestParameter[T comparable] struct {
	Input    string
	Expected T
}

func getTestName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func ParseTest[T comparable](t *testing.T, parseFunc func(string) (*T, error), parameters []TestParameter[T]) {
	testName := getTestName(parseFunc)

	for _, test := range parameters {
		result, err := parseFunc(test.Input)
		if err != nil {
			t.Errorf("%s(%q) => err %s", testName, test.Input, err)
		}

		if *result != test.Expected {
			t.Errorf("%s(%q) => %v, expected %v", testName, test.Input, result, test.Expected)
		}
	}

	result, err := parseFunc("")
	if err == nil {
		t.Errorf("%s(\"\") => %q, expected err", testName, result)
	}
}

type TestParameterArr[T comparable] struct {
	Input    string
	Expected []T
}

func ParseTestArr[T comparable](t *testing.T, parseFunc func(string) ([]T, error), parameters []TestParameterArr[T]) {
	testName := getTestName(parseFunc)

	for _, test := range parameters {
		result, err := parseFunc(test.Input)
		if err != nil {
			t.Errorf("%s(%q) => err %s", testName, test.Input, err)
		}

		if !slices.Equal(result, test.Expected) {
			t.Errorf("%s(%q) => %v, expected %v", testName, test.Input, result, test.Expected)
		}
	}

	result, err := parseFunc("")
	if err == nil {
		t.Errorf("%s(\"\") => %v, expected err", testName, result)
	}
}
