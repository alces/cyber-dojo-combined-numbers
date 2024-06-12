package combine_numbers

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestWithoutItem(t *testing.T) {
    assert.Equal(t, []int{1}, withoutItem([]int{1,2}, 1))
}
