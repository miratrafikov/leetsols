func maxDistToClosest(seats []int) int {
    n := len(seats)
    distanceToClosestOccupiedSeatFromTheLeft := make([]int, n)
    for i := 0; i < n; i++ {
        distance := n
        if seats[i] == 1 {
            distance = 0
        } else if i > 0 {
            distance = distanceToClosestOccupiedSeatFromTheLeft[i - 1] + 1
        }
        distanceToClosestOccupiedSeatFromTheLeft[i] = distance
    }
    distanceToClosestOccupiedSeatFromTheRight := make([]int, n)
    for i := n - 1; i >= 0; i-- {
        distance := n
        if seats[i] == 1 {
            distance = 0   
        } else if i < n - 1 {
            distance = distanceToClosestOccupiedSeatFromTheRight[i + 1] + 1
        }
        distanceToClosestOccupiedSeatFromTheRight[i] = distance
    }
    var maxDistance int
    for i := 0; i < n; i++ {
        distance := min(
            distanceToClosestOccupiedSeatFromTheLeft[i],
            distanceToClosestOccupiedSeatFromTheRight[i],
        )
        if distance > maxDistance {
            maxDistance = distance
        }
    }
    return maxDistance
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
