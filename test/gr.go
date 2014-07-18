package main

import (
    "fmt"
    "os"
    "strconv"
    "time"
)

func wf(cnt int) {
    file, _ := os.Create("d:/temp/a_" + strconv.Itoa(cnt) + ".log")
    file.WriteString(time.Now().String())
    file.Close()
    fmt.Printf("#%d over", cnt)
}

func main() {

    for i := 0; i < 10; i++ {
        go wf(i)
    }

    time.Sleep(time.Minute * 100)
    fmt.Printf("Go.........")
}
