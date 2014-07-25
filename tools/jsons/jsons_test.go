package jsons

import (
    "fmt"
    "testing"
)

func TestJsonSEncode(t *testing.T) {
    obj := make(map[string]interface{})

    ss, e := ToJsonString(obj, false)
    if e == nil {
        fmt.Println(ss)
    } else {
        t.Error(e)
    }
}

func TestJsonsDecode(t *testing.T) {
    jsons := `
    {
  "app_id": "004174d3ac5b9000",
  "user_credential": "basic weiwy3@lenovo.com:000000",
  "developer_kid": "82b5de4eca6fc69699cdcdff07944014",
  "developer_key": "29614d94afb302a87d782f4c6a3889c6"
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
func BenchmarkJsonsDecode(t *testing.B) {
    for i := 0; i < t.N; i++ {
        jsons := `
    {
  "app_id": "004174d3ac5b9000",
  "user_credential": "basic weiwy3@lenovo.com:000000",
  "developer_kid": "82b5de4eca6fc69699cdcdff07944014",
  "developer_key": "29614d94afb302a87d782f4c6a3889c6"
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
