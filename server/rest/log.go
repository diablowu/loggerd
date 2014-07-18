package rest

import (
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "time"

    "net/http"
)

func Add(w http.ResponseWriter, req *http.Request) {
    ba, err := ioutil.ReadAll(req.Body)
    if err != nil {
        io.WriteString(w, "error.")
    }
    // body := string(ba)

    var f interface{}

    json.Unmarshal(ba, &f)

    m := f.(map[string]interface{})

    file, err := os.OpenFile("d:/temp/db", os.O_APPEND|os.O_WRONLY, 0600)

    defer func() {
        if err := file.Close(); err != nil {
            panic(err)
        }
    }()

    for k, v := range m {
        line := k + ":" + v.(string)
        // file.WriteString()
        fmt.Println(line)
    }

    // io.WriteString(w, time.Now().Format(time.ANSIC))
}

func Get(w http.ResponseWriter, req *http.Request) {
    go doGet()
}

func afterDoGet() {
    fmt.Println("after do get")
}

func doGet() {
    fmt.Println("Get...")
    time.Sleep(time.Second * 5)
    fmt.Println("Get over...")

}
