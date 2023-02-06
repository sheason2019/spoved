package compile

import (
	common "github.com/sheason2019/spoved/libs/idl-lib/common"
)

type GetCompileRecordResponse struct {
	Pagination common.Pagination `json:"pagination" form:"pagination"`
	Records    CompileRecord     `json:"records" form:"records"`
}

type CompileRecord struct {
	ProjectId  int    `json:"projectId" form:"projectId"`
	Image      string `json:"image" form:"image"`
	Version    string `json:"version" form:"version"`
	CreateAt   int    `json:"createAt" form:"createAt"`
	Operator   string `json:"operator" form:"operator"`
	Branch     string `json:"branch" form:"branch"`
	Output     string `json:"output" form:"output"`
	StatusCode int    `json:"statusCode" form:"statusCode"`
}
