package sim

import(
    "../defs"
    "fmt"
    "encoding/json"
    "io/ioutil"
    )

func findCycle(l []float64, i int) int {
    val := l[i]
    for j:=i-1; j>=0; j-- {
        if l[j] == val{
            return  j
        }
    }
    return -1
}

func convergencePoints(k float64,
                        state float64,
                        p []defs.Point,
                        verbose bool) ([]defs.Point,bool) {
    if(verbose){ fmt.Println("k: ",k,"\ts: ",state,"\tdefs.Limit: ",defs.Limit) }

    trace := [defs.Limit]float64{}
    trace[0] = state
    setStates := make(map[float64]bool)
    setStates[state] = true

    for i := 1; i<defs.Limit; i++{
        if( state < 0 || 1 < state){
            return p, true
        }
        state = k * state * (1 - state)
        trace[i] = state
        if setStates[state] {
            if verbose { fmt.Println("convergance", state, "\tsteps: ",i) }
            idx := findCycle(trace[:], i)
            for j := idx; j < i; j++ {
                p = append(p, defs.Point{k, trace[j]})
            }
            return p, true
        }
        setStates[state] = true
    }
    return p, false
}

func endPoints(k float64,
                state float64,
                p []defs.Point,
                verbose bool) []defs.Point {

    for i := 1; i<defs.Limit; i++{
        state = k * state * (1 - state)
    }

    if( state < 0 || 1 < state){
        return p
    }

    return append(p, defs.Point{k, state})
}

type simSett struct {
    Kmin    float64 `json:"Kmin"`
    Kmax    float64 `json:"Kmax"`
    Kstep   float64 `json:"kDelta"`
    Smin    float64 `json:"stateMin"`
    Smax    float64 `json:"stateMax"`
    Sstep   float64 `json:"stateStep"`
    State   float64 `json:"state"`
}

func runSimThread(st simSett, p [2][]defs.Point) [2][]defs.Point {
    var cycles bool
    for k := st.Kmin; k<=st.Kmax; k+=st.Kstep {
        p[0], cycles = convergencePoints(k,st.State,p[0],false)
        if !cycles {
            for s := st.Smin; s<=st.Smax && s!=st.State; s+=st.Sstep{
                p[1] = endPoints(k,s,p[1],false)
            }
        }
    }

    return p
}

func readSettings() (simSett) {
    file, err := ioutil.ReadFile("internal/settings.json")
    if err != nil {
        panic(err)
    }
    s := simSett{}
    _ = json.Unmarshal([]byte(file), &s)
    if err != nil {
        panic(err)
    }
    return s
}

func RunSimulation() [][]defs.Point {
    simSettings := readSettings()
    scale := defs.Cores
    p := make([][2][]defs.Point,scale)

    var Kstep float64 = simSettings.Kstep
    simSettings.Kstep *= float64(scale)

    c := make(chan int)
    for i := 0; i < scale; i++{
        go func(a int, s simSett) {
            p[a] = runSimThread(s, p[a])
            c <- a
        }(i, simSettings)
        simSettings.Kmin += Kstep
    }
    var done int
    var ret [2][]defs.Point
    for i := 0; i < scale; i++{
        done = <-c
        ret[0] = append(ret[0], p[done][0]...)
        ret[1] = append(ret[1], p[done][1]...)
    }
    return ret[:]
}
