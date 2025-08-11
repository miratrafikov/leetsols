func productQueries(n int, queries [][]int) []int {
    moduloDivider := int(math.Pow(float64(10), float64(9))) + 7
    maxN := int(math.Pow(float64(10), float64(9)))

    // pots содержит степени двойки от нулевой до той, которая выдаст число, точно большее, чем n
    pots := []int{1,2,4,8}
    for pots[len(pots)-1] < maxN {
        pots = append(pots, 2 * pots[len(pots)-1])
    }

    // potDec содержит наиболее эффективную факторизацию n на степени двойки
    potDec := []int{}
    i := len(pots) - 1
    for n != 0 {
        if pots[i] <= n {
            n -= pots[i]
            potDec = append(potDec, pots[i])
        }

        if n == 0 {
            break
        }

        i--
    }

    // переворачиваем potDec
    for i := 0; i < len(potDec) / 2; i++ {
        potDec[i], potDec[len(potDec)-1-i] = potDec[len(potDec)-1-i], potDec[i]
    }

    answers := make([]int, 0, len(queries))
    for _, q := range queries {
        answer := potDec[q[0]]
        for i := q[0] + 1; i <= q[1]; i++ {
            answer = (answer * potDec[i]) % moduloDivider
        }

        answers = append(answers, answer)
    }

    return answers
}
