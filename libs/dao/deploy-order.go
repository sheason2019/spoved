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
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "spoved-volume",
									MountPath: "/code",
									SubPath:   "repos/" + userName + "/" + projName + "/" + do.CompileOrder.Version,
								},
								{
									Name:      "spoved-volume",
									MountPath: "/data",
									SubPath:   "datas/" + userName + "/" + projName,
								},
							},
							Command: bootCommand,
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
