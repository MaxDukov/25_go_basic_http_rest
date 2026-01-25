package service

import (
	"fmt"
	"strings"
	"unicode"

	"local/pkg/morse"
)

func MorzeConvert(input string) (string, error) {
	if input == "" {
		return input, nil
	}

	// Удаляем все пробелы для проверки
	inputWithoutSpaces := strings.ReplaceAll(input, " ", "")
	
	// Проверяем, является ли строка кодом Морзе
	// Код Морзе содержит только точки и тире (после удаления пробелов)
	isMorse := true
	for _, char := range inputWithoutSpaces {
		if char != '.' && char != '-' {
			isMorse = false
			break
		}
	}

	if isMorse {
		return morse.ToText(input), nil
	}

	// Проверяем на невалидные символы перед конвертацией в код Морзе
	for _, char := range input {
		// Пропускаем пробелы
		if char == ' ' {
			continue
		}

		// Проверяем, есть ли символ в DefaultMorse
		upperChar := unicode.ToUpper(char)
		if _, exists := morse.DefaultMorse[upperChar]; !exists {
			return "", fmt.Errorf("invalid character: %c", char)
		}
	}

	return morse.ToMorse(input), nil
}
