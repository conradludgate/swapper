package main

import (
	"fmt"
	"sort"
	"math/rand"
	"github.com/fatih/color"
)

type Game struct {
	Tiles []int
	Mode Mode
	Size int
}

// Creates a new game. Size is relative to the difficulty
// There are 4 modes that I had on the original game
func NewGame(mode Mode, size int) *Game {
	g := &Game{mode(size), mode, size}

	rand.Shuffle(len(g.Tiles), func(i, j int) {
	    g.Tiles[i], g.Tiles[j] = g.Tiles[j], g.Tiles[i]
	})
	
	return g
}

func (g *Game) Won() bool {
	win := g.Mode(g.Size)

	for i, v := range g.Tiles {
		if v != win[i] {
			return false
		}
	}

	return true
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

func (g *Game) Swap(i, j int) bool {
	if i < 0 || i >= len(g.Tiles) {
		return false
	}

	if j < 0 || j >= len(g.Tiles) {
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
			swaps = append(swaps, j)
		}
	}

	sort.Ints(swaps)

	return swaps
}

func (g *Game) String() string {
	return g.Selected(-1) // Nothing selected
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

	win := g.Mode(g.Size)

	for i, v := range g.Tiles {
		var c func(format string, a ...interface{}) string
		if i == swaps[si] && v == win[i] {
			si++
			c = color.MagentaString
		} else if i == swaps[si] {
			si++
			c = color.RedString
		} else if i == j {
			c = color.YellowString
		} else if v == win[i] {
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