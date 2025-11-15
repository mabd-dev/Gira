package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func SaveMockApiResponse(filename string, data []byte) error {
	return SaveToJsonFile("api/testdata/mockApiResponses/"+filename, data)
}

func SaveToJsonFile(filename string, data []byte) error {
	fmt.Printf("saving file to %s\n", filename)
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
