package combine_numbers

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var withoutItemResults = []struct{
    slice    []int
    index    int
    expected []int
    message  string
} {
    {
        []int{1,2}, 1, []int{1}, "remove last",
    },
    {
        []int{1,2}, 0, []int{1}, "remove first",
    },
}

func TestWithoutItem(t *testing.T) {
    for _, r := range withoutItemResults {
        assert.Equal(t, r.expected, withoutItem(r.slice, r.index))
    }
}
