package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// 默认发送json格式请求
func HttpPost(requestUrl string, requestParams map[string]interface{}, requestHeaders map[string]string) ([]byte, int64, error) {
	return HttpPostJson(requestUrl, requestParams, requestHeaders)
}

// get请求
func HttpGet(requestUrl string, requestParams map[string]interface{}, requestHeaders map[string]string) ([]byte, int64, error) {

	// 参数处理 start
	Url, err := url.Parse(requestUrl)
	if err != nil {
		log.Printf("get request [%s] err %s\n", requestUrl, err)
	}

	params := url.Values{}
	if len(requestParams) > 0 {
		for k, v := range requestParams {
			params.Set(k, fmt.Sprintf("%v", v))
		}
	}

	Url.RawQuery = params.Encode()

	requestUrl = Url.String()

	log.Printf("get request [%s] start:\n", requestUrl)
	// 参数处理 end

	// 头部处理 start
	client := &http.Client{}
	requestGet, _ := http.NewRequest("GET", requestUrl, nil)

	if len(requestHeaders) > 0 {
		for headerKey, headerValue := range requestHeaders {
			requestGet.Header.Add(headerKey, headerValue)
		}
	}
	// 头部处理 end

	// 获取内容
	resp, err := client.Do(requestGet)
	if err != nil {
		log.Printf("get request [%s] failed, err:[%s]\n", requestUrl, err.Error())
		return nil, 500, err
	}
	// 关闭请求
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Printf("get request [%s] failed, err:[code:%d]\n", requestUrl, resp.StatusCode)
		return nil, int64(resp.StatusCode), err
	}

	log.Printf("get request [%s]resp status code:[%d]\n", requestUrl, resp.StatusCode)
	log.Printf("get request [%s]resp body data:[%s]\n", requestUrl, string(bodyContent))
	return bodyContent, 200, nil
}



// form请求post格式
func HttpPostForm(requestUrl string, requestParams map[string]interface{}, requestHeaders map[string]string) ([]byte, int64, error) {
	// 参数处理 start
	Url, err := url.Parse(requestUrl)
	if err != nil {
		log.Printf("get request [%s] err %s\n", requestUrl, err)
	}
	params := url.Values{}
	if len(requestParams) > 0 {
		for k, v := range requestParams {
			params.Set(k, fmt.Sprintf("%v", v))
		}
	}

	Url.RawQuery = params.Encode()

	requestUrl = Url.String()

	log.Printf("post request [%s] start:\n", requestUrl)
	// 参数处理 end

	client := &http.Client{}
	requestGet, _ := http.NewRequest("POST", requestUrl, nil)
	// 头部处理 start
	if len(requestHeaders) > 0 {
		for headerKey, headerValue := range requestHeaders {
			requestGet.Header.Add(headerKey, headerValue)
		}
	}
	requestGet.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 头部处理 end

	resp, err := client.Do(requestGet)
	if err != nil {
		log.Printf("post form request failed, err:[%s]", err.Error())
		return nil, 500, err
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Printf("get request [%s] failed, err:[code:%d]\n", requestUrl, resp.StatusCode)
		return nil, int64(resp.StatusCode), err
	}

	log.Printf("get request [%s]resp status code:[%d]\n", requestUrl, resp.StatusCode)
	log.Printf("get request [%s]resp body data:[%s]\n", requestUrl, string(bodyContent))
	return bodyContent, 200, nil
}

func HttpPostJson(requestUrl string, requestParams map[string]interface{}, requestHeaders map[string]string) ([]byte, int64, error) {
	// 参数处理 start
	jsonData, _ := json.Marshal(requestParams)
	// 参数处理 end

	client := &http.Client{}
	requestGet, _ := http.NewRequest("POST", requestUrl, bytes.NewReader(jsonData))

	// 头部处理 start
	if len(requestHeaders) > 0 {
		for headerKey, headerValue := range requestHeaders {
			requestGet.Header.Add(headerKey, headerValue)
		}
	}
	requestGet.Header.Set("Content-Type", "application/json")
	// 头部处理 end

	resp, err := client.Do(requestGet)
	if err != nil {
		log.Printf("post form request failed, err:[%s]", err.Error())
		return nil, 500, err
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Printf("get request [%s] failed, err:[code:%d]\n", requestUrl, resp.StatusCode)
		return nil, int64(resp.StatusCode), err
	}

	log.Printf("get request [%s]resp status code:[%d]\n", requestUrl, resp.StatusCode)
	log.Printf("get request [%s]resp body data:[%s]\n", requestUrl, string(bodyContent))
	return bodyContent, 200, nil
}
