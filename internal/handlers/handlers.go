package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"local/internal/service"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	htmlPath := filepath.Join(".", "index.html")
	htmlFile, err := os.Open(htmlPath)
	if err != nil {
		http.Error(w, "Не могу открыть index.html", http.StatusInternalServerError)
		return
	}
	defer htmlFile.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.Copy(w, htmlFile)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка при получении файла: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка при чтении файла: "+err.Error(), http.StatusInternalServerError)
		return
	}

	input := string(fileContent)
	result, err := service.MorzeConvert(input)
	if err != nil {
		http.Error(w, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	outputFilename := time.Now().UTC().String() + ext
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		http.Error(w, "Ошибка создания файла: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(result)
	if err != nil {
		http.Error(w, "Ошибка записи в файл: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment; filename="+header.Filename)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
