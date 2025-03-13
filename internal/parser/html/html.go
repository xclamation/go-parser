package html

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

// HTMLParser структура для парсера HTML
type HTMLParser struct{}

// NewHTMLParser создает новый экземпляр HTMLParser
func NewHTMLParser() *HTMLParser {
	return &HTMLParser{}
}

// Parse извлекает весь текст из HTML-файла
func (p *HTMLParser) Parse(filePath string) (string, error) {
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close() // Закрываем файл при завершении функции

	// Создаем goquery-документ
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return "", fmt.Errorf("error parsing HTML: %v", err)
	}

	// Извлекаем весь текст из документа
	text := doc.Text()

	// Форматируем текст
	formattedText := formatText(strings.Split(text, "\n"))

	return formattedText, nil
}

// formatText форматирует текст, добавляя отступы для абзацев и разделов
func formatText(lines []string) string {
	var buf strings.Builder
	var lastIsNewline bool

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			if !lastIsNewline {
				buf.WriteByte('\n')
				lastIsNewline = true
			}
		} else {
			// Добавляем отступ для абзаца
			if buf.Len() > 0 && !lastIsNewline {
				buf.WriteString("\n    ") // 4 пробела для отступа
			}
			buf.WriteString(line)
			lastIsNewline = false
		}
	}

	// Убираем лишние пробелы и возвращаем текст
	result := strings.TrimRightFunc(buf.String(), unicode.IsSpace)
	return strings.ReplaceAll(result, "\n ", "\n")
}
