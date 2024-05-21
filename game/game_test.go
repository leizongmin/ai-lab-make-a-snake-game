package game

import (
	"testing"

	"github.com/nsf/termbox-go"
)

func TestHandleKey(t *testing.T) {
	gameState := NewGameState(20, 10)
	gameState.HandleKey('a', termbox.KeyArrowLeft)
	if gameState.Direction != 'L' {
		t.Errorf("Expected direction 'L', got '%c'", gameState.Direction)
	}
}

func TestCalculateSleepDuration(t *testing.T) {
	gameState := NewGameState(20, 10)
	duration := gameState.CalculateSleepDuration()
	if duration <= 0 {
		t.Errorf("Expected positive duration, got %d", duration)
	}
}

func TestUpdateGame(t *testing.T) {
	gameState := NewGameState(20, 10)
	initialLength := len(gameState.SnakePosition)
	gameState.UpdateGame()
	if len(gameState.SnakePosition) != initialLength {
		t.Errorf("Expected snake length to remain the same, got %d", len(gameState.SnakePosition))
	}
}
