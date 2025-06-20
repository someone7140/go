package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.74

import (
	"context"
	"wasurena-task-api/graph/model"
	"wasurena-task-api/service"
)

// ExecuteScheduleCheckBatch is the resolver for the executeScheduleCheckBatch field.
func (r *mutationResolver) ExecuteScheduleCheckBatch(ctx context.Context, token string) (bool, error) {
	return service.CheckDailyNotify(ctx, token)
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.CategoryInput) (bool, error) {
	return service.CreateTaskCategoryService(ctx, input)
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, id string, input model.CategoryInput) (bool, error) {
	return service.UpdateTaskCategoryService(ctx, id, input)
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (bool, error) {
	return service.DeleteTaskCategoryService(ctx, id)
}

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.TaskInput) (bool, error) {
	return service.CreateTaskService(ctx, input)
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, id string, input model.TaskInput) (bool, error) {
	return service.UpdateTaskService(ctx, id, input)
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (bool, error) {
	return service.DeleteTaskDefinitionService(ctx, id)
}

// CreateTaskExecute is the resolver for the createTaskExecute field.
func (r *mutationResolver) CreateTaskExecute(ctx context.Context, input model.NewTaskExecute) (bool, error) {
	return service.CreateTaskExecuteService(ctx, input)
}

// DeleteTaskExecute is the resolver for the deleteTaskExecute field.
func (r *mutationResolver) DeleteTaskExecute(ctx context.Context, taskExecuteID string) (bool, error) {
	return service.DeleteTaskExecuteService(ctx, taskExecuteID)
}

// CreateUserAccount is the resolver for the createUserAccount field.
func (r *mutationResolver) CreateUserAccount(ctx context.Context, input model.NewUserAccount) (*model.UserAccountResponse, error) {
	return service.CreateUserAccount(ctx, input)
}

// UpdateUserAccount is the resolver for the updateUserAccount field.
func (r *mutationResolver) UpdateUserAccount(ctx context.Context, input model.UpdateUserAccountInput) (*model.UserAccountResponse, error) {
	return service.UpdateUserAccount(ctx, input)
}

// GetUserRegisterToken is the resolver for the getUserRegisterToken field.
func (r *queryResolver) GetUserRegisterToken(ctx context.Context, lineAuthCode string) (*model.CreateUserRegisterTokenResponse, error) {
	return service.GetUserRegisterTokenFromLineAuthCode(ctx, lineAuthCode)
}

// GetRegisteredUser is the resolver for the getRegisteredUser field.
func (r *queryResolver) GetRegisteredUser(ctx context.Context, lineAuthCode string) (*model.UserAccountResponse, error) {
	return service.GetUserAccountFromLineAuthCode(ctx, lineAuthCode)
}

// GetUserAccountFromAuthHeader is the resolver for the getUserAccountFromAuthHeader field.
func (r *queryResolver) GetUserAccountFromAuthHeader(ctx context.Context) (*model.UserAccountResponse, error) {
	return service.GetUserAccountFromContext(ctx)
}

// GetTaskCategories is the resolver for the getTaskCategories field.
func (r *queryResolver) GetTaskCategories(ctx context.Context) ([]*model.TaskCategoryResponse, error) {
	return service.GetTaskCategoriesService(ctx)
}

// GetTaskCategoryByID is the resolver for the getTaskCategoryById field.
func (r *queryResolver) GetTaskCategoryByID(ctx context.Context, categoryID string) (*model.TaskCategoryResponse, error) {
	return service.GetTaskCategoryByIDService(ctx, categoryID)
}

// GetTaskDefinitions is the resolver for the getTaskDefinitions field.
func (r *queryResolver) GetTaskDefinitions(ctx context.Context) ([]*model.TaskDefinitionResponse, error) {
	return service.GetTaskDefinitionService(ctx)
}

// GetTaskDefinitionByID is the resolver for the getTaskDefinitionById field.
func (r *queryResolver) GetTaskDefinitionByID(ctx context.Context, taskDefinitionID string) (*model.TaskDefinitionResponse, error) {
	return service.GetTaskDefinitionByIDService(ctx, taskDefinitionID)
}

// GetTaskCheckDisplayList is the resolver for the getTaskCheckDisplayList field.
func (r *queryResolver) GetTaskCheckDisplayList(ctx context.Context) ([]*model.TaskCheckDisplayResponse, error) {
	return service.GetTaskCheckDisplayListService(ctx)
}

// GetTaskExecuteListByDefinitionID is the resolver for the getTaskExecuteListByDefinitionId field.
func (r *queryResolver) GetTaskExecuteListByDefinitionID(ctx context.Context, taskDefinitionID string) ([]*model.TaskExecuteResponse, error) {
	return service.GetTaskExecuteListByDefinitionIDService(ctx, taskDefinitionID)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
