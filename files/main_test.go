package combine_numbers

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestJoinNumbers(t *testing.T) {
    assert.Equal(t, "", joinNumbers([]int{}))
    assert.Equal(t, "12345", joinNumbers([]int{1, 23, 45}))
}

var largestTestCases = []struct {
    slice    []int
    expected string
} {
    {[]int{5, 50, 56}, "56550"},
    {[]int{50, 2, 1, 9}, "95021"},
    {[]int{420, 42, 423}, "42423420"},
}

func TestLargest(t *testing.T) {
    for _, r := range largestTestCases {
        assert.Equal(t, r.expected, Largest(r.slice), fmt.Sprintf("%v", r.slice))
    }
}