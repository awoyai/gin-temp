package common

type PageResult struct {
	List       interface{} `json:"list"`
	TotalCount int64       `json:"total_count"`
	PageNum    int         `json:"page_num"`
	PageSize   int         `json:"page_size"`
}

type PageInfo struct {
	PageNum    int64   `json:"page_num"`
	PageSize   int64   `json:"page_size"` // 每页大小
	TotalCount int64 `json:"total_count"`
	TotalPage  int64 `json:"total_page"`
}
