package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	x, y, speed_x, speed_y, radius int32
}

func (b *Ball) draw() {
	rl.DrawCircle(b.x, b.y, float32(b.radius), rl.White)
}

func (b *Ball) update() {
	b.x += b.speed_x
	b.y += b.speed_y

	if b.y+b.radius >= screen_height || b.y-b.radius <= 0 {
		b.speed_y *= -1
	}
}
