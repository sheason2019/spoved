package deploy_service

import (
	"context"
	"fmt"

	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	"github.com/sheason2019/spoved/libs/idl-lib/deploy"
)

func ValidateDeployLimit(ctx context.Context, usr *dao.User, do *deploy.DeployOrder) error {
	client := dbc.DB

	proj := []dao.Project{}

	err := client.
		WithContext(ctx).
		Where("id = (?)", client.Table("compile_orders").Where("id = ?", do.CompileOrder.Id).Select("project_id")).
		Find(&proj).
		Error

	if err != nil {
		return fmt.Errorf("validate deploy limit error: find project error: %w", err)
	}

	if len(proj) == 0 {
		return fmt.Errorf("project record not found: user::%+v deployOrder::%+v", usr, do)
	}

	if proj[0].CreatorID != int(usr.ID) {
		return fmt.Errorf("validate error: post deploy order: creator id = %d, poster id = %d", proj[0].CreatorID, usr.ID)
	}

	return nil
}
