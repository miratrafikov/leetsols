func minimumRecolors(blocks string, k int) int {
    n := len(blocks)
    blacksInWindow := 0
    for i := 0; i < k; i++ {
        if blocks[i] == 'B' {
            blacksInWindow++
        }
    }
    maxBlacksInWindow := blacksInWindow
    for i := k; i < n; i++ {
        if maxBlacksInWindow == k {
            return 0
        }
        if blocks[i-k] == 'B' {
            blacksInWindow--
        }
        if blocks[i] == 'B' {
            blacksInWindow++
        }
        if blacksInWindow > maxBlacksInWindow {
            maxBlacksInWindow = blacksInWindow
        }
    }
    return k - maxBlacksInWindow
}
