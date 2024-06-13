package combine_numbers

import (
    "strconv"
    "strings"
)

func Largest(slice []int) (result string) {
    for _, p := range permutations(slice) {
        candidate := joinNumbers(p)
        if candidate > result {
            result = candidate
        }
    }
    
    return
}

func joinNumbers(slice []int) string {
    size := len(slice)
    buffer := make([]string, size)
    
    for i := 0; i < size; i++ {
        buffer[i] = strconv.Itoa(slice[i])
    }
    
    return strings.Join(buffer, "")
}