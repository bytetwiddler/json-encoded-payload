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
	var filesSlice []File
	inputFiles := []string{"Helo Foo Bar.docx", "Helo Foo Bar.pdf"}
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

	// take json array of encoded File objects and write to file

	// pint json array to screen
	filesArrayJson, err := json.Marshal(filesSlice)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonString := string(filesArrayJson)
	fmt.Println(jsonString)

	var fromJsonFilesSlice []File
	// Unmarshal JSON to the interface.
	err = json.Unmarshal([]byte(jsonString), &fromJsonFilesSlice)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(77)
	}
	for _, newfile := range fromJsonFilesSlice {
		newfilename := fmt.Sprintf("new_%v", newfile.FileName)

		fileBytes, _ := base64.URLEncoding.DecodeString(newfile.EncodedData)
		err = ioutil.WriteFile(newfilename, fileBytes, 0644)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(66)
		}
	}
}
