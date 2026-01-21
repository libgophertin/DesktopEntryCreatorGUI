package functions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/libgophertin/DesktopEntryCreatorGUI/models"
)

func GenerateDesktopEntry(m *models.DesktopEntryModel, w fyne.Window) {
	_, err := os.Stat(m.ExecPath)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}

	_, err = os.Stat(m.IconPath)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}

	content := `[Desktop Entry]
Type=Application
Version=1.0
Name=%s
Exec=%s
Icon=%s
Terminal=false
Categories=%s
`
	finalContent := fmt.Sprintf(content, m.AppName, m.ExecPath, m.IconPath, m.Categories)
	homePath, err := os.UserHomeDir()
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	fileName := strings.ReplaceAll(strings.ToLower(m.AppName), " ", "-")
	filePath := filepath.Join(homePath, ".local", "share", "applications", fileName+".desktop")

	err = os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	err = os.WriteFile(filePath, []byte(finalContent), 0644)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}

	dialog.ShowInformation("Success", "Desktop Entry created successfully!", w)
}
