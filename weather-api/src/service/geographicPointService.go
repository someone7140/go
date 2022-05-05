package service

import (
	"context"
	"fmt"
	"sort"
	"weather-api/src/db/repository"
	"weather-api/src/openWeatherApi"
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
	err := repository.AddGeographicPoint(s.dbEngine, userId, r.Name, r.Lat, r.Lon, r.DisplayOrder)

	return &pb.RegsiterGeographicPointResponse{}, err

}

// 地点の更新
func (s *GeographicPointService) UpdateGeographicPoint(ctx context.Context, r *pb.UpdateGeographicPointRequest) (*pb.RegsiterGeographicPointResponse, error) {
	if ctx.Err() == context.Canceled {
		return &pb.RegsiterGeographicPointResponse{}, fmt.Errorf("client cancelled: abandoning")
	}
	// contextからユーザID取得
	userId := GetUserIdFromContext(ctx)
	err := repository.UpdateGeographicPoint(s.dbEngine, userId, r.Id, r.Name, r.Lat, r.Lon, r.DisplayOrder)

	return &pb.RegsiterGeographicPointResponse{}, err

}

// 地点の削除
func (s *GeographicPointService) DeleteGeographicPoint(ctx context.Context, r *pb.DeleteGeographicPointRequest) (*pb.RegsiterGeographicPointResponse, error) {
	if ctx.Err() == context.Canceled {
		return &pb.RegsiterGeographicPointResponse{}, fmt.Errorf("client cancelled: abandoning")
	}
	// contextからユーザID取得
	userId := GetUserIdFromContext(ctx)
	err := repository.DeleteGeographicPoint(s.dbEngine, userId, r.Id)

	return &pb.RegsiterGeographicPointResponse{}, err

}

// 地点毎の天気一覧
func (s *GeographicPointService) GetWeatherListByGeographicPoint(ctx context.Context, r *pb.GetWeatherListByGeographicPointRequest) (*pb.GetWeatherListByGeographicPointResponse, error) {
	if ctx.Err() == context.Canceled {
		return &pb.GetWeatherListByGeographicPointResponse{}, fmt.Errorf("client cancelled: abandoning")
	}
	weatherList := []*pb.WeatherByGeographicPoint{}
	// contextからユーザID取得
	userId := GetUserIdFromContext(ctx)
	// DBからユーザ登録の地点を取得
	points, err := repository.GetGeographicPointsByUserId(s.dbEngine, userId)
	if err != nil {
		return nil, err
	}
	if len(points) == 0 {
		return &pb.GetWeatherListByGeographicPointResponse{
			WeatherByGeographicPoint: weatherList,
		}, nil
	}

	// OpenWeatherAPIから地点毎の天気情報を取得
	for _, point := range points {
		result, err := openWeatherApi.GetWeatherInfoByGeographicPoint(
			point.Lat, point.Lon, point.Id, point.Name, point.DisplayOrder,
		)
		if err != nil {
			return nil, err
		}
		weatherList = append(weatherList, result...)
	}
	//　ソート（第一優先Unixtime、第二優先DisplayOrder）
	sort.Slice(weatherList, func(i, j int) bool {
		if weatherList[i].UnixTime < weatherList[j].UnixTime {
			return true
		}
		if weatherList[i].UnixTime == weatherList[j].UnixTime {
			return weatherList[i].DisplayOrder < weatherList[j].DisplayOrder
		}
		return false
	})

	return &pb.GetWeatherListByGeographicPointResponse{
		WeatherByGeographicPoint: weatherList,
	}, nil

}
