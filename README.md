# logistic_map
Visualisation of logistic map in Go.

The program draws the diagram that a sequence defined by - 's(n+1) = k * s(n) * ( 1 - s(n) )'
- converges to (blue). If given the simulation length the sequence does not converge or have
a periodic behaviour - the program plots end values for many starting conditions (red).


to install dependencies:

`go get gonum.org/v1/plot/...`

to run:

`go run cmd/main.go`

the output will be `out.png`

to change the simulation settings access `internal/settings.json`

The solution for the difference equation changes with k. The values of k are
defined in the settings:
-Kmin
-Kmax
-Kstep
where the values of k drawn will be from Kmin to Kmax separated by Kstep

State is the initial state of the simulation. In the case for that value of K
it does not converge. The end states of different starting condition s are found.
The s are defined analogously to k and can be changed in settings file.
