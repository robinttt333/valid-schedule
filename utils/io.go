package utils

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func Read(path string) ([]string, error) {
	if path == "" {
		return nil, errors.New("file path cannot be empty")
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("error while reading file %s", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("error while reading file %s", err.Error())
	}
	return res, nil
}
