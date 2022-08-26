package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	_, err := fmt.Print(inputPrefix)
	if err != nil {
		return "", err
	}
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.Trim(input, "\r\n ")
	return input, nil
}
