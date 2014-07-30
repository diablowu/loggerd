package oss

import (
    "github.com/wubo19842008/loggerd/tools/jsons"
    "github.com/wubo19842008/loggerd/tools/oss/utils"
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

func Login(uc map[string]interface{}) (location string, err error) {
    req := utils.ReqBody{}
    req.Method = "POST"
    req.Url = gw_connector

    body, _ := jsons.ToJsonString(uc, false)
    req.Body = []byte(body)

    resp := utils.RespBody{}
    resp, err = utils.RequestBody(req)

    retObj, _ := jsons.FromJsonString(string(resp.Body))

    conn, _ := retObj["connector"].(string)

    reqMap := make(map[string]interface{}, 3)
    reqMap["connector"] = conn
    reqMap["developer_kid"] = uc["developer_kid"]
    reqMap["developer_key"] = uc["developer_key"]
    var reqStr string
    reqStr, err = jsons.ToJsonString(reqMap, false)

    req = utils.ReqBody{}
    req.Method = "POST"
    req.Url = gw_session
    req.Body = []byte(reqStr)

    resp, err = utils.RequestBody(req)

    location = resp.Response.Header.Get("location")
    return
}
