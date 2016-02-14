package main

type Player struct {
	name   string
	symbol string
}

func (p Player) Play(x, y int, tic *Tictactoe) (string, error) {
	return tic.Play(x, y, p.symbol)
}
