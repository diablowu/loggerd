package main

import (
    "fmt"
    "github.com/qiniu/api/io"
)

func main() {
    fmt.Println(glgserver.Port)

    ggs := glgserver.Create(0, "127.0.0.1:9999", true)

    ggs.Start()

}
