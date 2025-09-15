package kind

type AttributeKind struct {
	Kind
}

var (
	AttributeLight  = AttributeKind{Kind{ID: 1, NameJa: "光", NameEn: "LIGHT"}}
	AttributeDark   = AttributeKind{Kind{ID: 2, NameJa: "闇", NameEn: "DARK"}}
	AttributeEarth  = AttributeKind{Kind{ID: 3, NameJa: "地", NameEn: "EARTH"}}
	AttributeWater  = AttributeKind{Kind{ID: 4, NameJa: "水", NameEn: "WATER"}}
	AttributeFire   = AttributeKind{Kind{ID: 5, NameJa: "炎", NameEn: "FIRE"}}
	AttributeWind   = AttributeKind{Kind{ID: 6, NameJa: "風", NameEn: "WIND"}}
	AttributeDivine = AttributeKind{Kind{ID: 7, NameJa: "神", NameEn: "DIVINE"}}
)
