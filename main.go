package main

import (
	// "fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Unix()))
	mat := MakeMatrix(4, 0)
	image := MakeImage(500, 500)
	for i := 0; i < 20; i++ {
		mat.AddEdge(rand.Float64()*500, 500-rand.Float64()*500, 0.0, rand.Float64()*500, 500-rand.Float64()*500, 0.0)
	}
	image.DrawLines(mat, Color{r: 0, b: 0, g: 0})
	image.Display()
	image.SavePPM("mat.ppm")
}
