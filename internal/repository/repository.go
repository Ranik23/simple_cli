package repository

import (
	"cli/internal/entity/file"
	"io"
	"log/slog"
	"os"
	"time"
	"github.com/lmittmann/tint"
)

type UserRepository interface {
	ReadFromFile(filePath string) (*file.File, error)
	WriteTo(io.Writer, *file.File) error
}


type Repository struct {
	logger *slog.Logger
}


func NewRepository() *Repository {
	return &Repository{
		logger: slog.New(tint.NewHandler(os.Stdout,&tint.Options{
			TimeFormat: time.TimeOnly,
		})),
	}	
}


func (r *Repository) ReadFromFile(filePath string) (*file.File, error) {

	File, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(File)

	if err != nil {
		return nil, err
	}

	newFile := &file.File{
		Name: File.Name(),
		Size: fileInfo.Size(),
		Mode: fileInfo.Mode(),
		ModificationTime: fileInfo.ModTime(),
		Content: string(bytes),
	}

	return newFile, nil
}


func (r *Repository) WriteTo(writer io.Writer, file *file.File) error {

	_, err := writer.Write([]byte(file.Content + "\n"))

	if err != nil {
		return err
	}
	return nil
}