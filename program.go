package main

import (
	"fmt"
	"os"
)

func main() {
	printUsage()
	for {
		input, err := getInput()
		if err != nil {
			printError(err)
			os.Exit(1)
		}
		switch true {
		case addCommandRegexp.MatchString(input):
			matches := addCommandRegexp.FindStringSubmatch(input)
			err := writeName(matches[1])
			if err != nil {
				printError(err)
				os.Exit(1)
			}
			fmt.Println(nameAddedMessage)
		case deleteCommandRegexp.MatchString(input):
			matches := deleteCommandRegexp.FindStringSubmatch(input)
			err := deleteName(matches[1])
			if err != nil {
				printError(err)
				os.Exit(1)
			}
			fmt.Println(nameDeletedMessage)
		case searchCommandRegexp.MatchString(input):
			matches := searchCommandRegexp.FindStringSubmatch(input)
			found, err := searchName(matches[1])
			if err != nil {
				printError(err)
				os.Exit(1)
			}
			if found {
				fmt.Println(nameFoundMessage)
			} else {
				fmt.Println(nameNotFoundMessage)
			}
		case clearCommand == input:
			err := clearNames()
			if err != nil {
				printError(err)
				os.Exit(1)
			}
			fmt.Println(namesClearedMessage)
		case quitCommand == input:
			os.Exit(0)
		default:
			fmt.Println(unknownCommandMessage)
		}
	}
}
