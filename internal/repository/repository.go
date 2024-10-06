package repository

import (
	"cli/internal/entity/file"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"time"
	"github.com/lmittmann/tint"
	"github.com/fatih/color"
)

var (
	ErrUnSupportedType = fmt.Errorf("unsupported type")
)

type UserRepository interface {
	ReadFromFile(filePath string) (*file.File, error)
	WriteTo(io.Writer, interface{}) error
	GetEntries(dir string) (*[]fs.DirEntry, error)
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


func (r *Repository) WriteTo(writer io.Writer, entity interface{}) error {

	if entity, ok := entity.(fs.DirEntry); ok {

		red := color.New(color.FgCyan).SprintfFunc()

		var err error

		if entity.IsDir() {
			_, err = writer.Write([]byte(red(entity.Name()) + " "))
		} else {
			_, err = writer.Write([]byte(entity.Name() + " "))
		}

		if err != nil {
			return err
		}
		return nil
	}

	if entity, ok := entity.(file.File); ok {
		_, err := writer.Write([]byte(entity.Content + "\n"))

		if err != nil {
			return err
		}
		return nil
	}
	return ErrUnSupportedType
}


func (r *Repository) GetEntries(dir string) (*[]fs.DirEntry, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return &entries, err
}
