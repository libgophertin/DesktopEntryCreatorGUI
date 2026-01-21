package main

import "github.com/libgophertin/DesktopEntryCreatorGUI/ui"

//func main() {
//	//categoryLabel := widget.NewLabel("Enter category:")
//	//category := widget.NewEntry()
//}

func main() {
	myApp := &ui.App{}

	myApp.NewUI()
	myApp.Run()
}
