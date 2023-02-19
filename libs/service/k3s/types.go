package k3s_service

import (
	"context"
	"fmt"
	"time"

	batch_v1 "k8s.io/api/batch/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type JobContext struct {
	context.Context
	Job *batch_v1.Job
}

// 等待Job执行完成
func (ctx *JobContext) Wait() error {
	for {
		time.Sleep(2 * time.Second)

		job, err := clientSet.BatchV1().Jobs("default").Get(ctx, ctx.Job.Name, v1.GetOptions{})
		if err != nil {
			return err
		}

		ctx.Job = job

		fmt.Printf("Status %+v\n", job.Status)

		// 成功数达到预期值时结束监听
		if job.Status.Succeeded >= *job.Spec.Completions {
			break
		}
		// 超出失败重试次数时结束监听
		if job.Status.Failed > *job.Spec.BackoffLimit {
			break
		}
	}

	if ctx.Job.Status.Succeeded == *ctx.Job.Spec.Completions {
		return nil
	} else {
		return fmt.Errorf(
			"error: Job Completions: %d, Job Succeeded: %d, Job Failed: %d",
			*ctx.Job.Spec.Completions,
			ctx.Job.Status.Succeeded,
			ctx.Job.Status.Failed,
		)
	}
}
