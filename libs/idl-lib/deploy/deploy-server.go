package deploy

import (
	"github.com/gin-gonic/gin"
)

type DeployApi interface {
	GetDeployOrders(ctx *gin.Context, payload GetDeployOrdersPayload) GetDeployOrdersResponse
	PostDeployOrder(ctx *gin.Context, order DeployOrder) DeployOrder
	GetOptionalImages(ctx *gin.Context) GetOptionalImagesResponse
}
type _deployApiDefinition struct {
	GET_DEPLOY_ORDERS_PATH   string
	POST_DEPLOY_ORDER_PATH   string
	GET_OPTIONAL_IMAGES_PATH string
}

var DeployApiDefinition = _deployApiDefinition{
	GET_DEPLOY_ORDERS_PATH:   "/DeployApi.DeployOrders",
	POST_DEPLOY_ORDER_PATH:   "/DeployApi.DeployOrder",
	GET_OPTIONAL_IMAGES_PATH: "/DeployApi.OptionalImages",
}
