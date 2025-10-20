func finalValueAfterOperations(operations []string) int {
    val := 0
    for _, op := range operations {
        if op[0] == '-' || op[2] == '-' {
            val--
        } else {
            val++
        }
    }
    return val   
}
