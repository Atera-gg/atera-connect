package main

import (
	"atera_connect/pkg/functions"
	"atera_connect/pkg/ui"
)

func main() {
	if !functions.AskForMacOSPassword() {
		return
	}
	ui.StartApplication()
}
