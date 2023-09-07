package main

import {
	"fmt",
	"github.com/impelfin/go/simpletetris/tetris",
	"github.com/impelfin/go/simpletetris/screen"
}
 
func main() {
	fmt.Println("Tetris")

	game := tetris.New()
	scr := screen.New()

	scr.RenderAsci(game.GetBoard())
}
