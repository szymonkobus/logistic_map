package main

import(
    "fmt"
    "image/color"

    "gonum.org/v1/plot"
  	"gonum.org/v1/plot/plotter"
  	"gonum.org/v1/plot/vg"
  	"gonum.org/v1/plot/vg/draw"
)

const limit int = 5000
type point struct {x float64; y float64}

func compare_end(l [limit]float64, i int, rng int) bool {
    val := l[i]
    for j:=i-1; j>=0 && j>=i-rng; j-- {
        if l[j] == val{
            return true
        }
    }
    return false
}

func convergence_points(k float64,
                        state float64,
                        p []point,
                        verbose bool) []point {

    trace := [limit]float64{}
    trace[0] = state
    var last int
    if(verbose){ fmt.Println("k: ",k,"\ts: ",state,"\tlimit: ",limit) }

    for i := 1; i<limit; i++{
        state = k * state * (1 - state)
        trace[i] = state
        if(verbose){ fmt.Println("s: ", state) }
        if( compare_end(trace, i, 10) ){
            if(verbose){ fmt.Println("convergance", state, "\tsteps: ",i) }
            last = i
            break
        }
    }

    p = append(p, point{k, trace[last]})
    return p
}

func run_simulations() []point {
    var p []point

    state := 0.3
    mink := 0.1
    maxk := 3.9
    step := 0.01

    for k := mink; k < maxk; k+=step {
        p = convergence_points(k,state,p,false)
    }

    return p
}

func points2plotter(p []point) plotter.XYs {
    r := make(plotter.XYs, len(p))
    for i, cord := range p {
        r[i].X = cord.x
        r[i].Y = cord.y
    }
    return r
}

func plot_points(p []point) {
    data := points2plotter(p)
    diagram, err := plot.New()
	if err != nil {
		panic(err)
	}
    diagram.Title.Text = "Convergance points"
    diagram.X.Label.Text = "k"
    diagram.Y.Label.Text = "value"
    diagram.Add(plotter.NewGrid())

    scatter, err := plotter.NewScatter(data)
	if err != nil {
		panic(err)
	}
    scatter.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
    scatter.GlyphStyle.Shape = draw.CircleGlyph{}
    scatter.GlyphStyle.Radius = vg.Points(1)
    fmt.Println(scatter)


    diagram.Add(scatter)
    if err := diagram.Save(7*vg.Inch, 7*vg.Inch, "out.png"); err != nil {
		panic(err)
	}
}

func main() {
    //var p = []point {{0,0}, {1,1}, {0.5,0.75},} //sample points
    p := run_simulations()
    //points := convergence_points()
    plot_points(p)
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
