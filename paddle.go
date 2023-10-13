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

func (p *Paddle) update() {
	if rl.IsKeyDown(rl.KeyUp) {
		p.y -= p.speed * delta
	}
	if rl.IsKeyDown(rl.KeyDown) {
		p.y += p.speed * delta
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
