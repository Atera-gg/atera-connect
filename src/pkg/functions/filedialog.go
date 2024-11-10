package functions

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func ShowMacOSFileDialog() (string, error) {
	cmd := exec.Command("osascript", "-e", `set filePath to POSIX path of (choose file with prompt "Wählen Sie die Konfigurationsdatei aus:")`)
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output.String()), nil
}

func LoadConfigFileContent(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
