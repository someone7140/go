package openWeatherApi

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"weather-api/src/pb"
	"weather-api/src/util"

	jsoniter "github.com/json-iterator/go"
)

func GetWeatherInfoByGeographicPoint(lat float64, lon float64, pointId string, pointName string, displayOrder int32) ([]*pb.WeatherByGeographicPoint, error) {
	requestUrl := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/forecast?lat=%s&lon=%s&appid=%s&units=metric",
		strconv.FormatFloat(lat, 'f', -1, 64),
		strconv.FormatFloat(lon, 'f', -1, 64),
		os.Getenv("OPEN_WEATHER_API_KEY"),
	)
	weatherResultStr := util.SendGetHTTPRequest(requestUrl)
	if weatherResultStr == "" {
		return nil, errors.New("Failed get openWeatherApi")
	}

	return getWeatherResponseStrJSONConvert(weatherResultStr, lat, lon, pointId, pointName, displayOrder)
}

func getWeatherResponseStrJSONConvert(jsonStr string, lat float64, lon float64, pointId string, pointName string, displayOrder int32) ([]*pb.WeatherByGeographicPoint, error) {

	weatherList := []*pb.WeatherByGeographicPoint{}
	// Json文字列からオブジェクトに変換
	weatherJson := jsoniter.Get([]byte(jsonStr))
	// list項目を取得
	weatherJsonList := weatherJson.Get("list")

	for i := 0; i < weatherJsonList.Size(); i++ {
		response := weatherJsonList.Get(i)

		weatherInfo := &pb.WeatherByGeographicPoint{}

		weatherInfo.PointId = pointId
		weatherInfo.PointName = pointName
		weatherInfo.Lat = lat
		weatherInfo.Lon = lon
		weatherInfo.DisplayOrder = displayOrder
		weatherInfo.UnixTime = response.Get("dt").ToInt64()
		weatherInfo.WeatherIcon = response.Get("weather").Get(0).Get("icon").ToString()
		weatherInfo.TempFeelsLike = response.Get("main").Get("feels_like").ToFloat64()
		weatherInfo.TempMin = response.Get("main").Get("temp_min").ToFloat64()
		weatherInfo.TempMax = response.Get("main").Get("temp_max").ToFloat64()
		weatherInfo.Clouds = response.Get("clouds").Get("all").ToFloat64()
		// 降水量の項目存在判定のため、まずはstringで取得
		rainFallStr := response.Get("rain").Get("3h").ToString()
		if rainFallStr != "" {
			weatherInfo.RainFall = response.Get("rain").Get("3h").ToFloat64()
		}
		weatherInfo.Humidity = response.Get("main").Get("humidity").ToFloat64()
		weatherInfo.WindSpeed = response.Get("wind").Get("speed").ToFloat64()
		weatherInfo.Pressure = response.Get("main").Get("pressure").ToFloat64()

		weatherList = append(weatherList, weatherInfo)
	}
	return weatherList, nil
}
