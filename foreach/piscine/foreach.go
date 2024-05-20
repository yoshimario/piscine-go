package piscine

// ForEach applies function f to each element of the slice a.
func ForEach(f func(int), a []int) {
    for _, value := range a {
        f(value)
    }
}
