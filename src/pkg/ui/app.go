package ui

import (
	"atera_connect/pkg/functions"
	"fmt"
	"image/png"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var configFilePath string

func StartApplication() {
	myApp := app.NewWithID("com.atera.connect")
	myWindow := myApp.NewWindow("Atera Connect")

	header := widget.NewLabelWithStyle("Atera Connect - WireGuard Verwaltung", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	statusLabel := widget.NewLabel("Status: Getrennt")
	statusLabel.Alignment = fyne.TextAlignCenter

	configPathEntry := widget.NewEntry()
	configPathEntry.SetPlaceHolder("Pfad zur WireGuard-Konfigurationsdatei...")
	configContent := widget.NewMultiLineEntry()
	configContent.SetPlaceHolder("Konfigurationsinhalt wird hier angezeigt...")

	openConfigFile := func() {
		path, err := functions.ShowMacOSFileDialog()
		if err == nil && path != "" {
			configFilePath = path
			configPathEntry.SetText(configFilePath)
			content, err := functions.LoadConfigFileContent(configFilePath)
			if err == nil {
				configContent.SetText(content)
			} else {
				dialog.ShowError(fmt.Errorf("Fehler beim Lesen der Datei: %v", err), myWindow)
			}
		} else {
			dialog.ShowError(fmt.Errorf("Fehler beim Öffnen der Datei: %v", err), myWindow)
		}
	}

	connect := func() {
		if configFilePath == "" {
			dialog.ShowInformation("Fehler", "Bitte wähle eine Konfigurationsdatei aus.", myWindow)
			return
		}
		output, err := functions.Connect(configFilePath)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Fehler beim Verbinden: %v\n%s", err, output), myWindow)
		} else {
			statusLabel.SetText("Status: Verbunden")
			dialog.ShowInformation("Erfolgreich", "Verbindung hergestellt.", myWindow)
		}
	}

	disconnect := func() {
		if configFilePath == "" {
			dialog.ShowInformation("Fehler", "Bitte wähle eine Konfigurationsdatei aus.", myWindow)
			return
		}
		output, err := functions.Disconnect(configFilePath)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Fehler beim Trennen: %v\n%s", err, output), myWindow)
		} else {
			statusLabel.SetText("Status: Getrennt")
			dialog.ShowInformation("Erfolgreich", "Verbindung getrennt.", myWindow)
		}
	}

	connectButton := widget.NewButtonWithIcon("Verbinden", theme.ConfirmIcon(), connect)
	disconnectButton := widget.NewButtonWithIcon("Trennen", theme.CancelIcon(), disconnect)
	openFileButton := widget.NewButtonWithIcon("Datei öffnen", theme.FolderOpenIcon(), openConfigFile)

	content := container.NewVBox(
		header,
		statusLabel,
		widget.NewSeparator(),
		configPathEntry,
		openFileButton,
		configContent,
		container.NewGridWithColumns(2, connectButton, disconnectButton),
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(500, 500))
	myWindow.ShowAndRun()
}

func loadLogo(path string) fyne.Resource {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Fehler beim Laden des Logos: %v\n", err)
		return nil
	}
	defer file.Close()

	_, err = png.Decode(file)
	if err != nil {
		fmt.Printf("Fehler beim Dekodieren des Logos: %v", err)
		return nil
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Fehler beim Abrufen der Dateiinformationen: %v\n", err)
		return nil
	}

	buffer := make([]byte, fileInfo.Size())
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Printf("Fehler beim Lesen der Bilddaten: %v\n", err)
		return nil
	}

	return fyne.NewStaticResource("logo", buffer)
}
