package database

import (
	"fmt"
	"github.com/nell209/AutumnRefactor/graph/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Branch{},
		&model.Company{},
		&model.Project{},
		&model.KanbanFilter{},
		&model.Position{},
		&model.Task{},
		//&model.TaskFilter{},
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to automigrate but got error %v", err))
	}
}
