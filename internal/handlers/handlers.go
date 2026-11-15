package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler отдаёт HTML-страницу с формой загрузки файла.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// UploadHandler принимает файл из формы, конвертирует его содержимое
// и возвращает результат конвертации.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "ошибка разбора формы", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка получения файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "ошибка конвертации", http.StatusInternalServerError)
		return
	}

	// Имя файла — текущее время в UTC плюс расширение исходного файла.
	// Двоеточия заменяем: с ними имя не создать в Windows.
	filename := strings.ReplaceAll(time.Now().UTC().String(), ":", "-") + filepath.Ext(header.Filename)

	out, err := os.Create(filename)
	if err != nil {
		http.Error(w, "ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	if _, err := out.WriteString(result); err != nil {
		http.Error(w, "ошибка записи файла", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(result))
}
