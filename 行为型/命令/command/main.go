package main

import "command/internal"

func main() {
	tv := &internal.Tv{}

	onCommand := &internal.OnCommand{
		Device: tv,
	}

	offCommand := &internal.OffCommand{
		Device: tv,
	}

	onButton := &internal.Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton := &internal.Button{
		Command: offCommand,
	}
	offButton.Press()
}
