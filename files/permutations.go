package combine_numbers

func permutations(slice []int) [][]int {
    size := len(slice)
    
    if size <= 1 {
        return [][]int{slice}
    }
    
    result := [][]int{}
    
    for i := 0; i < size; i++ {
        head := []int{slice[i]}
        tail := withoutItem(slice, i)
        
        for _, p := range permutations(tail) {
            result = append(result, append(head, p))
        }
    }    
    
    return result
}

func withoutItem(slice []int, index int) []int {
    if index >= len(slice) {
        return slice
    }
    
    return append(slice[:index], slice[index + 1:]...)
}
