package kind

var (
	SpellTypeNormal = SpellKind{
		Kind: Kind{
			ID:     1,
			NameJa: "通常",
			NameEn: "Normal",
		},
	}
	SpellTypeContinuous = SpellKind{
		Kind: Kind{
			ID:     2,
			NameJa: "永続",
			NameEn: "Continuous",
		},
	}
	SpellTypeEquip = SpellKind{
		Kind: Kind{
			ID:     3,
			NameJa: "装備",
			NameEn: "Equip",
		},
	}
	SpellTypeField = SpellKind{
		Kind: Kind{
			ID:     4,
			NameJa: "フィールド",
			NameEn: "Field",
		},
	}
	SpellTypeQuickPlay = SpellKind{
		Kind: Kind{
			ID:     5,
			NameJa: "速攻",
			NameEn: "Quick-Play",
		},
	}
	SpellTypeRitual = SpellKind{
		Kind: Kind{
			ID:     6,
			NameJa: "儀式",
			NameEn: "Ritual",
		},
	}
	SpellTypeUnknown = SpellKind{
		Kind: Kind{
			ID:     -1,
			NameJa: "不明",
			NameEn: "Unknown",
		},
	}
)

type SpellKind struct {
	Kind
}

func (s *SpellKind) FromSelectFullKindInfoRow(row SelectFullKindInfoRow) SpellKind {
	var kind SpellKind

	switch row.ID {
	case SpellTypeNormal.ID:
		kind = SpellTypeNormal
	case SpellTypeContinuous.ID:
		kind = SpellTypeContinuous
	case SpellTypeEquip.ID:
		kind = SpellTypeEquip
	case SpellTypeField.ID:
		kind = SpellTypeField
	case SpellTypeQuickPlay.ID:
		kind = SpellTypeQuickPlay
	case SpellTypeRitual.ID:
		kind = SpellTypeRitual
	default:
		kind = SpellTypeUnknown
	}

	return kind
}
