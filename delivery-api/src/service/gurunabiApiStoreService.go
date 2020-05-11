package service

import (
	"encoding/json"
	"os"
	"sort"
	"strconv"

	"deliverly/src/model"

	"github.com/thoas/go-funk"
)

// GetStoreSearchURL ぐるなび店舗取得API送信用URL
func GetStoreSearchURL(request model.StoreInfoRequest) string {
	if request.Latitude == 0 || request.Longitude == 0 {
		return ""
	}
	url := "https://api.gnavi.co.jp/RestSearchAPI/v3/" +
		"?keyid=" + os.Getenv("GURUNABI_API_KEY") +
		"&latitude=" + strconv.FormatFloat(request.Latitude, 'f', -1, 64) +
		"&longitude=" + strconv.FormatFloat(request.Longitude, 'f', -1, 64) +
		"&hit_per_page=100"
	if request.Range != 0 {
		url += "&range=" + strconv.FormatInt(request.Range, 10)
	}
	if request.CategoryL != "" {
		url += "&category_l=" + request.CategoryL
	}
	return url
}

// ResponseStoreJSONConvert ぐるなびの店情報レスポンスをStructの配列にうつす
func ResponseStoreJSONConvert(jsonStr string, provideType string) model.StoreInfos {
	var storeInfoList model.StoreInfos
	if jsonStr != "" {
		var responseMap map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &responseMap)
		if err != nil {
			return storeInfoList
		}
		if responseMap["rest"] != nil {
			restList := responseMap["rest"].([]interface{})
			for _, r := range restList {
				var s model.StoreInfo
				s.ID = r.(map[string]interface{})["id"].(string)
				s.Name = r.(map[string]interface{})["name"].(string)
				s.Latitude, err = strconv.ParseFloat(r.(map[string]interface{})["latitude"].(string), 32)
				if err != nil {
					continue
				}
				s.Longitude, err = strconv.ParseFloat(r.(map[string]interface{})["longitude"].(string), 32)
				if err != nil {
					continue
				}
				s.Category = r.(map[string]interface{})["category"].(string)
				s.URL = r.(map[string]interface{})["url"].(string)
				s.Image = r.(map[string]interface{})["image_url"].(map[string]interface{})["shop_image1"].(string)
				s.Opentime = r.(map[string]interface{})["opentime"].(string)
				s.Holiday = r.(map[string]interface{})["holiday"].(string)
				s.Pr = r.(map[string]interface{})["pr"].(map[string]interface{})["pr_short"].(string)
				s.Type = provideType
				storeInfoList = append(storeInfoList, s)
			}
		}
	}
	return storeInfoList
}

// MergeStoreList デリバリー可リストとテイクアウト可リストをマージ
func MergeStoreList(
	delivelyStoreList model.StoreInfos,
	takeoutStoreList model.StoreInfos) model.StoreInfos {
	// デリバリーのみ
	delivelyOnlyList := model.StoreInfos(funk.Filter(delivelyStoreList, func(d model.StoreInfo) bool {
		return funk.Find(takeoutStoreList, func(t model.StoreInfo) bool {
			return d.ID == t.ID
		}) == nil
	}).([]model.StoreInfo))
	// テイクアウトのみ
	takeoutOnlyList := model.StoreInfos(funk.Filter(takeoutStoreList, func(d model.StoreInfo) bool {
		return funk.Find(delivelyStoreList, func(t model.StoreInfo) bool {
			return d.ID == t.ID
		}) == nil
	}).([]model.StoreInfo))
	// 両方
	bothListFilterOnly := funk.Filter(delivelyStoreList, func(d model.StoreInfo) bool {
		return funk.Find(takeoutStoreList, func(t model.StoreInfo) bool {
			return d.ID == t.ID
		}) != nil
	}).([]model.StoreInfo)
	bothList := model.StoreInfos(funk.Map(bothListFilterOnly, func(b model.StoreInfo) model.StoreInfo {
		b.Type = "all"
		return b
	}).([]model.StoreInfo))
	// ソートしてreturn
	allList := append(append(delivelyOnlyList, takeoutOnlyList...), bothList...)
	sort.Sort(allList)
	return allList
}
