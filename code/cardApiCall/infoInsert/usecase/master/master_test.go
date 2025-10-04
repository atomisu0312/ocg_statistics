package master_test

import (
	"context"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/app"
	"atomisu.com/ocg-statics/infoInsert/config"
	"atomisu.com/ocg-statics/infoInsert/usecase/master"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestInsertCardInfo(t *testing.T) {
	t.Run("正常系01: 仮テストの実行", func(t *testing.T) {
		config.BeforeEachForUnitTest()
		defer config.AfterEachForUnitTest()

		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)
		masterUseCase := do.MustInvoke[master.MasterUseCase](injector)

		cardID, err := masterUseCase.InsertCardInfo(context.Background(), 4560)
		assert.NoError(t, err)
		assert.NotEqual(t, int64(0), cardID)
	})
}
