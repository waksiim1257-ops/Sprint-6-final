package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert определяет, что передано — обычный текст или код Морзе, —
// и конвертирует данные в противоположный формат.
func Convert(data string) (string, error) {
	data = strings.TrimSpace(data)
	if data == "" {
		return "", errors.New("получены пустые данные")
	}

	// Код Морзе состоит только из точек, тире и пробелов.
	if strings.Trim(data, ".- \t\r\n") == "" {
		return morse.ToText(data), nil
	}

	return morse.ToMorse(data), nil
}
