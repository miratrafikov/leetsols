func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    amountToCoinsNeeded := make([]int, amount+1)
    for i := 0; i <= amount; i++ {
        amountToCoinsNeeded[i] = -1
    }
    for i := 0; i < len(coins); i++ {
        if coins[i] <= amount {
            amountToCoinsNeeded[coins[i]] = 1
        }
    }
    for i := 1; i <= amount; i++ {
        if amountToCoinsNeeded[i] == -1 {
            minCoins := -2
            for j := 0; j < len(coins); j++ {
                if i - coins[j] > 0 && amountToCoinsNeeded[i-coins[j]] != -1 {
                    if minCoins == -2 {
                        minCoins = 1 + amountToCoinsNeeded[i-coins[j]]
                    } else {
                        minCoins = min(minCoins, 1+amountToCoinsNeeded[i-coins[j]])
                    }
                }
            }
            if minCoins != -2 {
                amountToCoinsNeeded[i] = minCoins
            }
        }
    }
    return amountToCoinsNeeded[amount]
}
