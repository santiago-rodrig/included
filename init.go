package main

import (
	"os"
	"path/filepath"
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("unable to get user HOME directory")
	}
	filePath = filepath.Join(homeDir, ".included", "names.txt")
	createFolder()
}
