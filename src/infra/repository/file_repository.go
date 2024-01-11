package repository

import (
	"os"
	"path/filepath"

	"github.com/DTunnel0/upload-go/src/domain/entity"
)

type FileRepository interface {
	SaveFile(entity.File) error
}

type fileRepository struct{}

func NewFileRepository() FileRepository {
	return &fileRepository{}
}

func (r *fileRepository) SaveFile(file entity.File) error {
	err := os.MkdirAll(file.Path, os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(file.Path, file.Name))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(file.Content)
	if err != nil {
		return err
	}

	return nil
}
