package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Paddle struct {
	side                       byte
	x, y, width, height, speed float32
}

func (p *Paddle) draw() {
	rl.DrawRectangle(int32(p.x), int32(p.y), int32(p.width), int32(p.height), rl.White)
}

func (p *Paddle) poll_input() {
	// code to check if the mouse cursor is inside the game window by raysan5 from this reddit comment:
	// https://www.reddit.com/r/raylib/comments/1587p0l/comment/jt8kkx7/?utm_source=reddit&utm_medium=web2x&context=3
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{X: 0, Y: 0, Width: float32(rl.GetScreenWidth()), Height: float32(rl.GetScreenHeight())}) {
		p.y = float32(rl.GetMouseY()) - (p.height / 2)
	}
}

func (p *Paddle) update() {
	p.limit_movement()
}

func (p *Paddle) limit_movement() {
	if p.y <= 0 {
		p.y = 0
	}
	if p.y+paddle_height >= screen_height {
		p.y = screen_height - paddle_height
	}
}
