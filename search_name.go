package main

import (
	"bytes"
	"io"
	"os"
)

func searchName(name string) (bool, error) {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return false, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return false, err
	}
	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		if bytes.Equal(line, []byte(name)) {
			return true, nil
		}
	}
	return false, nil
}
