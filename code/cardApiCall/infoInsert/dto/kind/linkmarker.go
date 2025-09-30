package kind

import "slices"

type LinkMarker struct {
	MarkerVal   int
	DirectionEn string
}

var (
	Top         = LinkMarker{MarkerVal: 1, DirectionEn: "Top"}
	TopRight    = LinkMarker{MarkerVal: 2, DirectionEn: "Top-Right"}
	Right       = LinkMarker{MarkerVal: 4, DirectionEn: "Right"}
	BottomRight = LinkMarker{MarkerVal: 8, DirectionEn: "Bottom-Right"}
	Bottom      = LinkMarker{MarkerVal: 16, DirectionEn: "Bottom"}
	BottomLeft  = LinkMarker{MarkerVal: 32, DirectionEn: "Bottom-Left"}
	Left        = LinkMarker{MarkerVal: 64, DirectionEn: "Left"}
	LeftTop     = LinkMarker{MarkerVal: 128, DirectionEn: "Left-Top"}
)

func ConvertLinkMarkerStringToLinkMarker(linkMarkerStringList []string) int {
	result := 0
	if slices.Contains(linkMarkerStringList, Top.DirectionEn) {
		result += Top.MarkerVal
	}
	if slices.Contains(linkMarkerStringList, TopRight.DirectionEn) {
		result += TopRight.MarkerVal
	}
	if slices.Contains(linkMarkerStringList, Right.DirectionEn) {
		result += Right.MarkerVal
	}
	if slices.Contains(linkMarkerStringList, BottomRight.DirectionEn) {
		result += BottomRight.MarkerVal
	}
	if slices.Contains(linkMarkerStringList, Bottom.DirectionEn) {
		result += Bottom.MarkerVal
	}
	if slices.Contains(linkMarkerStringList, BottomLeft.DirectionEn) {
		result += BottomLeft.MarkerVal
	}
	if slices.Contains(linkMarkerStringList, Left.DirectionEn) {
		result += Left.MarkerVal
	}
	if slices.Contains(linkMarkerStringList, LeftTop.DirectionEn) {
		result += LeftTop.MarkerVal
	}
	return result
}
