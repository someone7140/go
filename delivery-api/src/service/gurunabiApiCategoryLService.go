package service

import (
	"encoding/json"
	"os"
	"sort"

	"../model"
)

// GetCategoryLURL ぐるなび大カテゴリーAPI送信用URL
func GetCategoryLURL() string {
	url := "https://api.gnavi.co.jp/master/CategoryLargeSearchAPI/v3/" +
		"?keyid=" + os.Getenv("GURUNABI_API_KEY")
	return url
}

// ResponseCategoryLJSONConvert ぐるなびの店情報レスポンスをStructの配列にうつす
func ResponseCategoryLJSONConvert(jsonStr string) model.CategoryLInfos {
	var categoryLList model.CategoryLInfos
	if jsonStr != "" {
		var responseMap map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &responseMap)
		if err != nil {
			return categoryLList
		}
		if responseMap["category_l"] != nil {
			responseList := responseMap["category_l"].([]interface{})
			for _, r := range responseList {
				var c model.CategoryLInfo
				c.Code = r.(map[string]interface{})["category_l_code"].(string)
				c.Name = r.(map[string]interface{})["category_l_name"].(string)
				categoryLList = append(categoryLList, c)
			}
		}
	}
	sort.Sort(categoryLList)
	return categoryLList
}
