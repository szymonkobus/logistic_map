package plt

import(
    "../defs"
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
    "gonum.org/v1/plot/vg/draw"
    "image/color"
)

func points2plotter(p []defs.Point) plotter.XYs {
    r := make(plotter.XYs, len(p))
    for i, cord := range p {
        r[i].X = cord.X
        r[i].Y = cord.Y
    }
    return r
}

func PlotPoints(p []defs.Point) {
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


    diagram.Add(scatter)
    if err := diagram.Save(7*vg.Inch, 7*vg.Inch, "out.png"); err != nil {
		panic(err)
	}
}
