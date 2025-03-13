package djvu

import (
	"fmt"
	"strings"

	"code.sajari.com/docconv/v2"
	"github.com/otiai10/gosseract/v2"
)

// DJVUParser структура для парсера DJVU
type DJVUParser struct{}

// NewDJVUParser создает новый экземпляр DJVUParser
func NewDJVUParser() *DJVUParser {
	return &DJVUParser{}
}

// Parse извлекает текст из DJVU-файла
func (p *DJVUParser) Parse(filePath string) (string, error) {
	// Используем docconv для извлечения текста из DJVU
	res, err := docconv.ConvertPath(filePath)
	if err != nil {
		return "", fmt.Errorf("error converting DJVU file: %v", err)
	}

	// Если текст извлечен, возвращаем его
	if res.Body != "" {
		return res.Body, nil
	}

	// Если текст не извлечен, пробуем извлечь текст из изображений с помощью Tesseract OCR
	text, err := extractTextFromImages(filePath)
	if err != nil {
		return "", fmt.Errorf("error extracting text from images: %v", err)
	}

	return text, nil
}

// extractTextFromImages извлекает текст из изображений с помощью Tesseract OCR
func extractTextFromImages(filePath string) (string, error) {
	// Используем gosseract для распознавания текста
	client := gosseract.NewClient()
	defer client.Close()

	// Указываем путь к файлу
	if err := client.SetImage(filePath); err != nil {
		return "", fmt.Errorf("error setting image: %v", err)
	}

	// Распознаем текст
	text, err := client.Text()
	if err != nil {
		return "", fmt.Errorf("error recognizing text: %v", err)
	}

	// Убираем лишние пробелы и форматируем текст
	text = strings.TrimSpace(text)
	return text, nil
}
