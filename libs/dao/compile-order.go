package dao

import (
	"fmt"

	"gorm.io/gorm"
	batch_v1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CompileOrder struct {
	gorm.Model

	// 构建时使用的镜像名称
	Image      string
	Version    string
	StatusCode int // 0表示执行中，1表示成功，-1表示失败
	Branch     string

	// 构建时所指定的环境变量
	Env map[string]string `gorm:"serializer:json"`
	// 是否为线上版本
	Production bool

	Operator   User `gorm:"foreignKey:OperatorID"`
	OperatorID int

	Project   Project `gorm:"foreignKey:ProjectID"`
	ProjectID int

	DeployOrders []DeployOrder
}

type JobOptions struct {
	AtiveDeadlineSeconds int64
	BackoffLimit         int32
	Completions          int32
}

// 默认的Job设定
var defaultJobOption = JobOptions{
	// 120s的超时时间
	AtiveDeadlineSeconds: 120,
	// 默认失败后进行2次重试
	BackoffLimit: 2,
	// 完成Job需要执行的次数
	Completions: 1,
}

// 生成拉取代码的Job
func (co *CompileOrder) GenerateGitCloneJob(option *JobOptions) *batch_v1.Job {
	if option == nil {
		option = &defaultJobOption
	}

	return &batch_v1.Job{
		ObjectMeta: v1.ObjectMeta{
			Name: "git-clone-co-id-" + fmt.Sprint(co.ID),
			Labels: map[string]string{
				"owner":       co.Operator.Username,
				"projectName": co.Project.ProjectName,
			},
		},
		Spec: batch_v1.JobSpec{
			Completions: &option.Completions,
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "git-clone-co-id-" + fmt.Sprint(co.ID),
							Image:   "bitnami/git",
							Command: []string{"/bin/bash", "-c", cloneSh(co.Project.GitUrl, co.Branch)},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "spoved-volumn",
									MountPath: "/code",
									SubPath:   "repos/" + co.Project.Creator.Username + "/" + co.Project.ProjectName + "/" + co.Version,
								},
								{
									Name:      "spoved-volumn",
									MountPath: "/root/.ssh",
									SubPath:   "account/" + co.Project.Creator.Username + "/.ssh",
								},
							},
						},
					},
					RestartPolicy: "Never",
					Volumes: []corev1.Volume{
						{
							Name: "spoved-volumn",
							VolumeSource: corev1.VolumeSource{
								PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
									ClaimName: "spoved-data-pvc",
								},
							},
						},
					},
					ActiveDeadlineSeconds: &option.AtiveDeadlineSeconds,
				},
			},
			BackoffLimit: &option.BackoffLimit,
		},
	}
}

func cloneSh(url, branch string) string {
	return fmt.Sprintf(`
	rm -rf /code
	GIT_SSH_COMMAND="ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no" git clone -b %s %s /code --progress
	`, branch, url)
}
