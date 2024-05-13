package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func main() {
	fmt.Println("Switch case in golang")

	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// diceNumber := r.Intn(6) + 1

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You rolled a 1. Move 1 step forward.")
	case 2:
		fmt.Println("You rolled a 2. Move 2 steps forward.")
	case 3:
		fmt.Println("You rolled a 3. Move 3 steps forward.")
		//fallthrough
	case 4:
		fmt.Println("You rolled a 4. Move 4 steps forward.")
	case 5:
		fmt.Println("You rolled a 5. Move 5 steps forward.")
	case 6:
		fmt.Println("You rolled a 6. Move 6 steps forward.")
	default:
		fmt.Println("Unexpected dice outcome.") // Shouldn't happen
	}
}
