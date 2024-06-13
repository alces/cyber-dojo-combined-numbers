package combine_numbers

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestJoinNumbers(t *testing.T) {
    assert.Equal(t, "", joinNumbers([]int{}))
    assert.Equal(t, "12345", joinNumbers([]int{1, 23, 45}))
}