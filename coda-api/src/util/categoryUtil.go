package util

import (
	"math"
	"sort"
	"strings"

	postModel "coda-api/src/model/post"

	"github.com/thoas/go-funk"
)

// GetCategoriesFromAttributeInput 属性情報のインプットからカテゴリーを取得
func GetCategoriesFromAttributeInput(
	gender string,
	silhouette string,
	height int,
	genres []string,
) []string {
	var categories []string
	category := ""
	// 性別・体型
	category += map[bool]string{true: "none", false: gender}[gender == ""] + "-" + map[bool]string{true: "none", false: silhouette}[silhouette == ""]
	// 身長
	if height < 1 {
		category += "-none"
	} else {
		category += "-" + GetHeightCategory(height, gender)
	}
	// ジャンル
	if len(genres) == 0 {
		category += "-none"
		categories = append(categories, category)
	} else {
		funk.ForEach(genres, func(genre string) {
			categories = append(categories, category+"-"+genre)
		})
	}
	return categories
}

// GetMatchCategories 設定されたカテゴリーにマッチ度をつけて返す
func GetMatchCategories(categories []string) map[string]int {
	matchCategories := map[string]int{}

	// 性別の判定
	genderArray := []string{"male", "female"}
	settingGenderArray := funk.Map(categories, func(category string) string {
		return getGenderFromCategory(category)
	}).([]string)
	funk.ForEach(genderArray, func(gender string) {
		if funk.ContainsString(settingGenderArray, gender) {
			matchCategories[gender] = 50
		} else {
			matchCategories[gender] = 0
		}
	})

	// 体型の判定
	tempMatchCategories := map[string]int{}
	silhouetteMap := map[string]int{"standard": 100, "chubby": 101, "thick": 102}
	for silhouetteKey, silhouetteValue := range silhouetteMap {
		// 設定されている体型からポイントの最大値を取得
		settingSilhouettesPoints := funk.Map(categories, func(category string) int {
			diff := math.Abs(float64(silhouetteValue - silhouetteMap[getSilhouetteFromCategory(category)]))
			if diff == 0 {
				return 30
			} else if diff == 1 {
				return 20
			} else {
				return 0
			}
		}).([]int)
		maxPoint := funk.MaxInt(settingSilhouettesPoints)
		for matchCategoryKey, matchCategoryValue := range matchCategories {
			tempMatchCategories[matchCategoryKey+"-"+silhouetteKey] = matchCategoryValue + maxPoint
		}
	}
	matchCategories = tempMatchCategories

	// 身長の判定
	tempMatchCategories = map[string]int{}
	heightMap := map[string]int{"low": 100, "standard": 101, "high": 102}
	for heightKey, heightValue := range heightMap {
		// 設定されている身長からポイントの最大値を取得
		settingHeightsPoints := funk.Map(categories, func(category string) int {
			diff := math.Abs(float64(heightValue - silhouetteMap[getHeightFromCategory(category)]))
			if diff == 0 {
				return 15
			} else if diff == 1 {
				return 10
			} else {
				return 0
			}
		}).([]int)
		maxPoint := funk.MaxInt(settingHeightsPoints)
		for matchCategoryKey, matchCategoryValue := range matchCategories {
			tempMatchCategories[matchCategoryKey+"-"+heightKey] = matchCategoryValue + maxPoint
		}
	}
	matchCategories = tempMatchCategories

	// ジャンルの判定
	tempMatchCategories = map[string]int{}
	genreArray := []string{"street", "beautiful", "casual"}
	settingGenreArray := funk.Map(categories, func(category string) string {
		return GetGenreFromCategory(category)
	}).([]string)
	funk.ForEach(genreArray, func(genre string) {
		if funk.ContainsString(settingGenreArray, genre) {
			for matchCategoryKey, matchCategoryValue := range matchCategories {
				tempMatchCategories[matchCategoryKey+"-"+genre] = matchCategoryValue + 5
			}
		} else {
			for matchCategoryKey, matchCategoryValue := range matchCategories {
				tempMatchCategories[matchCategoryKey+"-"+genre] = matchCategoryValue
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

// getGenderFromCategory カテゴリーから性別を取得
func getGenderFromCategory(category string) string {
	slice := strings.Split(category, "-")
	return slice[0]
}

// getSilhouetteFromCategory カテゴリーから体型を取得
func getSilhouetteFromCategory(category string) string {
	slice := strings.Split(category, "-")
	if len(slice) > 1 {
		return slice[1]
	}
	return "none"
}

// getHeightFromCategory カテゴリーから身長を取得
func getHeightFromCategory(category string) string {
	slice := strings.Split(category, "-")
	if len(slice) > 2 {
		return slice[2]
	}
	return "none"
}

// GetGenreFromCategory カテゴリーからジャンルを取得
func GetGenreFromCategory(category string) string {
	slice := strings.Split(category, "-")
	if len(slice) > 3 {
		return slice[3]
	}
	return "none"
}

// GetPostCategoryMergePostGenre カテゴリーと投稿で指定したジャンルをマージして取得
func GetPostCategoryMergePostGenre(category string, genre string) string {
	returnCategory := ""
	slice := strings.Split(category, "-")
	returnCategory = returnCategory + slice[0] + "-"
	if len(slice) > 1 {
		returnCategory = returnCategory + slice[1] + "-"
	} else {
		returnCategory = returnCategory + "none-"
	}
	if len(slice) > 2 {
		returnCategory = returnCategory + slice[2] + "-"
	} else {
		returnCategory = returnCategory + "none-"
	}
	returnCategory = returnCategory + genre
	return returnCategory
}

// GetPostResponseSortedByMatchePoint マッチ度が高いカテゴリーでソートして結果を返す
func GetPostResponseSortedByMatchePoint(
	posts []postModel.PostInfoWithCategory, targetCategories map[string]int, limit int,
) []postModel.PostResponseForRecommend {
	sort.Slice(posts, func(i, j int) bool {
		matchPointI := targetCategories[posts[i].Category]
		matchPointJ := targetCategories[posts[j].Category]
		if matchPointI == matchPointJ {
			return posts[i].PostDate.After(posts[j].PostDate)
		}
		return matchPointI > matchPointJ
	})
	getCount := map[bool]int{true: limit, false: len(posts)}[len(posts) > limit]
	var postsResponseForRecommend []postModel.PostResponseForRecommend
	for i := 0; i < getCount; i++ {
		postsResponseForRecommend = append(postsResponseForRecommend, postModel.PostResponseForRecommend{
			ID:            posts[i].ID,
			ContentURL:    posts[i].ContentURL,
			Genre:         GetGenreFromCategory(posts[i].Category),
			PostDate:      posts[i].PostDate,
			FavoriteCount: posts[i].FavoriteCount,
			FavoritedFlg:  posts[i].FavoritedFlg,
		})
	}
	return postsResponseForRecommend
}

// GetHeightCategory 身長の数値から区分を取得
func GetHeightCategory(height int, gender string) string {
	if gender == "male" {
		if height <= 165 {
			return "low"
		} else if height <= 179 {
			return "standard"
		} else {
			return "high"
		}
	} else {
		if height <= 150 {
			return "low"
		} else if height <= 169 {
			return "standard"
		} else {
			return "high"
		}
	}
}
