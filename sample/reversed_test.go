package sample

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"multiple elements", []int{1, 2, 3}, []int{3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Reverse(%v) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
