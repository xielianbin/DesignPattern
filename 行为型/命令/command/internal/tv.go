package internal

import "fmt"

type Tv struct {
	isRunning bool
}

func (t *Tv) On() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) Off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
