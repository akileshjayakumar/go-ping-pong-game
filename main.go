package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"golang.org/x/image/colornames"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/font/basicfont"
	"github.com/faiface/pixel/text"
)

const (
	windowWidth      = 1024
	windowHeight     = 768
	paddleWidth      = 20
	paddleHeight     = 100
	ballSize         = 20
	initialBallSpeed = 300
	acceleration     = 1.05
)

type Paddle struct {
	Rect     pixel.Rect
	Color    pixel.RGBA
	Speed    float64
	UpKey    pixelgl.Button
	DownKey  pixelgl.Button
	Position pixel.Vec
}

func NewPaddle(x, y float64, color pixel.RGBA, upKey, downKey pixelgl.Button) *Paddle {
	return &Paddle{
		Rect:     pixel.R(x, y, x+paddleWidth, y+paddleHeight),
		Color:    color,
		Speed:    400,
		UpKey:    upKey,
		DownKey:  downKey,
		Position: pixel.V(x, y),
	}
}

func (p *Paddle) Update(dt float64, win *pixelgl.Window) {
	if win.Pressed(p.UpKey) && p.Rect.Max.Y < windowHeight {
		p.Position.Y += p.Speed * dt
	}
	if win.Pressed(p.DownKey) && p.Rect.Min.Y > 0 {
		p.Position.Y -= p.Speed * dt
	}
	p.Rect = p.Rect.Moved(p.Position.Sub(p.Rect.Center()))
}

func (p *Paddle) Draw(imd *imdraw.IMDraw) {
	imd.Color = p.Color
	imd.Push(p.Rect.Min, p.Rect.Max)
	imd.Rectangle(0)
}

type Ball struct {
	Circle   pixel.Circle
	Color    pixel.RGBA
	Velocity pixel.Vec
}

func NewBall(x, y float64, color pixel.RGBA, speed float64) *Ball {
	// Random initial direction
	direction := rand.Float64() * 2 * math.Pi
	vx := speed * math.Cos(direction)
	vy := speed * math.Sin(direction)
	return &Ball{
		Circle: pixel.C(pixel.V(x, y), ballSize/2),
		Color:  color,
		Velocity: pixel.V(vx, vy),
	}
}

func (b *Ball) Update(dt float64, leftPaddle, rightPaddle *Paddle) {
	b.Circle = b.Circle.Moved(b.Velocity.Scaled(dt))

	if b.Circle.Center.Y-b.Circle.Radius < 0 || b.Circle.Center.Y+b.Circle.Radius > windowHeight {
		b.Velocity.Y = -b.Velocity.Y
	}

	// Check for collision with paddles
	if b.Circle.Center.X-b.Circle.Radius < leftPaddle.Rect.Max.X && 
		b.Circle.Center.Y > leftPaddle.Rect.Min.Y && 
		b.Circle.Center.Y < leftPaddle.Rect.Max.Y {
		b.Velocity.X = -b.Velocity.X * acceleration
		b.Velocity.Y *= acceleration
	}
	if b.Circle.Center.X+b.Circle.Radius > rightPaddle.Rect.Min.X && 
		b.Circle.Center.Y > rightPaddle.Rect.Min.Y && 
		b.Circle.Center.Y < rightPaddle.Rect.Max.Y {
		b.Velocity.X = -b.Velocity.X * acceleration
		b.Velocity.Y *= acceleration
	}

	// Reset ball if it goes out of screen bounds
	if b.Circle.Center.X < 0 || b.Circle.Center.X > windowWidth {
		b.Circle.Center = pixel.V(windowWidth/2, windowHeight/2)
		// Randomize initial direction again
		direction := rand.Float64() * 2 * math.Pi
		b.Velocity = pixel.V(initialBallSpeed*math.Cos(direction), initialBallSpeed*math.Sin(direction))
	}
}

func (b *Ball) Draw(imd *imdraw.IMDraw) {
	imd.Color = b.Color
	imd.Push(b.Circle.Center)
	imd.Circle(b.Circle.Radius, 0)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Ping Pong",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	leftPaddle := NewPaddle(50, windowHeight/2, pixel.ToRGBA(colornames.Blue), pixelgl.KeyW, pixelgl.KeyS)
	rightPaddle := NewPaddle(windowWidth-50-paddleWidth, windowHeight/2, pixel.ToRGBA(colornames.Red), pixelgl.KeyUp, pixelgl.KeyDown)
	ball := NewBall(windowWidth/2, windowHeight/2, pixel.ToRGBA(colornames.White), initialBallSpeed)

	imd := imdraw.New(nil)
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(10, windowHeight-30), atlas)
	txt.Color = colornames.White

	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		imd.Clear()
		win.Clear(colornames.Black)

		leftPaddle.Update(dt, win)
		rightPaddle.Update(dt, win)
		ball.Update(dt, leftPaddle, rightPaddle)

		leftPaddle.Draw(imd)
		rightPaddle.Draw(imd)
		ball.Draw(imd)

		imd.Draw(win)
		txt.Clear()
		fmt.Fprintf(txt, "Ping Pong Game\n")
		txt.Draw(win, pixel.IM.Scaled(txt.Orig, 2))
		win.Update()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	pixelgl.Run(run)
}
