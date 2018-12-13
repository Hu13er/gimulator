package game

import (
	"math/rand"
	"time"
)

type Position string

const (
	UpperPos Position = "upperPlace"
	LowerPos Position = "lowerPlace"

	HeightOfMap = 13
	WidthOfMap  = 11
)

const WorldType = "World"

type World struct {
	Width      int       `json:"width"`
	Height     int       `json:"height"`
	Moves      []Move    `json:"moves"`
	Turn       string    `json:"turn"`
	BallPos    State     `json:"ball_pos"`
	Winner     string    `json:"winner"`
	Player1    Player    `json:"player1"`
	Player2    Player    `json:"player2"`
	LastAction time.Time `json:"last_action"`
}

func NewWorld(playerName1, playerName2 string) World {
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(2)

	player1 := NewPlayer(playerName1, LowerPos)
	player2 := NewPlayer(playerName2, UpperPos)
	if rnd == 1 {
		player1 = NewPlayer(playerName2, LowerPos)
		player2 = NewPlayer(playerName1, UpperPos)
	}

	world := World{
		Width:      WidthOfMap,
		Height:     HeightOfMap,
		Moves:      initMoves,
		Turn:       player1.Name,
		BallPos:    State{X: 6, Y: 7},
		Player1:    player1,
		Player2:    player2,
		LastAction: time.Now(),
	}

	return world
}

func (w *World) updateTimer(playerName string) {
	switch playerName {
	case w.Player1.Name:
		w.Player1.UpdateTimer(w.LastAction)
	case w.Player2.Name:
		w.Player2.UpdateTimer(w.LastAction)
	}
}

func (w *World) otherPlayer(playerName string) string {
	if playerName == w.Player1.Name {
		return w.Player2.Name
	}
	return w.Player1.Name
}

func (w *World) setLastAction() {
	w.LastAction = time.Now()
}

type Action struct {
	PlayerName string `json:"player_name"`
	From       State  `json:"from"`
	To         State  `json:"to"`
}

const PlayerIntroType = "PlayerIntro"

type PlayerIntro struct{}

type Player struct {
	Duration time.Duration `json:"duration"`
	Name     string        `json:"name"`
	Side     Side          `json:"side"`
}

func NewPlayer(name string, position Position) Player {
	var side Side
	if position == UpperPos {
		side = UpperSide
	} else {
		side = LowerSide
	}

	return Player{
		Name:     name,
		Side:     side,
		Duration: time.Minute * 5,
	}
}

func (p *Player) UpdateTimer(t time.Time) {
	p.Duration -= time.Now().Sub(t)
}

type Side struct {
	Pos        Position `json:"pos"`
	WinStates  []State  `json:"win_states"`
	LoseStates []State  `json:"lose_states"`
}

type State struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p *State) Equal(point State) bool {
	if p.X == point.X && p.Y == point.Y {
		return true
	}
	return false
}

type Move struct {
	A State `json:"a"`
	B State `json:"b"`
}

func (e *Move) Equal(edge Move) bool {
	if e.A.Equal(edge.A) && e.B.Equal(edge.B) {
		return true
	}
	if e.B.Equal(edge.A) && e.A.Equal(edge.B) {
		return true
	}
	return false
}
