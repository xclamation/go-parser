package doc

import (
	"code.sajari.com/docconv/v2"
	"fmt"
)

// DOCParser структура для парсера DOC
type DOCParser struct{}

// NewDOCParser создает новый экземпляр DOCParser
func NewDOCParser() *DOCParser {
	return &DOCParser{}
}

// Parse извлекает текст из DOC-файла
func (p *DOCParser) Parse(filePath string) (string, error) {
	// Используем docconv для конвертации DOC в текст
	res, err := docconv.ConvertPath(filePath)
	if err != nil {
		return "", fmt.Errorf("error converting DOC file: %v", err)
	}

	// Возвращаем извлеченный текст
	return res.Body, nil
}
