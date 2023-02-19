package k3s_service

import (
	"context"

	"github.com/sheason2019/spoved/libs/dao"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 根据Compile工单创建一个Job Clone代码
func GitClone(ctx context.Context, co *dao.CompileOrder) error {
	job := co.GenerateGitCloneJob(nil)

	job, err := clientSet.BatchV1().Jobs("default").Create(ctx, job, v1.CreateOptions{})
	if err != nil {
		return err
	}

	gcCtx := JobContext{
		Context: ctx,
		Job:     job,
	}

	return gcCtx.Wait()
}
