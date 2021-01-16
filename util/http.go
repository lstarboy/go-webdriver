package util

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ExtParams struct {
	Headers    map[string]string //请求头
	RetryCount int               //失败重试次数
}

//处理map到encode字符串
func EncodeUri(params map[string]interface{}) string {
	urls := url.Values{}
	if params != nil && len(params) > 0 {
		for k, v := range params {
			urls.Add(k, ToString(v))
		}
	}
	return urls.Encode()
}

//发起GET请求
func GetRequest(ctx *context.Context, reqUrl string, params map[string]interface{}, timeOut time.Duration, extParams ExtParams) (resBody string, resStatus int, resErr error) {
	if strings.Contains(reqUrl, "?") == false && params != nil && len(params) > 0 {
		reqUrl += "?"
	}
	reqUrl += EncodeUri(params)
	return DoRequest(ctx, reqUrl, "", "GET", timeOut, extParams)
}

//发起POST请求
func PostRequest(ctx *context.Context, reqUrl string, params map[string]interface{}, timeOut time.Duration, extParams ExtParams) (resBody string, resStatus int, resErr error) {
	if extParams.Headers == nil || len(extParams.Headers) == 0 {
		extParams.Headers = map[string]string{}
	}
	if _, ok := extParams.Headers["Content-Type"]; !ok {
		extParams.Headers["Content-Type"] = "application/x-www-form-urlencoded;charset=UTF-8"
	}
	return DoRequest(ctx, reqUrl, EncodeUri(params), "POST", timeOut, extParams)
}

//POST发送json数据
func PostJsonRequest(ctx *context.Context, reqUrl string, params map[string]interface{}, timeOut time.Duration, extParams ExtParams) (resBody string, resStatus int, resErr error) {
	if extParams.Headers == nil || len(extParams.Headers) == 0 {
		extParams.Headers = map[string]string{}
		extParams.Headers["Content-Type"] = "application/json;charset=UTF-8"
	} else {
		extParams.Headers["Content-Type"] = "application/json;charset=UTF-8"
	}
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return "", 0, errors.New("json解析错误:" + err.Error())
	}
	return DoRequest(ctx, reqUrl, string(jsonParams), "POST", timeOut, extParams)
}

//发起Http请求
func DoRequest(ctx *context.Context, reqUrl string, params string, method string, timeOut time.Duration, extParams ExtParams) (resBody string, resStatus int, resErr error) {
	reqTimeOut := timeOut
	retryCount := extParams.RetryCount
	if reqTimeOut <= 0 {
		reqTimeOut = 30 * time.Second //默认30秒超时
	}
	if retryCount == 0 {
		retryCount = 1 //默认重试一次
	}
	//添加常用http请求直接close链接，防止占用描述符链接资源
	if extParams.Headers == nil {
		extParams.Headers = map[string]string{
			"Connection": "close",
		}
	} else {
		if _, ok := extParams.Headers["Connection"]; !ok {
			extParams.Headers["Connection"] = "close"
		}
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		//MaxIdleConnsPerHost: 20,
		//DisableKeepAlives:false,
	}
	client := &http.Client{
		Timeout:   reqTimeOut,
		Transport: tr,
	}
	var res *http.Response
	var req *http.Request
	var err error
	for i := 0; i < retryCount; i++ {
		req, err = http.NewRequest(method, reqUrl, strings.NewReader(params))
		if err != nil {
			if res != nil {
				_ = res.Body.Close()
			}
			continue
		}
		if extParams.Headers != nil && len(extParams.Headers) > 0 {
			for k, v := range extParams.Headers {
				req.Header.Set(k, v)
			}
		}
		res, err = client.Do(req)
		if err != nil {
			if res != nil {
				_ = res.Body.Close()
			}
			continue
		} else {
			break
		}
	}
	statusCode := 0 //响应状态码
	if res != nil {
		statusCode = res.StatusCode
	}
	if err != nil {
		//请求异常
		return "", statusCode, err
	}

	defer func() {
		if res != nil {
			_ = res.Body.Close()
		}
	}()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//解析数据异常
		return "", statusCode, err
	}
	return string(body), statusCode, nil
}
