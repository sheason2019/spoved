package compile

import (
	common "github.com/sheason2019/spoved/libs/idl-lib/common"
)

type CompileOrder struct {
	Id         int    `json:"id" form:"id"`
	ProjectId  int    `json:"projectId" form:"projectId"`
	Image      string `json:"image" form:"image"`
	Version    string `json:"version" form:"version"`
	CreateAt   int    `json:"createAt" form:"createAt"`
	Operator   string `json:"operator" form:"operator"`
	Branch     string `json:"branch" form:"branch"`
	StatusCode int    `json:"statusCode" form:"statusCode"`
}

type GetCompileOrdersPayload struct {
	ProjectId int `json:"projectId" form:"projectId"`
	Page      int `json:"page" form:"page"`
	PageSize  int `json:"pageSize" form:"pageSize"`
}

type GetCompileOrdersResponse struct {
	Records    []CompileOrder    `json:"records" form:"records"`
	Pagination common.Pagination `json:"pagination" form:"pagination"`
}
