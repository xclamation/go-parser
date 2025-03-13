package pdf

import (
	"fmt"

	"code.sajari.com/docconv/v2"
)

// PDFParser структура для парсера PDF
type PDFParser struct{}

// NewPDFParser создает новый экземпляр PDFParser
func NewPDFParser() *PDFParser {
	return &PDFParser{}
}

// Parse извлекает текст из PDF-файла
func (p *PDFParser) Parse(filePath string) (string, error) {
	res, err := docconv.ConvertPath(filePath)
	if err != nil {
		return "", fmt.Errorf("error converting PDF file: %v", err)
	}
	return res.Body, nil
}
