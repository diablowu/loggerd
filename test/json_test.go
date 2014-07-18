package test

import (
    // "github.com/wubo19842008/loggerd/server/entity"
    "io/ioutil"
    // "bufio"

    "encoding/json"
    "fmt"
    // "io/ioutil"
    "os"
    "testing"
)

func TestDecode(t *testing.T) {

    file, _ := os.Open("d:/temp/d.js")

    defer func() {
        if err := file.Close(); err != nil {
            panic(err)
        }
    }()

    ba, _ := ioutil.ReadAll(file)

    var vvv interface{}
    json.Unmarshal(ba, &vvv)
    // bio := bufio.NewReader(file)
    // decoder := json.NewDecoder(file)
    // var le entity.LogEntity
    // decoder.Decode(&le)
    // fmt.Println(v)

    m := vvv.(map[string]interface{})

    for k, v := range m {
        switch vv := v.(type) {
        case string:
            fmt.Println(k, "is string", vv)
        case int:
            fmt.Println(k, "is int", vv)
        case []interface{}:
            fmt.Println(k, "is an array:")
            for i, u := range vv {
                fmt.Println(i, u)
            }
        case bool:
            fmt.Println(k, "is boolean", vv)
        case float64:
            fmt.Println(k, "is float64", vv)
        default:
            fmt.Println(k, "is of a type I don't know how to handle")
        }
    }

}
