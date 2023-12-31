package result

import (
	"errors"
	"testing"
)

func TestResultWorkflow(t *testing.T) {
	testError := errors.New("test")

	type output any
	type args[O output] struct {
		o   O
		err error
	}

	for _, tt := range []struct {
		name string
		args args[output]
		want *Result[output]
	}{
		{"test#1", args[output]{1, nil}, &Result[output]{state: Success, value: 1}},
		{"test#2", args[output]{nil, testError}, &Result[output]{state: Failure, fault: testError}},
		{"test#3", args[output]{nil, nil}, &Result[output]{state: Success, value: nil}},
		{"test#4", args[output]{1, testError}, &Result[output]{state: Failure, fault: testError, value: 1}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			result := GetResult(func() (output, error) { return tt.args.o, tt.args.err })
			if !result.Equals(*tt.want) {
				t.Errorf("GetResult() = %v, want %v", result, tt.want)
			}
		})
	}

}
