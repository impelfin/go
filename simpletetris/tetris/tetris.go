package tetris
const BOARD_HEIGHT = 16
const BOARD_WIDTH = 10 

type game struct {
	board [][]int
}

func (g *game)GetBoard(params) [][]int {
	return g.board
}

func (g *game) init() {
	g.board = make([][]int, BOARD_HEIGHT )
	for y := 0; y < BOARD_HEIGHT; y++ {
		g.board[y] = make([]int, BOARD_WIDTH)
		for x := 0; x < BOARD_WIDTH; x++ {
			g.board[y][x] = 0
		}
	}
}

func New() *game{
	g := &game{}
	g.init()
	return g
}