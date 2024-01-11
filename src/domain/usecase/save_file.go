package usecase

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"

	"github.com/DTunnel0/upload-go/src/domain/entity"
	"github.com/DTunnel0/upload-go/src/infra/repository"
)

type SaveFileUseCase struct {
	Repository repository.FileRepository
}

func NewSaveFileUseCase(repositiry repository.FileRepository) SaveFileUseCase {
	return SaveFileUseCase{repositiry}
}

func (s *SaveFileUseCase) Execute(filename string, conent []byte) (string, error) {
	fileExtension := getFileExtension(filename)
	filePath := getFilePath(fileExtension)
	randomName := generateRandomName(fileExtension)

	file := entity.NewFile(randomName, filePath, conent)
	error := s.Repository.SaveFile(*file)

	return file.Name, error
}

func getFileExtension(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}

func generateRandomName(extension string) string {
	randomName := generateMD5Hash()
	return randomName + "." + extension
}

func generateMD5Hash() string {
	hasher := md5.New()
	hasher.Write([]byte(time.Now().String()))
	return hex.EncodeToString(hasher.Sum(nil))
}

func getFilePath(fileExtension string) string {
	switch fileExtension {
	case "jpg", "png", "gif":
		return "uploads/images"
	case "mp4", "avi", "mkv":
		return "uploads/videos"
	default:
		return "uploads/files"
	}
}
