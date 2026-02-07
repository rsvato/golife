package main

import (
	"fmt"
	"time"

	golife "github.com/rsvato/golife/lib"
)

func main() {
	board := `
		.....
		..#..
		..#..
		..#..
		.....
		`
	f := *golife.ReadStrings(board)
	fmt.Println(f)
	for i := 0; i < 5; i++ {
		fmt.Print("\033[H\033[2J")
		f = f.Step()
		fmt.Println(f)
		time.Sleep(1 * time.Second)
	}

}
