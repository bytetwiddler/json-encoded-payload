Example of embedding binary base64 encoded data in a json object.

The main.go code will take the 2 included binary files in the project root and convert them to base64 URLEncoded json objects.  It will then take each JSON object and conver it back to a file prepeinding "new_" to the original fileName.

Here is a single file json object:
```
{
    "fileName": "Helo Foo Bar.docx",
    "encodingType": "Base64",
    "encodedData": "UEsDBBQABgAIAAAAIQDfpNJsWgEAACAFAAATAAgCW0NvbnRl...FNwAAAPAAAAAAAAAAAAAAAAADgfAAB3b3JkL3N0eWxlcy54bWxQSwUGAAAAAAsACwDBAgAAjyoAAAAA"
  }
  ```

Here is a an array of files json object:
```
[
  {
    "fileName": "Helo Foo Bar.docx",
    "encodingType": "Base64",
    "encodedData": "UEsDBBQABgAIAAAAIQDfpNJsWgEAACAFAAATAAgCW0NvbnRlbnRfVHlwZXNdLnhtb...AAADgfAAB3b3JkL3N0eWxlcy54bWxQSwUGAAAAAAsACwDBAgAAjyoAAAAA"
  },
  {
    "fileName": "Helo Foo Bar.pdf",
    "encodingType": "Base64",
    "encodedData": "JVBERi0xLjUKJcOkw7zDtsOfCjIgMCBvYmoKPDwvTGVuZ3RoIDMgMCBSL0ZpbHRl...HhyZWYKODYyNwolJUVPRgo="
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
    "encodedData": "UEsDBBQABgAIAAAAIQDfpNJsWgEAACAFAAATAAgCW0NvbnRlbnRfVHlwZXNdLnhtb...AAADgfAAB3b3JkL3N0eWxlcy54bWxQSwUGAAAAAAsACwDBAgAAjyoAAAAA"
  },
  {
    "fileName": "Helo Foo Bar.pdf",
    "encodingType": "Base64",
    "encodedData": "JVBERi0xLjUKJcOkw7zDtsOfCjIgMCBvYmoKPDwvTGVuZ3RoIDMgMCBSL0ZpbHRl...HhyZWYKODYyNwolJUVPRgo="
  }
]
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
