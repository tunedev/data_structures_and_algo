package bubblesort

import (
    "reflect"
    "testing"
)

func TestBubbleSort(t *testing.T) {
    testCases := []struct {
        name string
        input []int
        expected []int
    }{
        {
            name: "sorts an empty array",
            input: []int{},
            expected: []int{},
        },
        {
            name: "sorts an array with one element",
            input: []int{1},
            expected: []int{1},
        },
        {
            name: "sorts an array with multiple elements",
            input: []int{5, 3, 1, 2, 4},
            expected: []int{1, 2, 3, 4, 5},
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            BubbleSort(tc.input)
            if !reflect.DeepEqual(tc.input, tc.expected) {
                t.Errorf("BubbleSort(%v) = %v; expected %v", tc.input, tc.input, tc.expected)
            }
        })
    }
}
