package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xclamation/go-parser/internal/parser"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: parser <file_type> <file_path>")
	}

	fileType := os.Args[1]
	filePath := os.Args[2]

	var p parser.Parser
	switch fileType {
	case "html":
		p = parser.NewHTMLParser()
	case "docx":
		p = parser.NewDOCXParser()
	case "pdf":
		p = parser.NewPDFParser()
	case "djvu":
		p = parser.NewDJVUParser()
	default:
		log.Fatalf("Unsupported file type: %s", fileType)
	}

	// Вызываем Parse
	result, err := p.Parse(filePath)
	if err != nil {
		log.Fatalf("Error parsing file: %v", err)
	}

	// Вывод результата
	fmt.Println(result)
}
