package main

import "regexp"

var filePath string

const dirName = ".included"

const usageMessage = `
USAGE:
	add <name>	-- Adds a name
	del <name>	-- Deletes a name
	has <name>	-- Searches a name
	q		-- Exits the program
	h		-- Prints this message

`

const nameFoundMessage = "INCLUDED"
const nameNotFoundMessage = "NOT INCLUDED"
const nameAddedMessage = "ADDED"
const nameDeletedMessage = "DELETED"
const unknownCommandMessage = "UNKOWN COMMAND"
const namesClearedMessage = "CLEARED"

const addCommand = "add"
const deleteCommand = "del"
const searchCommand = "has"
const clearCommand = "clc"
const quitCommand = "q"

var addCommandRegexp = regexp.MustCompile(`^` + addCommand + ` ([a-zA-Z0-9\_\.]+)$`)
var deleteCommandRegexp = regexp.MustCompile(`^` + deleteCommand + ` ([a-zA-Z0-9\_\.]+)$`)
var searchCommandRegexp = regexp.MustCompile(`^` + searchCommand + ` ([a-zA-Z0-9\_\.]+)$`)

const inputPrefix = "INPUT: "
