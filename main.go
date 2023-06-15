package main

import (
	"fmt"

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

var player_score = 0
var cpu_score = 0

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
	random_direction(&ball)

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

	ball.check_collission(&p1)
	ball.check_collission(&p2.paddle)
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	ball.draw()
	p1.draw()
	p2.draw()
	rl.DrawLine(screen_width/2, 0, screen_width/2, screen_height, rl.White)
	rl.DrawText(fmt.Sprint(cpu_score), (screen_width/4)-20, 20, 80, rl.White)
	rl.DrawText(fmt.Sprint(player_score), 3*(screen_width/4)-20, 20, 80, rl.White)

	rl.EndDrawing()
}
