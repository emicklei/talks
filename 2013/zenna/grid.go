package main

import (
	"github.com/ajstarks/svgo"
	. "github.com/emicklei/zenna/svgf"
	. "github.com/emicklei/zenna/xy"
	"os"
)

func main() {
	canvas := svg.New(os.Stdout)
	canvas.Start(800, 600)
	canvas.Translate(400, 300)

	painter := NewSVGPainter("stroke-width:0.5px;stroke:red", canvas)
	grid := Grid{Center: P(0, 0), Width: 200, Height: 200, Dx: 10, Dy: 10}

	painter.Paint(grid)

	canvas.Gend()
	canvas.End()
}
