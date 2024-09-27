package main

import (
	"ethical-ddos/cmd/packages/logging"
	"os"
)

func init() {
	os.Setenv("KEY_WORD", "cmd")
}

func main() {
	info := logging.NewLogger(logging.INFO)

	info.Info("something")
}
