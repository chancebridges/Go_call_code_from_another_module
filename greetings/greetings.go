package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// If no name was give, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
