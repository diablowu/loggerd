package jsons

import (
    "encoding/json"
    "strings"
)

//format a json map to json string
func ToJsonString(obj map[string]interface{}, pretty bool) (jsons string, e error) {
    var ba []byte

    if pretty {
        ba, e = json.MarshalIndent(obj, "", "\t")
    } else {
        ba, e = json.Marshal(obj)
    }

    if e == nil {
        jsons = string(ba)
    }

    return
}

func FromJsonString(jsstring string) (obj map[string]interface{}, e error) {

    ss := strings.TrimSpace(jsstring)
    if len(ss) > 0 {
        var vs interface{}
        e = json.Unmarshal([]byte(ss), &vs)
        if e == nil {
            obj = vs.(map[string]interface{})
        }
    }
    return
}
