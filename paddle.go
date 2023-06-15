package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Paddle struct {
	x, y, width, height, speed int32
}

func (p *Paddle) draw() {
	rl.DrawRectangle(p.x, p.y, p.width, p.height, rl.White)
}

func (p *Paddle) update() {
	if rl.IsKeyDown(rl.KeyUp) {
		p.y -= p.speed
	}
	if rl.IsKeyDown(rl.KeyDown) {
		p.y += p.speed
	}

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
