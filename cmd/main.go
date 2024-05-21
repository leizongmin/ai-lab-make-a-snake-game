package main

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
	"snake-game/game"
)

func clearScreen() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

func hideCursor() {
	termbox.HideCursor()
}

func processInput(gameState *game.GameState) {
	for !gameState.GameExited {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			gameState.HandleKey(ev.Ch, ev.Key)
		}
	}
}

func drawScreen(gameState *game.GameState) {
	clearScreen()

	screenWidth, screenHeight := gameState.ScreenWidth, gameState.ScreenHeight

	// Set background color for top and bottom rows and the rest of the screen
	for i := 0; i < screenWidth; i++ {
		termbox.SetCell(i, 0, ' ', termbox.ColorBlack, termbox.ColorWhite)
		termbox.SetCell(i, screenHeight-1, ' ', termbox.ColorBlack, termbox.ColorWhite)
		for j := 1; j < screenHeight-1; j++ {
			termbox.SetCell(i, j, ' ', termbox.ColorDefault, termbox.ColorBlue)
		}
	}

	if gameState.GameOver {
		msg := "Game Over"
		x := (screenWidth - len(msg)) / 2
		y := screenHeight / 2
		for _, c := range msg {
			termbox.SetCell(x, y, c, termbox.ColorRed, termbox.ColorDefault)
			x++
		}
		termbox.Flush()
		return
	}

	if gameState.GamePaused {
		msg := "Paused"
		x := (screenWidth - len(msg)) / 2
		y := screenHeight / 2
		for _, c := range msg {
			termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorDefault)
			x++
		}
	}

	for _, pos := range gameState.SnakePosition {
		termbox.SetCell(pos[1], pos[0], '●', termbox.ColorGreen, termbox.ColorBlue)
	}
	termbox.SetCell(gameState.FoodPosition[1], gameState.FoodPosition[0], '●', termbox.ColorRed, termbox.ColorBlue)

	// Display score
	scoreMsg := fmt.Sprintf("Score: %d", gameState.Score)
	for i, c := range scoreMsg {
		termbox.SetCell(i, screenHeight-1, c, termbox.ColorBlack, termbox.ColorWhite)
	}

	// Display control instructions
	controlMsg := "Use ASDW or arrow keys to move, Space to pause/start, X to exit."
	for i, c := range controlMsg {
		if i < screenWidth {
			termbox.SetCell(i, 0, c, termbox.ColorBlack, termbox.ColorWhite)
		}
	}

	termbox.Flush()
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	screenWidth, screenHeight := termbox.Size()
	gameState := game.NewGameState(screenWidth, screenHeight)

	hideCursor()

	go processInput(gameState)

	for !gameState.GameExited {
		gameState.UpdateGame()
		drawScreen(gameState)
		time.Sleep(gameState.CalculateSleepDuration()) // Use dynamic sleep duration based on snake's length
	}
}
