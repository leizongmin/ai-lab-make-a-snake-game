package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

var (
	snakePosition = [][2]int{{5, 5}}
	foodPosition  = [2]int{10, 10}
	direction     = 'R'
	gameOver      = false
	gamePaused    = false
	gameExited    = false
	score         = 0
	screenWidth   = 0
	screenHeight  = 0
)

func clearScreen() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

func hideCursor() {
	termbox.HideCursor()
}

func processInput() {
	for !gameExited {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			handleKey(ev.Ch, ev.Key)
		}
	}
}

func handleKey(char rune, key termbox.Key) {
	if char == 'x' || char == 'X' {
		gameExited = true
		return
	}

	if key == termbox.KeySpace {
		gamePaused = !gamePaused
		if gameOver {
			gameOver = false
			snakePosition = [][2]int{{5, 5}}
			foodPosition = [2]int{10, 10}
			direction = 'R'
			score = 0
		}
		return
	}

	if !gamePaused && !gameOver {
		switch char {
		case 'a', 'A':
			direction = 'L'
		case 'd', 'D':
			direction = 'R'
		case 'w', 'W':
			direction = 'U'
		case 's', 'S':
			direction = 'D'
		}
		switch key {
		case termbox.KeyArrowLeft:
			direction = 'L'
		case termbox.KeyArrowRight:
			direction = 'R'
		case termbox.KeyArrowUp:
			direction = 'U'
		case termbox.KeyArrowDown:
			direction = 'D'
		}
	}
}

func calculateSleepDuration(length int) time.Duration {
	// Ensure the sleep duration decreases as the snake's length increases, but set a minimum limit to prevent the game from becoming unplayable.
	speed := 400 - length*20
	if speed < 100 {
		speed = 100
	}
	return time.Duration(speed) * time.Millisecond
}

func updateGame() {
	if gamePaused || gameOver {
		return
	}

	// Update snake position based on direction
	head := snakePosition[0]
	switch direction {
	case 'L':
		head[1]--
	case 'R':
		head[1]++
	case 'U':
		head[0]--
	case 'D':
		head[0]++
	}
	snakePosition = append([][2]int{head}, snakePosition[:len(snakePosition)-1]...)

	// Check for game over conditions
	if head[0] < 2 || head[0] >= screenHeight-2 || head[1] < 0 || head[1] >= screenWidth {
		gameOver = true
	}

	// Check if snake eats the food
	if head[0] == foodPosition[0] && head[1] == foodPosition[1] {
		score++
		snakePosition = append(snakePosition, foodPosition)
		// Generate new food position
		for {
			newFoodPosition := [2]int{rand.Intn(screenHeight-4) + 2, rand.Intn(screenWidth)}
			if !contains(snakePosition, newFoodPosition) {
				foodPosition = newFoodPosition
				break
			}
		}
	}
}

func contains(slice [][2]int, item [2]int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func drawScreen() {
	clearScreen()

	// Set background color for top and bottom rows and the rest of the screen
	for i := 0; i < screenWidth; i++ {
		termbox.SetCell(i, 0, ' ', termbox.ColorBlack, termbox.ColorWhite)
		termbox.SetCell(i, screenHeight-1, ' ', termbox.ColorBlack, termbox.ColorWhite)
		for j := 1; j < screenHeight-1; j++ {
			termbox.SetCell(i, j, ' ', termbox.ColorDefault, termbox.ColorBlue)
		}
	}

	if gameOver {
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

	if gamePaused {
		msg := "Paused"
		x := (screenWidth - len(msg)) / 2
		y := screenHeight / 2
		for _, c := range msg {
			termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorDefault)
			x++
		}
	}

	for _, pos := range snakePosition {
		termbox.SetCell(pos[1], pos[0], '●', termbox.ColorGreen, termbox.ColorBlue)
	}
	termbox.SetCell(foodPosition[1], foodPosition[0], '●', termbox.ColorRed, termbox.ColorBlue)

	// Display score
	scoreMsg := fmt.Sprintf("Score: %d", score)
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
	rand.Seed(time.Now().UnixNano())
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	screenWidth, screenHeight = termbox.Size()

	hideCursor()

	go processInput()

	for !gameExited {
		updateGame()
		drawScreen()
		time.Sleep(calculateSleepDuration(len(snakePosition))) // Use dynamic sleep duration based on snake's length
	}
}
