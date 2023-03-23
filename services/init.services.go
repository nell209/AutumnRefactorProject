package services

import (
	"context"
	"fmt"
	"github.com/nell209/AutumnRefactor/models"
)

// Init create a mock company and branch for quick testing purposes
func (s *Service) Init(ctx context.Context) error {
	hashedPassword, err := HashPassword("pa55word")
	if err != nil {
		return err
	}
	role := models.ADMIN
	user := models.User{
		FirstName: "Nelson",
		LastName:  "Test",
		Email:     "test@test.com",
		Password:  &hashedPassword,
		Role:      &role,
	}
	company := models.Company{
		Name: "My Company",
	}
	if err := s.db.Create(&company).Error; err != nil {
		return err
	}

	// create branch and assign to company
	branch := models.Branch{
		Name:      "My Branch",
		CompanyID: company.ID,
		Users:     []*models.User{&user},
	}
	if err := s.db.Create(&branch).Error; err != nil {
		return err
	}

	// create a project and assign to a branch
	project := models.Project{
		Name:     "My Project",
		BranchID: branch.ID,
	}
	if err := s.db.Create(&project).Error; err != nil {
		return err
	}

	tasks := generateABunchOfTasks(15, project.ID)
	if err := s.db.Create(&tasks).Error; err != nil {
		return err
	}

	columns := generateABunchOfColumns(4, branch.ID)
	if err := s.db.Create(&columns).Error; err != nil {
		return err
	}

	// todo assign the tasks to columns

	//const branchId = "87d4ee3c-1dbf-4941-a107-f956227493d9"
	//hashedPassword, err := HashPassword("pa55word")
	//if err != nil {
	//	return err
	//}
	//
	//tasks := generateABunchOfTasks(15)
	//project := model.Project{
	//	Name:     "Test Project",
	//	BranchID: branchId,
	//	Tasks:    tasks,
	//}
	//role := model.ADMIN
	//user := model.User{
	//	FirstName: "Nelson",
	//	LastName:  "Test",
	//	Email:     "test@test.com",
	//	Password:  &hashedPassword,
	//	Role:      &role,
	//}
	//
	//columns := generateABunchOfColumns(5, branchId)
	//branch := model.Branch{
	//	Name:     "My Branch",
	//	Users:    []*model.User{&user},
	//	Projects: []*model.Project{&project},
	//	Columns:  columns,
	//}
	//
	//if err := s.db.Create(&company).Error; err != nil {
	//	return err
	//}
	//// todo later send the tasks to columns
	return nil
}

func generateABunchOfColumns(numberOfColumns int, branchId string) []*models.Column {
	var columns []*models.Column
	for i := 0; i < numberOfColumns; i++ {
		column := models.Column{
			Name:     fmt.Sprintf("Column #%d", i),
			BranchID: branchId,
			Position: i,
		}
		columns = append(columns, &column)
	}
	return columns
}

func generateABunchOfTasks(num int, projectId string) []*models.Task {
	var tasks []*models.Task
	for i := 0; i < num; i++ {
		task := models.Task{
			Name:      fmt.Sprintf("Task %d", i),
			ProjectID: &projectId,
		}
		tasks = append(tasks, &task)
	}
	return tasks
}
