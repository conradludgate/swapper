package main

import (
	"time"
	"math/rand"
	"sort"
	"github.com/fatih/color"
	"fmt"
)

type Game struct {
	Tiles []int
}

func Abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func Diff(i, j int) int {
	return Abs(i - j)
}

// Creates a new game. Size is relative to the difficulty
// There are 4 modes that I had on the original game
func NewGame(size int) *Game {
	g := &Game{make([]int, size)}
	// g := &Game{make([]int, 2*size+1)} // Hard modes

	for i := 0; i < size; i++ {
		g.Tiles[i] = i + 1 // Easy mode
		// g.Tiles[i] = i // Medium mode
	}

	// for i := size; i < 2*size+1; i++ {
	// 	g.Tiles[i] = 2*size+1 - i // Hard mode (corresponds with easy mode)
	// 	g.Tiles[i] = 2*size - i // Extra hard mode (corresponds with medium mode)
	// }

	rand.Shuffle(len(g.Tiles), func(i, j int) {
	    g.Tiles[i], g.Tiles[j] = g.Tiles[j], g.Tiles[i]
	})
	
	return g
}

func (g *Game) Swap(i, j int) bool {
	if i >= len(g.Tiles) {
		return false
	}

	if j >= len(g.Tiles) {
		return false
	}

	if (Diff(i, j) != g.Tiles[i] && Diff(i, j) != g.Tiles[j]) {
		return false
	}

	g.Tiles[i], g.Tiles[j] = g.Tiles[j], g.Tiles[i]
	return true
}

func (g *Game) CanSwapWith(i int) []int {
	swaps := []int{}

	if i < 0 {
		return swaps
	}

	if i >= len(g.Tiles) {
		return swaps
	}

	v := g.Tiles[i]
	if i - v >= 0 {
		swaps = append(swaps, i - v)
	}

	if i + v < len(g.Tiles) {
		swaps = append(swaps, i + v)
	}

	for j, v := range g.Tiles {
		if i == j {
			continue
		}

		if Diff(i, j) == v {
			swaps = append(swaps, i + v)
		}
	}

	sort.Ints(swaps)

	return swaps
}

func (g *Game) String() string {
	return g.Selected(-1)
}

func (g *Game) Selected(j int) string {
	swaps := g.CanSwapWith(j)
	swaps = append(swaps, -1)
	si := 0

	out := "  1 "
	for i := 1; i < len(g.Tiles); i++ {
		out += fmt.Sprintf(",  %2d ", i + 1)
	}

	out += "\n"

	for i, v := range g.Tiles {
		var c func(format string, a ...interface{}) string
		if i == swaps[si] {
			si++
			c = color.RedString
		} else if i == j {
			c = color.GreenString
		} else {
			c = fmt.Sprintf
		}

		if i != 0 {
			out += ", "
		}

		out += c("[%2d]", v)
	}

	return out + "\n"
}

func main() {
	rand.Seed(time.Now().Unix())

	g := NewGame(16)
	fmt.Println(g)
	fmt.Println(g.Selected(6))
}