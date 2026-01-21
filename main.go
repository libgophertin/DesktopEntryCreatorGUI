package main

import "github.com/libgophertin/DesktopEntryCreatorGUI/ui"

func main() {
	myApp := &ui.App{}

	myApp.NewUI()
	myApp.Run()
}
