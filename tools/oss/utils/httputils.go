package utils

import (
    "bytes"
    "io/ioutil"
    "net/http"
)

const KEY_CONTENT_TYPE = "Content-type"

type ReqBody struct {
    Url         string
    Method      string
    Header      map[string]string
    Body        []byte
    ContentType string "application/json"
}

type RespBody struct {
    Status     string
    StatusCode int
    Body       []byte
}

//请求一个url同时发送一个json的body
func RequestBody(req ReqBody) (resp RespBody, err error) {
    bss := bytes.NewReader(req.Body)
    var request *http.Request
    request, err = http.NewRequest(req.Method, req.Url, bss)
    if err != nil {
        return
    }

    for k, v := range req.Header {
        request.Header.Set(k, v)
    }

    var response *http.Response
    response, err = http.DefaultClient.Do(request)

    if err != nil {
        return
    }

    resp.Status = response.Status
    resp.StatusCode = response.StatusCode
    resp.Body, err = ioutil.ReadAll(response.Body)
    defer response.Body.Close()

    if err != nil {
        return
    }

    return
}
