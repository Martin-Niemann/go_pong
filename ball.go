package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	x, y, radius, speed_x, speed_y float32
}

func (b *Ball) draw() {
	rl.DrawCircle(int32(b.x), int32(b.y), b.radius, rl.White)
}

func (b *Ball) update() {
	b.x += b.speed_x * delta
	b.y += b.speed_y * delta

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
		X: b.x,
		Y: b.y}

	var paddle_rectangle rl.Rectangle = rl.Rectangle{
		X:      p.x,
		Y:      p.y,
		Width:  p.width,
		Height: p.height}

	if rl.CheckCollisionCircleRec(ball_vector, ball.radius, paddle_rectangle) {
		if p.side == 0 {
			if b.speed_x > 0 {
				b.speed_x *= -1.1
				b.speed_y = (b.y - p.y) / (p.height / 2) * -b.speed_x
			}
		}
		if p.side == 1 {
			if b.speed_x < 0 {
				b.speed_x *= -1.1
				b.speed_y = (b.y - p.y) / (p.height / 2) * b.speed_x
			}
		}
	}
}

func reset_ball(ball *Ball) {
	ball.speed_x = ball_speed
	ball.speed_y = ball_speed
	ball.x = screen_width / 2
	ball.y = screen_height / 2

	random_direction(ball)
}

func random_direction(ball *Ball) {
	var speed_directions = [...]int32{-1, 1}
	ball.speed_x *= float32(speed_directions[rl.GetRandomValue(0, 1)])
	ball.speed_y *= float32(speed_directions[rl.GetRandomValue(0, 1)])
}
