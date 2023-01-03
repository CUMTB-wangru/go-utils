package gutil

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HttpDO http请求
func HttpDO(method string, url string, body io.Reader,
	header map[string][]string) (httpStatus int, resp []byte, err error) {
	return HttpDOTimeOut(method, url, body, header, 0)

}

func HttpDOTimeOut(method string, url string, body io.Reader,
	header map[string][]string, millisecond int) (httpStatus int, resp []byte, err error) {
	client := http.DefaultClient
	if millisecond > 0 {
		client.Timeout = time.Millisecond * time.Duration(millisecond)
	}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	if header != nil {
		request.Header = header
	}

	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer func() {
		_ = response.Body.Close()
	}()

	//if response.StatusCode == http.StatusOK {
	//	resp, err = ioutil.ReadAll(response.Body)
	//}
	resp, err = ioutil.ReadAll(response.Body)
	return response.StatusCode, resp, err

}

// HttpPost http POST 请求
func HttpPost(url string, body io.Reader,
	header map[string][]string) (httpStatus int, resp []byte, err error) {
	return HttpDO(http.MethodPost, url, body, header)
}

// HttpGet http GET 请求
func HttpGet(url string, body io.Reader,
	header map[string][]string) (httpStatus int, resp []byte, err error) {
	return HttpDO(http.MethodGet, url, body, header)
}

// HttpPostJson http Post Json 请求
func HttpPostJson(url string, body []byte,
	header map[string][]string) (httpStatus int, resp []byte, err error) {
	if header == nil {
		header = make(map[string][]string)
	}
	header["Content-Type"] = []string{"application/json"}
	return HttpDO(http.MethodPost, url, bytes.NewReader(body), header)
}

// HttpGetJson http GET Json 请求
func HttpGetJson(url string, body []byte,
	header map[string][]string) (httpStatus int, resp []byte, err error) {
	if header == nil {
		header = make(map[string][]string)
	}
	header["Content-Type"] = []string{"application/json"}
	return HttpDO(http.MethodGet, url, bytes.NewReader(body), header)
}

// HttpPostForm http Form 请求
func HttpPostForm(postUrl string, body map[string][]string,
	header map[string][]string) (httpStatus int, resp []byte, err error) {
	if header == nil {
		header = make(map[string][]string)
	}
	header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	return HttpDO(http.MethodPost,
		postUrl, strings.NewReader(url.Values(body).Encode()), header)
}

// HttpPostTimeOut http POST 请求
func HttpPostTimeOut(url string, body io.Reader,
	header map[string][]string, millisecond int) (httpStatus int, resp []byte, err error) {
	return HttpDOTimeOut(http.MethodPost, url, body, header, millisecond)
}

// HttpGetTimeOut http GET 请求
func HttpGetTimeOut(url string, body io.Reader,
	header map[string][]string, millisecond int) (httpStatus int, resp []byte, err error) {
	return HttpDOTimeOut(http.MethodGet, url, body, header, millisecond)
}

// HttpPostJsonTimeOut http Post Json 请求
func HttpPostJsonTimeOut(url string, body []byte,
	header map[string][]string, millisecond int) (httpStatus int, resp []byte, err error) {
	if header == nil {
		header = make(map[string][]string)
	}
	header["Content-Type"] = []string{"application/json"}
	return HttpDOTimeOut(http.MethodPost,
		url, bytes.NewReader(body), header, millisecond)
}

// HttpGetJsonTimeOut http GET Json 请求
func HttpGetJsonTimeOut(url string, body []byte,
	header map[string][]string, millisecond int) (httpStatus int, resp []byte, err error) {
	if header == nil {
		header = make(map[string][]string)
	}
	header["Content-Type"] = []string{"application/json"}
	return HttpDOTimeOut(http.MethodGet, url, bytes.NewReader(body), header, millisecond)
}

// HttpPostFormTimeOut http Form 请求
func HttpPostFormTimeOut(postUrl string, body map[string][]string,
	header map[string][]string, millisecond int) (httpStatus int, resp []byte, err error) {
	if header == nil {
		header = make(map[string][]string)
	}
	header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	return HttpDOTimeOut(http.MethodPost, postUrl,
		strings.NewReader(url.Values(body).Encode()), header, millisecond)
}
