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

	if b.x+b.radius >= screen_width {
		cpu_score++
		reset_ball(b)
	}
	if b.x+b.radius <= 0 {
		player_score++
		reset_ball(b)
	}
}

func (b *Ball) check_collission(p *Paddle) {
	var ball_vector rl.Vector2 = rl.Vector2{
		X: float32(b.x),
		Y: float32(b.y)}

	var paddle_rectangle rl.Rectangle = rl.Rectangle{
		X:      float32(p.x),
		Y:      float32(p.y),
		Width:  float32(p.width),
		Height: float32(p.height)}

	if rl.CheckCollisionCircleRec(ball_vector, float32(ball.radius), paddle_rectangle) {
		b.speed_x *= -1
	}
}

func reset_ball(ball *Ball) {
	ball.x = screen_width / 2
	ball.y = screen_height / 2

	random_direction(ball)
}

func random_direction(ball *Ball) {
	var speed_directions = [...]int32{-1, 1}
	ball.speed_x *= speed_directions[rl.GetRandomValue(0, 1)]
	ball.speed_y *= speed_directions[rl.GetRandomValue(0, 1)]
}
