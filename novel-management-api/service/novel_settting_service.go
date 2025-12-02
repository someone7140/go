package service

import (
	"context"
	"main/db/db_model"
	"main/graph/graphql_model"
	"main/repository"
	"slices"

	"github.com/google/uuid"
)

func RegisterNovelSettings(userAccountID string, inputs []graphql_model.NovelSettingRegisterInput) (*bool, error) {
	ctx := context.Background()
	result := true
	var registerSettings []*db_model.NovelSetting
	if len(inputs) == 0 {
		return &result, nil
	}

	settingsFromDB, err := repository.GetMyNovelSettings(ctx, userAccountID)
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
			isContains := slices.ContainsFunc(settingsFromDB, func(s *db_model.NovelSetting) bool {
				return s.ID == id
			})
			if !isContains {
				continue
			}
		}

		registerSettings = append(registerSettings, &db_model.NovelSetting{
			ID:                 id,
			Name:               input.Name,
			NovelID:            input.NovelID,
			OwnerUserAccountID: userAccountID,
			ParentSettingID:    input.ParentSettingID,
			DisplayOrder:       input.DisplayOrder,
			Attributes:         input.Attributes,
			Description:        input.Description,
		})
	}

	err = repository.RegisterNovelSettings(ctx, registerSettings)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetNovelSettingsByNovelID(userAccountID string, novelID string) ([]graphql_model.NovelSettingResponse, error) {
	ctx := context.Background()
	settingsFromDB, err := repository.GetNovelSettingsByNovelID(ctx, userAccountID, novelID)
	if err != nil {
		return nil, err
	}

	var results []graphql_model.NovelSettingResponse
	for _, settings := range settingsFromDB {
		results = append(results, graphql_model.NovelSettingResponse{
			ID:              settings.ID,
			Name:            settings.Name,
			NovelID:         settings.NovelID,
			ParentSettingID: settings.ParentSettingID,
			DisplayOrder:    settings.DisplayOrder,
			Attributes:      settings.Attributes,
			Description:     settings.Description,
		})
	}
	return results, err
}

func GetNovelSettingsByParentID(userAccountID string, parentID string) ([]graphql_model.NovelSettingResponse, error) {
	ctx := context.Background()
	settingsFromDB, err := repository.GetNovelSettingsByParentID(ctx, userAccountID, parentID)
	if err != nil {
		return nil, err
	}

	var results []graphql_model.NovelSettingResponse
	for _, settings := range settingsFromDB {
		results = append(results, graphql_model.NovelSettingResponse{
			ID:              settings.ID,
			Name:            settings.Name,
			NovelID:         settings.NovelID,
			ParentSettingID: settings.ParentSettingID,
			DisplayOrder:    settings.DisplayOrder,
			Attributes:      settings.Attributes,
			Description:     settings.Description,
		})
	}
	return results, err
}

func DeleteNovelSettingByID(userAccountID string, id string) (*bool, error) {
	ctx := context.Background()
	result := true

	// 該当のidを親とする設定を削除する
	children, err := repository.GetNovelSettingsByParentID(ctx, userAccountID, id)
	if err != nil {
		return nil, err
	}
	if len(children) > 0 {
		var ids []string
		for _, child := range children {
			ids = append(ids, child.ID)
		}
		err = repository.DeleteNovelSettings(ctx, userAccountID, ids)
		if err != nil {
			return nil, err
		}
	}

	err = repository.DeleteNovelSettings(ctx, userAccountID, []string{id})
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func DeleteNovelSettingByIDs(userAccountID string, ids []string) (*bool, error) {
	ctx := context.Background()
	result := true
	err := repository.DeleteNovelSettings(ctx, userAccountID, ids)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
