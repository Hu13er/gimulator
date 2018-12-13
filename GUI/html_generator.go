package main

import (
	"fmt"
	"strconv"
	"time"
)

type worldDrawer struct {
	World
	width, height int
}

func (w worldDrawer) DrawField() string {
	var (
		html    = ""
		delta   = min(w.width/(WidthOfMap+1), w.height/(HeightOfMap+1))
		marginx = (w.width - delta*(WidthOfMap-1)) / 2
		marginy = (w.height - delta*(HeightOfMap-1)) / 2
	)
	fmt.Println(w.width, w.height)
	fmt.Println(marginx, marginy)
	fmt.Println(delta)

	grid := make([][]State, WidthOfMap+1)
	for i := 0; i < WidthOfMap+1; i++ {
		grid[i] = make([]State, HeightOfMap+1)
	}

	for x := 0; x < WidthOfMap; x++ {
		for y := 0; y < HeightOfMap; y++ {
			xx := marginx + x*delta
			yy := marginy + y*delta
			html += newCircle(xx, yy, 5, "yellow")
			grid[x+1][y+1] = State{xx, yy}
		}
	}

	for _, move := range initMoves {
		html += newLine(grid[move.A.X][move.A.Y].X, grid[move.A.X][move.A.Y].Y, grid[move.B.X][move.B.Y].X, grid[move.B.X][move.B.Y].Y, "red")
	}

	return html
}

func (w worldDrawer) genUpperSpec() (string, string) {
	if w.Player1.Side.Pos == UpperPos {
		return w.Player1.Name, timeConverter(w.Player1.Duration)
	}
	if w.Player2.Side.Pos == UpperPos {
		return w.Player2.Name, timeConverter(w.Player2.Duration)
	}
	return "No Player", "00:00"
}

func (w worldDrawer) genLowerSpec() (string, string) {
	if w.Player1.Side.Pos == LowerPos {
		return w.Player1.Name, timeConverter(w.Player1.Duration)
	}
	if w.Player2.Side.Pos == LowerPos {
		return w.Player2.Name, timeConverter(w.Player2.Duration)
	}
	return "No Player", "00:00"
}

func (w worldDrawer) genTurn() string {
	if w.Turn == "" {
		return "Turn: -"
	}
	return "Turn: " + w.Turn
}

func timeConverter(duration time.Duration) string {
	min := int(duration.Minutes())
	sec := int(duration.Seconds()) - min*60

	return strconv.Itoa(min) + ":" + strconv.Itoa(sec)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func newCircle(cx, cy, r int, col string) string {
	return fmt.Sprintf(`<circle cx="%d" cy="%d" r="%d" stroke="black" stroke-width="1" fill="%s" />`, cx, cy, r, col)
}

func newLine(x1, y1, x2, y2 int, col string) string {
	return fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="%s" stroke-width="2" />`, x1, y1, x2, y2, col)
}

