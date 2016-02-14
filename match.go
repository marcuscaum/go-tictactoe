package main

type Match struct {
  players [2]Player
  game *Tictactoe
}

type Game interface {
  Play(x, y int, symbol string) (string, error)
}

func initMatch(playerNames [2]string, game *Tictactoe) Match {
  var players [2]Player
  symbols := [2]string{"x", "o"}
  for i, p := range playerNames {
    players[i] = Player{p, symbols[i]}
  }
  return Match{players, game}
}

func (m Match) start() string {
  // It manages players turns
  // Ask player to make the move
  // Returns winner player
}