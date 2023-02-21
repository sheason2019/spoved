package project

import (
	common "github.com/sheason2019/spoved/libs/idl-lib/common"
)

type GetProjectPayload struct {
	Username    string `json:"username" form:"username"`
	ProjectName string `json:"projectName" form:"projectName"`
}

type Project struct {
	Id          int    `json:"id" form:"id"`
	Owner       string `json:"owner" form:"owner"`
	ProjectName string `json:"projectName" form:"projectName"`
	GitUrl      string `json:"gitUrl" form:"gitUrl"`
	Describe    string `json:"describe" form:"describe"`
}

type GetProjectsResponse struct {
	Projects   []Project         `json:"projects" form:"projects"`
	Pagination common.Pagination `json:"pagination" form:"pagination"`
}
