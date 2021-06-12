package request

import (
	"bytes"
	"easygo/pkg/wrapper"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewPost(url string, reqBody *bytes.Buffer, content_type string) (string, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", content_type)
	req.Header.Set("Host", url)
	req.Header.Set("Content-Length", fmt.Sprint(reqBody.Len()))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", wrapper.ErrBadRequest
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	return string(resBody), nil
}
