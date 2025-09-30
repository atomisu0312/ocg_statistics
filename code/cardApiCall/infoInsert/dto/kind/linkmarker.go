package kind

import "strings"

var linkMarkerMap = map[string]int32{
	"top":          1,
	"top-right":    2,
	"right":        4,
	"bottom-right": 8,
	"bottom":       16,
	"bottom-left":  32,
	"left":         64,
	"top-left":     128, // 標準表記は Top-Left にする（既存データとの整合は要確認）
}

// ConvertLinkMarkerStringToLinkMarker は方向文字列リストをマスク(int32)に変換する。
// 受け取る文字列は小文字化してハイフン区切りで正規化する（例: "Top-Right" -> "top-right"）。
func ConvertLinkMarkerStringToLinkMarkerValInt(linkMarkerStringList []string) int32 {
	var mask int32
	for _, s := range linkMarkerStringList {
		key := strings.ToLower(strings.TrimSpace(s))
		// 正規化: "top-right" / "top right" / "Top-Right" 等を許容したければここで置換する
		key = strings.ReplaceAll(key, " ", "-")
		if v, ok := linkMarkerMap[key]; ok {
			mask |= v
		}
	}
	return mask
}
