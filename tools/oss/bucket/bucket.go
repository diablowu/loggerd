package bucket

import (
    "encoding/json"

    "github.com/wubo19842008/loggerd/tools/oss/utils"
)

const (
    gw_url = "http://cosx.lenovows.com/v2/service"
)

type BucketCreateResp struct {
    Id                  string   `json:"_id`
    CannedAccessControl string   `json:"canned_access_control"`
    LastModified        string   `json:"last_modified"`
    Location            string   `json:"location"`
    Name                string   `json:"name"`
    Tags                []string `json:"tags"`
    TimeCreate          string   `json:"time_created"`
}

func CreateBucket(location string, bucketName string) (bcr BucketCreateResp, err error) {
    req := utils.ReqBody{}
    header := make(map[string]string)

    header["Authorization"] = location
    header["X-Lenovows-OSS-Access-Control"] = "public-read-write"
    req.Url = gw_url + "?bucket=" + bucketName
    req.Method = "POST"

    resp, _ := utils.RequestBody(req)

    err = json.Unmarshal(resp.Body, &bcr)
    return
}
