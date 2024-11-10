package functions

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

var AdminPassword string

func AskForMacOSPassword() bool {
	cmd := exec.Command("osascript", "-e", `display dialog "Bitte geben Sie das Administrator-Passwort für Atera-Connect ein:" default answer "" with hidden answer with title "Atera-Connect"`)
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		return false
	}
	result := output.String()
	parts := strings.Split(result, "text returned:")
	if len(parts) > 1 {
		AdminPassword = strings.TrimSpace(parts[1])
	} else {
		return false
	}
	if !verifyAdminPassword() {
		fmt.Println("Das eingegebene Passwort ist falsch. Bitte versuchen Sie es erneut.")
		return false
	}
	return true
}

func verifyAdminPassword() bool {
	cmd := exec.Command("sudo", "-k", "-S", "true")
	cmd.Stdin = bytes.NewBufferString(AdminPassword + "\n")
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	err := cmd.Run()
	return err == nil
}
