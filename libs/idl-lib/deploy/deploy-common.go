package deploy

import (
	common "github.com/sheason2019/spoved/libs/idl-lib/common"
	compile "github.com/sheason2019/spoved/libs/idl-lib/compile"
)

type DeployOrder struct {
	Id           int                  `json:"id" form:"id"`
	Image        string               `json:"image" form:"image"`
	CreateAt     int                  `json:"createAt" form:"createAt"`
	Operator     string               `json:"operator" form:"operator"`
	CompileOrder compile.CompileOrder `json:"compileOrder" form:"compileOrder"`
	StatusCode   int                  `json:"statusCode" form:"statusCode"`
}

type GetDeployOrdersPayload struct {
	ProjectId int `json:"projectId" form:"projectId"`
	Page      int `json:"page" form:"page"`
	PageSize  int `json:"pageSize" form:"pageSize"`
}

type GetDeployOrdersResponse struct {
	Records    []DeployOrder     `json:"records" form:"records"`
	Pagination common.Pagination `json:"pagination" form:"pagination"`
}

type GetOptionalImagesResponse struct {
	Images []string `json:"images" form:"images"`
}
