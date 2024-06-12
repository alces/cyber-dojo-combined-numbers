package combine_numbers

func permutations(slice []int) (result [][]int) {
    size := len(slice)
    
    if size <= 1 {
        return append(result, slice)
    }
        
    for i := 0; i < size; i++ {
        for _, p := range permutations(withoutItem(slice, i)) {
            result = append(result, append([]int{slice[i]}, p...))
        }
    }    
    
    return
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
