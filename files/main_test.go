package combine_numbers

import (
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
    message  string
} {
    {[]int{50, 2, 1, 9}, "95021", "[50, 2, 1, 9]"},
}

func TestLargest(t *testing.T) {
    for _, r := range largestTestCases {
        assert.Equal(t, r.expected, Largest(r.slice), r.message)
    }
}