package main

import (
	"fmt"
	"time"

	"github.com/rsvato/golife/lib"
)

func main() {
	board := `
		.....
		..#..
		..#..
		..#..
		.....
		`
	f := *lib.ReadStrings(board)
	fmt.Println(f)
	for i := 0; i < 5; i++ {
		fmt.Print("\033[H\033[2J")
		f = f.Step()
		fmt.Println(f)
		time.Sleep(1 * time.Second)
	}

}
