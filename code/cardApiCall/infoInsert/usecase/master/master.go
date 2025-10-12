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

// insertCardInfoListWithWorkers は、指定されたワーカー数で並行処理を行います
func (m *masterUseCaseImpl) insertCardInfoListWithWorkers(ctx context.Context, startId int64, delta int64, maxWorkers int) ([]int64, error) {
	if delta <= 0 {
		return nil, nil
	}

	// Job channel に ID を流し、固定数のワーカーで処理する
	jobs := make(chan int64)
	results := make(chan int64)
	var wg sync.WaitGroup

	// ワーカー起動
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for cardID := range jobs {
				if _, err := m.InsertCardInfo(ctx, cardID); err != nil {
					// 失敗IDを results に出す
					results <- cardID
				}
			}
		}()
	}

	// ジョブ送信を別ゴルーチンで行い、終了後に jobs を閉じる
	go func() {
		defer close(jobs)
		for id := startId; id < startId+delta; id++ {
			jobs <- id
		}
	}()

	// 結果（失敗したcardId）収集用ゴルーチン（閉じるタイミング管理）
	var collectWg sync.WaitGroup
	failed := []int64{}
	collectWg.Add(1)
	go func() {
		defer collectWg.Done()
		for id := range results {
			failed = append(failed, id)
		}
	}()

	// 全ワーカー完了後に results を閉じる
	wg.Wait()
	close(results)
	collectWg.Wait()

	return failed, nil
}
