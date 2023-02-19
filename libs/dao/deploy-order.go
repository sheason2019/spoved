package dao

import (
	"gorm.io/gorm"
	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeployOrder struct {
	gorm.Model
	// 部署时使用的镜像名称
	Image      string
	StatusCode int

	// 构建时所指定的环境变量
	Env map[string]string `gorm:"serializer:json"`

	Operator   User `gorm:"foreignKey:OperatorID"`
	OperatorID int

	CompileOrder   CompileOrder `gorm:"foreignKey:CompileOrderID"`
	CompileOrderID int
}

func (do *DeployOrder) GenerateDeployment(deployName string) *appv1.Deployment {
	userName := do.CompileOrder.Project.Creator.Username
	projName := do.CompileOrder.Project.ProjectName

	sa := ""
	if userName == "root" {
		sa = "spoved-operator"
	}

	var bootCommand []string
	if do.Image != "root/spoved-nginx" {
		bootCommand = []string{"sh", "/code/start.sh"}
	}

	// 设置持久卷
	// 将拉取并编译完成的代码挂载到/code目录下
	volumeMounts := []v1.VolumeMount{
		{
			Name:      "spoved-volume",
			MountPath: "/code",
			SubPath:   "repos/" + userName + "/" + projName + "/" + do.CompileOrder.Version,
		},
	}
	// 如果是Root用户，还需要将Spoved的相关数据挂载到/data目录下
	if userName == "root" {
		volumeMounts = append(volumeMounts, v1.VolumeMount{
			Name:      "spoved-volume",
			MountPath: "/data",
		})
	}

	deployment := &appv1.Deployment{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      deployName,
			Namespace: "default",
		},
		Spec: appv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &meta_v1.LabelSelector{
				MatchLabels: map[string]string{
					"owner":       userName,
					"projectName": projName,
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: meta_v1.ObjectMeta{
					Labels: map[string]string{
						"owner":       userName,
						"projectName": projName,
					},
				},
				Spec: v1.PodSpec{
					ServiceAccountName: sa,
					Containers: []v1.Container{
						{
							Name:            projName,
							Image:           do.Image,
							ImagePullPolicy: "IfNotPresent",
							Ports: []v1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
							VolumeMounts: volumeMounts,
							Command:      bootCommand,
						},
					},
					Volumes: []v1.Volume{
						{
							Name: "spoved-volume",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "spoved-data-pvc",
								},
							},
						},
					},
				},
			},
		},
	}

	return deployment
}

func int32Ptr(i int32) *int32 { return &i }
