package main

import (
	"fmt"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkWinner(winner string, err error) {
	checkErr(err)
	fmt.Println(winner)
}

func main() {
	tic := Tictactoe{}
	p1 := Player{"Marcus", "x"}
	// p2 := Player{"Lucas", "o"}

	winner, err := p1.Play(0, 1, &tic)
	checkWinner(winner, err)

	winner, err = p1.Play(1, 1, &tic)
	checkWinner(winner, err)

	winner, err = p1.Play(2, 1, &tic)
	checkWinner(winner, err)

	// winner, err = p2.Play(1, 2, &tic)
	// checkWinner(winner, err)

	fmt.Println(tic)
}
