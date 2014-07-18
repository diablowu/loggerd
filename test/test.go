package main

import (
    "github.com/wubo19842008/loggerd/server/rest"
    "net/http"
)

func main() {
    http.HandleFunc("/user/get", rest.Get)
    http.HandleFunc("/user/add", rest.Add)

    http.ListenAndServe("0.0.0.0:80", nil)

}
