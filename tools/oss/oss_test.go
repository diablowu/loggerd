package oss

import (
    httputils "github.com/wubo19842008/loggerd/tools/oss/utils"
    "testing"
)

var oss_uc = map[string]interface{}{
    "app_id":          "",
    "user_credential": "",
    "developer_kid":   "",
    "developer_key":   "",
}

func TestGetConn(t *testing.T) {
    conn, err := GetConn(oss_uc)
    if err == nil {
        t.Log("connector is " + conn)
    } else {
        t.Error(err)
    }
}

func TestCreateSession(t *testing.T) {
    conn, _ := GetConn(oss_uc)
    location, err := CreateSession(oss_uc, conn)
    if err == nil {
        t.Log("location is %s", location)
    } else {
        t.Error(err)
    }
}

func TestHTTP(t *testing.T) {
    // httputils.
}
