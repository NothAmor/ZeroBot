package safe

import (
	"fmt"
	"testing"
	"time"
)

func TestGoTickets(t *testing.T) {
	GoTickets, _ := NewGoTickets(10)
	for i := 0; i < 100; i++ {
		GoTickets.Take()
		index := i
		Go(func() {
			fmt.Printf("i:%d \n", index)
			time.Sleep(3 * time.Second)
			defer GoTickets.Return()
		})
	}
}
