package compile_service

import (
	"context"
	"fmt"

	"github.com/sheason2019/spoved/libs/dao"
)

// 检验用户是否有创建编译工单的权限
func ValidateOperator(ctx context.Context, proj *dao.Project, operator *dao.User) error {
	// 目前仅项目创建者拥有对项目的操作权限
	if proj.CreatorID != int(operator.ID) {
		return fmt.Errorf("缺少创建编译工单的权限")
	}
	return nil
}
