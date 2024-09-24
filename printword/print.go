package print

import "fmt"

type W struct {
	word string
}

func (w *W) Print(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(w.word)
	}
}
