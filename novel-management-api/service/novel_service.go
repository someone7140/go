package service

import (
	"context"
	"main/db/db_model"
	"main/graph/graphql_model"
	"main/repository"
	"time"

	"github.com/google/uuid"
)

func AddNovel(userAccountID string, title string, description *string) (*bool, error) {
	ctx := context.Background()
	uid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	novelDbModel := db_model.Novel{
		ID:                 uid.String(),
		Title:              title,
		Description:        description,
		OwnerUserAccountID: userAccountID,
		CreatedAt:          time.Now(),
	}
	err = repository.AddNovel(ctx, novelDbModel)
	if err != nil {
		return nil, err
	}
	result := true

	return &result, nil
}

func EditNovel(userAccountID string, novelID string, title string, description *string) (*bool, error) {
	ctx := context.Background()
	err := repository.EditNovel(ctx, novelID, userAccountID, title, description)
	if err != nil {
		return nil, err
	}
	result := true

	return &result, nil
}

func DeleteNovel(userAccountID string, novelID string) (*bool, error) {
	ctx := context.Background()
	err := repository.DeleteNovel(ctx, userAccountID, novelID)
	if err != nil {
		return nil, err
	}
	result := true

	return &result, nil
}

func GetMyNovels(userAccountID string) ([]graphql_model.NovelResponse, error) {
	ctx := context.Background()
	novels, err := repository.GetMyNovels(ctx, userAccountID)
	novelResponse := []graphql_model.NovelResponse{}
	if err != nil {
		return novelResponse, err
	}

	for _, n := range novels {
		novelResponse = append(novelResponse, graphql_model.NovelResponse{
			ID:          n.ID,
			Title:       n.Title,
			Description: n.Description,
		})
	}
	return novelResponse, nil
}

func GetMyNovelByID(userAccountID string, novelID string) (*graphql_model.NovelResponse, error) {
	ctx := context.Background()
	novel, err := repository.GetMyNovelByID(ctx, userAccountID, novelID)
	if err != nil {
		return nil, err
	}

	return &graphql_model.NovelResponse{
		ID:          novel.ID,
		Title:       novel.Title,
		Description: novel.Description,
	}, nil
}
