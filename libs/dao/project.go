package dao

import (
	"fmt"

	"gorm.io/gorm"
)

// User holds the schema definition for the User entity.
type Project struct {
	gorm.Model
	ProjectName string `gorm:"index"`
	Describe    string
	GitUrl      string

	CompileOrders []CompileOrder `gorm:"foreignKey:ProjectID"`
	Creator       User           `gorm:"foreignKey:CreatorID"`
	CreatorID     int
}

func (p *Project) DirPath() string {
	return fmt.Sprintf("/%s/%s", p.Creator.Username, p.ProjectName)
}
