package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nell209/AutumnRefactor/graph/generated"
	"github.com/nell209/AutumnRefactor/graph/model"
)

// Init is the resolver for the Init field.
func (r *mutationResolver) Init(ctx context.Context) (*model.Response, error) {
	if err := r.service.Init(ctx); err != nil {
		return nil, err
	}
	return &model.Response{
		Code:    200,
		Message: "Success!",
	}, nil
}

// AddPrerequisitesToTask is the resolver for the AddPrerequisitesToTask field.
func (r *mutationResolver) AddPrerequisitesToTask(ctx context.Context, prerequisiteIDs []string, taskUID string, typeArg string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// AddUsersToTask is the resolver for the AddUsersToTask field.
func (r *mutationResolver) AddUsersToTask(ctx context.Context, taskUID string, users []*string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// ArchiveProject is the resolver for the ArchiveProject field.
func (r *mutationResolver) ArchiveProject(ctx context.Context, projectUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// CancelTask is the resolver for the CancelTask field.
func (r *mutationResolver) CancelTask(ctx context.Context, taskUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// CompleteTask is the resolver for the CompleteTask field.
func (r *mutationResolver) CompleteTask(ctx context.Context, taskUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateBranch is the resolver for the CreateBranch field.
func (r *mutationResolver) CreateBranch(ctx context.Context, branch model.BranchInput, companyUID string) (*model.Branch, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateCompany is the resolver for the CreateCompany field.
func (r *mutationResolver) CreateCompany(ctx context.Context, company model.CompanyInput) (*model.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateProject is the resolver for the CreateProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, branchUID string, project model.ProjectInput) (*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateTask is the resolver for the CreateTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, branchUID string, task model.TaskInput) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateTaskPriority is the resolver for the CreateTaskPriority field.
func (r *mutationResolver) CreateTaskPriority(ctx context.Context, branchUID string, taskPriority model.ColumnInput) (*model.Column, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateUserWithTempPass is the resolver for the CreateUserWithTempPass field.
func (r *mutationResolver) CreateUserWithTempPass(ctx context.Context, branchID string, branchRole string, user model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteBranch is the resolver for the DeleteBranch field.
func (r *mutationResolver) DeleteBranch(ctx context.Context, branchUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteTask is the resolver for the DeleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, taskUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteTaskFilter is the resolver for the DeleteTaskFilter field.
func (r *mutationResolver) DeleteTaskFilter(ctx context.Context, taskFilterUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteTaskPriority is the resolver for the DeleteTaskPriority field.
func (r *mutationResolver) DeleteTaskPriority(ctx context.Context, taskPriorityUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteTasks is the resolver for the DeleteTasks field.
func (r *mutationResolver) DeleteTasks(ctx context.Context, taskUids []string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteTimeTracker is the resolver for the DeleteTimeTracker field.
func (r *mutationResolver) DeleteTimeTracker(ctx context.Context, uid string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// ForceCompleteTask is the resolver for the ForceCompleteTask field.
func (r *mutationResolver) ForceCompleteTask(ctx context.Context, taskUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// GenerateTemporarySignUp is the resolver for the GenerateTemporarySignUp field.
func (r *mutationResolver) GenerateTemporarySignUp(ctx context.Context, newEmail string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// HoistPublished is the resolver for the HoistPublished field.
func (r *mutationResolver) HoistPublished(ctx context.Context, columnUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// Move is the resolver for the Move field.
func (r *mutationResolver) Move(ctx context.Context, afterUID *string, columnUID string, toBeMovedUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoveTaskPriority is the resolver for the MoveTaskPriority field.
func (r *mutationResolver) MoveTaskPriority(ctx context.Context, position int, taskPriorityUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoveToTheBottomOfAColumn is the resolver for the MoveToTheBottomOfAColumn field.
func (r *mutationResolver) MoveToTheBottomOfAColumn(ctx context.Context, columnUID string, toBeMovedUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// PublishProject is the resolver for the PublishProject field.
func (r *mutationResolver) PublishProject(ctx context.Context, projectUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// RemovePrerequisitesFromTask is the resolver for the RemovePrerequisitesFromTask field.
func (r *mutationResolver) RemovePrerequisitesFromTask(ctx context.Context, prerequisiteIDs []string, taskUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// RemoveUser is the resolver for the RemoveUser field.
func (r *mutationResolver) RemoveUser(ctx context.Context, branchUID string, userUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// RemoveUsersFromTask is the resolver for the RemoveUsersFromTask field.
func (r *mutationResolver) RemoveUsersFromTask(ctx context.Context, taskUID string, users []*string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// SetTaskDate is the resolver for the SetTaskDate field.
func (r *mutationResolver) SetTaskDate(ctx context.Context, date *string, taskUID string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// SetTasksStatus is the resolver for the SetTasksStatus field.
func (r *mutationResolver) SetTasksStatus(ctx context.Context, status string, taskIDs []string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// StartTask is the resolver for the StartTask field.
func (r *mutationResolver) StartTask(ctx context.Context, taskUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// UnCompleteTask is the resolver for the UnCompleteTask field.
func (r *mutationResolver) UnCompleteTask(ctx context.Context, taskUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateAllBranchTaskPriorities is the resolver for the UpdateAllBranchTaskPriorities field.
func (r *mutationResolver) UpdateAllBranchTaskPriorities(ctx context.Context, taskPriorities []*model.ColumnUpdate) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateBranch is the resolver for the UpdateBranch field.
func (r *mutationResolver) UpdateBranch(ctx context.Context, branch model.BranchInput, branchUID string) (*model.Branch, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateDefaultKanbanForBranch is the resolver for the UpdateDefaultKanbanForBranch field.
func (r *mutationResolver) UpdateDefaultKanbanForBranch(ctx context.Context, branchUID string, defaultKanban *string) (*model.Branch, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateKanbanSettings is the resolver for the UpdateKanbanSettings field.
func (r *mutationResolver) UpdateKanbanSettings(ctx context.Context, branchUID string, kelownaMode *bool, prairieMode *bool) (*model.Branch, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateUser is the resolver for the UpdateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, user model.UserUpdate, userUID string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateProject is the resolver for the UpdateProject field.
func (r *mutationResolver) UpdateProject(ctx context.Context, project model.ProjectInput, projectUID string, users []*string) (*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateTask is the resolver for the UpdateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, task model.TaskUpdate, taskUID string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateTaskUrgency is the resolver for the UpdateTaskUrgency field.
func (r *mutationResolver) UpdateTaskUrgency(ctx context.Context, taskUID string, urgency bool) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateUserPassword is the resolver for the UpdateUserPassword field.
func (r *mutationResolver) UpdateUserPassword(ctx context.Context, password string, userUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateKanbanFilter is the resolver for the CreateKanbanFilter field.
func (r *mutationResolver) CreateKanbanFilter(ctx context.Context, branchUID string, columns []string, filterName string, completedMode *string) (*model.KanbanFilter, error) {
	panic(fmt.Errorf("not implemented"))
}

// SetDefaultUserFilter is the resolver for the SetDefaultUserFilter field.
func (r *mutationResolver) SetDefaultUserFilter(ctx context.Context, branchUID string, filterUID string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// Branches is the resolver for the Branches field.
func (r *queryResolver) Branches(ctx context.Context) ([]*model.Branch, error) {
	panic(fmt.Errorf("not implemented"))
}

// BranchTaskFilters is the resolver for the BranchTaskFilters field.
func (r *queryResolver) BranchTaskFilters(ctx context.Context, branchUID string) ([]*model.TaskFilter, error) {
	panic(fmt.Errorf("not implemented"))
}

// BranchTaskPriorities is the resolver for the BranchTaskPriorities field.
func (r *queryResolver) BranchTaskPriorities(ctx context.Context, branchUID string) ([]*model.Column, error) {
	panic(fmt.Errorf("not implemented"))
}

// CompanyBranches is the resolver for the CompanyBranches field.
func (r *queryResolver) CompanyBranches(ctx context.Context, companyUID string) ([]*model.Branch, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetBranchProjects is the resolver for the GetBranchProjects field.
func (r *queryResolver) GetBranchProjects(ctx context.Context, branchUID string, from *string, next *float64, options *model.ProjectSearchOptions) (*model.ProjectConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetBranchTasks is the resolver for the GetBranchTasks field.
func (r *queryResolver) GetBranchTasks(ctx context.Context, branchUID string, filterUID *string, completedMode *string) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetBranchUsers is the resolver for the GetBranchUsers field.
func (r *queryResolver) GetBranchUsers(ctx context.Context, branchID string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetColumnTasks is the resolver for the GetColumnTasks field.
func (r *queryResolver) GetColumnTasks(ctx context.Context, branchUID string, columnUID *string, from *string, next *float64, options model.TaskSearchOptions) (*model.TasksConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetCompanyByUID is the resolver for the GetCompanyByUid field.
func (r *queryResolver) GetCompanyByUID(ctx context.Context, companyUID string) (*model.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetProjectByID is the resolver for the GetProjectByID field.
func (r *queryResolver) GetProjectByID(ctx context.Context, projectUID string) (*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetProjectTasks is the resolver for the GetProjectTasks field.
func (r *queryResolver) GetProjectTasks(ctx context.Context, projectUID string) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetTaskByID is the resolver for the GetTaskByID field.
func (r *queryResolver) GetTaskByID(ctx context.Context, taskUID string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetUnPositionedTasks is the resolver for the GetUnPositionedTasks field.
func (r *queryResolver) GetUnPositionedTasks(ctx context.Context, branchUID string, columnUID string) (*model.TasksConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Me is the resolver for the Me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetKanbanFilters is the resolver for the GetKanbanFilters field.
func (r *queryResolver) GetKanbanFilters(ctx context.Context, branchUID string) ([]*model.KanbanFilter, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetKanbanFilterByUID is the resolver for the GetKanbanFilterByUid field.
func (r *queryResolver) GetKanbanFilterByUID(ctx context.Context, filterUID string) (*model.KanbanFilter, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateUser(ctx context.Context, branchID string, branchRole string, user model.UserInput) (*model.User, error) {
	// ensure user is Authorized to do this
	panic(fmt.Errorf("not implemented"))
}
