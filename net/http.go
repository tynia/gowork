package net

import (
	"net/http"
	"io/ioutil"
	"io"
	"encoding/json"
	"bytes"
	"strings"
	"time"
	"fmt"
	e "gowork/error"
	"gowork/extern/logging"
)

var (
	HTTP_METHOD_GET    = "GET"
	HTTP_METHOD_PUT    = "PUT"
	HTTP_METHOD_POST   = "POST"
	HTTP_METHOD_PATCH  = "PATCH"
	HTTP_METHOD_DELETE = "DELETE"

	CONTENT_NONE = ""
	CONTENT_JSON = "application/json"
	CONTENT_YAML = "application/yaml"
)

type HTTPRequest struct {
	Method string
	URL    string
	UserID string
	X_Auth_Token string
	Form interface{}
	Body interface{}
}

type HTTPResponse struct {
	Code int     `json:"code, omitempty"`
	Msg  string  `json:"msg, omitempty"`
	Data interface{} `json:"data, omitempty"`
}

type IRequest interface {
	DoRequest(resp interface{}) *e.WError
}

// warpper
type Responser http.ResponseWriter
type Requester http.Request
type HandlerFunc func(http.ResponseWriter, *http.Request)
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func Serve(srvPort string, handler http.Handler) {
	logging.Info("[Serve] Try to listen on port: %s", srvPort)
	go func() {
		err := http.ListenAndServe(srvPort, handler)
		if err != nil {
			logging.Error("[Serve] Listen failed, error = %s", err.Error())
			return
		}
	}()
}

func Health(srvPort string, handler http.Handler) {
	logging.Info("[Health] Try to monitor health condition on port: %s", srvPort)
	go func() {
		err := http.ListenAndServe(srvPort, handler)
		if err != nil {
			logging.Error("[Health] monitor failed, error = %s", err.Error())
			return
		}
	}()
}

func HandleFunc(addr string, handler HandlerFunc) {
	http.HandleFunc(addr, handler)
}

///////////////////////////////////////////////////////////////////////
type Request struct {
	ID      string
	URL     string
	Method  string
	Content string
	Req     *http.Request
}

func NewRequest(id string, target string, method string, content string, contentType string) *Request {
	req := &Request{
		ID:      id,
		URL:     target,
		Method:  method,
		Content: content,
	}

	request, err := http.NewRequest(req.Method, req.URL, strings.NewReader(req.Content))
	if err != nil {
		logging.Error("[NewRequest] Fatal error when create request, error = %s, url = %s", err.Error(), target)
		return nil
	}

	if contentType != CONTENT_NONE {
		request.Header.Set("Content-Type", contentType)
	}

	req.Req = request

	return req
}

func (req *Request) AddHeader(key string, value string) {
	req.Req.Header.Add(key, value)
}

func (req *Request) DoRequest(resp interface{}) (*http.Response, *e.WError) {
	logging.Debug("[Request.DoRequest, id: %s] start connection[to: %s, method: %s, content: %s]", req.ID, req.URL, req.Method, req.Content)
	client := &http.Client{}
	client.Timeout = time.Duration(4 * time.Second)

	response, err := client.Do(req.Req)
	if err != nil {
		logging.Error("[Request.DoRequest] Failed to talk with remote server, error = %s, id: %s", err.Error(), req.ID)
		return nil, e.NewWError(1003, "Failed to request to remote server[url: %s], error = %s", req.URL, err.Error())
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		logging.Error("[Request.DoRequest] Error returns with code: %d, msg: %s", response.StatusCode, response.Status)
		return nil, e.NewWError(1003, "Error returns with code: %d, msg: %s", response.StatusCode, response.Status)
	}

	if resp != nil {
		defer response.Body.Close()

		text, err := ioutil.ReadAll(response.Body)
		if err != nil {
			logging.Error("[Request.DoRequest, id: %s] Failed to read body of response: %v", req.ID, response.Body)
			return nil, e.NewWError(1003, err.Error())
		}

		err = json.Unmarshal(text, &resp)
		if err != nil {
			logging.Error("[Request.DoRequest, id: %s] json unmarshal, error: %s", req.ID, err.Error())
			return nil, e.NewWError(1003, err.Error())
		}

		return nil, nil
	}

	return response, nil
}

func Serialize(source map[string]interface{}) string {
	text, err := json.Marshal(&source)
	if err != nil {
		logging.Error("[Serialize] Failed to convert map to json byte, error: %s", err.Error())
		return ""
	}

	buff := bytes.NewBuffer(text)

	return string(buff.Bytes())
}

func Call(name string, url string, method string, content string, contentType string, resp interface{}) (string, *e.WError) {
	req := NewRequest(name, url, method, content, contentType)
	if req == nil {
		logging.Error("[Call, id: %s] Failed to create Request", name)
		return "", e.NewWError(1003, "Failed to create Request")
	}

	if resp != nil {
		_, err := req.DoRequest(resp)
		if err != nil {
			logging.Error("[Call, id: %s] Failed to do Request, error = %s", name, err.Error())
			return "", err
		}

		return "", nil
	}

	response, err := req.DoRequest(nil)
	if err != nil {
		logging.Error("[Call, id: %s] Failed to do Request, error = %s", name, err.Error())
		return "", err
	}

	text, ee := ioutil.ReadAll(response.Body)
	if ee != nil {
		logging.Error("[Call, id: %s] Failed to read body of response: %v", name, response.Body)
		return "", e.NewWError(1003, ee.Error())
	}

	return string(text), nil
}

func GetRequestBody(req *http.Request, v interface{}) (interface{}, *e.WError) {
	buff := bytes.NewBufferString("")
	_, err := io.Copy(buff, req.Body)
	if err != nil {
		logging.Error("[GetRequestBody] Failed to copy req body error, request = %v, error = %s", req.Body, err.Error())
		return nil, e.NewWError(e.ERR_CODE_IO, "Failed to copy req body, error = %s", err.Error())
	}
	text := buff.String()
	req.Body = ioutil.NopCloser(strings.NewReader(text))

	err = json.Unmarshal([]byte(text), v)
	if err != nil {
		logging.Error("[GetRequestBody] Failed to unmarshal json body error, text = %s, error = %s", text, err.Error())
		return nil, e.NewWError(e.ERR_CODE_IO, "Failed to unmarshal json body, error = %s", err.Error())
	}

	return v, nil
}

func GetResponseData(err *e.WError, data interface{}) []byte {
	resp := &HTTPResponse{}
	if err == nil {
		resp.Code = 0;
		resp.Msg = "ok"
	} else {
		resp.Code = err.Code()
		resp.Msg = err.Error()
	}
	resp.Data = data

	r, we := json.Marshal(resp)
	if we != nil {
		logging.Error("[GetResponseData] Failed, %s, data = %v", err.Error(), data)
		panic(fmt.Sprintf("[GetResponseData] Failed, %s", we.Error()))
	}

	return r
}

func LogGetResponseData(req *http.Request, err *e.WError, data interface{}) []byte {
	ret := GetResponseData(err, data)

	body := ""
	b_body, ee := ioutil.ReadAll(req.Body)
	if ee == nil {
		body = string(b_body)
	}

	logReq := HTTPRequest{
		Method: req.Method,
		URL: req.RequestURI,
		Form: req.Form,
		Body: body,
	}
	logging.Info("HANDLE_LOG: url = %s, request = %#v", req.RequestURI, logReq)
	logging.Debug("HANDLE_RESPONSE: response = %s", string(ret))

	return ret
}