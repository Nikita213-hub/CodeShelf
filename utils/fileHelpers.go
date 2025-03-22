package utils

import (
	"os"
	"strconv"
	"time"
)

func NewFile(fileType string) (*os.File, error) {
	fileName := strconv.Itoa(int(time.Now().Unix()))
	file, err := os.Create("../snippets/" + fileName + "." + fileType)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	return file, nil
}

func WriteToFile(filePointer *os.File, payload string) error {
	err := os.WriteFile(filePointer.Name(), []byte(payload), 0666)
	if err != nil {
		return err
	}
	return nil
}

func GetFileContent(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
