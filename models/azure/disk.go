package azure

type Disk struct {
	Status string `json:"status"`
	SizeGb int64  `json:"size"`
}
