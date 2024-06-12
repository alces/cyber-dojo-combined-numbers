package combine_numbers

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var permutationsResults = []struct {
    slice    []int
    expected [][]int
} {
    {
        []int{1}, [][]int{{1}},
    },
    {
        []int{1,2}, [][]int{{1,2}, {2,1}},
    },
    {
        []int{1,2,3}, [][]int{{1,2,3}, {1,3,2}, {2,1,3}, {2,3,1}, {3,1,2}, {3,2,1}},
    },
}

func TestPermutations(t *testing.T) {
    for _, r := range permutationsResults {
        actual := permutations(r.slice)
        
        assert.Len(t, actual, len(r.expected))
        
        for _, e := range r.expected {
            assert.Contains(t, actual, e)
        }
    }
}    

var withoutItemResults = []struct{
    slice    []int
    index    int
    expected []int
    message  string
} {
    {
        []int{1,2,3}, 2, []int{1,2}, "remove last",
    },
    {
        []int{1,2,3}, 0, []int{2,3}, "remove first",
    },
    {
        []int{1,2,3,4,5}, 2, []int{1,2,4,5}, "remove middle",
    },
    {
        []int{1,2,3}, 3, []int{1,2,3}, "index out of bound",
    },
}

func TestWithoutItem(t *testing.T) {
    for _, r := range withoutItemResults {
        assert.Equal(t, r.expected, withoutItem(r.slice, r.index))
    }
}
