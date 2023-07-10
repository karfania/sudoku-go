package models

import (
	"github.com/bradhe/stopwatch"
)

type Sudoku struct {
	gameBoard Board
	time      stopwatch.Watch
}
