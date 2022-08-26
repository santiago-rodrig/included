package main

import "os"

func clearNames() error {
	return os.Remove(filePath)
}
