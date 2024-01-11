package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/DTunnel0/upload-go/src/domain/usecase"
)

type FileHandler struct {
	usecase usecase.SaveFileUseCase
}

func NewFileHandler(usecase usecase.SaveFileUseCase) FileHandler {
	return FileHandler{usecase}
}

func (h *FileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Filename string `json:"filename"`
		Error    string `json:"error,omitempty"`
	}

	clientIP := r.RemoteAddr
	log.Printf("IP: %s\n", clientIP)

	file, header, err := r.FormFile("file")
	filename := header.Filename
	log.Printf("Filename: %s\n", filename)

	if err != nil {
		response := Response{Error: "Erro ao obter o arquivo: " + err.Error()}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	defer file.Close()

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, file)
	if err != nil {
		response := Response{Error: "Erro ao ler o arquivo: " + err.Error()}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	resultFilename, err := h.usecase.Execute(filename, buffer.Bytes())
	if err != nil {
		response := Response{Error: "Erro ao executar o caso de uso: " + err.Error()}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := Response{Filename: resultFilename}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
