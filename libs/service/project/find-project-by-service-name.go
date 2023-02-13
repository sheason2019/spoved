package project_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	"gorm.io/gorm"
)

func FindProjectByServiceName(ctx context.Context, serviceName string) (*dao.Project, error) {
	client := dbc.GetClient()

	proj := &dao.Project{
		ServiceName: serviceName,
	}

	err := client.WithContext(ctx).Where(proj).Find(proj).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return proj, nil
}
