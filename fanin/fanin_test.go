package fallin

import (
	"fmt"
	"testing"
)

func TestFallIn(t *testing.T) {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
