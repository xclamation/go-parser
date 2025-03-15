# Go Parser Project

Этот проект представляет собой парсер документов различных форматов (`PDF`, `DOC`, `DOCX`, `HTML`, `DjVu`) на языке Go. Парсер извлекает текст из файлов и предоставляет его в удобном для обработки виде.

## Зависимости

Для работы проекта необходимы следующие зависимости:

### Системные зависимости

#### wv (wvWare)  
Утилита для конвертации `.doc` файлов в текстовый формат.  
Руководство по установке для Windows: [wvWare Installation Guide for Windows](https://sourceforge.net/projects/wvware/files/latest/download)  
Руководство по установке для Linux: [wvWare Installation Guide for Linux](https://wiki.debian.org/WvWare)

#### Poppler
Используется для работы с PDF-файлами.  
Руководство по установке: [Poppler Installation Guide](https://poppler.freedesktop.org/)

#### DjVuLibre
Используется для работы с DjVu-файлами.  
Руководство по установке: [DjVuLibre Installation Guide](http://djvu.sourceforge.net/)

#### Tesseract OCR
Используется для извлечения текста из изображений в DJVU и других форматах.  
Руководство по установке: [Tesseract Installation Guide](https://tesseract-ocr.github.io/tessdoc/Installation.html)

**Языковые пакеты (например, английский и русский):**  
Руководство по установке: [Tesseract Language Data](https://github.com/tesseract-ocr/tessdata)

### Go-зависимости
Все Go-зависимости указаны в файле `go.mod` и будут автоматически установлены при сборке проекта.

---

## Установка и запуск

### С использованием Docker

1. Установите Docker и Docker Compose, следуя [официальной документации](https://docs.docker.com/get-docker/).
2. Соберите и запустите контейнер:

```bash
docker-compose up --build
```

3. Остановка контейнера:

```bash
docker-compose down
```

---

### Без Docker

1. **Установите Go:**  
   Скачайте и установите Go с [официального сайта](https://go.dev/dl/).

2. **Установите системные зависимости:**  
   Убедитесь, что все системные зависимости установлены (см. раздел **Зависимости**).

3. **Скачайте зависимости и соберите проект:**  
   Перейдите в корневую директорию проекта и выполните:

```bash
go mod download
go build -o go-parser ./cmd/main.go
```

4. **Запустите проект:**  
   Выполните:

```bash
./go-parser
```

---

## Структура проекта

```
/go-parser-project
│
├── Dockerfile
├── docker-compose.yaml
├── go.mod
├── go.sum
├── cmd
│   └── main.go
├── internal
│   └──parser
        ├── djvu
        │   └── djvu.go
        ├── doc
        │   └── doc.go
        ├── docx
        │   └── docx.go
        ├── html
        │   └── html.go
        ├── pdf
        │   └── pdf.go
        └── parser.go
└── 
```