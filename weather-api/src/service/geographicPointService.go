package service

import (
	"context"
	"fmt"
	"weather-api/src/db/repository"
	"weather-api/src/pb"

	"xorm.io/xorm"
)

type GeographicPointService struct {
	dbEngine *xorm.Engine
}

// 地点の追加
func (s *GeographicPointService) AddGeographicPoint(ctx context.Context, r *pb.AddGeographicPointRequest) (*pb.RegsiterGeographicPointResponse, error) {
	if ctx.Err() == context.Canceled {
		return &pb.RegsiterGeographicPointResponse{}, fmt.Errorf("client cancelled: abandoning")
	}
	// contextからユーザID取得
	userId := GetUserIdFromContext(ctx)
	err := repository.AddGeographicPoint(s.dbEngine, userId, r.Name, r.Lat, r.Lon)

	return &pb.RegsiterGeographicPointResponse{}, err

}
