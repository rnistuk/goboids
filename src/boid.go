package src

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Boid struct {
	Position       Vector
	Velocity       Vector
	targetDistance float64
}

func NewBoid() Boid {
	return Boid{
		Vector{},
		Vector{},
		0.0,
	}
}

// TODO: boid should not need to know about screen width or height, I need to refactor Boid to be only logic
func (b Boid) Draw(render *sdl.Renderer, screenWidth int32, screenHeight int32) {
	err := render.SetDrawColor(5, 55, 8, 255)
	if err != nil {
		return
	}
	X := screenWidth/2 + int32(b.Position.X/1.0)
	Y := screenHeight/2 + int32(b.Position.Y/1.0)
	r := sdl.Rect{X, Y, 2, 2}
	_ = render.FillRect(&r)

	render.SetDrawColor(64, 5, 5, 0)
	drawCircle(render, X, Y, int32(Parameters["near"]))
}

func (b Boid) toString() string {
	return fmt.Sprintf("P: %s    V: %s", b.Position.toString(), b.Velocity.toString())
}

func (b *Boid) UpdateVelocity(bs Boids) {
	b.Velocity = b.Velocity.Add(NCohesionRule(b, bs))
	b.Velocity = b.Velocity.Add(SeparationRule(b, bs))
	b.Velocity = b.Velocity.Add(AlignmentRule(b, bs))
	b.Velocity = b.Velocity.Add(HomeRule(b, bs))
	b.Velocity = b.Velocity.Add(LimitSpeedRule(b, bs))
	b.Velocity = b.Velocity.Add(MinimumSpeedRule(b, bs))
}

func (b *Boid) UpdatePosition() {
	b.Position = b.Position.Add(b.Velocity)
}

func drawCircle(renderer *sdl.Renderer, centreX int32, centreY int32, radius int32) {
	diameter := (radius * 2)
	x := (radius - 1)
	y := int32(0)
	tx := int32(1)
	ty := int32(1)
	error := (tx - diameter)

	for x >= y {
		//  Each of the following renders an octant of the circle
		renderer.DrawPoint(centreX+x, centreY-y)
		renderer.DrawPoint(centreX+x, centreY+y)
		renderer.DrawPoint(centreX-x, centreY-y)
		renderer.DrawPoint(centreX-x, centreY+y)
		renderer.DrawPoint(centreX+y, centreY-x)
		renderer.DrawPoint(centreX+y, centreY+x)
		renderer.DrawPoint(centreX-y, centreY-x)
		renderer.DrawPoint(centreX-y, centreY+x)
		if error <= 0 {
			y += 1
			error += ty
			ty += 2
		}
		if error > 0 {
			x -= 1
			tx += 2
			error += (tx - diameter)
		}
	}
}
