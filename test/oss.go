package main

import (
    // "strings"
    // "strconv"
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
)

const (
    app_id          = "004174d3ac5b9000"
    user_credential = "basic weiwy3@lenovo.com:000000"
    developer_kid   = "82b5de4eca6fc69699cdcdff07944014"
    developer_key   = "29614d94afb302a87d782f4c6a3889c6"
)

func main() {
    resp, _ := http.Post("http://airx.lenovows.com/v1/app/connector", "", bytes.NewBufferString("s"))

    ba, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(ba)
}
