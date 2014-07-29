package bucket

import (
    "encoding/json"
    "fmt"

    "testing"
)

func TestDecode(t *testing.T) {

    var obj BucketCreateResp
    // fmt.Println(SS)
    err := json.Unmarshal([]byte(SS), &obj)
    if err != nil {
        // panic(err)
        fmt.Println(err)
    } else {
        if len(obj.Tags) > 0 {
            for _, v := range obj.Tags {
                fmt.Println(v)
            }
        }
        // fmt.Println(obj.CannedAccessControl)
    }

}
