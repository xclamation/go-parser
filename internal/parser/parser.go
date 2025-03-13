package parser

import (
	"github.com/xclamation/go-parser/internal/parser/djvu"
	"github.com/xclamation/go-parser/internal/parser/docx"
	"github.com/xclamation/go-parser/internal/parser/html"
	"github.com/xclamation/go-parser/internal/parser/pdf"
)

// Parser интерфейс для всех парсеров
type Parser interface {
	Parse(filePath string) (string, error)
}

// NewHTMLParser создает парсер для HTML
func NewHTMLParser() Parser {
	return &html.HTMLParser{}
}

// NewDOCXParser создает парсер для DOCX
func NewDOCXParser() Parser {
	return &docx.DOCXParser{}
}

// NewDOCXParser создает парсер для PDF
func NewPDFParser() Parser {
	return &pdf.PDFParser{}
}

// NewDOCXParser создает парсер для PDF
func NewDJVUParser() Parser {
	return &djvu.DJVUParser{}
}
