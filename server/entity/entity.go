package entity

import (
// "time"
)

type LogEntity struct {
    Level     string `json:"level"`
    TransType string
    TransFrom string
    TransTo   string
    // start     time.Time
    // end       time.Time
    Request  string `json:"req"`
    Response string
    // create    time.Time
    // update    time.Time
}
