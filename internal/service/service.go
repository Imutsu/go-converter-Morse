package service

import  (
	"strings"
	"errors"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDetect(input string) (string, error) {
	// убрать пробелы по краям
	input = strings.TrimSpace(input)
	
	// проверяем строку
	if input == "" {
		return "", errors.New("empty input")
	}
	
	// проверить морзе это или текст
	isMorse := true

	for _, r:= range input {
		if r != '.' && r != '-' && r != ' ' {
			isMorse = false
			break
		}
	} 

	// возвращаем результат
	if isMorse {
		result := morse.ToText(input)
		return result, nil
	}

	result := morse.ToMorse(input)
	return result, nil
}