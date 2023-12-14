func maxDistToClosest(seats []int) int {
	n := len(seats)
	maxDistance := 0
	lastOccupiedSeatPosition := -1
	for position := 0; position < n; position++ {
		if seats[position] == 1 {
            var maxDistanceFromCurrentSeat int
			if lastOccupiedSeatPosition == -1 {
				maxDistanceFromCurrentSeat = position
			} else {
                // for N consecutive empty seats, max achievable distance is (N + 1) / 2
				maxDistanceFromCurrentSeat = (position - lastOccupiedSeatPosition) / 2
			}
            maxDistance = max(maxDistance, maxDistanceFromCurrentSeat)
			lastOccupiedSeatPosition = position
		}
	}
    maxDistanceFromLastSeatToTheRight := n - 1 - lastOccupiedSeatPosition
	maxDistance = max(maxDistance, maxDistanceFromLastSeatToTheRight)
	return maxDistance
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
