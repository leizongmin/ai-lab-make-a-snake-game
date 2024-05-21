package game

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type GameState struct {
	SnakePosition [][2]int
	FoodPosition  [2]int
	Direction     rune
	GameOver      bool
	GamePaused    bool
	GameExited    bool
	Score         int
	ScreenWidth   int
	ScreenHeight  int
}

func NewGameState(screenWidth, screenHeight int) *GameState {
	return &GameState{
		SnakePosition: [][2]int{{5, 5}},
		FoodPosition:  [2]int{10, 10},
		Direction:     'R',
		GameOver:      false,
		GamePaused:    false,
		GameExited:    false,
		Score:         0,
		ScreenWidth:   screenWidth,
		ScreenHeight:  screenHeight,
	}
}

func (g *GameState) HandleKey(char rune, key termbox.Key) {
	if char == 'x' || char == 'X' {
		g.GameExited = true
		return
	}

	if key == termbox.KeySpace {
		g.GamePaused = !g.GamePaused
		if g.GameOver {
			g.GameOver = false
			g.SnakePosition = [][2]int{{5, 5}}
			g.FoodPosition = [2]int{10, 10}
			g.Direction = 'R'
			g.Score = 0
		}
		return
	}

	if !g.GamePaused && !g.GameOver {
		switch char {
		case 'a', 'A':
			g.Direction = 'L'
		case 'd', 'D':
			g.Direction = 'R'
		case 'w', 'W':
			g.Direction = 'U'
		case 's', 'S':
			g.Direction = 'D'
		}
		switch key {
		case termbox.KeyArrowLeft:
			g.Direction = 'L'
		case termbox.KeyArrowRight:
			g.Direction = 'R'
		case termbox.KeyArrowUp:
			g.Direction = 'U'
		case termbox.KeyArrowDown:
			g.Direction = 'D'
		}
	}
}

func (g *GameState) CalculateSleepDuration() time.Duration {
	speed := 400 - len(g.SnakePosition)*20
	if speed < 100 {
		speed = 100
	}
	return time.Duration(speed) * time.Millisecond
}

func (g *GameState) UpdateGame() {
	if g.GamePaused || g.GameOver {
		return
	}

	head := g.SnakePosition[0]
	switch g.Direction {
	case 'L':
		head[1]--
	case 'R':
		head[1]++
	case 'U':
		head[0]--
	case 'D':
		head[0]++
	}
	g.SnakePosition = append([][2]int{head}, g.SnakePosition[:len(g.SnakePosition)-1]...)

	if head[0] < 2 || head[0] >= g.ScreenHeight-2 || head[1] < 0 || head[1] >= g.ScreenWidth {
		g.GameOver = true
	}

	if head[0] == g.FoodPosition[0] && head[1] == g.FoodPosition[1] {
		g.Score++
		g.SnakePosition = append(g.SnakePosition, g.FoodPosition)
		for {
			newFoodPosition := [2]int{rand.Intn(g.ScreenHeight-4) + 2, rand.Intn(g.ScreenWidth)}
			if !g.Contains(newFoodPosition) {
				g.FoodPosition = newFoodPosition
				break
			}
		}
	}
}

func (g *GameState) Contains(item [2]int) bool {
	for _, v := range g.SnakePosition {
		if v == item {
			return true
		}
	}
	return false
}
