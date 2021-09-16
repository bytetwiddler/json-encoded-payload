Example of embedding binary base64 encoded data in a json object.

The main.go code will take the 2 included 'Helo Foo Bar.docx' and 'Helo Foo Bar.pdf' files in the project root and convert them to base64 URLEncoded json objects.  It will then print out a JSON array of those files with a fileName, encodingType, encodedData properties.  It then takes that JSON object and convert it back to a new files prepending "new_" to the original fileName.  You should be left with 2 docx and 2 pdf files that are the same except for their names.

Here is a single file json object:
```
{
    "fileName": "Helo Foo Bar.docx",
    "encodingType": "Base64",
    "encodedData": "UEsDBBQABgAIAA...bWxQSwUGAAAAAAsACwDBAgAAjyoAAAAA"
}
```

Here is a an array of files json object:
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
$ ls *docx *pdf
'Helo Foo Bar.docx'  'Helo Foo Bar.pdf'

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

New file should now be created with "new_" prepended to the old file names:
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
