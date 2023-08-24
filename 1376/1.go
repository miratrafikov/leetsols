type Employee struct {
    Subordinates []*Employee
    TimeToInformSubordinates int
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    tree := buildHierarchyTree(n, headID, manager, informTime)
    return timeToInformSubordinates(tree)
}

func buildHierarchyTree(n int, headID int, manager []int, informTime []int) Employee {
    employees := make([]Employee, n)
    for i := 0; i < n; i++ {
        employees[i].TimeToInformSubordinates = informTime[i]
        if i != headID {
            employees[manager[i]].Subordinates = append(
                employees[manager[i]].Subordinates,
                &employees[i],
            )
        }
    }
    return employees[headID]
}

func timeToInformSubordinates(employee Employee) int {
    timesForSubordinatesToInformTheirSubordinates := make([]int, len(employee.Subordinates))
    for i := 0; i < len(employee.Subordinates); i++ {
        timesForSubordinatesToInformTheirSubordinates[i] = timeToInformSubordinates(*employee.Subordinates[i])
    }
    return employee.TimeToInformSubordinates + maxInSlice(timesForSubordinatesToInformTheirSubordinates)
}

func maxInSlice(s []int) int {
    if len(s) == 0 {
        return 0
    }
    max := s[0]
    for i := 1; i < len(s); i++ {
        if s[i] > max {
            max = s[i]
        }
    }
    return max
}
