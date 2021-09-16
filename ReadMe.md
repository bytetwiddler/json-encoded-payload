A simple example of embedding binary base64 encoded data in a json object. 

NOTE: **NOT Production ready and never will be**.  Just a simple example of what it takes will little to no error handling, everything in main.

The *main.go* code will take the 2 included 'Helo Foo Bar.docx' and 'Helo Foo Bar.pdf' files in the project root and convert them to base64 URLEncoded json objects.  It will then print out a JSON array of those files with a fileName, encodingType, encodedData properties.  It then takes that JSON objects and use them to create 2 new files prepending "new_" to the original fileName. You should be left with 2 docx and 2 pdf files that have the exact same binary content.

Here is a single file json object:
```
{
    "fileName": "Helo Foo Bar.docx",
    "encodingType": "Base64",
    "encodedData": "UEsDBBQABgAIAA...bWxQSwUGAAAAAAsACwDBAgAAjyoAAAAA"
}
```

Here is a json array of each file base64 URLEncoded:
```
[
  {
    "fileName": "Helo Foo Bar.docx",
    "encodingType": "Base64",
    "encodedData": "UEsDBBQABgAIAAA...QSwUGACwDBAgAAjyoAAAAA"
  },
  {
    "fileName": "Helo Foo Bar.pdf",
    "encodingType": "Base64",
    "encodedData": "JVBERi0xLjUKJc...HhyZWYKODYyNwolJUVPRgo="
  }
]
```

Example usage:
```
$ go run main.go
[
  {
    "fileName": "Helo Foo Bar.docx",
    "encodingType": "Base64",
    "encodedData": "UEsDBBQABgAIAAA...AAAsACwDBAgAAjyoAAAAA"
  },
  {
    "fileName": "Helo Foo Bar.pdf",
    "encodingType": "Base64",
    "encodedData": "JVBERi0xLjUKJc...HhyZWYKODYyNwolJUVPRgo="
  }
]
```

2 new files should now be created with "new_" prepended to the old file names:
```
$ ls *docx *pdf
'Helo Foo Bar.docx'  'Helo Foo Bar.pdf'  'new_Helo Foo Bar.docx'  'new_Helo Foo Bar.pdf'
```

Are the files the same? (MD5 thinks so)
```
$ md5sum *docx *pdf
57956932603f9e194be4296fafecedf9  Helo Foo Bar.docx
57956932603f9e194be4296fafecedf9  new_Helo Foo Bar.docx
155cac160bd06870efc2c5f01645d8a1  Helo Foo Bar.pdf
155cac160bd06870efc2c5f01645d8a1  new_Helo Foo Bar.pdf
```
***Running or building the example***

**Requires: golang 1.16+** - it probably works with much, much earlier versions but it was only tested with 1.16. 

**How To run:**
```go run main.go```

Or if you have *make* installed on a POSIX compliant or *nix platform:
```make run```

**Compiling:**
```
*nix/posix:
go build -o example main.go

windows:
go build -o example.exe main.go
```
or again with **make** with the following targets:
```
To compile for your OS only:
$ make

To compile specific to the target platform for 386,amd64 and arm64 depending upon platform
$ make darwin || linux || windows 

cross compiles darwin, linux and windows binaries on 386, amd64, arm64
$ make all

Create zip files for each platform that contain the binaries for that platform 
$ make zips    
GOOS=darwin GOARCH=amd64 go build -o 'example-darwin-amd64' main.go
GOOS=darwin GOARCH=arm64 go build -o 'example-darwin-arm64' main.go
GOOS=linux GOARCH=amd64 go build -o 'example-linux-amd64' main.go
GOOS=linux GOARCH=386 go build -o 'example-linux-386' main.go
GOOS=windows GOARCH=amd64 go build -o 'example-windows-amd64.exe' main.go
GOOS=windows GOARCH=386 go build -o 'example-windows-386.exe' main.go
zip windows.zip example-windows-*exe
  adding: example-windows-386.exe (deflated 42%)
  adding: example-windows-amd64.exe (deflated 44%)
zip linux.zip example-linux-*
  adding: example-linux-386 (deflated 41%)
  adding: example-linux-amd64 (deflated 43%)
zip mac.zip example-darwin-*
  adding: example-darwin-amd64 (deflated 43%)
  adding: example-darwin-arm64 (deflated 48%)

Run main.go
$ make run   

To remove compiled binaries and files created executing the program
$ make clean 
```
