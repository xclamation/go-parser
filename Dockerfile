# Используем базовый образ Ubuntu 22.04
FROM ubuntu:22.04

# Устанавливаем переменные окружения для избежания вопросов при установке пакетов
ENV DEBIAN_FRONTEND=noninteractive

# Устанавливаем необходимые пакеты
RUN apt-get update && apt-get install -y \
    wget \
    build-essential \
    cmake \
    pkg-config \
    wv \                      # wv для работы с DOC \
    libpoppler-dev \          # Poppler для работы с PDF \
    libpoppler-cpp-dev \      # Poppler C++ bindings \
    libdjvulibre-dev \        # DjVuLibre для работы с DjVu \
    libleptonica-dev \        # Leptonica для Tesseract \
    libtesseract-dev \        # Tesseract OCR \
    tesseract-ocr-eng \       # Английский язык для Tesseract \
    tesseract-ocr-rus \       # Русский язык для Tesseract \
    git \
    && rm -rf /var/lib/apt/lists/*

# Устанавливаем Go 1.24.0
RUN wget https://go.dev/dl/go1.24.0.linux-amd64.tar.gz -O /tmp/go.tar.gz \
    && tar -C /usr/local -xzf /tmp/go.tar.gz \
    && rm /tmp/go.tar.gz

# Добавляем Go в PATH
ENV PATH="/usr/local/go/bin:${PATH}"

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем только необходимые файлы
COPY go.mod go.sum ./
COPY cmd ./cmd
COPY internal ./internal
COPY parser ./parser

# Собираем проект
RUN go build -o go-parser ./cmd/main.go

# Команда по умолчанию
CMD ["./go-parser"]