package main

import (
    "fmt"
    "time"
)

func main() {
    go req()

}

func req() {
    fmt.Println(time.Now().String())
}
