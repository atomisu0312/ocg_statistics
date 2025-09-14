# 原則
- 魔法カードとトラップカードについては、Cardと魔法種別ID(or罠ID)を付加するだけで基本的には問題ない
- useCaseにおいては、魔法の種別(string) -> IDに変換する処理が必要

# テストコードの流れ
## 0.セットアップ
```Go
	t.Run("正常系01 トラップカードの新規登録処理", func(t *testing.T) {
		// セットアップ
		dbConn, card, cleanup := setupTest(t) // 中身については参考に記載
		defer cleanup()
```

## 1.定数の準備
```Go
		dbConn, card, cleanup := setupTest(t)
		defer cleanup()

		// Test data
		trapTypeID := int32(2) // ＜ーここ

		// トランザクションの整備
		ctx := context.Background()
```

## 2.トランザクションの準備
```Go
		// トランザクションの整備
		ctx := context.Background()
		tr := transaction.NewTx(dbConn.DB)

```
## 3.トランザクション境界内で実行
```Go
		var insertedTrap sqlc_gen.Trap

		// トランザクション境界の中で実行(useCaseではこの中にbaseCard挿入処理を入れる)
		err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {
			trapRepo := repository.NewTrapRepository(q)

			trap, err := trapRepo.InsertTrap(ctx, card.ID, trapTypeID)
			if err != nil {
				return fmt.Errorf("error inserting trap: %w", err)
			}
			insertedTrap = trap
			return nil
		})
```

## 4.アサーション
```Go
		require.NoError(t, err, "Transaction should execute without error")

		// Verification
		assert.NotZero(t, insertedTrap.CardID, "Inserted trap should have a non-zero card ID")
		assert.Equal(t, trapTypeID, insertedTrap.TrapTypeID.Int32, "The trap's type ID should match the input")
```

# 初期セットアップについて
## 概要
やっていることは以下の処理の通りです。
- DBの疎通確認
- 初期データの挿入(CSVで代替できんじゃね)
- クリーンアップ関数の返却

クリーンアップ関数でやっているのは、以下の処理
- DBコネクションのクローズ
- テーブルのトランケート

## 実例
```Go
// UTの際に用いるベースとなるカード
var baseCard = sqlc_gen.InsertCardParams{
	NameJa:     sql.NullString{String: "テストカード", Valid: true},
	NameEn:     sql.NullString{String: "Test Card", Valid: true},
	CardTextJa: sql.NullString{String: "テストカードの説明", Valid: true},
	CardTextEn: sql.NullString{String: "Test Card Description", Valid: true},
	NeuronID:   sql.NullInt64{Int64: 1, Valid: true},
	OcgApiID:   sql.NullInt64{Int64: 1, Valid: true},
}

// ベースとなるカードをとりあえず挿入
func insertBaseCard(db *sql.DB) (sqlc_gen.Card, error) {
	ctx := context.Background()
	cardRepo := repository.NewCardRepository(sqlc_gen.New(db))
	card, err := cardRepo.InsertCard(ctx, baseCard)
	if err != nil {
		return sqlc_gen.Card{}, fmt.Errorf("error creating card: %w", err)
	}
	return card, nil
}

// テストの共通セットアップ処理
// 1. DBの疎通確認
// 2. ベースカードを挿入
// 3. クリーンアップ関数の返却
func setupTest(t *testing.T) (*config.DbConn, sqlc_gen.Card, func()) {
	// テスト前処理
	config.BeforeEachForUnitTest()

	// DIコンテナ内の依存関係を設定
	injector := do.New()
	do.Provide(injector, config.TestDbConnection)
	dbConn := do.MustInvoke[*config.DbConn](injector)

	// ベースカードを挿入
	card, err := insertBaseCard(dbConn.DB)

	// エラーが発生した場合はクリーンアップを行う
	if err != nil {
		dbConn.DB.Close()
		config.AfterEachForUnitTest()
		t.Fatalf("Failed to insert base card: %v", err)
	}

	// クリーンアップ関数を返す
	cleanup := func() {
		dbConn.DB.Close()
		config.AfterEachForUnitTest()
	}

	return dbConn, card, cleanup
}
```