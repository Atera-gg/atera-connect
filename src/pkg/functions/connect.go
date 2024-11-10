package functions

import "atera_connect/pkg/wg"

func Connect(configFilePath string) (string, error) {
	return wg.RunCommandWithStoredPassword(AdminPassword, []string{"wg-quick", "up", configFilePath})
}

func Disconnect(configFilePath string) (string, error) {
	return wg.RunCommandWithStoredPassword(AdminPassword, []string{"wg-quick", "down", configFilePath})
}
