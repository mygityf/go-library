package httplib

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	// html内容类型
	ContentTypeTextHtml = "text/html;charset=utf-8"
	// json内容类型
	ContentTypeJson = "application/json;charset=utf-8"
	// png内容类型
	ContentTypePng = "image/png"
	// jpg内容类型
	ContentTypeJpeg = "image/jpeg"
	// pdf内容类型
	ContentTypePdf = "application/pdf"
	// 二进制
	ContentTypeOctetStream = "application/octet-stream"
)

// DoPostHttp post请求
func DoPostHttp(url, contentType, reqString string, reqTimeout time.Duration) ([]byte, error) {
	client := http.Client{
		Timeout: reqTimeout,
	}
	resp, err := client.Post(url, contentType, strings.NewReader(reqString))
	if err != nil {
		return nil, err
	}
	return readResponse(resp)
}

// DoPostHttp post请求
func DoPostHttpWithHeader(
	url, contentType, reqString string, reqTimeout time.Duration, headers http.Header,
) ([]byte, error) {
	client := http.Client{
		Timeout: reqTimeout,
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(reqString))
	if err != nil {
		return nil, errors.New("http.NewRequest err:" + err.Error())
	}
	req.Header.Set("Content-Type", contentType)
	for key, value := range headers {
		if len(value) <= 0 {
			continue
		}
		req.Header.Set(key, value[0])
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("client.Do err:" + err.Error())
	}
	return readResponse(resp)
}

// DoGetHttp get请求
func DoGetHttp(url string, reqTimeout time.Duration) ([]byte, error) {
	client := http.Client{
		Timeout: reqTimeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return readResponse(resp)
}

// 取结果
func readResponse(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return nil, errors.New("resp is nil")
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("%s: %s", resp.Status, body)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
