type TemperatureAndDay struct {
    T int
    D int
}

type MonoDecStack []TemperatureAndDay

func newMonoDecStack() MonoDecStack {
    s := make([]TemperatureAndDay, 0)
    return MonoDecStack(s)
}

func (s *MonoDecStack) push(v TemperatureAndDay) {
    *s = append(*s, v)
}

func (s *MonoDecStack) pop() TemperatureAndDay {
    n := len(*s)
    v := (*s)[n - 1]
    *s = (*s)[:n - 1]
    return v
}

func (s *MonoDecStack) PopAndPush(v TemperatureAndDay) int {
    for len(*s) > 0 {
        popped := (*s).pop()
        if popped.T > v.T {
            (*s).push(popped)
            (*s).push(v)
            return popped.D
        }
    }
    (*s).push(v)
    return -1
}

func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    mds := newMonoDecStack()
    for i := n - 1; i >= 0; i-- {
        temperatureAndDay := TemperatureAndDay{
            T: temperatures[i],
            D: i,
        }
        targetDay := mds.PopAndPush(temperatureAndDay)
        if targetDay == -1 {
            temperatures[i] = 0
        } else {
            temperatures[i] = targetDay - i
        }
    }
    return temperatures
}
