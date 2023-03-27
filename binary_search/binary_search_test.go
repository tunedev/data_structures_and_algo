package binarysearch

import (
	"reflect"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	testCases := []struct{
		name string
		haystack []int
		needle int
		expected bool
	}{
		{
			name: "Test case 1: Search for an element in an array with one element",
			haystack: []int{5},
			needle: 5,
			expected: true,
		},
		{
			name: "Test case 2: Search for an element in an array with multiple elements",
			haystack: []int{1,2,3,4,5},
			needle: 3,
			expected: true,
		},
		{
			name: "Test case 3: Search for an element that is not in the array",
			haystack: []int{1,2,3,4,5},
			needle: 6,
			expected: false,
		},
		{
			name: "Test case 4: Search for an element at the beginning of the array",
			haystack: []int{1,2,3,4,5},
			needle: 1,
			expected: true,
		},
		{
			name: "Test case 5: Search for an element at the end of the array",
			haystack: []int{1,2,3,4,5},
			needle: 5,
			expected: true,
		},
		{
			name: "Test case 6: Search for an element in an empty array",
			haystack: []int{},
			needle: 5,
			expected: false,
		},
	}

	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T){
			actual := BinarySearch(tc.haystack, tc.needle)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf(`
				BinarySearch(%v, %v): 
				expected: %v
				actual: %v
				`, tc.haystack, tc.needle, tc.expected, actual)
			}
		})
	}
}