package usecase

import (
	"cli/internal/repository"
	"fmt"
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"strings"
	"time"
)

var (
	ErrOutOfBound = fmt.Errorf("out of bound")
)

type UseCase interface {
	Print(filePath string) error
	СountStrings(filePath string) (int, error)
	CountWordsOnEachString(filePath string) (*[]int, error)
	CountWords(filePath string, lowBound, highBound int) (int, error)
	Ls() error
}

type UserOperator struct {
	repository repository.UserRepository
	logger     slog.Logger
}

func NewUserOperator(repository repository.UserRepository) *UserOperator {
	return &UserOperator{
		repository: repository,
		logger: *slog.New(tint.NewHandler(os.Stdout, &tint.Options{
			TimeFormat: time.TimeOnly,
		})),
	}
}

func (us UserOperator) Print(filePath string) error {

	file, err := us.repository.ReadFromFile(filePath)
	if err != nil {
		return err
	}
	err = us.repository.WriteTo(os.Stdout, file)

	if err != nil {
		return err
	}

	return nil
}

func (us UserOperator) CountStrings(filePath string) (int, error) {

	file, err := us.repository.ReadFromFile(filePath)

	if err != nil {
		return -1, err
	}

	numberOfStrings := len(strings.Split(file.Content, "\n")) // вот тут огромное потребление памяти

	return numberOfStrings, nil
}

func (us UserOperator) CountWordsOnEachString(filePath string) (*[]int, error) {

	file, err := us.repository.ReadFromFile(filePath)

	if err != nil {
		return nil, err
	}

	strs := strings.Split(file.Content, "\n")

	number := len(strs)

	var answer []int

	for i := 0; i < number; i++ {

		newString := strings.Replace(strings.Replace(strs[i], "\n", " ", -1), "  ", " ", -1)

		n := len(strings.Split(newString, " "))

		answer = append(answer, n)
	}

	return &answer, nil
}

func (us UserOperator) CountWords(filePath string, lowBound, highBound int) (int, error) {

	file, err := us.repository.ReadFromFile(filePath)

	if err != nil {
		return -1, err
	}

	strs := strings.Split(file.Content, "\n")

	number := len(strs)

	if number == 0 {
		return 0, nil
	}

	if lowBound == -1 && highBound == -1 {
		lowBound = 0
		highBound = number - 1
	}

	if lowBound < 0 || highBound >= number {
		return -1, ErrOutOfBound
	}

	answer := 0

	for i := lowBound; i <= highBound; i++ {

		newString := strings.Replace(strings.Replace(strs[i], "\n", " ", -1), "  ", " ", -1)

		n := len(strings.Split(newString, " "))

		answer += n
	}

	return answer, nil
}


func (us UserOperator) Ls() error {

	entries, err := us.repository.GetEntries(".")

	if err != nil {
		return nil
	}

	for _, entry := range *entries {
		err = us.repository.WriteTo(os.Stdout, entry)

		if err != nil {
			return err
		}
	}

	return nil
}
