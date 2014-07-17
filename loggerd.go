package main

import (
    "fmt"
    "net/http"
)

type User struct {
    id   int64
    name string
}

func main() {
    // uu := new User{1,"asdf"}

    uu := &User{1, "asdfd"}

    fmt.Printf("%p", uu)

    http.HandleFunc("/logger/add", func(w http.ResponseWriter, r *http.Request) {
        fmt.Print("add ...")
    })

    err := http.ListenAndServe("0.0.0.0:8080", nil)

    if err != nil {
        panic(err)
    }

}
