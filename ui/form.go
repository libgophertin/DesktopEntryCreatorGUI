package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/libgophertin/DesktopEntryCreatorGUI/models"
)

func (a *App) BuildForm() fyne.CanvasObject {
	w := a.MainWindow
	m := models.DesktopEntryModel{}

	appName := widget.NewEntry()

	execPathButton := widget.NewButton("Select file", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			if reader == nil {
				return
			}

			m.ExecPath = reader.URI().Path()
		}, w)
	})

	iconPathButton := widget.NewButton("Select file", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			if reader == nil {
				return
			}

			m.IconPath = reader.URI().Path()
		}, w)
	})

	categories := widget.NewEntry()

	m.AppName = appName.Text
	m.Categories = categories.Text

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "App name", Widget: appName},
			{Text: "Exec", Widget: execPathButton},
			{Text: "Icon", Widget: iconPathButton},
			{Text: "Categories", Widget: categories},
		},
	}

	return form
}
