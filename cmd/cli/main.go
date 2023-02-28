package main

import (
	"fmt"
	"time"
)

type texter struct {
	text string
}

func newTexter(text string) texter {
	return texter{text: text}
}

func (t texter) scroll() {
	clearScreen()
	for {
		t.text += " "
		var shifted string
		char := []byte(t.text)
		shifted = string(append(char[len(char)-1:], char[0:len(char)-1]...))
		t.text = shifted
		t.print()
		time.Sleep(time.Duration(90) * time.Millisecond)
		clearScreen()
	}
}

func (t texter) print() {
	fmt.Print(t.text)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	t := newTexter("jokes on you!")
	t.scroll()
}
