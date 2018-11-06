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
	Mode Mode
	Size int
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

	win := g.Mode(g.Size)

	for i, v := range g.Tiles {
		var c func(format string, a ...interface{}) string
		if i == swaps[si] {
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

func main() {
	rand.Seed(time.Now().Unix())

	g := NewGame(Easy, 16)
	swaps := 0

	selected := -1
	for !g.Won() {
		if selected == -1 {
			fmt.Println(g)

			for selected <= 0 || selected > len(g.Tiles) {
				fmt.Print("Select a tile to swap > ")
				fmt.Scanln(&selected)
			}
			
			selected--

			if len(g.CanSwapWith(selected)) == 0 {
				fmt.Println("No swaps available with this tile")
				selected = -1
			}

		} else {
			fmt.Println(g.Selected(selected))

			choices := g.CanSwapWith(selected)
			choice := -1
			for choice == -1 {
				fmt.Print("Select a tile to swap with (Choose 0 to cancel selection)> ")
				fmt.Scanln(&choice)

				choice--

				if choice == -1 {
					break
				}

				match := false
				for _, v := range choices {
					if choice == v {
						match = true
						break
					}
				}

				if !match {
					choice = -1
				}
			}

			if g.Swap(selected, choice) {
				swaps++
			}

			selected = -1
		}

		fmt.Println()
	}

	fmt.Println(g)

	fmt.Printf("Congratulations! You won in %d swaps!", swaps + 1)
}