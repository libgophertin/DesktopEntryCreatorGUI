package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type App struct {
	App        fyne.App
	MainWindow fyne.Window
}

func (a *App) NewUI() {
	a.App = app.New()
	a.MainWindow = a.App.NewWindow("Desktop Entry Creator")
}

func (a *App) Run() {
	content := a.BuildForm()
	a.MainWindow.SetContent(content)
	//w.SetContent(container.NewVBox(appNameLabel, appName, execFileLabel, execPathButton, iconFileLabel, iconPathButton, categoryLabel, category))
	a.MainWindow.Resize(fyne.NewSize(400, 600))
	a.MainWindow.ShowAndRun()
}
