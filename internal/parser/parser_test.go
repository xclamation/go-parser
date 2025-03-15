package parser

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xclamation/go-parser/pkg/consts"
	"io"
	"os"
	"path/filepath"
	"testing"
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
	text, err := HtmlParser.Parse(filepath.Join(PATH, "example.html"))

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
	text, err := DOCXParser.Parse(filepath.Join(PATH, "example.docx"))

	// Проверка на ошибку
	require.NoError(t, err, "Failed to parse DOCX file")

	// Чтение ожидаемого результата из файла
	file, err := os.Open(filepath.Join(PATH, "docx.txt"))
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
	text, err := DOCParser.Parse(filepath.Join(PATH, "example.doc"))

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
	text, err := PDFParser.Parse(filepath.Join(PATH, "example.pdf"))

	// Проверка на ошибку
	require.NoError(t, err, "Failed to parse PDF file")

	// Чтение ожидаемого результата из файла
	file, err := os.Open(filepath.Join(PATH, "pdf.txt"))
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
	text, err := DJVUParser.Parse(filepath.Join(PATH, "example.djvu"))

	// Проверка на ошибку
	require.NoError(t, err, "Failed to parse DJVU file")

	// Чтение ожидаемого результата из файла
	file, err := os.Open(filepath.Join(PATH, "djvu.txt"))
	require.NoError(t, err, "Failed to open file")
	defer file.Close()

	expectedDJVU, err := io.ReadAll(file)
	require.NoError(t, err, "Failed to read file")

	// Проверка результата
	assert.Equal(t, string(expectedDJVU), text, "Parsed DJVU does not match expected")
}
