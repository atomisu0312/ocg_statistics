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
)

type TrapKind struct {
	Kind
}

func (t *TrapKind) FromSelectFullKindInfoRow(row SelectFullKindInfoRow) TrapKind {
	var k *Kind
	kind := TrapKind{
		Kind: k.FromSelectFullKindInfoRow(row),
	}

	// 冗長に見えるかもしれないが、明示的に「通常」「永続」「カウンター」を返すことを強調する意味でこのような表記にしている
	switch row.ID {
	case 1:
		kind = TrapTypeNormal
	case 2:
		kind = TrapTypeContinuous
	case 3:
		kind = TrapTypeCounter
	}

	return kind
}
