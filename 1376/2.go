type Employee struct {
    Subordinates []*Employee
    TimeToInformSubordinates int
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    employeeIdToTimeToInformAllSubordinates := make([]int, n)
    for i := 0; i < n; i++ {
        employeeIdToTimeToInformAllSubordinates[i] = -1
    }
    maxTimeToGetInformedAcrossTheCompany := 0
    for i := 0; i < n; i++ {
        timeForEmployeAndSubordinatesToGetInformedByManager := getTimeForEmployeAndSubordinatesToGetInformedByManager(
            i,
            &manager,
            &informTime,
            &employeeIdToTimeToInformAllSubordinates,
        )
        maxTimeToGetInformedAcrossTheCompany = max(
            maxTimeToGetInformedAcrossTheCompany,
            timeForEmployeAndSubordinatesToGetInformedByManager,
        )
    }
    return maxTimeToGetInformedAcrossTheCompany
}

func getTimeForEmployeAndSubordinatesToGetInformedByManager(
    employeeId int,
    employeeIdToManagerId *[]int,
    employeeIdToTimeToInformImmediateSubordinates *[]int,
    employeeIdToTimeToInformAllSubordinates *[]int,
) int {
    managerId := (*employeeIdToManagerId)[employeeId]
    if managerId == -1 {
        return 0
    }
    if (*employeeIdToTimeToInformAllSubordinates)[employeeId] != -1 {
        return (*employeeIdToTimeToInformAllSubordinates)[employeeId]
    }
    (*employeeIdToTimeToInformAllSubordinates)[employeeId] =
        (*employeeIdToTimeToInformImmediateSubordinates)[managerId] +
        getTimeForEmployeAndSubordinatesToGetInformedByManager(
            managerId,
            employeeIdToManagerId,
            employeeIdToTimeToInformImmediateSubordinates,
            employeeIdToTimeToInformAllSubordinates,
        )
    return (*employeeIdToTimeToInformAllSubordinates)[employeeId]
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
