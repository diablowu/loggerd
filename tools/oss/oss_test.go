package oss

import (
    "testing"
)

var oss_uc = map[string]interface{}{
    "app_id":          "",
    "user_credential": "",
    "developer_kid":   "82b5de4eca6fc69699cdcdff07944014",
    "developer_key":   "29614d94afb302a87d782f4c6a3889c6",
}

// func TestHTTP(t *testing.T) {
//     body, err := jsons.ToJsonString(oss_uc, false)
//     if err != nil {
//         fmt.Println(err)
//     }
//     // fmt.Println(body)
//     resp, _ := http.Post(gw_connector, JSON_ContentType, bytes.NewBufferString(body))
//     // fmt.Println(resp.Status)
//     // for k, v := range resp.Header {
//     //     fmt.Printf("key is %s value is %s \n", k, v)
//     // }

//     ba, _ := ioutil.ReadAll(resp.Body)
//     defer resp.Body.Close()
//     retObj, _ := jsons.FromJsonString(string(ba))
//     // session := retObj["session"]
//     // uid := retObj["uid"]
//     cct := retObj["connector"]

//     sessionReqBody := make(map[string]interface{}, 3)
//     sessionReqBody["connector"] = cct
//     sessionReqBody["developer_kid"] = oss_uc["developer_kid"]
//     sessionReqBody["developer_key"] = oss_uc["developer_key"]

//     sessionReqStr, _ := jsons.ToJsonString(sessionReqBody, false)

//     sessionResp, _ := http.DefaultClient.Post(gw_session, JSON_ContentType, bytes.NewBufferString(sessionReqStr))

//     location := sessionResp.Header.Get("location")
//     fmt.Println(location)

//     buckectName := "mslife"

//     // http.Post(gw_srv+"?bucket="+buckectName, contentType, "")

//     newBktReq, _ := http.NewRequest("POST", gw_srv+"?bucket="+buckectName, nil)

//     newBktReq.Header.Set("Authorization", location)
//     newBktReq.Header.Set("X-Lenovows-OSS-Access-Control", "private")

//     newBktResp, _ := http.DefaultClient.Do(newBktReq)

//     fmt.Println(newBktResp.Status)
//     // ba, _ = ioutil.ReadAll(newBktResp.Body)
//     // bktResp := string(ba)
//     defer newBktResp.Body.Close()
//     // fmt.Println(bktResp)

//     newObjReq, _ := http.NewRequest("POST", gw_create_obj, nil)
//     newObjReq.Header.Set("Authorization", location)
//     newObjResp, _ := http.DefaultClient.Do(newObjReq)

//     ba, _ = ioutil.ReadAll(newObjResp.Body)
//     defer newObjResp.Body.Close()

//     nobjResp := string(ba)
//     fmt.Println(nobjResp)
// }

func TestLogin(t *testing.T) {
    location, err := Login(oss_uc)
    if err == nil {
        t.Logf("location is\n %s\n", location)
    } else {
        t.Error(err)
    }
}
