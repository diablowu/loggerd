package bucket

import (
    "testing"
)

func TestCreat(t *testing.T) {
    lc := "http://iam.lenovows.com/v1/session/004181578ea9e000"
    resp, err := CreateBucket(lc, "mspublic")

    if err == nil {
        t.Log(resp)
    } else {
        t.Error(err)
    }
}
