package utils

import (
	"bytes"
	"encoding/json"
	"os"
)

func SaveToJsonFile(filename string, data []byte) error {
	// Create or overwrite JSON file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Adjust indentation for better readability
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, data, "", "    ")

	// Write the formatted JSON to the file
	if _, err = file.Write(prettyJSON.Bytes()); err != nil {
		return err
	}

	return nil
}
