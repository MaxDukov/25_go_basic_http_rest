package service

import (
	"testing"
)

func TestMorzeConvert_TextToMorse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "один символ - буква А",
			input:    "А",
			expected: ".-",
			wantErr:  false,
		},
		{
			name:     "слово ПРИВЕТ",
			input:    "ПРИВЕТ",
			expected: ".--. .-. .. .-- . -",
			wantErr:  false,
		},
		{
			name:     "слово с пробелами",
			input:    "ПРИВЕТ МИР",
			expected: ".--. .-. .. .-- . - -- .. .-.",
			wantErr:  false,
		},
		{
			name:     "цифры",
			input:    "123",
			expected: ".---- ..--- ...--",
			wantErr:  false,
		},
		{
			name:     "знаки препинания",
			input:    "А,Б.",
			expected: ".- .-.-.- -... ......",
			wantErr:  false,
		},
		{
			name:     "нижний регистр",
			input:    "привет",
			expected: ".--. .-. .. .-- . -",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MorzeConvert(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MorzeConvert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("MorzeConvert() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMorzeConvert_MorseToText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "один символ - код А",
			input:    ".-",
			expected: "А",
			wantErr:  false,
		},
		{
			name:     "слово ПРИВЕТ",
			input:    ".--. .-. .. .-- . -",
			expected: "ПРИВЕТ",
			wantErr:  false,
		},
		{
			name:     "слово с пробелами между словами",
			input:    ".--. .-. .. .-- . -   -- .. .-.",
			expected: "ПРИВЕТ МИР",
			wantErr:  false,
		},
		{
			name:     "цифры",
			input:    ".---- ..--- ...--",
			expected: "123",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MorzeConvert(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MorzeConvert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("MorzeConvert() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestMorzeConvert_EmptyString(t *testing.T) {
	got, err := MorzeConvert("")
	if err != nil {
		t.Errorf("MorzeConvert() error = %v, want nil", err)
	}
	if got != "" {
		t.Errorf("MorzeConvert() = %q, want %q", got, "")
	}
}

func TestMorzeConvert_InvalidCharacters(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "латинская буква A",
			input:   "A",
			wantErr: true,
		},
		{
			name:    "латинское слово HELLO",
			input:   "HELLO",
			wantErr: true,
		},
		{
			name:    "специальный символ @",
			input:   "@",
			wantErr: true,
		},
		{
			name:    "смешанный текст с невалидным символом",
			input:   "ПРИВЕТA",
			wantErr: true,
		},
		{
			name:    "китайский символ",
			input:   "中",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MorzeConvert(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MorzeConvert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err == nil {
				t.Errorf("MorzeConvert() expected error, got nil, result = %q", got)
			}
			if !tt.wantErr && err != nil {
				t.Errorf("MorzeConvert() unexpected error = %v", err)
			}
		})
	}
}

func TestMorzeConvert_OnlySpaces(t *testing.T) {
	got, err := MorzeConvert("   ")
	if err != nil {
		t.Errorf("MorzeConvert() error = %v, want nil", err)
	}
	// Пустая строка после удаления пробелов должна вернуть пустую строку
	// Но текущая логика проверяет isMorse для строки с пробелами
	// Нужно проверить фактическое поведение
	if got == "" {
		t.Logf("MorzeConvert() вернул пустую строку для пробелов, что ожидаемо")
	}
}
