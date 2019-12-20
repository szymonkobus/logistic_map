package plt

import(
    "../defs"
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
    "gonum.org/v1/plot/vg/draw"
    "image/color"
)

func points2plotter(p [][]defs.Point) []plotter.XYs {
    var r []plotter.XYs
    for _, points := range p{
        pltr := make(plotter.XYs, len(points))
        for i, cord := range points {
            pltr[i].X = cord.X
            pltr[i].Y = cord.Y
        }
        r = append(r, pltr)
    }
    return r
}

func PlotPoints(p [][]defs.Point) {
    plotterL := points2plotter(p)
    diagram, err := plot.New()
	if err != nil {
		panic(err)
	}
    diagram.Title.Text = "Convergance points"
    diagram.X.Label.Text = "k"
    diagram.Y.Label.Text = "value"
    diagram.Add(plotter.NewGrid())

    var size float64 = 20

    for i := 0; i < len(plotterL); i++{
        scatter, err := plotter.NewScatter(plotterL[i])
    	if err != nil {
    		panic(err)
    	}
        red_lvl:= uint8(255*float64(i)/float64(len(plotterL)-1))
        scatter.GlyphStyle.Color = color.RGBA{R: red_lvl, B: 128, A: 255}
        scatter.GlyphStyle.Shape = draw.CircleGlyph{}
        scatter.GlyphStyle.Radius = vg.Points(size*1)
        diagram.Add(scatter)
    }


    if err := diagram.Save(vg.Length(size*7)*vg.Inch, vg.Length(size*7)*vg.Inch, "out.png"); err != nil {
		panic(err)
	}
}
