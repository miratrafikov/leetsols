func rotate(m [][]int)  {
    n := len(m)
    // mirror along diagonal
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            m[i][j], m[j][i] = m[j][i], m[i][j]
        }
    }
    // mirror left-right
    for i := 0; i < n; i++ {
        for j := 0; j < n / 2; j++ {
            m[i][j], m[i][n-j-1] = m[i][n-j-1], m[i][j]
        }
    }
    return
}

