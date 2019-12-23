package main

import(
    "../internal/sim"
    "../internal/plt"

    "fmt"
    "time"
)

func main() {
    fmt.Println("Running simulation.")

    start := time.Now()
    p := sim.RunSimulation()
    t := time.Now()
    elapsed_sim := t.Sub(start)

    fmt.Println("Plotting points.")
    start = time.Now()
    plt.PlotPoints(p)
    t = time.Now()
    elapsed_plot := t.Sub(start)

    fmt.Println("Time simulation:",elapsed_sim,"\tTime plot",elapsed_plot)
}
