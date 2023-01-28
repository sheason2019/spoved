package common

type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	ItemCounts int `json:"itemCounts"`
}
