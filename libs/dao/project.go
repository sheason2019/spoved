package dao

import (
	"fmt"

	"gorm.io/gorm"
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// User holds the schema definition for the User entity.
type Project struct {
	gorm.Model
	ProjectName string `gorm:"index"`
	Describe    string
	GitUrl      string
	ServiceName string

	CompileOrders []CompileOrder `gorm:"foreignKey:ProjectID"`
	Creator       User           `gorm:"foreignKey:CreatorID"`
	CreatorID     int
}

func (p *Project) DirPath() string {
	return fmt.Sprintf("/%s/%s", p.Creator.Username, p.ProjectName)
}

func (p *Project) GenerateService(svcName string) *v1.Service {
	svc := v1.Service{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      svcName,
			Namespace: "default",
			Labels: map[string]string{
				"owner":       p.Creator.Username,
				"projectName": p.ProjectName,
			},
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Port:     80,
					Protocol: "TCP",
				},
			},
			Selector: map[string]string{
				"owner":       p.Creator.Username,
				"projectName": p.ProjectName,
			},
		},
	}

	return &svc
}
