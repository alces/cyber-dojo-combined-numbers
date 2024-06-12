package combine_numbers

func permutations(slice []int) [][]int {
    if len(slice) <= 1 {
        return [][]int{slice}
    }
    
    return [][]int{}
}

func withoutItem(slice []int, index int) []int {
    if index >= len(slice) {
        return slice
    }
    
    return append(slice[:index], slice[index + 1:]...)
}
