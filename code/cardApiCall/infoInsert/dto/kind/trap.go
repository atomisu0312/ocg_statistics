package kind

var (
	TrapTypeNormal = TrapKind{
		Kind: Kind{
			ID:     1,
			NameJa: "通常",
			NameEn: "Normal",
		},
	}
	TrapTypeContinuous = TrapKind{
		Kind: Kind{
			ID:     2,
			NameJa: "永続",
			NameEn: "Continuous",
		},
	}
	TrapTypeCounter = TrapKind{
		Kind: Kind{
			ID:     3,
			NameJa: "カウンター",
			NameEn: "Counter",
		},
	}
	TrapTypeUnknown = TrapKind{
		Kind: Kind{
			ID:     -1,
			NameJa: "不明",
			NameEn: "Unknown",
		},
	}
)

type TrapKind struct {
	Kind
}

func (t *TrapKind) FromSelectFullKindInfoRow(row SelectFullKindInfoRow) TrapKind {
	var kind TrapKind

	// 冗長に見えるかもしれないが、明示的に「通常」「永続」「カウンター」を返すことを強調する意味でこのような表記にしている
	switch row.ID {
	case TrapTypeNormal.ID:
		kind = TrapTypeNormal
	case TrapTypeContinuous.ID:
		kind = TrapTypeContinuous
	case TrapTypeCounter.ID:
		kind = TrapTypeCounter
	default:
		kind = TrapTypeUnknown
	}

	return kind
}
