// package oss

import (
    "bytes"
    "github.com/wubo19842008/loggerd/tools/jsons"
    "io/ioutil"
    "net/http"
    // "strings"
    // "strconv"
    // "bytes"
    "fmt"
    // "io/ioutil"
    // "net/http"
)


const (
    JSON_ContentType = "application/json"
)
const (
    gw_connector  = "http://airx.lenovows.com/v1/app/connector"
    gw_session    = "http://airx.lenovows.com/v1/app/session"
    gw_srv        = "http://cosx.lenovows.com/v2/service"
    gw_create_obj = "http://cosx.lenovows.com/v2/bucket/mslife"
    gw_upload     = "http://cosx.lenovows.com/v2/object/data/mslife/00417f3cf6800200"
)

func GetConn(uc map[string]interface{}) (conn string, err error) {
    body, err := jsons.ToJsonString(uc, false)
    if err != nil {
        return
    }

    resp, reqErr := http.Post(gw_connector, JSON_ContentType,
        bytes.NewBufferString(body))
    if reqErr != nil {
        err = reqErr
        return
    }

    ba, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        return
    }

    retObj, err := jsons.FromJsonString(string(ba))
    if err != nil {
        return
    }
    conn, _ = retObj["connector"].(string)

    return
}

func CreateSession(uc map[string]interface{}, conn string) (location string, err error) {
    req := make(map[string]interface{}, 3)
    req["connector"] = conn
    req["developer_kid"] = uc["developer_kid"]
    req["developer_key"] = uc["developer_key"]
    var reqStr string
    reqStr, err = jsons.ToJsonString(req, false)

    if err != nil {
        return
    }

    var resp *http.Response
    resp, err = http.DefaultClient.Post(
        gw_session,
        JSON_ContentType,
        bytes.NewBufferString(reqStr))

    if err != nil {
        return
    }

    location = resp.Header.Get("location")
    return
}

func main() {

    body, err := jsons.ToJsonString(oss_uc, false)
    if err != nil {
        fmt.Println(err)
    }
    // fmt.Println(body)
    resp, _ := http.Post(gw_connector, JSON_ContentType, bytes.NewBufferString(body))
    // fmt.Println(resp.Status)
    // for k, v := range resp.Header {
    //     fmt.Printf("key is %s value is %s \n", k, v)
    // }

    ba, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    retObj, _ := jsons.FromJsonString(string(ba))
    // session := retObj["session"]
    // uid := retObj["uid"]
    cct := retObj["connector"]

   

    sessionReqBody := make(map[string]interface{}, 3)
    sessionReqBody["connector"] = cct
    sessionReqBody["developer_kid"] = oss_uc["developer_kid"]
    sessionReqBody["developer_key"] = oss_uc["developer_key"]

    sessionReqStr, _ := jsons.ToJsonString(sessionReqBody, false)

    sessionResp, _ := http.DefaultClient.Post(gw_session, JSON_ContentType, bytes.NewBufferString(sessionReqStr))

    location := sessionResp.Header.Get("location")
    fmt.Println(location)

    buckectName := "mslife"

    // http.Post(gw_srv+"?bucket="+buckectName, contentType, "")

    newBktReq, _ := http.NewRequest("POST", gw_srv+"?bucket="+buckectName, nil)

    newBktReq.Header.Set("Authorization", location)
    newBktReq.Header.Set("X-Lenovows-OSS-Access-Control", "private")

    newBktResp, _ := http.DefaultClient.Do(newBktReq)

    fmt.Println(newBktResp.Status)
    // ba, _ = ioutil.ReadAll(newBktResp.Body)
    // bktResp := string(ba)
    defer newBktResp.Body.Close()
    // fmt.Println(bktResp)

    newObjReq, _ := http.NewRequest("POST", gw_create_obj, nil)
    newObjReq.Header.Set("Authorization", location)
    newObjResp, _ := http.DefaultClient.Do(newObjReq)

    ba, _ = ioutil.ReadAll(newObjResp.Body)
    defer newObjResp.Body.Close()

    nobjResp := string(ba)
    fmt.Println(nobjResp)

}
