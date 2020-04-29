package model

// CategoryLInfo ぐるなび大カテゴリー情報の項目
type CategoryLInfo struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// CategoryLInfos ぐるなび大カテゴリー情報のスライス
type CategoryLInfos []CategoryLInfo

// Len インターフェースの実装
func (s CategoryLInfos) Len() int {
	return len(s)
}

// Swap インターフェースの実装
func (s CategoryLInfos) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less インターフェースの実装
func (s CategoryLInfos) Less(i, j int) bool {
	return s[i].Code < s[j].Code
}
