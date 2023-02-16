package k3s_service

import (
	"context"
	"fmt"
	"time"

	appv1 "k8s.io/api/apps/v1"
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
		job, err := clientSet.BatchV1().Jobs("default").Get(ctx, ctx.Job.Name, v1.GetOptions{})
		if err != nil {
			return err
		}

		fmt.Println("job", job)
		ctx.Job = job

		time.Sleep(time.Second)
	}
}

type DeployContext struct {
	context.Context
	Deployment *appv1.Deployment
}

// 等待Job执行完成
func (ctx *DeployContext) Wait() error {
	for {
		deploy, err := clientSet.AppsV1().Deployments("defualt").Get(ctx, ctx.Deployment.Name, v1.GetOptions{})
		if err != nil {
			return err
		}

		fmt.Println("deploy", deploy)
		ctx.Deployment = deploy

		time.Sleep(time.Second)
	}
}
