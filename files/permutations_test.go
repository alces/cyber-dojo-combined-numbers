package combine_numbers

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
    assert.Equal(t, [][]int{{1}}, permutations([]int{1})
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
