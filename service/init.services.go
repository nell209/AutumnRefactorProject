package service

import (
	"context"
	"fmt"
	"github.com/nell209/AutumnRefactor/graph/model"
)

// Init create a mock company and branch for quick testing purposes
func (s *Service) Init(ctx context.Context) error {
	hashedPassword, err := HashPassword("pa55word")
	if err != nil {
		return err
	}
	role := model.ADMIN
	user := model.User{
		FirstName: "Nelson",
		LastName:  "Test",
		Email:     "test@test.com",
		Password:  &hashedPassword,
		Role:      &role,
	}
	company := model.Company{
		Name: "My Company",
	}
	if err := s.db.Create(&company).Error; err != nil {
		return err
	}

	// create branch and assign to company
	branch := model.Branch{
		Name:      "My Branch",
		CompanyID: company.ID,
		Users:     []*model.User{&user},
	}
	if err := s.db.Create(&branch).Error; err != nil {
		return err
	}

	// create a project and assign to a branch
	project := model.Project{
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

func generateABunchOfColumns(numberOfColumns int, branchId string) []*model.Column {
	var columns []*model.Column
	for i := 0; i < numberOfColumns; i++ {
		column := model.Column{
			Name:     fmt.Sprintf("Column #%d", i),
			BranchID: branchId,
			Position: i,
		}
		columns = append(columns, &column)
	}
	return columns
}

func generateABunchOfTasks(num int, projectId string) []*model.Task {
	var tasks []*model.Task
	for i := 0; i < num; i++ {
		task := model.Task{
			Name:      fmt.Sprintf("Task %d", i),
			ProjectID: projectId,
		}
		tasks = append(tasks, &task)
	}
	return tasks
}
