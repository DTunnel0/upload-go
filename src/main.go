package main

import (
	"log"
	"net/http"

	"github.com/DTunnel0/upload-go/src/domain/usecase"
	"github.com/DTunnel0/upload-go/src/infra/handler"
	"github.com/DTunnel0/upload-go/src/infra/middleware"
	"github.com/DTunnel0/upload-go/src/infra/repository"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	file_repository := repository.NewFileRepository()
	save_file_use_case := usecase.NewSaveFileUseCase(file_repository)
	save_file_handler := handler.NewFileHandler(save_file_use_case)

	http.HandleFunc("/upload", middleware.ValidateTokenMiddleware(save_file_handler.UploadFile))

	log.Println("Server listening on: 0.0.0.0:8888")
	http.ListenAndServe("0.0.0.0:8888", nil)
}
