package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const screen_width = 1280
const screen_height = 800
const paddle_height = 120
const paddle_width = 25
const paddle_edge_padding = 60
const circle_radius = 20

var ball Ball
var p1 Paddle
var p2 CPU

func main() {
	rl.InitWindow(screen_width, screen_height, "Pung")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	initialize_gameobjects()

	for !rl.WindowShouldClose() {
		update()
		draw()
	}
}

func initialize_gameobjects() {
	ball = Ball{
		x:       screen_width / 2,
		y:       screen_height / 2,
		speed_x: 7,
		speed_y: 7,
		radius:  circle_radius,
	}

	p1 = Paddle{
		x:      screen_width - (paddle_edge_padding + paddle_width),
		y:      (screen_height / 2) - (paddle_height / 2),
		width:  paddle_width,
		height: paddle_height,
		speed:  6,
	}

	p2 = CPU{
		paddle: Paddle{x: paddle_edge_padding,
			y:      (screen_height / 2) - (paddle_height / 2),
			width:  paddle_width,
			height: paddle_height,
			speed:  6},
	}
}

func update() {
	ball.update()
	p1.update()
	p2.update(&ball)

	check_collissions(&p1)
	check_collissions(&p2.paddle)
}

func check_collissions(p *Paddle) {
	if rl.CheckCollisionCircleRec(rl.Vector2{X: float32(ball.x), Y: float32(ball.y)}, float32(ball.radius), rl.Rectangle{X: float32(p.x), Y: float32(p.y), Width: float32(p.width), Height: float32(p.height)}) {
		ball.speed_x *= -1
	}
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	ball.draw()
	p1.draw()
	p2.draw()
	rl.DrawLine(screen_width/2, 0, screen_width/2, screen_height, rl.White)

	rl.EndDrawing()
}
