package postUseCase

import (
	placeNote "placeNote/src/gen/proto"
	modelDb "placeNote/src/model/db"
	"placeNote/src/placeNoteUtil"
	"placeNote/src/repository"

	"github.com/bufbuild/connect-go"
)

// AddPostPlace 場所の追加
func AddPostPlace(req *placeNote.AddPostPlaceRequest, userAccountId string) (*string, *connect.Error) {
	uid, err := placeNoteUtil.GenerateUUID()
	if err != nil {
		return nil, err
	}
	// 登録処理
	err = repository.AddPostPlaceRepository(modelDb.PostPlacesEntity{
		ID:                  uid,
		Name:                req.Name,
		CreateUserAccountId: userAccountId,
		LonLat:              getLatLonForEntity(req.LatLon),
		PrefectureCode:      req.PrefectureCode,
		CategoryIdList:      &req.CategoryIdList,
		Detail:              req.Detail,
		UrlList:             &req.UrlList,
	})
	if err != nil {
		return nil, err
	}

	return &uid, err
}

// UpdatePostPlace 場所の追加
func UpdatePostPlace(req *placeNote.UpdatePostPlaceRequest, userAccountId string) *connect.Error {
	// 更新処理
	err := repository.UpdatePostPlaceRepository(modelDb.PostPlacesEntity{
		ID:                  req.Id,
		Name:                req.Name,
		CreateUserAccountId: userAccountId,
		LonLat:              getLatLonForEntity(req.LatLon),
		PrefectureCode:      req.PrefectureCode,
		CategoryIdList:      &req.CategoryIdList,
		Detail:              req.Detail,
		UrlList:             &req.UrlList,
	})
	if err != nil {
		return err
	}

	return err
}

// DeletePostPlace 場所の削除
func DeletePostPlace(req *placeNote.DeletePostPlaceRequest, userAccountId string) *connect.Error {
	// 削除処理
	err := repository.DeletePostPlaceRepository(req.Id, userAccountId)
	if err != nil {
		return err
	}

	return err
}

// GetPostPlaceList ユーザーの場所一覧を取得
func GetPostPlaceList(userAccountId string) (*connect.Response[placeNote.GetPostPlaceListResponse], *connect.Error) {

	placeEntities, err := repository.GetUserPostPlaceListByUserAccountIdRepository(userAccountId)
	if err != nil {
		return nil, err
	}
	var places []*placeNote.PostPlaceResponse

	for _, entity := range placeEntities {
		places = append(places, &placeNote.PostPlaceResponse{
			Id:             entity.ID,
			Name:           entity.Name,
			Address:        entity.Address,
			LatLon:         getLatLonForResponse(entity.LonLat),
			PrefectureCode: entity.PrefectureCode,
			CategoryIdList: *entity.CategoryIdList,
			Detail:         entity.Detail,
			UrlList:        *entity.UrlList,
		})
	}

	return connect.NewResponse(&placeNote.GetPostPlaceListResponse{
		PlaceList: places,
	}), nil
}

// GetPostPlaceById ID指定でユーザーの場所を取得
func GetPostPlaceById(userAccountId string, placeId string) (*connect.Response[placeNote.PostPlaceResponse], *connect.Error) {

	entity, err := repository.GetUserPostPlaceByIdRepository(userAccountId, placeId)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&placeNote.PostPlaceResponse{
		Id:             entity.ID,
		Name:           entity.Name,
		Address:        entity.Address,
		LatLon:         getLatLonForResponse(entity.LonLat),
		PrefectureCode: entity.PrefectureCode,
		CategoryIdList: *entity.CategoryIdList,
		Detail:         entity.Detail,
		UrlList:        *entity.UrlList,
	}), nil
}

// getLatLonForEntity DB保存用の緯度経度を取得
func getLatLonForEntity(latLon *placeNote.LatLon) *[]float64 {
	lonLat := []float64{}
	if latLon != nil {
		lonLat = []float64{latLon.Lon, latLon.Lat}
	}
	return &lonLat
}

// getLatLonForResponse レスポンスの緯度経度を取得
func getLatLonForResponse(lonLat *[]float64) *placeNote.LatLon {
	if lonLat != nil && len(*lonLat) > 1 {
		return &placeNote.LatLon{
			Lat: (*lonLat)[0],
			Lon: (*lonLat)[1],
		}
	}
	return nil
}
