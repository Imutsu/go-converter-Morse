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

func UploadHandlers(w http.ResponseWriter, r *http.Request) {
	// проверить метод
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// парсинг
	// макс размер примерно 10MB
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		 http.Error(w, err.Error(), http.StatusInternalServerError)
		 return
	}

	// получение файла
	file, header, err := r.FormFile("myfile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// чтение файла
	// data - это []byte
	data, err := io.ReadAll(file) // пометка для себя, тут чтение файла в io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// превращение в строку
	input := string(data)

	//вызов AutoDetect
	// после этого получим полный результат
	result, err:= service.AutoDetect(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)

	fileName := time.Now().UTC().String() + ext
	newFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// запись результата
	_, err = newFile.WriteString(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// теперь важно отправить клиенту
	fmt.Fprint(w, result)
}


func IndexHandler (w http.ResponseWriter, r *http.Request) {
	//проверка метода, важно, чтобы был GET
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//отправка файла клиенту
	http.ServeFile(w, r, "index.html") // где w - канал, через которые отправляются данные
	// r - объект запроса
}
