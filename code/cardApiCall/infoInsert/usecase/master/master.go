package master

import (
	"context"
	"runtime"
	"sync"

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
	InsertCardInfoList(ctx context.Context, startId int64, delta int64) ([]int64, error)
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

func (m *masterUseCaseImpl) InsertCardInfoList(ctx context.Context, startId int64, delta int64) ([]int64, error) {
	// 並行度を制御（CPUコア数の2倍、最大10）
	maxWorkers := runtime.NumCPU() * 2
	if maxWorkers > 10 {
		maxWorkers = 10
	}

	return m.insertCardInfoListWithWorkers(ctx, startId, delta, maxWorkers)
}

// InsertCardInfoListWithWorkers は、指定されたワーカー数で並行処理を行います
func (m *masterUseCaseImpl) insertCardInfoListWithWorkers(ctx context.Context, startId int64, delta int64, maxWorkers int) ([]int64, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	failedCardIDs := []int64{}

	// セマフォで並行度を制御
	semaphore := make(chan struct{}, maxWorkers)

	for neuronCardID := startId; neuronCardID < startId+delta; neuronCardID++ {
		wg.Add(1)
		go func(cardID int64) {
			defer wg.Done()

			// セマフォを取得（並行度制御）
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			_, err := m.InsertCardInfo(ctx, cardID)
			if err != nil {
				mu.Lock()
				failedCardIDs = append(failedCardIDs, cardID)
				mu.Unlock()
			}
		}(neuronCardID)
	}

	wg.Wait()
	return failedCardIDs, nil
}
