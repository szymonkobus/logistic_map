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

/*
  var x int = 9
  y := 11
  z := y - x
  fmt.Println(z)

  a := []int{1,2,3}
  fmt.Println(a)

  v := make(map[string]int)
  v["1"]=2
  fmt.Println(v)
*/


/*
    set := make(map[string]bool) // New empty set
    set["Foo"] = true            // Add
    for k := range set {         // Loop
    fmt.Println(k)
    }
    delete(set, "Foo")    // Delete
    size := len(set)      // Size
    exists := set["Foo"]  // Membership
*/
