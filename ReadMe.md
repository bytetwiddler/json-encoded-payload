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
