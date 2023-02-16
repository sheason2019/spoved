package k3s_service

import (
	"context"
	"fmt"

	"github.com/sheason2019/spoved/libs/dao"
	batch_v1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 根据Compile工单创建一个Job Clone代码
func GitClone(ctx context.Context, co *dao.CompileOrder) error {
	job := &batch_v1.Job{
		ObjectMeta: v1.ObjectMeta{
			Name: "git-clone-co-id-" + fmt.Sprint(co.ID),
			Labels: map[string]string{
				"owner":       co.Operator.Username,
				"projectName": co.Project.ProjectName,
			},
		},
		Spec: batch_v1.JobSpec{
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
				},
			},
		},
	}

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

func cloneSh(url, branch string) string {
	return fmt.Sprintf(`
	GIT_SSH_COMMAND="ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no" git clone -b %s %s /code --progress
	`, branch, url)
}
