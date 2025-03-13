package docx

import (
	"fmt"

	"code.sajari.com/docconv/v2"
)

// DOCXParser структура для парсера DOCX
type DOCXParser struct{}

// NewDOCXParser создает новый экземпляр DOCXParser
func NewDOCXParser() *DOCXParser {
	return &DOCXParser{}
}

// Parse извлекает текст из DOCX-файла
func (p *DOCXParser) Parse(filePath string) (string, error) {
	// Используем docconv для конвертации DOCX в текст
	res, err := docconv.ConvertPath(filePath)
	if err != nil {
		return "", fmt.Errorf("error converting DOCX file: %v", err)
	}

	// Возвращаем извлеченный текст
	return res.Body, nil
}
