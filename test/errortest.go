package main1

import (
    "runtime"
    "time"
    // "errors"
    "fmt"
)

type Coffee struct {
    fee float32
}

func (c Coffee) getTotal() float32 {
    return c.fee
}

type SoftDrink interface {
    getTotal() float32
}
type Mocha struct {
    Coffee
}

func main() {
    go sale(Mocha{Coffee{12.33}})
    go sale(Mocha{Coffee{12.33}})
    go sale(Mocha{Coffee{12.33}})
    go sale(Mocha{Coffee{12.33}})
    go sale(Mocha{Coffee{12.33}})
    sale(Mocha{Coffee{12.33}})
    // time.Sleep(1 * time.Hour)

}

func sale(sd SoftDrink) {
    runtime.Gosched()
    fmt.Println(sd.getTotal())
    time.Sleep(5 * time.Second)
}
