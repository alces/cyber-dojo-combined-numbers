package combine_numbers

func withoutItem(slice []int, index int) []int {
    if index >= len(slice) {
        return slice
    }
    
    return append(slice[:index], slice[index + 1:]...)
}
