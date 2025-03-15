package parser

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

const (
	PATH         string = "../../examples"
	expectedHTML string = `
Test Page
Header 1
   Lorem ipsum dolor sit amet, consectetur adipisicing elit. Aliquam distinctio inventore labore molestiae nemo nobis, omnis quaerat reprehenderit! Adipisci consequuntur cum cumque, dolorem illo illum itaque voluptates. Blanditiis culpa, veniam.
   Paragraph 2
   Header 2
   Paragraph 3`
	expectedDOC string = ""
)

func TestHTMLParser(t *testing.T) {
	HtmlParser := NewHTMLParser()
	text, err := HtmlParser.Parse(filepath.Join(PATH, "example.html"))
	if err != nil {
		t.Fatalf("Failed to parse HTML file: %v", err)
	}

	if text != expectedHTML {
		t.Errorf("\nExpected:\n'%s'\ngot:\n'%s'", expectedHTML, text)
	}
}

func TestDOCXParser(t *testing.T) {
	DOCXParser := NewDOCXParser()
	text, err := DOCXParser.Parse(filepath.Join(PATH, "example.docx"))
	if err != nil {
		t.Fatalf("Failed to parse DOCX file: %v", err)
	}

	file, err := os.Open(filepath.Join(PATH, "docx.txt"))
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close() // Закрываем файл после завершения

	expectedDOCX, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if text != string(expectedDOCX) {
		t.Errorf("\nExpected:\n'%s'\ngot:\n'%s'", expectedDOCX, text)
	}
}

func TestDOCParser(t *testing.T) {
	DOCParser := NewDOCParser()
	text, err := DOCParser.Parse(filepath.Join(PATH, "example.doc"))
	if err != nil {
		t.Fatalf("Failed to parse DOC file: %v", err)
	}

	if text != expectedDOC {
		t.Errorf("\nExpected:\n'%s'\ngot:\n'%s'", expectedDOC, text)
	}
}

func TestPDFParser(t *testing.T) {
	PDFParser := NewPDFParser()
	text, err := PDFParser.Parse(filepath.Join(PATH, "example.pdf"))
	if err != nil {
		t.Fatalf("Failed to parse PDF file: %v", err)
	}

	file, err := os.Open(filepath.Join(PATH, "pdf.txt"))
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	expectedPDF, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if text != string(expectedPDF) {
		t.Errorf("\nExpected:\n'%s'\ngot:\n'%s'", expectedPDF, text)
	}
}

func TestDJVUParser(t *testing.T) {
	DJVUParser := NewDJVUParser()
	text, err := DJVUParser.Parse(filepath.Join(PATH, "example.djvu"))
	if err != nil {
		t.Fatalf("Failed to parse DJVU file: %v", err)
	}

	file, err := os.Open(filepath.Join(PATH, "djvu.txt"))
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	expectedDJVU, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if text != string(expectedDJVU) {
		t.Errorf("\nExpected:\n'%s'\ngot:\n'%s'", expectedDJVU, text)
	}
}
