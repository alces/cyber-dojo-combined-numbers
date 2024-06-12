package combine_numbers

func withoutItem(slice []int, index int) []int {
    if index == 0 {
        return slice[index + 1:]
    } else if index == len(slice) - 1 {
        return slice[:index]
    }
    
    return append(slice[:index], slice[index + 1:]...)
}
