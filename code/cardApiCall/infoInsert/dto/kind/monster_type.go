package kind

type MonsterKind struct {
	Kind
}

var (
	MonsterTypeNormal  = MonsterKind{Kind{ID: 1, NameJa: "通常", NameEn: "Normal"}}
	MonsterTypeEffect  = MonsterKind{Kind{ID: 2, NameJa: "効果", NameEn: "Effect"}}
	MonsterTypeToon    = MonsterKind{Kind{ID: 3, NameJa: "トゥーン", NameEn: "Toon"}}
	MonsterTypeSpirit  = MonsterKind{Kind{ID: 4, NameJa: "スピリット", NameEn: "Spirit"}}
	MonsterTypeUnion   = MonsterKind{Kind{ID: 5, NameJa: "ユニオン", NameEn: "Union"}}
	MonsterTypeGemini  = MonsterKind{Kind{ID: 6, NameJa: "デュアル", NameEn: "Gemini"}}
	MonsterTypeTuner   = MonsterKind{Kind{ID: 7, NameJa: "チューナー", NameEn: "Tuner"}}
	MonsterTypeReverse = MonsterKind{Kind{ID: 8, NameJa: "リバース", NameEn: "Reverse"}}
)
