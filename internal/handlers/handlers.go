package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	input := string(fileContent)
	converted, err := service.AutoDetectAndConvert(input)
	if err != nil {
		http.Error(w, "Conversion error", http.StatusInternalServerError)
		return
	}

	timestamp := time.Now().UTC().String()
	originalExt := filepath.Ext(header.Filename)
	filename := "converted_" + timestamp + originalExt

	outputFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(converted)
	if err != nil {
		http.Error(w, "Unable to write file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, converted)
}
