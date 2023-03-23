package database

import (
	"fmt"
	"github.com/nell209/AutumnRefactor/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Company{},
		&models.Branch{},
		&models.User{},
		&models.Project{},
		&models.KanbanFilter{},
		&models.Task{},
		//&model.TaskFilter{},
		&models.TemporaryAccount{},
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to automigrate but got error %v", err))
	}
}
