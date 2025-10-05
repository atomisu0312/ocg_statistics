package cardrecord

import (
	"atomisu.com/ocg-statics/infoInsert/dto"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/util"
)

// StandardCard はUsecaseの関数の引数として機能する
type StandardCard struct {
	CardID         int64    `json:"cardId"`
	DescEn         string   `json:"descEn"`
	DescJa         string   `json:"descJa"`
	NameEn         string   `json:"nameEn"`
	NameJa         string   `json:"nameJa"`
	NeuronID       int64    `json:"neuronId"`
	TcgID          int64    `json:"tcgId"`
	Def            int32    `json:"def"`
	Atk            int32    `json:"atk"`
	Type           string   `json:"type"`
	Level          int32    `json:"level"`
	Race           string   `json:"race"`
	LinkMarkers    []string `json:"linkMarkers"`
	Attribute      string   `json:"attribute"`
	LinkVal        int32    `json:"linkVal"`
	TypeLines      []string `json:"typeLines"`
	PendulumScale  int32    `json:"pendulumScale"`
	PendulumTextJa string   `json:"pendulumTextJa"`
	PendulumTextEn string   `json:"pendulumTextEn"`
}

func (s *StandardCard) ToInsertCardParamsExceptMonster() sqlc_gen.InsertCardParams {
	return sqlc_gen.InsertCardParams{
		NameEn:     util.ParseAsSqlNullString(s.NameEn),
		NameJa:     util.ParseAsSqlNullString(s.NameJa),
		CardTextEn: util.ParseAsSqlNullString(s.DescEn),
		CardTextJa: util.ParseAsSqlNullString(s.DescJa),
		NeuronID:   util.ParseAsSqlNullInt64WithTreatZeroAsNull(s.NeuronID),
		OcgApiID:   util.ParseAsSqlNullInt64WithTreatZeroAsNull(s.TcgID),
	}
}

type AbstractCardSelectResult struct {
	dto.AbstractSelectResult
	NeuronID   int64  `db:"neuron_id" json:"neuronId"`
	OcgApiID   int64  `db:"ocg_api_id" json:"ocgApiId"`
	NameJa     string `db:"name_ja" json:"nameJa"`
	NameEn     string `db:"name_en" json:"nameEn"`
	CardTextJa string `db:"card_text_ja" json:"cardTextJa"`
	CardTextEn string `db:"card_text_en" json:"cardTextEn"`
}

func GenerateStandardCardFromNeuronAndTCGAPIResult(neuronExtractedData *NeuronExtractedData, tcgAPICard *TcgApiCard) StandardCard {
	return StandardCard{
		CardID:         neuronExtractedData.CardID,
		DescEn:         tcgAPICard.Desc,
		DescJa:         neuronExtractedData.CardTextJa,
		NameEn:         tcgAPICard.Name,
		NameJa:         neuronExtractedData.CardNameJa,
		NeuronID:       neuronExtractedData.CardID,
		TcgID:          tcgAPICard.ID,
		Def:            tcgAPICard.Def,
		Atk:            tcgAPICard.Atk,
		Type:           tcgAPICard.Type,
		Level:          tcgAPICard.Level,
		Race:           tcgAPICard.Race,
		LinkMarkers:    tcgAPICard.LinkMarkers,
		Attribute:      tcgAPICard.Attribute,
		LinkVal:        tcgAPICard.LinkVal,
		TypeLines:      tcgAPICard.TypeLines,
		PendulumScale:  tcgAPICard.Scale,
		PendulumTextJa: neuronExtractedData.PendulumTextJa,
		PendulumTextEn: tcgAPICard.PendulumText,
	}
}

// NeuronExtractedData は、NeuronUseCaseの抽出データです。
type NeuronExtractedData struct {
	CardID         int64
	CardNameEn     string
	CardNameJa     string
	CardTextJa     string
	PendulumTextJa string
}

// TcgApiCard は、TcgApiUseCaseの抽出データです。
type TcgApiCard struct {
	Desc                  string   `json:"desc"`
	Name                  string   `json:"name"`
	ID                    int64    `json:"id"`
	Def                   int32    `json:"def"`
	Atk                   int32    `json:"atk"`
	Type                  string   `json:"type"`
	Level                 int32    `json:"level"`
	Race                  string   `json:"race"`
	LinkMarkers           []string `json:"linkMarkers"`
	Attribute             string   `json:"attribute"`
	LinkVal               int32    `json:"linkVal"`
	TypeLines             []string `json:"typeline"`
	HumanReadableCardType string   `json:"humanReadableCardType"`
	Scale                 int32    `json:"scale"`
	PendulumText          string   `json:"pendulumText"`
}

type CardPatternSelectResult struct {
	CardID    int64 `db:"card_id" json:"cardId"`
	NeuronID  int64 `db:"neuron_id" json:"neuronId"`
	OcgApiID  int64 `db:"ocg_api_id" json:"ocgApiId"`
	IsMonster bool  `db:"is_monster" json:"isMonster"`
	IsSpell   bool  `db:"is_spell" json:"isSpell"`
	IsTrap    bool  `db:"is_trap" json:"isTrap"`
}

func (c *CardPatternSelectResult) FromSelectCardPatternByCardIDRow(row sqlc_gen.SelectCardPatternByCardIDRow) *CardPatternSelectResult {
	return &CardPatternSelectResult{
		CardID:    row.CardID,
		NeuronID:  row.NeuronID.Int64,
		OcgApiID:  row.OcgApiID.Int64,
		IsMonster: row.IsMonster,
		IsSpell:   row.IsSpell,
		IsTrap:    row.IsTrap,
	}
}
