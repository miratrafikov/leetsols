func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    for i := 0; i < n; i++ {
        j := i + 1
        for j < n && temperatures[j] <= temperatures[i] {
            j++
        }
        if j < n {
            temperatures[i] = j - i
        } else {
            temperatures[i] = 0
        }
    }
    return temperatures
}
