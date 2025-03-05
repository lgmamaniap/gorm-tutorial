package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadJsonFile[T any](path string, store T) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error open json file ", err)
		return err
	}

	fmt.Printf("Successfully Opened %s\n", path)
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(byteValue, &store); err != nil {
		return err
	}

	return nil
}
