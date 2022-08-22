package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("blank name")
	}

	// message := fmt.Sprintf("hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi %v. Welcome !!!",
		"Hello %v. Nice to see you !!!",
		"Hey buddy. Long time %v",
	}
	return formats[rand.Intn(len(formats))]
}
