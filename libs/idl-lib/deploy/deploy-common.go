package deploy

import (
	common "github.com/sheason2019/spoved/libs/idl-lib/common"
)

type DeployOrder struct {
	Id             int    `json:"id" form:"id"`
	ProjectId      int    `json:"projectId" form:"projectId"`
	Image          string `json:"image" form:"image"`
	Versoin        string `json:"versoin" form:"versoin"`
	CreateAt       int    `json:"createAt" form:"createAt"`
	Operator       string `json:"operator" form:"operator"`
	CompileOrderId int    `json:"compileOrderId" form:"compileOrderId"`
	StatusCode     int    `json:"statusCode" form:"statusCode"`
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
