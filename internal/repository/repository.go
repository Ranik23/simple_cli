package repository

import (
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
	ReadFromFile(string) ([]byte, error)
	WriteTo(io.Writer, []byte, bool) error
	GetEntries(string) (*[]fs.DirEntry, error)
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


func (r *Repository) ReadFromFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}


func (r *Repository) WriteTo(writer io.Writer, data []byte, isDir bool) error {
	var output []byte
	var err error

	if isDir {
		color := color.New(color.FgCyan).SprintFunc()
		output = []byte(color(string(data)))
	} else {
		output = data
	}

	_, err = writer.Write([]byte(string(output) + "\n"))
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetEntries(dir string) (*[]fs.DirEntry, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return &entries, err
}
