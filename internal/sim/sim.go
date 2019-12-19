package sim

import(
    "../defs"
    "fmt"
    )

func findCycle(l []float64, i int) (bool, int) {
    val := l[i]
    for j:=i-1; j>=0; j-- {
        if l[j] == val{
            return true, j
        }
    }
    return false, -1
}

func convergencePoints(k float64,
                        state float64,
                        p []defs.Point,
                        verbose bool) []defs.Point {

    trace := [defs.Limit]float64{}
    trace[0] = state
    if(verbose){ fmt.Println("k: ",k,"\ts: ",state,"\tdefs.Limit: ",defs.Limit) }

    for i := 1; i<defs.Limit; i++{
        if( state < 0 || 1 < state){
            return p
        }
        state = k * state * (1 - state)
        trace[i] = state
        if(verbose){ fmt.Println("s: ", state) }
        if found, idx := findCycle(trace[:], i); found {
            if(verbose){ fmt.Println("convergance", state, "\tsteps: ",i) }
            for j := idx; j < i; j++ {
                p = append(p, defs.Point{k, trace[j]})
            }
            break
        }
    }

    return p
}

func RunSimulation() []defs.Point {
    var p []defs.Point

    state := 0.3
    var mink float64 = 1
    maxk := 3.99
    step := 0.005

    for k := mink; k < maxk; k+=step {
        p = convergencePoints(k,state,p,false)
    }

    return p
}
