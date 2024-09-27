package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}


func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		fmt.Println("Could not open file")
		fmt.Println(err)
		return nil, errors.New("failed to open file")
	}

	// Allows us to read line by line from an open file
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("failed to read line in file")
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	// Writes to a file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data) 

	if err != nil {
		file.Close()
		return errors.New("failed to convert data to JSON")
	}

	file.Close()
	return nil
}


func New(InputFilePath, OutputFilePath string) FileManager {
	return FileManager{
		InputFilePath: InputFilePath,
		OutputFilePath: OutputFilePath,
	}
}