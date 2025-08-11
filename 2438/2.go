func productQueries(n int, queries [][]int) []int {
    m := 1_000_000_007

    // pd (powers of two decomposition) содержит минимальное число степеней двойки, которые нужно сложить, чтобы получить n
    pd := make([]int, 0, 30)
    for i := 0; i <= 30; i++ {
        if (n >> i) & 1 == 1 {
            pd = append(pd, 1 << i)
        }
    }

    // products содержит префиксные произведения: a_k = a_1 * a_2 * ... * a_(k-1)
    products := make([][]int, len(pd))
    for i := 0; i < len(pd); i++ {
        products[i] = make([]int, len(pd))
        acc := 1
        for j := i; j < len(pd); j++ {
            acc = (acc * pd[j]) % m
            products[i][j] = acc
        }
    }

    answers := make([]int, 0, len(queries))
    for _, q := range queries {
        answers = append(answers, products[q[0]][q[1]])
    }

    return answers
}
