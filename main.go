package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screen_width = 1280
const screen_height = 800
const paddle_height = 120
const paddle_width = 10
const paddle_edge_padding = 60
const circle_radius = 6
const ball_speed = 500
const paddle_speed = 300

var delta float32 = 1
var ball Ball
var p1 Paddle
var p2 CPU

var player_score = 0
var cpu_score = 0

func main() {
	rl.InitWindow(screen_width, screen_height, "Pung")
	rl.SetWindowState(rl.FlagVsyncHint)
	defer rl.CloseWindow()

	initialize_gameobjects()

	for !rl.WindowShouldClose() {
		update()
		draw()
	}
}

func initialize_gameobjects() {
	ball = Ball{
		x:               screen_width / 2,
		y:               screen_height / 2,
		speed_x:         ball_speed,
		speed_y:         ball_speed,
		radius:          circle_radius,
		temp_ball_speed: ball_speed,
	}
	random_direction(&ball)

	p1 = Paddle{
		side:   0,
		x:      screen_width - (paddle_edge_padding + paddle_width),
		y:      (screen_height / 2) - (paddle_height / 2),
		width:  paddle_width,
		height: paddle_height,
		speed:  paddle_speed,
	}

	p2 = CPU{
		paddle: Paddle{
			side:   1,
			x:      paddle_edge_padding,
			y:      (screen_height / 2) - (paddle_height / 2),
			width:  paddle_width,
			height: paddle_height,
			speed:  paddle_speed},
	}
}

func update() {
	delta = rl.GetFrameTime()

	ball.update()
	p1.update()
	p2.update(&ball)

	ball.check_collission(&p1)
	ball.check_collission(&p2.paddle)
}

func draw() {
	rl.BeginDrawing()
	rl.DrawFPS(10, 10)
	fmt.Print("\033[H\033[2J")
	print(ball.speed_y)
	rl.ClearBackground(rl.Black)

	ball.draw()
	p1.draw()
	p2.draw()
	rl.DrawLine(screen_width/2, 0, screen_width/2, screen_height, rl.White)
	rl.DrawText(fmt.Sprint(cpu_score), (screen_width/4)-20, 20, 80, rl.White)
	rl.DrawText(fmt.Sprint(player_score), 3*(screen_width/4)-20, 20, 80, rl.White)

	rl.EndDrawing()
}
