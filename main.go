package main

import (
	"time"
	"math/rand"
	"github.com/fatih/color"
	"fmt"
)

func main() {
	// Seed the random
	rand.Seed(time.Now().Unix())

	// Get the difficulty
	fmt.Println("What difficulty would you like to play?")
	
	diff := 0
	for diff < 1 || diff > 4 {
		fmt.Print("1 for Easy, 2 for Medium, 3 for Hard and 4 for Extra Hard > ")
		fmt.Scan(&diff)
	}

	fmt.Println("You have selected", 
		[]string{
			color.GreenString("Easy"),
			color.YellowString("Medium"), 
			color.MagentaString("Hard"), 
			color.RedString("Extra Hard"),
		}[diff-1], "mode")
	fmt.Println()

	// Get the game size
	size := 0
	for size <= 0 {
		fmt.Print("What size game board would you like to play with? ")
		fmt.Scan(&size)
	}

	fmt.Println()

	// Make the new game
	g := NewGame([]Mode{Easy, Medium, Hard, ExtraHard}[diff-1], size)
	swaps := 0

	// Intro
	fmt.Println("Your goal is to move the tiles into this layout")
	fmt.Println(&Game{g.Mode(size), g.Mode, size})

	fmt.Println("You can only swap tiles if the distance between them matches one of the values on the tiles")
	fmt.Println("Good luck")

	fmt.Println()

	// Main game loop
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
		fmt.Printf("Swaps: %3d\n", swaps)
	}

	fmt.Println(g)

	fmt.Printf("Congratulations! You won in %d swaps!", swaps + 1)
}