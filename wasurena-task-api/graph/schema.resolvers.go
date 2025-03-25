package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.68

import (
	"context"
	"fmt"
	"wasurena-task-api/graph/model"
	"wasurena-task-api/service"
)

// ExecuteScheduleCheckBatch is the resolver for the executeScheduleCheckBatch field.
func (r *mutationResolver) ExecuteScheduleCheckBatch(ctx context.Context, token string) (bool, error) {
	return service.CheckDailyNotify(ctx, token)
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (bool, error) {
	return service.CreateCategoryService(ctx, input)
}

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (bool, error) {
	return service.CreateTaskService(ctx, input)
}

// CreateTaskExecute is the resolver for the createTaskExecute field.
func (r *mutationResolver) CreateTaskExecute(ctx context.Context, input model.NewTaskExecute) (bool, error) {
	return service.CreateTaskExecuteService(ctx, input)
}

// CreateUserAccount is the resolver for the createUserAccount field.
func (r *mutationResolver) CreateUserAccount(ctx context.Context, input model.NewUserAccount) (*model.UserAccountResponse, error) {
	return service.CreateUserAccount(ctx, input)
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// GetUserRegisterToken is the resolver for the getUserRegisterToken field.
func (r *queryResolver) GetUserRegisterToken(ctx context.Context, lineAuthCode string) (*model.CreateUserRegisterTokenResponse, error) {
	return service.GetUserRegisterToken(ctx, lineAuthCode)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
