package main

import "errors"

var (
	InvalidPlayError = errors.New("Invalid play")
)

type Tictactoe struct {
	board [3][3]string
}

func (t Tictactoe) String() string {
	var output string

	for _, i := range t.board {
		for p, a := range i {

			if a == "" {
				output += " "
			} else {
				output += a
			}

			if p != 2 {
				output += " | "
			}

		}
		output += "\n"
	}

	return output
}

func (t Tictactoe) checkLineWin(winner chan string, done chan bool) {
	for _, i := range t.board {
		if i[0] != "" && i[0] == i[1] && i[1] == i[2] {
			winner <- i[0]
		}
	}

	done <- true
	t.checkDone(done, winner)
}

func (t Tictactoe) checkColumnWin(winner chan string, done chan bool) {
	for i, _ := range t.board {
		if t.board[0][i] != "" && t.board[0][i] == t.board[1][i] && t.board[1][i] == t.board[2][i] {
			winner <- t.board[0][i]
		}
	}

	done <- true
	t.checkDone(done, winner)
}

func (t Tictactoe) checkDiagonalWin(winner chan string, done chan bool) {
	if t.board[0][0] != "" && t.board[0][0] == t.board[1][1] && t.board[1][1] == t.board[2][2] {
		winner <- t.board[0][0]
	}

	if t.board[0][2] != "" && t.board[0][2] == t.board[1][1] && t.board[1][1] == t.board[2][0] {
		winner <- t.board[0][2]
	}

	done <- true
	t.checkDone(done, winner)
}


func (t Tictactoe) checkDone(done chan bool, winner chan string){
	// Check if go routine is finished, then close the channel to avoid deadlocks
	if len(done) == 3 {
		close(winner)
	}
}

func (t Tictactoe) WhoWon() string {
	winner := make(chan string, 1)
	done := make(chan bool, 3)

	go t.checkLineWin(winner, done)
	go t.checkColumnWin(winner, done)
	go t.checkDiagonalWin(winner, done)	


	msg := <-winner

	if msg != "" {
		return msg
	}
	// Check if it is a draw
	gameFinished := true
	for _, i := range t.board {
		for _, j := range i {
			if j == "" {
				gameFinished = false
			}
		}
	}

	if gameFinished {
		return "Draw"
	} else {
		return "Make your move..."
	}
}

func (t *Tictactoe) Play(x, y int, symbol string) (string, error) {
	if t.board[x][y] == "" {
		t.board[x][y] = symbol
	} else {
		return "", InvalidPlayError
	}

	return t.WhoWon(), nil
}
