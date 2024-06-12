package combine_numbers

import (
    "fmt"
)

func permutations(slice []int) [][]int {
    size := len(slice)
    
    if size <= 1 {
        return [][]int{slice}
    }
    
    result := [][]int{}
    
    for i := 0; i < size; i++ {
        head := []int{slice[i]}
        tail := withoutItem(slice, i)
        
        fmt.Printf("slice: %v, i: %v, head: %v, tail: %v\n", slice, i, head, tail)
        
        for _, p := range permutations(tail) {
            result = append(result, append(head, p...))
        }
    }    
    
    return result
}

func withoutItem(slice []int, index int) []int {
    size := len(slice)
    result := make([]int, size)
    
    copy(result, slice)
    
    if index >= size {
        return result
    }
    
    return append(result[:index], result[index + 1:]...)
}
