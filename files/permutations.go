package combine_numbers

func withoutItem(slice []int, index int) []int {
    return append(slice[:index], slice[index + 1:]...)
}
