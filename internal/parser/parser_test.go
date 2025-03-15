package parser

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xclamation/go-parser/pkg/consts"
)

const (
	PATH string = "../../examples"
)

type ParserSuite struct {
	suite.Suite
}

func TestParserSuite(t *testing.T) {
	suite.RunSuite(t, new(ParserSuite))
}

func (s *ParserSuite) TestHTMLParser(t provider.T) {
	t.Epic("Parser")
	t.Feature("HTML Parsing")
	t.Title("Test HTML Parser")
	t.Description("This test checks if the HTML parser correctly parses an HTML file.")

	HtmlParser := NewHTMLParser()
	text, err := HtmlParser.Parse(filepath.Join(PATH, "example_eng.html"))

	// Проверка на ошибку
	require.NoError(t, err, "Failed to parse HTML file")

	// Проверка результата
	assert.Equal(t, consts.ExpectedHTML, text, "Parsed HTML does not match expected")
}

func (s *ParserSuite) TestDOCXParser(t provider.T) {
	t.Epic("Parser")
	t.Feature("DOCX Parsing")
	t.Title("Test DOCX Parser")
	t.Description("This test checks if the DOCX parser correctly parses a DOCX file.")

	DOCXParser := NewDOCXParser()
	text, err := DOCXParser.Parse(filepath.Join(PATH, "example_ru.docx"))

	// Проверка на ошибку
	require.NoError(t, err, "Failed to parse DOCX file")

	// Чтение ожидаемого результата из файла
	file, err := os.Open(filepath.Join(PATH, "txt/docx_ru.txt"))
	require.NoError(t, err, "Failed to open file")
	defer file.Close()

	expectedDOCX, err := io.ReadAll(file)
	require.NoError(t, err, "Failed to read file")

	// Проверка результата
	assert.Equal(t, string(expectedDOCX), text, "Parsed DOCX does not match expected")
}

func (s *ParserSuite) TestDOCParser(t provider.T) {
	t.Epic("Parser")
	t.Feature("DOC Parsing")
	t.Title("Test DOC Parser")
	t.Description("This test checks if the DOC parser correctly parses a DOC file.")

	DOCParser := NewDOCParser()
	text, err := DOCParser.Parse(filepath.Join(PATH, "example_ru.doc"))

	// Проверка на ошибку
	require.NoError(t, err, "Failed to parse DOC file")

	// Проверка результата
	assert.Equal(t, consts.ExpectedDOC, text, "Parsed DOC does not match expected")
}

func (s *ParserSuite) TestPDFParser(t provider.T) {
	t.Epic("Parser")
	t.Feature("PDF Parsing")
	t.Title("Test PDF Parser")
	t.Description("This test checks if the PDF parser correctly parses a PDF file.")

	PDFParser := NewPDFParser()
	text, err := PDFParser.Parse(filepath.Join(PATH, "example_ru.pdf"))

	// Проверка на ошибку
	require.NoError(t, err, "Failed to parse PDF file")

	// Чтение ожидаемого результата из файла
	file, err := os.Open(filepath.Join(PATH, "txt/pdf_ru.txt"))
	require.NoError(t, err, "Failed to open file")
	defer file.Close()

	expectedPDF, err := io.ReadAll(file)
	require.NoError(t, err, "Failed to read file")

	// Проверка результата
	assert.Equal(t, string(expectedPDF), text, "Parsed PDF does not match expected")
}

func (s *ParserSuite) TestDJVUParser(t provider.T) {
	t.Epic("Parser")
	t.Feature("DJVU Parsing")
	t.Title("Test DJVU Parser")
	t.Description("This test checks if the DJVU parser correctly parses a DJVU file.")

	DJVUParser := NewDJVUParser()
	text, err := DJVUParser.Parse(filepath.Join(PATH, "example_eng.djvu"))

	// Проверка на ошибку
	require.NoError(t, err, "Failed to parse DJVU file")

	// Чтение ожидаемого результата из файла
	file, err := os.Open(filepath.Join(PATH, "txt/djvu_eng.txt"))
	require.NoError(t, err, "Failed to open file")
	defer file.Close()

	expectedDJVU, err := io.ReadAll(file)
	require.NoError(t, err, "Failed to read file")

	// Проверка результата
	assert.Equal(t, string(expectedDJVU), text, "Parsed DJVU does not match expected")
}

func (s *ParserSuite) TestParserLangRU(t provider.T) {
	t.Epic("Parser")
	t.Feature("Russian Language Parsing")
	t.Title("Test Russian Language Parser")
	t.Description("This test checks if the parser correctly parses a Russian Language all of file types.")

	HTMLParser := NewHTMLParser()
	textHTML, err := HTMLParser.Parse(filepath.Join(PATH, "example_ru.html"))
	require.NoError(t, err, "Failed to parse DJVU file")

	fileHtml, err := os.Open(filepath.Join(PATH, "txt/html_ru.txt"))
	require.NoError(t, err, "Failed to open file"+" html_ru.txt")
	defer fileHtml.Close()

	expectedHTML, err := io.ReadAll(fileHtml)
	require.NoError(t, err, "Failed to read file"+" html_ru.txt")

	assert.Equal(t, string(expectedHTML), textHTML, "Parsed HTML does not match expected")

	// ======================================================================================

	DOCXParser := NewDOCXParser()
	textDOCX, err := DOCXParser.Parse(filepath.Join(PATH, "example_ru.docx"))
	require.NoError(t, err, "Failed to parse DOCX file")

	fileDocx, err := os.Open(filepath.Join(PATH, "txt/docx_ru.txt"))
	require.NoError(t, err, "Failed to open file"+" docx_ru.txt")
	defer fileDocx.Close()

	expectedDOCX, err := io.ReadAll(fileDocx)
	require.NoError(t, err, "Failed to read file"+" docx_ru.txt")

	assert.Equal(t, string(expectedDOCX), textDOCX, "Parsed DOCX does not match expected")

	// ======================================================================================

	DOCParser := NewDOCParser()
	textDOC, err := DOCParser.Parse(filepath.Join(PATH, "example_ru.doc"))
	require.NoError(t, err, "Failed to parse DOC file")

	fileDoc, err := os.Open(filepath.Join(PATH, "txt/doc_ru.txt"))
	require.NoError(t, err, "Failed to open file"+" doc_ru.txt")
	defer fileDoc.Close()

	expectedDOC, err := io.ReadAll(fileDoc)
	require.NoError(t, err, "Failed to read file"+" doc_ru.txt")

	assert.Equal(t, string(expectedDOC), textDOC, "Parsed DOC does not match expected")

	// ======================================================================================

	PDFParser := NewPDFParser()
	textPDF, err := PDFParser.Parse(filepath.Join(PATH, "example_ru.pdf"))
	require.NoError(t, err, "Failed to parse PDF file")

	filePdf, err := os.Open(filepath.Join(PATH, "txt/pdf_ru.txt"))
	require.NoError(t, err, "Failed to open file"+" pdf_ru.txt")
	defer filePdf.Close()

	expectedPDF, err := io.ReadAll(filePdf)
	require.NoError(t, err, "Failed to read file"+" pdf_ru.txt")

	assert.Equal(t, string(expectedPDF), textPDF, "Parsed PDF does not match expected")

	// ======================================================================================

	DJVUParser := NewDJVUParser()
	textDJVU, err := DJVUParser.Parse(filepath.Join(PATH, "example_ru.djvu"))
	require.NoError(t, err, "Failed to parse DJVU file")

	fileDjvu, err := os.Open(filepath.Join(PATH, "txt/djvu_ru.txt"))
	require.NoError(t, err, "Failed to open file"+" djvu_ru.txt")
	defer fileDjvu.Close()

	expectedDJVU, err := io.ReadAll(fileDjvu)
	require.NoError(t, err, "Failed to read file"+" djvu_ru.txt")

	assert.Equal(t, string(expectedDJVU), textDJVU, "Parsed DJVU does not match expected")

}

func (s *ParserSuite) TestParserLangEN(t provider.T) {
	t.Epic("Parser")
	t.Feature("english Language Parsing")
	t.Title("Test english Language Parser")
	t.Description("This test checks if the parser correctly parses a english Language all of file types.")

	HTMLParser := NewHTMLParser()
	textHTML, err := HTMLParser.Parse(filepath.Join(PATH, "example_eng.html"))
	require.NoError(t, err, "Failed to parse DJVU file")

	fileHtml, err := os.Open(filepath.Join(PATH, "txt/html_eng.txt"))
	require.NoError(t, err, "Failed to open file"+" html_eng.txt")
	defer fileHtml.Close()

	expectedHTML, err := io.ReadAll(fileHtml)
	require.NoError(t, err, "Failed to read file"+" html_eng.txt")

	assert.Equal(t, string(expectedHTML), textHTML, "Parsed HTML does not match expected")

	// ======================================================================================

	DOCXParser := NewDOCXParser()
	textDOCX, err := DOCXParser.Parse(filepath.Join(PATH, "example_eng.docx"))
	require.NoError(t, err, "Failed to parse DOCX file")

	fileDocx, err := os.Open(filepath.Join(PATH, "txt/docx_eng.txt"))
	require.NoError(t, err, "Failed to open file"+" docx_eng.txt")
	defer fileDocx.Close()

	expectedDOCX, err := io.ReadAll(fileDocx)
	require.NoError(t, err, "Failed to read file"+" docx_eng.txt")

	assert.Equal(t, string(expectedDOCX), textDOCX, "Parsed DOCX does not match expected")

	// ======================================================================================

	DOCParser := NewDOCParser()
	textDOC, err := DOCParser.Parse(filepath.Join(PATH, "example_eng.doc"))
	require.NoError(t, err, "Failed to parse DOC file")

	fileDoc, err := os.Open(filepath.Join(PATH, "txt/doc_eng.txt"))
	require.NoError(t, err, "Failed to open file"+" doc_eng.txt")
	defer fileDoc.Close()

	expectedDOC, err := io.ReadAll(fileDoc)
	require.NoError(t, err, "Failed to read file"+" doc_eng.txt")

	assert.Equal(t, string(expectedDOC), textDOC, "Parsed DOC does not match expected")

	// ======================================================================================

	PDFParser := NewPDFParser()
	textPDF, err := PDFParser.Parse(filepath.Join(PATH, "example_eng.pdf"))
	require.NoError(t, err, "Failed to parse PDF file")

	filePdf, err := os.Open(filepath.Join(PATH, "txt/pdf_eng.txt"))
	require.NoError(t, err, "Failed to open file"+" pdf_eng.txt")
	defer filePdf.Close()

	expectedPDF, err := io.ReadAll(filePdf)
	require.NoError(t, err, "Failed to read file"+" pdf_eng.txt")

	assert.Equal(t, string(expectedPDF), textPDF, "Parsed PDF does not match expected")

	// ======================================================================================

	DJVUParser := NewDJVUParser()
	textDJVU, err := DJVUParser.Parse(filepath.Join(PATH, "example_eng.djvu"))
	require.NoError(t, err, "Failed to parse DJVU file")

	fileDjvu, err := os.Open(filepath.Join(PATH, "txt/djvu_eng.txt"))
	require.NoError(t, err, "Failed to open file"+" djvu_eng.txt")
	defer fileDjvu.Close()

	expectedDJVU, err := io.ReadAll(fileDjvu)
	require.NoError(t, err, "Failed to read file"+" djvu_eng.txt")

	assert.Equal(t, string(expectedDJVU), textDJVU, "Parsed DJVU does not match expected")

}
