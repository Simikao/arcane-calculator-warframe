package main

import (
	"errors"
	"fmt"
)

var neededPerRank = []int{0, 3, 6, 10, 15, 21}

func calculateRemainingArcanes(currentRank, spares int) (int, error) {
	if currentRank == 5 {
		return 0, errors.New("You can already reach rank 5, no more arcanes needed")
	}

	if currentRank < 0 || currentRank > 5 {
		return 0, errors.New("Invalid rank, rank should be between 0 and 5")
	}

	neededArcanes := neededPerRank[5] - neededPerRank[currentRank] - spares

	if neededArcanes < 0 {
		neededArcanes = 0
	}
	return neededArcanes, nil
}

func main() {
	var currentRank, spares int

	fmt.Println("Enter your current max arcane rank")
	if _, err := fmt.Scanf("%d", &currentRank); err != nil {
		fmt.Println("Invalid input, please enter a number")
		return
	}

	fmt.Println("Enter the number of rank 0 arcanes you have")
	if _, err := fmt.Scanf("%d", &spares); err != nil {
		fmt.Println("Invalid input, please enter a number")
		return
	}

	needed, err := calculateRemainingArcanes(currentRank, spares)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("You need to buy %d more rank 0 arcanes to reach rank 5.\n", needed)
}
