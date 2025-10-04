package master

import (
	"context"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/usecase"
	"atomisu.com/ocg-statics/infoInsert/usecase/neon"
	"atomisu.com/ocg-statics/infoInsert/usecase/neuron"
	"atomisu.com/ocg-statics/infoInsert/usecase/tcgapi"
	"github.com/samber/do"
)

// masterUseCaseImpl は、MasterUseCaseの実装です。
type masterUseCaseImpl struct {
	neuronUseCase neuron.NeuronUseCase
	tcgapiUseCase tcgapi.TcgUseCase
	neonUseCase   neon.NeonUseCase
}

// MasterUseCase は、MasterUseCaseのインターフェースです。
type MasterUseCase interface {
	InsertCardInfo(ctx context.Context, neuronCardID int64) (int64, error)
}

// NewMasterUseCase は、MasterUseCaseのコンストラクタです。
func NewMasterUseCase(i *do.Injector) (MasterUseCase, error) {
	return usecase.NewUseCase(i, func(u *usecase.UseCaseImpl) MasterUseCase {
		return &masterUseCaseImpl{
			neuronUseCase: do.MustInvoke[neuron.NeuronUseCase](i),
			tcgapiUseCase: do.MustInvoke[tcgapi.TcgUseCase](i),
			neonUseCase:   do.MustInvoke[neon.NeonUseCase](i),
		}
	})
}

func (m *masterUseCaseImpl) InsertCardInfo(ctx context.Context, neuronCardID int64) (int64, error) {
	neuronCardInfo, err := m.neuronUseCase.GetCardInfo(ctx, neuronCardID)
	if err != nil {
		return 0, err
	}
	tcgapiCardInfo, err := m.tcgapiUseCase.GetCardInfoByEnName(ctx, neuronCardInfo.CardNameEn)
	if err != nil {
		return 0, err
	}
	cardInfo := cardrecord.GenerateStandardCardFromNeuronAndTCGAPIResult(&neuronCardInfo, &tcgapiCardInfo)
	cardID, err := m.neonUseCase.InsertCardInfo(ctx, cardInfo)
	if err != nil {
		return 0, err
	}
	return cardID, nil
}
