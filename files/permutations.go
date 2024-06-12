package combine_numbers

func withoutItem(slice []int, index int) []int {
    if index == 0 {
        return slice[index + 1:]
    }
    
    return slice[:index]
}
