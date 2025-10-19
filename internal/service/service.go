package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDetectAndConvert(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", nil
	}

	if isMorseCode(input) {
		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}

func isMorseCode(input string) bool {
	allowedChars := ".- /"
	for _, char := range input {
		if !strings.ContainsRune(allowedChars, char) {
			return false
		}
	}
	return true
}
