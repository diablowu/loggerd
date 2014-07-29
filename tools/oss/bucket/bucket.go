package bucket

type BucketCreateResp struct {
    Id                  string   `json:"_id`
    CannedAccessControl string   `json:"canned_access_control"`
    LastModified        string   `json:"last_modified"`
    Location            string   `json:"location"`
    Name                string   `json:"name"`
    Tags                []string `json:"tags"`
    TimeCreate          string   `json:"time_created"`
}

func CreateBucket(location string, bucketName string) (resp BucketCreateResp, err error) {

}
