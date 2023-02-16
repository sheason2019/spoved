package k3s_service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	batch_v1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Build(ctx context.Context, co *dao.CompileOrder) error {
	job := &batch_v1.Job{
		ObjectMeta: v1.ObjectMeta{
			Name: "compile-build-co-id-" + fmt.Sprint(co.ID),
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
							Image:   co.Image,
							Command: []string{"sh", "/code/build.sh"},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "spoved-volumn",
									MountPath: "/code",
									SubPath:   "repos/" + co.Project.Creator.Username + "/" + co.Project.ProjectName + "/" + co.Version,
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
		return errors.WithStack(err)
	}
	buildCtx := JobContext{
		Context: ctx,
		Job:     job,
	}

	return buildCtx.Wait()
}
