package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMaster マスタデータの取得
func GetMaster(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"gender": []interface{}{
			map[string]interface{}{"value": "male", "label": "男性"},
			map[string]interface{}{"value": "female", "label": "女性"},
		},
		"silhouette": []interface{}{
			map[string]interface{}{"value": "standard", "label": "標準"},
			map[string]interface{}{"value": "chubby", "label": "ぽっちゃり"},
			map[string]interface{}{"value": "thick", "label": "太め"},
		},
		"height": []interface{}{
			map[string]interface{}{"value": "low", "label": "低い"},
			map[string]interface{}{"value": "standard", "label": "標準"},
			map[string]interface{}{"value": "high", "label": "高い"},
		},
		"genre": []interface{}{
			map[string]interface{}{"value": "street", "label": "ストリート"},
			map[string]interface{}{"value": "beautiful", "label": "キレイめ"},
			map[string]interface{}{"value": "casual", "label": "カジュアル"},
		},
		"complex": []interface{}{
			map[string]interface{}{"value": "spotsORscratches", "label": "肌、毛穴"},
			map[string]interface{}{"value": "shoulderWidth", "label": "肩幅"},
			map[string]interface{}{"value": "arm", "label": "腕"},
			map[string]interface{}{"value": "leg", "label": "脚"},
			map[string]interface{}{"value": "silhouette", "label": "体型"},
			map[string]interface{}{"value": "height", "label": "身長"},
		},
		"post_status": []interface{}{
			map[string]interface{}{"value": "notset", "label": "未設定"},
			map[string]interface{}{"value": "open", "label": "公開"},
			map[string]interface{}{"value": "close", "label": "未公開"},
		},
		"gather_status": []interface{}{
			map[string]interface{}{"value": "on", "label": "対象"},
			map[string]interface{}{"value": "off", "label": "非対象"},
		},
		"item_type": []interface{}{
			map[string]interface{}{"value": "apparel", "label": "アパレル系"},
			map[string]interface{}{"value": "accessory", "label": "アクセサリー系"},
			map[string]interface{}{"value": "cosmetics", "label": "コスメ系"},
			map[string]interface{}{"value": "hairCare", "label": "ヘアケア系"},
			map[string]interface{}{"value": "food", "label": "食品系"},
			map[string]interface{}{"value": "other", "label": "その他"},
		},
		"coordinate_category": []interface{}{
			map[string]interface{}{"value": "one_piece", "label": "ワンピース"},
			map[string]interface{}{"value": "dress", "label": "ドレス"},
			map[string]interface{}{"value": "tops", "label": "トップス"},
			map[string]interface{}{"value": "jacket", "label": "ジャケット"},
			map[string]interface{}{"value": "coat", "label": "コート"},
			map[string]interface{}{"value": "skirt", "label": "スカート"},
			map[string]interface{}{"value": "bottoms", "label": "ボトムス"},
			map[string]interface{}{"value": "denim", "label": "デニム"},
			map[string]interface{}{"value": "setup", "label": "セットアップ"},
			map[string]interface{}{"value": "jump_suit", "label": "ジャンプスーツ"},
			map[string]interface{}{"value": "accessory", "label": "アクセサリー"},
			map[string]interface{}{"value": "item", "label": "アイテム"},
			map[string]interface{}{"value": "shoes", "label": "シューズ"},
			map[string]interface{}{"value": "other", "label": "その他"},
		},
	})
}
