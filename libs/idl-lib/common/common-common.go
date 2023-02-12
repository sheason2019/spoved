package common

type Pagination struct {
	Page       int `json:"page" form:"page"`
	PageSize   int `json:"pageSize" form:"pageSize"`
	ItemCounts int `json:"itemCounts" form:"itemCounts"`
}
