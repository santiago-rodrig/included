package main

import (
	"bytes"
	"io"
	"os"
)

func deleteName(name string) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	lines := bytes.Split(data, []byte("\n"))
	newLines := make([][]byte, 0)
	for _, line := range lines {
		if !bytes.Equal(line, []byte(name)) {
			newLines = append(newLines, line)
		}
	}
	_, err = f.Write(bytes.Join(newLines, []byte("\n")))
	if err != nil {
		return err
	}
	_, err = f.Write([]byte("\n"))
	if err != nil {
		return err
	}
	return nil
}
