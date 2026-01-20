package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/ArtexOS/DesktopEntryCreatorGUI/config"
)

func main() {
	a := app.New()
	w := a.NewWindow("Выбор файла")
	var conf config.AppConfig

	appNameLabel := widget.NewLabel("App Name:")
	appName := widget.NewEntry()

	conf.AppName = appName.Text
	execFileLabel := widget.NewLabel("File is not selected")
	execPathButton := widget.NewButton("Select file", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			if reader == nil {
				return
			}

			conf.ExecPath = reader.URI().Path()
			execFileLabel.SetText(conf.ExecPath)
		}, w)
	})

	iconFileLabel := widget.NewLabel("File is not selected")
	iconPathButton := widget.NewButton("Select file", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			if reader == nil {
				return
			}

			conf.IconPath = reader.URI().Path()
			iconFileLabel.SetText(conf.IconPath)
		}, w)
	})

	categoryLabel := widget.NewLabel("Enter category:")
	category := widget.NewEntry()

	w.SetContent(container.NewVBox(appNameLabel, appName, execFileLabel, execPathButton, iconFileLabel, iconPathButton, categoryLabel, category))
	w.Resize(fyne.NewSize(400, 600))
	w.ShowAndRun()
}
