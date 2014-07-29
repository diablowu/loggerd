package utils

import (
    "fmt"
    "github.com/wubo19842008/loggerd/tools/jsons"
    "testing"
)

var oss_uc = map[string]interface{}{
    "app_id":          "",
    "user_credential": "",
    "developer_kid":   "",
    "developer_key":   "",
}

func TestReq(t *testing.T) {
    reqBody := ReqBody{}
    reqBody.Method = "POST"
    reqBody.Url = "http://airx.lenovows.com/v1/app/connector"

    jsonsstr, _ := jsons.ToJsonString(oss_uc, false)
    reqBody.Body = []byte(jsonsstr)

    // header := make(map[string]string)
    // header[""] = "gogog"
    // reqBody.Header = header

    // fmt.Println(reqBody)

    resp, err := RequestBody(reqBody)

    if err != nil {
        t.Error(err)
    }

    // fmt.Println(string(resp.Body))

    retObj, _ := jsons.FromJsonString(string(resp.Body))

    v, ok := retObj["connector"].(string)

    if !ok {
        t.Error("get connector error.")

    }

}
