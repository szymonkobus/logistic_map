package main

import(
      "fmt"
      "math"
)

func main() {
    limit := 5000
    k := 2.7        // k
    s := 0.99        //state

    var prev_s float64
    var change float64

    fmt.Println("k: ",k,"\ts: ",s,"\tlimit: ",limit)
    for i := 0; i<limit; i++{
        prev_s = s
        s = k * s * (1 - s)
        fmt.Println("s: ", s)
        change = math.Abs(prev_s - s)
        if(change <= 0.0000000000000005){
            fmt.Println("convergance", s, "\tsteps: ",i)
            break
        }
    }

//pytania:
// 1. jak sprawdzic convergance
// 2. co powinien potrafic:
// 2a) narysuj jeden przebieg
// 2b) narysuj wszytskie przebiegi



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

}
