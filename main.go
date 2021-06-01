package main

import (
	"log"
	"time"

	"github.com/h8gi/canvas"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

const (
	lifeSize     = 10
	initialCount = 50
	Width        = 500
	Height       = 400
)

type World struct {
	isLife []bool
}

func NewWorld() World {
	w := World{}
	w.isLife = make([]bool, Width*Height)
	return w
}

func (w World) IsLifeAtCoords(x int, y int) bool {
	if x < 0 || x >= Width || y < 0 || y >= Height {
		return false
	}

	return w.isLife[y*Width+x]
}

func (w *World) setAtCoords(x int, y int, life bool) {
	w.isLife[y*Width+x] = life
}

func (w World) LifeAroundCoords(x int, y int) int {
	life := 0
	life += w.scoreLifeAtCoords(x-1, y-1)
	life += w.scoreLifeAtCoords(x, y-1)
	life += w.scoreLifeAtCoords(x+1, y-1)

	life += w.scoreLifeAtCoords(x-1, y)
	life += w.scoreLifeAtCoords(x+1, y)

	life += w.scoreLifeAtCoords(x-1, y+1)
	life += w.scoreLifeAtCoords(x, y+1)
	life += w.scoreLifeAtCoords(x+1, y+1)
	return life
}

func (w World) scoreLifeAtCoords(x int, y int) int {
	score := 0
	if w.IsLifeAtCoords(x, y) {
		score = 1
	}
	return score
}

func (w *World) Cycle() {

	// Create a copy
	orig := NewWorld()
	orig.isLife = make([]bool, len(w.isLife)) // World{} //make([]bool, Width*Height)
	copy(orig.isLife, w.isLife)

	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			isLife := orig.IsLifeAtCoords(x, y)
			score := orig.LifeAroundCoords(x, y)
			if isLife {
				log.Printf("Life at (%d, %d) - neighbours = %d\n", x, y, score)
			}
			switch score {

			case 2:
				// Do nothing
			case 3:
				if !isLife {
					w.Spawn(x, y)
				}

			default:
				if isLife {
					w.Kill(x, y)
				}

			}
		}
	}
}

func (w *World) Kill(x int, y int) {
	log.Printf("Killing at (%d, %d)\n", x, y)
	w.setAtCoords(x, y, false)
}

func (w *World) Spawn(x int, y int) {
	log.Printf("Spawning at (%d, %d)\n", x, y)
	w.setAtCoords(x, y, true)
}

func main() {
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     Width * 2,
		Height:    Height * 2,
		FrameRate: 30,
		Title:     "Gome of life",
	})

	c.Setup(func(ctx *canvas.Context) {
		ctx.SetColor(colornames.White)
		ctx.Clear()
		ctx.SetColor(colornames.Green500)
		ctx.SetLineWidth(1)
	})

	w := NewWorld()
	w.Spawn(198, 198)
	w.Spawn(200, 198)
	w.Spawn(202, 198)
	w.Spawn(199, 199)
	w.Spawn(200, 199)
	w.Spawn(201, 199)
	w.Spawn(199, 203)
	w.Spawn(200, 203)
	w.Spawn(201, 203)
	w.Spawn(199, 207)
	w.Spawn(200, 207)
	w.Spawn(201, 207)

	w.Spawn(300, 200)
	w.Spawn(301, 200)
	w.Spawn(302, 200)
	w.Spawn(304, 200)
	w.Spawn(305, 200)
	w.Spawn(306, 200)
	w.Spawn(300, 201)
	w.Spawn(306, 201)
	w.Spawn(300, 201)
	w.Spawn(306, 202)
	w.Spawn(300, 203)
	w.Spawn(306, 203)
	w.Spawn(300, 204)
	w.Spawn(301, 204)
	w.Spawn(302, 204)
	w.Spawn(303, 204)
	w.Spawn(304, 204)
	w.Spawn(305, 204)
	w.Spawn(306, 204)

	log.Println("Creating ticker")
	ticker := time.NewTicker(time.Second)
	go func(world World) {
		for {
			<-ticker.C
			log.Println("Cycle")
			w.Cycle()
		}
	}(w)

	c.Draw(func(ctx *canvas.Context) {

		ctx.SetColor(colornames.White)
		ctx.Clear()
		ctx.SetColor(colornames.Green500)

		for y := 0; y < Height; y++ {
			ctx.Push()
			for x := 0; x < Width; x++ {
				if w.IsLifeAtCoords(x, y) {
					ctx.DrawPoint(float64(x*2), float64(y*2), 2)
					ctx.Fill()
				}
			}
			ctx.Pop()
		}
	})
}
