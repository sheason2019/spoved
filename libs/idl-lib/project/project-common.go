package project

import (
	common "github.com/sheason2019/spoved/libs/idl-lib/common"
)

type Project struct {
	Id          int    `json:"id"`
	Owner       string `json:"owner"`
	ProjectName string `json:"projectName"`
	GitUrl      string `json:"gitUrl"`
	Describe    string `json:"describe"`
}

type GetProjectsResponse struct {
	Projects   []Project         `json:"projects"`
	Pagination common.Pagination `json:"pagination"`
}
