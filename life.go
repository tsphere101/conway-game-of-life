package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	// Game of Life in Go

	// Create a new game
	game := NewGame(150, 50, "*", "-")

	// Create a new pattern builder
	patternBuilder := NewPatternBuilder()

	// Create a glider
	pattern := patternBuilder.Type("glider").Build()

	// Add the pattern to the game
	game.Add(pattern, 10, 10)

	game.Run()

}

// Step advances the game by one step
func (g *Game) Step() {
	g.state = g.Next()
}

type State [][]bool

// Game is a struct that holds the state of the game
type Game struct {
	state  State  // The current state of the game
	width  int    // The width of the game
	height int    // The height of the game
	life   string // The string representation of a live cell
	nolife string // The string representation of a dead cell
}

// NewGame creates a new game with the given state
func NewGame(width int, height int, life string, nolife string) *Game {
	state := make([][]bool, height)
	for i := range state {
		state[i] = make([]bool, width)
	}
	return &Game{state: state, width: width, height: height, life: life, nolife: nolife}
}

// Add add a new pattern to the game at the given coordinates
func (g *Game) Add(pattern Pattern, x, y int) {
	for i := range pattern {
		for j := range pattern[i] {
			g.state[x+i][y+j] = pattern[i][j]
		}
	}
}

// Next returns the next state of the game
func (g *Game) Next() [][]bool {
	next := make([][]bool, len(g.state))
	for i := range g.state {
		next[i] = make([]bool, len(g.state[i]))
		for j := range g.state[i] {
			next[i][j] = g.nextCellState(i, j)
		}
	}
	return next
}

// nextCellState returns the next state of the cell at (x, y)
func (g *Game) nextCellState(x, y int) bool {
	alive := g.state[x][y]
	neighbors := g.countNeighbors(x, y)
	if alive {
		return neighbors == 2 || neighbors == 3
	}
	return neighbors == 3
}

// countNeighbors returns the number of alive neighbors of the cell at (x, y)
func (g *Game) countNeighbors(x, y int) int {
	neighbors := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if g.isAlive(i, j) {
				neighbors++
			}
		}
	}
	return neighbors
}

// isAlive returns true if the cell at (x, y) is alive
func (g *Game) isAlive(x, y int) bool {
	if x < 0 || x >= len(g.state) {
		return false
	}
	if y < 0 || y >= len(g.state[x]) {
		return false
	}
	return g.state[x][y]
}

// String returns a string representation of the game
func (g *Game) String() string {
	// Let * represent alive cells and . represent dead cells
	var b strings.Builder
	for i := range g.state {
		for j := range g.state[i] {
			if g.state[i][j] {
				b.WriteString(g.life)
			} else {
				b.WriteString(g.nolife)
			}
		}

		// Add a newline after each row
		b.WriteString("\n")
	}

	return b.String()
}

// Run is the game loop
func (g *Game) Run() {
	for {
		// Print the game state
		fmt.Println(g)

		// Go to the next step after 17 milliseconds
		time.Sleep(17 * time.Millisecond)
		g.Step()

		// Clear the screen
		fmt.Print("\033[H\033[2J")
	}
}

// Rotate rotates the given slice of slices by 90 degrees
func Rotate(arr [][]bool, degrees int) [][]bool {
	// Create a new slice of slices
	var rotated [][]bool

	// Get the number of rows and columns
	rows := len(arr)
	cols := len(arr[0])

	// Create the new slice of slices
	for i := 0; i < cols; i++ {
		var row []bool
		for j := 0; j < rows; j++ {
			row = append(row, false)
		}
		rotated = append(rotated, row)
	}

	// Copy the values
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = arr[i][j]
		}
	}
	if degrees == 90 {
		return rotated
	}
	if degrees == 180 {
		return Rotate(rotated, 90)
	}
	if degrees == 270 {
		return Rotate(rotated, 180)
	}
	return arr
}

// PatternBuilder is a struct that build a pattern
type PatternBuilder struct {
	pattern [][]bool
}

// NewPatternBuilder creates a new PatternBuilder
func NewPatternBuilder() *PatternBuilder {
	return &PatternBuilder{}
}

// Type set the type of the pattern
func (p *PatternBuilder) Type(t string) *PatternBuilder {
	switch t {
	case "glider": // A glider
		p.pattern = [][]bool{
			{false, true, false},
			{false, false, true},
			{true, true, true},
		}
		break
	case "blinker": // A blinker
		p.pattern = [][]bool{
			{true},
			{true},
			{true},
		}
		break
	case "toad": // A toad
		p.pattern = [][]bool{
			{false, true, true, true},
			{true, true, true, false},
		}
		break
	case "beacon": // A beacon
		p.pattern = [][]bool{
			{true, true, false, false},
			{true, true, false, false},
			{false, false, true, true},
			{false, false, true, true},
		}
		break
	case "pulsar": // A pulsar
		p.pattern = [][]bool{
			{false, false, false, true, false, false, false, true, false, false, false},
			{false, false, false, true, false, false, false, true, false, false, false},
			{false, false, false, true, false, false, false, true, false, false, false},
			{true, true, true, false, true, true, true, false, true, true, true},
			{false, false, false, true, false, false, false, true, false, false, false},
			{false, false, false, true, false, false, false, true, false, false, false},
			{false, false, false, true, false, false, false, true, false, false, false},
			{true, true, true, false, true, true, true, false, true, true, true},
			{false, false, false, true, false, false, false, true, false, false, false},
			{false, false, false, true, false, false, false, true, false, false, false},
			{false, false, false, true, false, false, false, true, false, false, false},
		}
		break
	case "oscillator": // An oscillator
		p.pattern = [][]bool{
			{true, true, true},
		}
		break

	case "spaceship": // A spaceship
		p.pattern = [][]bool{
			{false, true, true, false},
			{true, false, false, true},
			{false, true, false, true},
			{false, false, true, true},
		}
		break

	}

	return p
}

// Face rotate the pattern to the given face
func (p *PatternBuilder) Face(face string) *PatternBuilder {
	// Rotate the pattern to the given face
	switch face {
	case "up":
		p.pattern = Rotate(p.pattern, 0)
		break
	case "right":
		p.pattern = Rotate(p.pattern, 90)
		break
	case "down":
		p.pattern = Rotate(p.pattern, 180)
		break
	case "left":
		p.pattern = Rotate(p.pattern, 270)
		break
	}
	return p
}

// Build builds the pattern
func (p *PatternBuilder) Build() [][]bool {
	return p.pattern
}

// Pattern is a struct that represents a pattern
type Pattern [][]bool
