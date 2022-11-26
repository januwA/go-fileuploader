package strategy

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

// 将文件保存到网络存储策略
type HttpStrategy struct {
	Filename      string
	FileFieldName string
	UploadUrl     *url.URL
}

// 创建一个携带文件的formdata
func (my *HttpStrategy) makeFormData(reader_in io.Reader, formdata_out io.Writer) (*multipart.Writer, error) {
	w := multipart.NewWriter(formdata_out)

	// w.WriteField("format", "json")
	part, err := w.CreateFormFile(my.FileFieldName, my.Filename)
	if err != nil {
		return nil, err
	}
	// part2 := bufio.NewWriter(part)
	// part2.ReadFrom(reader_in)
	// part2.Flush()

	f, _ := io.ReadAll(reader_in)
	part.Write(f)

	w.Close()

	return w, nil
}

// 创建发送文件的http请求
func (my *HttpStrategy) makeReq(reader_in io.Reader) (*http.Request, error) {
	buf := new(bytes.Buffer)
	formData, err := my.makeFormData(reader_in, buf)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", my.UploadUrl.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", formData.FormDataContentType())
	return req, nil
}

// result_out *[]byte
func (my *HttpStrategy) Save(reader_in io.Reader, result_out any) error {
	result := result_out.(*[]byte)

	req, err := my.makeReq(reader_in)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	*result = bodyBytes
	return nil
}
