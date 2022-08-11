package util

import (
	itemPostModel "coda-api/src/model/item"
	"math"
	"sort"

	"github.com/thoas/go-funk"
)

// GetItemCategoryFromAttributeInput 属性情報のインプットからカテゴリーを取得
func GetItemCategoryFromAttributeInput(
	gender string,
	silhouette string,
	complex string,
) string {
	category := ""
	// 性別
	category += map[bool]string{true: gender, false: "none"}[gender != ""] + "-"
	// 体型
	category += map[bool]string{true: silhouette, false: "none"}[silhouette != ""] + "-"
	// コンプレックス
	category += map[bool]string{true: complex, false: "none"}[complex != ""]
	return category
}

// GetMatchItemCategories 設定されたカテゴリーにアイテムのマッチ度をつけて返す
func GetMatchItemCategories(categories []string, registeredComplexes []string) map[string]int {
	matchCategories := map[string]int{}

	// 性別の判定
	genderArray := []string{"male", "female", "none"}
	settingGenderArray := funk.Map(categories, func(category string) string {
		return getGenderFromCategory(category)
	}).([]string)
	funk.ForEach(genderArray, func(gender string) {
		if gender == "none" {
			matchCategories[gender] = 0
		} else if funk.ContainsString(settingGenderArray, gender) {
			matchCategories[gender] = 50
		} else {
			matchCategories[gender] = 0
		}
	})

	// 体型の判定
	tempMatchCategories := map[string]int{}
	silhouetteMap := map[string]int{"standard": 100, "chubby": 101, "thick": 102, "none": 300}
	for silhouetteKey, silhouetteValue := range silhouetteMap {
		// 設定されている体型からポイントの最大値を取得
		settingSilhouettesPoints := funk.Map(categories, func(category string) int {
			if category == "none" {
				return 0
			} else {
				diff := math.Abs(float64(silhouetteValue - silhouetteMap[getSilhouetteFromCategory(category)]))
				if diff == 0 {
					return 30
				} else if diff == 1 {
					return 20
				} else {
					return 0
				}
			}
		}).([]int)
		maxPoint := funk.MaxInt(settingSilhouettesPoints)
		for matchCategoryKey, matchCategoryValue := range matchCategories {
			tempMatchCategories[matchCategoryKey+"-"+silhouetteKey] = matchCategoryValue + maxPoint
		}
	}
	matchCategories = tempMatchCategories

	// コンプレックスの判定
	tempMatchCategories = map[string]int{}
	complexArray := []string{"spotsORscratches", "shoulderWidth", "arm", "leg", "silhouette", "height", "none"}
	funk.ForEach(complexArray, func(complex string) {
		if complex != "none" && funk.ContainsString(registeredComplexes, complex) {
			for matchCategoryKey, matchCategoryValue := range matchCategories {
				tempMatchCategories[matchCategoryKey+"-"+complex] = matchCategoryValue + 20
			}
		} else {
			for matchCategoryKey, matchCategoryValue := range matchCategories {
				tempMatchCategories[matchCategoryKey+"-"+complex] = matchCategoryValue
			}
		}
	})
	matchCategories = tempMatchCategories

	// ポイントが30以上のみのカテゴリーを抽出
	filteredMatchCategories := map[string]int{}
	for categoryKey, categoryValue := range matchCategories {
		if categoryValue >= 30 {
			filteredMatchCategories[categoryKey] = categoryValue
		}
	}
	return filteredMatchCategories
}

// GetItemPostResponseSortedByMatchePoint マッチ度が高いカテゴリーでアイテムをソートして結果を返す
func GetItemPostResponseSortedByMatchePoint(
	itemPosts []itemPostModel.ItemPostResponse, targetCategories map[string]int, limit int,
) []itemPostModel.ItemPostResponse {
	sort.Slice(itemPosts, func(i, j int) bool {
		matchPointI := targetCategories[itemPosts[i].Category]
		matchPointJ := targetCategories[itemPosts[j].Category]
		if matchPointI == matchPointJ {
			return itemPosts[i].PostDate.After(itemPosts[j].PostDate)
		}
		return matchPointI > matchPointJ
	})
	getCount := map[bool]int{true: limit, false: len(itemPosts)}[len(itemPosts) > limit]
	var sortedItemPosts = []itemPostModel.ItemPostResponse{}
	for i := 0; i < getCount; i++ {
		sortedItemPosts = append(sortedItemPosts, itemPosts[i])
	}
	return sortedItemPosts
}

// GetItemPostResponseSortedByPostDate 投稿の最新順でアイテムをソートして結果を返す
func GetItemPostResponseSortedByPostDate(
	itemPosts []itemPostModel.ItemPostResponse, targetCategories map[string]int, limit int,
) []itemPostModel.ItemPostResponse {
	sort.Slice(itemPosts, func(i, j int) bool {
		return itemPosts[i].PostDate.After(itemPosts[j].PostDate)
	})
	getCount := map[bool]int{true: limit, false: len(itemPosts)}[len(itemPosts) > limit]
	var sortedItemPosts = []itemPostModel.ItemPostResponse{}
	for i := 0; i < getCount; i++ {
		sortedItemPosts = append(sortedItemPosts, itemPosts[i])
	}
	return sortedItemPosts
}
