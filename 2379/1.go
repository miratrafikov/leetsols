func minimumRecolors(blocks string, k int) int {
    whitesInWindow := 0
    for i := 0; i < k; i++ {
        if blocks[i] == 'W' {
            whitesInWindow++
        }
    }
    minWhitesInWindow := whitesInWindow
    for i := k; i < len(blocks); i++ {
        if minWhitesInWindow == 0 {
            return minWhitesInWindow
        }
        if blocks[i] == 'W' {
            whitesInWindow++
        }
        if blocks[i-k] == 'W' && whitesInWindow != 0 {
            whitesInWindow--
        }
        if whitesInWindow < minWhitesInWindow {
            minWhitesInWindow = whitesInWindow
        }
    }
    return minWhitesInWindow
}
