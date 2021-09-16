package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type File struct {
	FileName     string `json:"fileName"`
	EncodingType string `json:"encodingType"`
	EncodedData  string `json:"encodedData"`
}

func main() {
	// create array to hold File objects
	var filesSlice []File

	// array of file names
	inputFiles := []string{"Helo Foo Bar.docx", "Helo Foo Bar.pdf"}

	// for each file reading a byte slice
	for _, inFile := range inputFiles {
		// Open file on disk.
		f, err := os.Open(inFile)
		if err != nil {
			fmt.Printf("Problem reading file: \"%v\"\nerror: %v\n", inFile, err.Error())
			os.Exit(99)
		}

		// Read file into byte slice.
		reader := bufio.NewReader(f)
		content, _ := ioutil.ReadAll(reader)

		// Encode as url encoded base64.
		encoded := base64.URLEncoding.EncodeToString(content)

		file := File{FileName: inFile, EncodingType: "Base64", EncodedData: encoded}
		filesSlice = append(filesSlice, file)
	} // end range inputFiles

	// print json array to screen
	filesArrayJson, err := json.Marshal(filesSlice)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonString := string(filesArrayJson)
	fmt.Println(jsonString)

	// now lets take each json object and decode the EncodedData into new files
	var fromJsonFilesSlice []File
	// Unmarshal JSON to the interface.
	err = json.Unmarshal([]byte(jsonString), &fromJsonFilesSlice)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(77)
	}

	// for each File object get the FileName and EncodedData
	// base64 decode the EnocdedData into a byte slice and
	// write it to file prepending "new_" to the old file name
	for _, newfile := range fromJsonFilesSlice {
		newFilename := fmt.Sprintf("new_%v", newfile.FileName)

		fileBytes, _ := base64.URLEncoding.DecodeString(newfile.EncodedData)
		err = ioutil.WriteFile(newFilename, fileBytes, 0644)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(66)
		}
	}
}
