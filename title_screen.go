package main

import rl "github.com/gen2brain/raylib-go/raylib"

type TitleScreen struct {
	message string
}

func new_singeplayer_game() TitleScreen {
	return TitleScreen{"Welcome! Press Enter to start."}
}

func (t *TitleScreen) update() {
	if rl.IsKeyDown(rl.KeyEnter) {
		game_state = initialize_gameplay
	}
}

func (t *TitleScreen) draw() {
	rl.BeginDrawing()
	rl.DrawFPS(10, 10)
	rl.ClearBackground(rl.LightGray)

	rl.DrawText(t.message, screen_width/3, screen_height/2, 32, rl.Black)

	rl.EndDrawing()
}
