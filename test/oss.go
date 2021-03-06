package main

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

var OSS_UC = map[string]interface{}{
    "app_id":          "",
    "user_credential": "",
    "developer_kid":   "",
    "developer_key":   "",
}

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

func getConn(uc map[string]string) (conn string, err error) {
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
    conn = retObj["connector"]
    return
}

func main() {

    body, err := jsons.ToJsonString(oss_uc, false)
    if err != nil {
        fmt.Println(err)
    }
    // fmt.Println(body)
    resp, _ := http.Post(gw_connector, contentType, bytes.NewBufferString(body))
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

    // fmt.Println(session, uid, cct)

    //"{\"connector\":\"$Connector\",\"developer_kid\":\"$KeyId\",\"developer_key\":\"$Key
    //curl -X POST "http://airx.lenovows.com/v1/app/session" -d "{\"connector\":\"$Connector\",\"developer_kid\":\"$KeyId\",\"developer_key\":\"$Key" -v

    sessionReqBody := make(map[string]interface{}, 3)
    sessionReqBody["connector"] = cct
    sessionReqBody["developer_kid"] = oss_uc["developer_kid"]
    sessionReqBody["developer_key"] = oss_uc["developer_key"]

    sessionReqStr, _ := jsons.ToJsonString(sessionReqBody, false)

    sessionResp, _ := http.DefaultClient.Post(gw_session, contentType, bytes.NewBufferString(sessionReqStr))

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
