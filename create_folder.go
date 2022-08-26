package main

import (
	"os"
	"path"
)

func createFolder() error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	_, err = os.ReadDir(path.Join(homedir, dirName))
	if err != nil {
		err = os.Mkdir(path.Join(homedir, dirName), 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
