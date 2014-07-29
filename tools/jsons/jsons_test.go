package jsons

import (
    "fmt"
    "strings"
    "testing"
)

func TestEncode(t *testing.T) {
    obj := make(map[string]interface{})
    obj["app_id"] = "1111111111111111"

    ss, e := ToJsonString(obj, false)
    if e == nil {
        if strings.Index(ss, "1111111111111111") > 0 {
            t.Log("encode success.")
        } else {
            t.Error("encode error.")
        }
    } else {
        t.Error(e)
    }
}

func TestDecode(t *testing.T) {
    jsons := `
    {
  "app_id": "",
  "user_credential": "",
  "developer_kid": "",
  "developer_key": ""
}
`
    obj, err := FromJsonString(jsons)

    if err == nil && obj != nil {

        if v, ok := obj["app_id"]; ok {
            fmt.Printf("app_id is %s\n", v)
        } else {
            t.Error("decode error.")
        }
    } else {
        t.Error("decode error.")
    }

}
func BenchmarkJsonsDecode(t *testing.B) {
    for i := 0; i < t.N; i++ {
        jsons := `
    {
  "app_id": "",
  "user_credential": "",
  "developer_kid": "",
  "developer_key": ""
}
`
        obj, err := FromJsonString(jsons)

        if err != nil {
            t.Error(err)
        } else {
            if obj == nil {
                // fmt.Println(obj)
                v, ok := obj["app_id"]
                if !ok {
                    t.Error("decode error.")
                } else {
                    fmt.Printf("app_id is %s\n", v)
                }
            }
        }
    }
}
