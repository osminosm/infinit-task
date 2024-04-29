package main

import (
	"fmt"
	"os"
)

type ArgumentNotFoundError struct {
	argIndex int
}

func (e ArgumentNotFoundError) Error() string {
	return fmt.Sprintf("%s %d", "Argument could not be found at index", e.argIndex)
}

func getArgument(index int) (string, error) {
	if len(os.Args) > index+1 {
		return os.Args[index+1], nil
	}

	return "", ArgumentNotFoundError{argIndex: index}
}
