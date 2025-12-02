package service

import (
	"context"
	"main/db/db_model"
	"main/graph/graphql_model"
	"main/repository"
	"slices"

	"github.com/google/uuid"
)

func RegisterNovelContents(userAccountID string, inputs []graphql_model.NovelContentsRegisterInput) (*bool, error) {
	ctx := context.Background()
	result := true
	var registerContents []*db_model.NovelContent
	if len(inputs) == 0 {
		return &result, nil
	}

	contentsFromDB, err := repository.GetMyNovelContents(ctx, userAccountID)
	if err != nil {
		return nil, err
	}

	for _, input := range inputs {
		id := ""
		if input.ID == nil {
			uid, err := uuid.NewV7()
			if err != nil {
				return nil, err
			}
			id = uid.String()
		} else {
			id = *input.ID
			isContains := slices.ContainsFunc(contentsFromDB, func(c *db_model.NovelContent) bool {
				return c.ID == id
			})
			if !isContains {
				continue
			}
		}

		registerContents = append(registerContents, &db_model.NovelContent{
			ID:                 id,
			ChapterName:        input.ChapterName,
			NovelID:            input.NovelID,
			OwnerUserAccountID: userAccountID,
			ParentContentsID:   input.ParentContentsID,
			DisplayOrder:       input.DisplayOrder,
			Contents:           input.Contents,
			Description:        input.Description,
		})
	}

	err = repository.RegisterNovelContents(ctx, registerContents)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetNovelContentsByNovelID(userAccountID string, novelID string) ([]graphql_model.NovelContentsResponse, error) {
	ctx := context.Background()
	contentsFromDB, err := repository.GetNovelContentsByNovelID(ctx, userAccountID, novelID)
	if err != nil {
		return nil, err
	}

	var results []graphql_model.NovelContentsResponse
	for _, content := range contentsFromDB {
		results = append(results, graphql_model.NovelContentsResponse{
			ID:               content.ID,
			ChapterName:      content.ChapterName,
			NovelID:          content.NovelID,
			ParentContentsID: content.ParentContentsID,
			DisplayOrder:     content.DisplayOrder,
			Contents:         content.Contents,
			Description:      content.Description,
		})
	}
	return results, err
}

func DeleteNovelContentsByID(userAccountID string, id string) (*bool, error) {
	ctx := context.Background()
	result := true

	// 該当のidを親とする設定を削除する
	children, err := repository.GetNovelContentsByParentID(ctx, userAccountID, id)
	if err != nil {
		return nil, err
	}
	if len(children) > 0 {
		var ids []string
		for _, child := range children {
			ids = append(ids, child.ID)
		}
		err = repository.DeleteNovelContents(ctx, userAccountID, ids)
		if err != nil {
			return nil, err
		}
	}

	err = repository.DeleteNovelContents(ctx, userAccountID, []string{id})
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func DeleteNovelContentsByIDs(userAccountID string, ids []string) (*bool, error) {
	ctx := context.Background()
	result := true
	err := repository.DeleteNovelContents(ctx, userAccountID, ids)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
