package ui

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/libgophertin/DesktopEntryCreatorGUI/functions"
	"github.com/libgophertin/DesktopEntryCreatorGUI/models"
)

type exeFilter struct{}

func (f *exeFilter) Matches(uri fyne.URI) bool {
	ext := uri.Extension()
	allowedExts := map[string]bool{
		"":     true,
		".sh":  true,
		".bin": true,
	}
	return allowedExts[ext]
}

func (f *exeFilter) HelperText() string {
	return "Executables and Scripts (*, .sh, .bin)"
}

func (a *App) BuildForm() fyne.CanvasObject {
	w := a.MainWindow
	m := &models.DesktopEntryModel{}

	var execPathButton, iconPathButton *widget.Button

	appName := widget.NewEntry()

	execPathButton = widget.NewButton("Select File", func() {
		d := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			if reader == nil {
				return
			}

			execPathButton.SetText(reader.URI().Name())
			m.ExecPath = reader.URI().Path()
		}, w)

		d.SetFilter(&exeFilter{})
		d.Show()
	})

	iconPathButton = widget.NewButton("Select File", func() {
		d := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			if reader == nil {
				return
			}

			iconPathButton.SetText(reader.URI().Name())
			m.IconPath = reader.URI().Path()
		}, w)

		d.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".svg", ".jpg", ".jpeg", ".ico"}))
		d.Show()
	})

	categories := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "App Name", Widget: appName},
			{Text: "Executable", Widget: execPathButton},
			{Text: "Icon", Widget: iconPathButton},
			{Text: "Categories", Widget: categories},
		},
		OnSubmit: func() {
			if appName.Text == "" {
				dialog.ShowError(errors.New("App name cannot be empty"), w)
				return
			}
			m.AppName = appName.Text
			m.Categories = categories.Text
			functions.GenerateDesktopEntry(m, w)
		},
	}

	return form
}
