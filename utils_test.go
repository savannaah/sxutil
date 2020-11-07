package sxutil

import (
	"testing"
)

func TestContains(t *testing.T) {
	tt := []struct {
		name   string
		list   interface{}
		value  interface{}
		result bool
	}{
		{"", []string{"test1", "test2", "test3"}, "test4", false},
		{"", []string{"test1", "test2", "test3"}, "test3", true},
		{"", []int32{2, 3, 4}, int32(1), false},
		{"", []int32{4, 3, 2}, int32(4), true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := Contains(tc.list, tc.value)
			if actual != tc.result {
				t.Fatalf("Expected: %v; Actual: %v", tc.result, actual)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	tt := []struct {
		name   string
		a  interface{}
		b interface{}
		result bool
	}{
		{"", []string{"test1", "test2", "test3", "test3"}, []string{"test1", "test2", "test3"}, false},
		{"", []string{"test1", "test2", "test3"}, []string{"test3", "test2", "test1"}, true},
		{"", []string{"test1", "test2", "test3"}, []string{"test4", "test2", "test1"}, false},
		{"", []int32{2, 3, 4, 4}, []int32{2, 3, 4}, false},
		{"", []int32{4, 3, 2}, []int32{2, 4, 3}, true},
		{"", []int32{4, 3, 2}, []int32{2, 4, 4}, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := Equal(tc.a, tc.b)
			if actual != tc.result {
				t.Fatalf("Expected: %v; Actual: %v", tc.result, actual)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tt := []struct {
		name   string
		value  interface{}
		result interface{}
	}{
		{"", []string{"test1", "test2", "test3", "test3"}, []string{"test1", "test2", "test3"}},
		{"", []string{"test1", "test3", "test2", "test1"}, []string{"test3", "test2", "test1"}},
		{"", []int32{2, 3, 4}, []int32{2, 3, 4}},
		{"", []int32{4, 3, 2}, []int32{2, 4, 3}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := Unique(tc.value)
			if !Equal(actual, tc.result) {
				t.Fatalf("Expected: %v; Actual: %v", tc.result, actual)
			}
		})
	}
}

func TestMissing(t *testing.T) {
	tt := []struct {
		name   string
		a      interface{}
		b      interface{}
		result interface{}
	}{
		{"", []string{"test1", "test2"}, []string{"test1", "test2", "test3"}, []string{"test3"}},
		{"", []string{"test1", "test3"}, []string{"test3", "test1", "test2"}, []string{"test2"}},
		{"", []int32{2, 3}, []int32{2, 3, 4}, []int32{4}},
		{"", []int32{4, 3}, []int32{4, 3, 2}, []int32{2}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := Missing(tc.a, tc.b)
			if !Equal(actual, tc.result) {
				t.Fatalf("Expected: %v; Actual: %v", tc.result, actual)
			}
		})
	}
}

func TestUnmatched(t *testing.T) {
	tt := []struct {
		name   string
		a      interface{}
		b      interface{}
		result interface{}
	}{
		{"", []string{"test1", "test2", "test3"}, []string{"test1", "test2"}, []string{"test3"}},
		{"", []string{"test1", "test3", "test2"}, []string{"test3", "test1", "test5"}, []string{"test2", "test5"}},
		{"", []int32{2, 3, 4}, []int32{2, 3}, []int32{4}},
		{"", []int32{4, 3, 2}, []int32{4, 3, 5}, []int32{2, 5}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := Unmatched(tc.a, tc.b)
			if !Equal(actual, tc.result) {
				t.Fatalf("Expected: %v; Actual: %v", tc.result, actual)
			}
		})
	}
}

func TestCommon(t *testing.T) {
	tt := []struct {
		name   string
		a      interface{}
		b      interface{}
		result interface{}
	}{
		{"", []string{"test1", "test2", "test3"}, []string{"test1", "test2"}, []string{"test1","test2"}},
		{"", []string{"test1", "test3", "test2"}, []string{"test3", "test1", "test5"}, []string{"test1", "test3"}},
		{"", []int32{2, 3, 4}, []int32{2, 3}, []int32{2,3}},
		{"", []int32{4, 3, 2}, []int32{4, 3, 5}, []int32{4, 3}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := Common(tc.a, tc.b)
			if !Equal(actual, tc.result) {
				t.Fatalf("Expected: %v; Actual: %v", tc.result, actual)
			}
		})
	}
}

func TestCheckDecimalPlaces(t *testing.T) {
	tt := []struct {
		name   string
		a      int
		b      float64
		result bool
	}{
		{"", 3, 1.234, true},
		{"", 2, 1.234, false},
		{"", 2, 1.23, true},
		{"", 1, 1.2, true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := CheckDecimalPlaces(tc.a, tc.b)
			if actual != tc.result {
				t.Fatalf("Expected: %v; Actual: %v", tc.result, actual)
			}
		})
	}
}
