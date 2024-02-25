package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameplaySingeplayer struct {
	ball Ball
	p1   Paddle
	p2   CPU
}

func (g *GameplaySingeplayer) initialize_gameobjects() {
	g.ball = Ball{
		x:               screen_width / 2,
		y:               screen_height / 2,
		speed_x:         ball_speed,
		speed_y:         ball_speed,
		radius:          circle_radius,
		temp_ball_speed: ball_speed,
	}
	random_direction(&g.ball)

	g.p1 = Paddle{
		side:   0,
		x:      screen_width - (paddle_edge_padding + paddle_width),
		y:      (screen_height / 2) - (paddle_height / 2),
		width:  paddle_width,
		height: paddle_height,
		speed:  paddle_speed,
	}

	g.p2 = CPU{
		paddle: Paddle{
			side:   1,
			x:      paddle_edge_padding,
			y:      (screen_height / 2) - (paddle_height / 2),
			width:  paddle_width,
			height: paddle_height,
			speed:  paddle_speed},
	}
}

func (g *GameplaySingeplayer) update() {
	g.ball.update()
	g.p1.poll_input()
	g.p1.update()
	g.p2.update(&g.ball)

	g.ball.check_collission(&g.p1)
	g.ball.check_collission(&g.p2.paddle)

	/*fmt.Print("\033[H\033[2J")
	fmt.Println("Ball x: ", g.ball.x)
	fmt.Println("Ball y: ", g.ball.y)
	fmt.Println("P1 y: ", g.p1.y)
	fmt.Println("P2 y: ", g.p2.paddle.y)*/
}

func (g *GameplaySingeplayer) draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.DrawFPS(10, 10)

	g.ball.draw()
	g.p1.draw()
	g.p2.draw()
	rl.DrawLine(screen_width/2, 0, screen_width/2, screen_height, rl.White)
	rl.DrawText(fmt.Sprint(cpu_score), (screen_width/4)-20, 20, 80, rl.White)
	rl.DrawText(fmt.Sprint(player_score), 3*(screen_width/4)-20, 20, 80, rl.White)

	rl.EndDrawing()
}
