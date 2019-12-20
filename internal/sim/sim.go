package sim

import(
    "../defs"
    "fmt"
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

func RunSimulation() [][]defs.Point {
    var p [2][]defs.Point

    state := 0.5
    var kMin float64 = 2
    var kMax float64 = 3.99
    var kStep float64= 0.001
    var sMin float64 = 0.1
    var sMax float64 = 0.9
    var sStep float64= 0.1

    var cycles bool
    for k := kMin; k <= kMax; k+=kStep {
        p[0],cycles = convergencePoints(k,state,p[0],false)
        if !cycles {
            for s := sMin; s <= sMax && s!=state; s+=sStep{
                p[1] = endPoints(k,s,p[1],false)
            }
        }
    }

    return p[:]
}
