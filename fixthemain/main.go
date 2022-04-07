package main

import "github.com/01-edu/z01"

type Door struct {
	OPEN  bool
	CLOSE bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(state *Door) bool {
	PrintStr("Door Opening...\n")
	state.OPEN = true
	return true
}

func CloseDoor(state *Door) bool {
	PrintStr("Door Closing...\n")
	state.CLOSE = false
	return false
}

func IsDoorOpen(state *Door) bool {
	PrintStr("is the Door opened ?\n")
	state.OPEN = true
	return true
}

func IsDoorClose(state *Door) bool {
	PrintStr("is the Door closed ?\n")
	state.CLOSE = false
	return false
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
}
