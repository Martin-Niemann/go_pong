package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	x, y, radius, speed_x, speed_y, temp_ball_speed float32
}

func (b *Ball) draw() {
	rl.DrawCircle(int32(b.x), int32(b.y), b.radius, rl.White)
}

func (b *Ball) update() {
	b.x += b.speed_x * delta
	b.y += b.speed_y * delta

	if b.y+b.radius >= screen_height {
		if b.speed_y > 0 {
			b.speed_y *= -1
		}
	}

	if b.y-b.radius <= 0 {
		if b.speed_y < 0 {
			b.speed_y *= -1
		}
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

	if rl.CheckCollisionCircleRec(ball_vector, b.radius, paddle_rectangle) {
		b.temp_ball_speed *= 1.1

		// logic taken from this stackexchange post by Ricket: https://gamedev.stackexchange.com/a/4255
		var relative_intersect_y = (p.y + (p.height / 2)) - b.y
		var normalized_relative_intersect_y = relative_intersect_y / (p.height / 2)
		var bounce_angle = normalized_relative_intersect_y * max_bounce_angle

		// multiply by 1.5 so the speed of the ball after the first paddle hit matches the ball's initial speed
		b.speed_y = b.temp_ball_speed * 1.5 * -float32(math.Sin(float64(bounce_angle)))

		if p.side == 0 {
			b.speed_x = b.temp_ball_speed * 1.5 * -float32(math.Cos(float64(bounce_angle)))
		}
		if p.side == 1 {
			b.speed_x = b.temp_ball_speed * 1.5 * float32(math.Cos(float64(bounce_angle)))
		}
	}
}

func reset_ball(ball *Ball) {
	ball.temp_ball_speed = ball_speed
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
