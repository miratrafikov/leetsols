type Set struct {
    m map[string]bool
}

func (s Set) Add(v string) {
    s.m[v] = true
}

func (s Set) Contains(v string) bool {
    _, found := s.m[v]
    return found
}

func newSet() Set {
    return Set{m: make(map[string]bool)}
}

func destCity(paths [][]string) string {
    n := len(paths)
    departureCities := newSet()
    for i := 0; i < n; i++ {
        departureCity := paths[i][0]
        departureCities.Add(departureCity)
    }
    var answer string
    for i := 0; i < n; i++ {
        arrivalCity := paths[i][1]
        if !departureCities.Contains(arrivalCity) {
            answer = arrivalCity
            break
        }
    }
    return answer
}
