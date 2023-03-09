package stupidstuff

import (
	"fmt"
	"time"
)

type Texter struct {
	text string
}

func NewTexter(text string) Texter {
	return Texter{text: text}
}

func (t Texter) Scroll(duration int64) {
	ClearScreen()
	for {
		t.text += " "
		var shifted string
		char := []byte(t.text)
		shifted = string(append(char[len(char)-1:], char[0:len(char)-1]...))
		t.text = shifted
		t.print()
		time.Sleep(time.Duration(duration) * time.Millisecond)
		ClearScreen()
	}
}

func (t Texter) print() {
	fmt.Print(t.text)
}
