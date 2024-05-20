package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/nsf/termbox-go"
)

var (
	snakePosition = [][2]int{{5, 5}}
	foodPosition  = [2]int{10, 10}
	direction     = 'R'
	gameOver      = false
	gamePaused    = false
	score         = 0
	screenWidth   = 30
	screenHeight  = 20
)

func clearScreen() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

func hideCursor() {
	termbox.HideCursor()
}

func readInput() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
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
			case keyboard.KeyArrowLeft:
				direction = 'L'
			case keyboard.KeyArrowRight:
				direction = 'R'
			case keyboard.KeyArrowUp:
				direction = 'U'
			case keyboard.KeyArrowDown:
				direction = 'D'
			}
		}

		if key == keyboard.KeySpace {
			gamePaused = !gamePaused
			if gameOver {
				gameOver = false
				snakePosition = [][2]int{{5, 5}}
				foodPosition = [2]int{10, 10}
				direction = 'R'
				score = 0
			}
		}

		if key == keyboard.KeyEsc {
			break
		}
	}
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
	if head[0] < 0 || head[0] >= screenHeight || head[1] < 0 || head[1] >= screenWidth {
		gameOver = true
	}

	// Check if snake eats the food
	if head[0] == foodPosition[0] && head[1] == foodPosition[1] {
		score++
		snakePosition = append(snakePosition, foodPosition)
		// Generate new food position
		for {
			newFoodPosition := [2]int{rand.Intn(screenHeight), rand.Intn(screenWidth)}
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
		msg := "Pause"
		x := (screenWidth - len(msg)) / 2
		y := screenHeight / 2
		for _, c := range msg {
			termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorDefault)
			x++
		}
	}

	for _, pos := range snakePosition {
		termbox.SetCell(pos[1], pos[0], '●', termbox.ColorGreen, termbox.ColorDefault)
	}
	termbox.SetCell(foodPosition[1], foodPosition[0], '●', termbox.ColorRed, termbox.ColorDefault)

	// Display score
	scoreMsg := fmt.Sprintf("Score: %d", score)
	for i, c := range scoreMsg {
		termbox.SetCell(i, screenHeight-1, c, termbox.ColorBlack, termbox.ColorWhite)
	}

	termbox.Flush()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	hideCursor()
	fmt.Println("Welcome to the Snake Game!")
	fmt.Println("Use ASDW or arrow keys to move, Space to pause/start.")

	go readInput()

	for {
		updateGame()
		drawScreen()
		time.Sleep(time.Second / 2)
	}
}
