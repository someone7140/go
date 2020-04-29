package model

// StoreInfoRequest ぐるなび店情報のリクエスト項目
type StoreInfoRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Range     int64   `json:"range"`
	CategoryL string  `json:"category_l"`
}

// StoreInfo ぐるなび店情報の項目
type StoreInfo struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Category  string  `json:"category"`
	URL       string  `json:"url"`
	Image     string  `json:"image"`
	Opentime  string  `json:"opentime"`
	Holiday   string  `json:"holiday"`
	Pr        string  `json:"pr"`
	Type      string  `json:"type"`
}

// StoreInfos 店情報のスライス
type StoreInfos []StoreInfo

// Len インターフェースの実装
func (s StoreInfos) Len() int {
	return len(s)
}

// Swap インターフェースの実装
func (s StoreInfos) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less インターフェースの実装
func (s StoreInfos) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}
