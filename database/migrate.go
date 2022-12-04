package database

import (
	"fmt"
	"github.com/nell209/AutumnRefactor/graph/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Company{},
		&model.Branch{},
		&model.User{},
		&model.Project{},
		&model.KanbanFilter{},
		&model.Position{},
		&model.Task{},
		//&model.TaskFilter{},
		&model.TemporaryAccount{},
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to automigrate but got error %v", err))
	}
}
