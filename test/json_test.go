package test

import (
    // "bufio"
    "diablowu/golog/server/entity"
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

    // bio := bufio.NewReader(file)
    decoder := json.NewDecoder(file)
    var le entity.LogEntity
    decoder.Decode(&le)
    fmt.Println(le)

    for {

    }
}
