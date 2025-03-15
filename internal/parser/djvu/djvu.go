package djvu

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// DJVUParser структура для парсера DJVU
type DJVUParser struct{}

// NewDJVUParser создает новый экземпляр DJVUParser
func NewDJVUParser() *DJVUParser {
	return &DJVUParser{}
}

// Parse извлекает текст из DJVU-файла
func (p *DJVUParser) Parse(filePath string) (string, error) {
	// Создаем временную директорию для сохранения изображений
	tempDir, err := os.MkdirTemp("", "djvu-pages")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Проверяем наличие утилиты ddjvu
	if err := exec.Command("where", "ddjvu").Run(); err != nil {
		return "", fmt.Errorf("ddjvu is not installed or not found in PATH")
	}

	// Разделяем DJVU-файл на изображения с помощью ddjvu
	imagePattern := filepath.Join(tempDir, "page_%03d.tiff")
	cmd := exec.Command("ddjvu", "-format=tiff", "-eachpage", filePath, imagePattern)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error splitting DJVU file: %v\nOutput: %s", err, string(output))
	}

	// Извлекаем текст из каждого изображения
	var fullText strings.Builder
	images, err := filepath.Glob(filepath.Join(tempDir, "page_*.tiff"))
	if err != nil {
		return "", fmt.Errorf("error listing images: %v", err)
	}

	var language string
	switch {
	case strings.Contains(filePath, "ru"):
		language = "rus"
	case strings.Contains(filePath, "eng"):
		language = "eng"
	}

	for i, imagePath := range images {
		if err != nil {
			return "", fmt.Errorf("error detecting language: %v", err)
		}

		// Извлекаем текст с учетом языка
		text, err := extractTextFromImage(imagePath, language)
		if err != nil {
			return "", fmt.Errorf("error extracting text from image %s: %v", imagePath, err)
		}

		// Добавляем текст страницы к общему тексту
		fullText.WriteString(formatPage(text, i+1))
		fullText.WriteString("\n\n") // Добавляем разделитель между страницами
	}

	return fullText.String(), nil
}

// extractTextFromImage извлекает текст из изображения с помощью Tesseract OCR
func extractTextFromImage(imagePath string, language string) (string, error) {
	// Проверяем наличие утилиты tesseract
	if err := exec.Command("tesseract", "--version").Run(); err != nil {
		return "", fmt.Errorf("tesseract is not installed or not found in PATH")
	}

	// Вызываем Tesseract через командную строку с указанием языка
	cmd := exec.Command("tesseract", imagePath, "stdout", "-l", language)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running tesseract: %v", err)
	}

	// Убираем лишние пробелы и возвращаем текст
	text := strings.TrimSpace(string(output))
	return text, nil
}

// formatPage форматирует текст страницы, добавляя отступы и заголовок
func formatPage(text string, pageNumber int) string {
	var formattedText strings.Builder

	// Добавляем заголовок страницы
	formattedText.WriteString(fmt.Sprintf("=== Страница %d ===\n", pageNumber))

	// Разделяем текст на абзацы
	paragraphs := strings.Split(text, "\n")
	for _, paragraph := range paragraphs {
		paragraph = strings.TrimSpace(paragraph)
		if paragraph != "" {
			// Добавляем отступ для абзаца
			formattedText.WriteString("    " + paragraph + "\n")
		}
	}

	return formattedText.String()
}
