package main

import (
	"errors"
	"reflect"
	"testing"
)

// Helper function to check if two slices are equal
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// TestParseStringOfIntervalsToNumbersArrayOfIntervals tests the functionality of the ParseStringOfIntervalsToNumbersArrayOfIntervals function.
func TestParseStringOfIntervalsToNumbersArrayOfIntervals(t *testing.T) {
	tests := []struct {
		input  string
		output []int
		err    error
	}{
		{
			input:  "[25,30] [17,23] [5,24] [33,48] [2,18]",
			output: []int{25, 30, 17, 23, 5, 24, 33, 48, 2, 18},
			err:    nil,
		},
		{
			input:  "[1,10001]",
			output: nil,
			err:    errors.New("numbers cannot be bigger than 10,000"),
		},
		{
			input:  "[30,25]",
			output: nil,
			err:    errors.New("first number in the pair must be smaller than the second"),
		},
		{
			input:  "invalid string format",
			output: nil,
			err:    errors.New("input format is invalid"),
		},
		{
			input:  "[1,2] [3,4] [5,6] [7,8] [9,10] [11,12] [13,14] [15,16] [17,18] [19,20] [21,22] [23,24] [25,26] [27,28] [29,30] [31,32] [33,34] [35,36] [37,38] [39,40] [41,42]",
			output: nil,
			err:    errors.New("no more than 20 pairs are allowed"),
		},
	}

	for _, tt := range tests {
		result, err := ParseStringOfIntervalsToNumbersArrayOfIntervals(tt.input)
		if !reflect.DeepEqual(result, tt.output) || !reflect.DeepEqual(err, tt.err) {
			t.Errorf("ParseStringOfIntervalsToNumbersArrayOfIntervals(%q) => %v, %v; want %v, %v", tt.input, result, err, tt.output, tt.err)
		}
	}
}

// TestFindSmallestIntervalAndRemoveFromSlice tests the functionality of the FindSmallestIntervalAndRemoveFromSlice function.
func TestFindSmallestIntervalAndRemoveFromSlice(t *testing.T) {
	tests := []struct {
		input           []int
		expectedFirst   int
		expectedSecond  int
		expectedNewNums []int
	}{
		{[]int{25, 30, 2, 19, 14, 23, 4, 8}, 2, 19, []int{25, 30, 14, 23, 4, 8}},
		{[]int{}, 0, 0, []int{}},
		{[]int{3, 5}, 3, 5, []int{}},
		{[]int{10, 20, 30, 40, 50, 60}, 10, 20, []int{30, 40, 50, 60}},
	}

	for _, test := range tests {
		gotFirst, gotSecond, gotNewNums := FindSmallestIntervalAndRemoveFromSlice(test.input)

		if gotFirst != test.expectedFirst || gotSecond != test.expectedSecond || !equalSlices(gotNewNums, test.expectedNewNums) {
			t.Errorf("For input %v, expected (%d, %d, %v) but got (%d, %d, %v)",
				test.input, test.expectedFirst, test.expectedSecond, test.expectedNewNums,
				gotFirst, gotSecond, gotNewNums)
		}
	}
}

// TestIntervalMerger tests the functionality of the IntervalMerger function.
func TestIntervalMerger(t *testing.T) {
	tests := []struct {
		name         string
		firstInput   int
		secondInput  int
		newNumsInput []int
		wantFirst    int
		wantSecond   int
		wantNewNums  []int
	}{
		{
			name:         "Test case 1: Basic input",
			firstInput:   10,
			secondInput:  20,
			newNumsInput: []int{15, 25, 21, 23, 22, 24, 60, 70},
			wantFirst:    10,
			wantSecond:   25,
			wantNewNums:  []int{60, 70},
		},
		{
			name:         "Test case 2: No interval is merged",
			firstInput:   10,
			secondInput:  12,
			newNumsInput: []int{15, 25, 21, 23, 22, 24},
			wantFirst:    10,
			wantSecond:   12,
			wantNewNums:  []int{15, 25, 21, 23, 22, 24},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirst, gotSecond, gotNewNums := IntervalMerger(tt.firstInput, tt.secondInput, tt.newNumsInput)
			if gotFirst != tt.wantFirst || gotSecond != tt.wantSecond || !reflect.DeepEqual(gotNewNums, tt.wantNewNums) {
				t.Errorf("IntervalMerger() = (%v, %v, %v), want (%v, %v, %v)", gotFirst, gotSecond, gotNewNums, tt.wantFirst, tt.wantSecond, tt.wantNewNums)
			}
		})
	}
}

func TestMERGE(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "[2,5] [10,15]",
			output: "[2,5] [10,15]",
		},
		{
			input:  "[2,7] [4,10]",
			output: "[2,10]",
		},
		{
			input:  "[2,7] [2,7]",
			output: "[2,7]",
		},
		{
			input:  "",
			output: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := MERGE(tt.input)
			if got != tt.output {
				t.Errorf("Expected: %s, Got: %s", tt.output, got)
			}
		})
	}
}
