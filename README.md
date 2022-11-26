## 简单的文件上传模型

你可以实现自己的`IStrategy`在你的项目中运行。


## Example

```go
package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/januwA/go-fileuploader"
	fu_strategy "github.com/januwA/go-fileuploader/strategy"
)

func main() {
	f, _ := os.Open("./upload.jpg")
	defer f.Close()
  
	fu := new(fileuploader.FileUploader)

	// LocalStrategy
	fu.SetStrategy(&fu_strategy.LocalStrategy{
		Dir:  "./",
		Name: "out",
		Ext:  ".jpg",
	})
	var r1 string
	fu.Save(f, &r1)
	fmt.Printf("LocalStrategy result: %s\n", r1)

	// HttpStrategy
	UploadUrl, _ := url.Parse("http://127.0.0.1:7777/api/upload")
	fu.SetStrategy(&fu_strategy.HttpStrategy{
		UploadUrl:     UploadUrl,
		FileFieldName: "file",
		Filename:      "out.jpg",
	})

	var r2 []byte
	fu.Save(f, &r2)
	fmt.Printf("HttpStrategy result: %s\n", string(r2))
}

```