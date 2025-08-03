package main

import (
	"fmt"
	"strings"
)

func main() {
	// タブが含まれた文字列のサンプル
	textWithTabs := "Hello\tWorld\tThis\tis\ta\ttest\tstring"
	fmt.Println("元の文字列:")
	fmt.Printf("'%s'\n", textWithTabs)

	// 方法1: strings.ReplaceAll を使用
	textWithoutTabs1 := strings.ReplaceAll(textWithTabs, "\t", "")
	fmt.Println("\n方法1 - strings.ReplaceAll:")
	fmt.Printf("'%s'\n", textWithoutTabs1)

	// 方法2: strings.Replace を使用（複数回置換）
	textWithoutTabs2 := strings.Replace(textWithTabs, "\t", "", -1)
	fmt.Println("\n方法2 - strings.Replace:")
	fmt.Printf("'%s'\n", textWithoutTabs2)

	// 方法3: strings.Fields を使用（空白文字で分割して結合）
	fields := strings.Fields(textWithTabs)
	textWithoutTabs3 := strings.Join(fields, " ")
	fmt.Println("\n方法3 - strings.Fields:")
	fmt.Printf("'%s'\n", textWithoutTabs3)

	// 方法4: 正規表現を使用（regexpパッケージが必要）
	// import "regexp" が必要
	// re := regexp.MustCompile(`\t+`)
	// textWithoutTabs4 := re.ReplaceAllString(textWithTabs, "")

	// 実際のスクレイピングでの使用例
	fmt.Println("\n--- スクレイピングでの使用例 ---")
	scrapedText := "カード名\t青眼の白龍\t攻撃力\t3000\t守備力\t2500"
	fmt.Printf("スクレイピングされたテキスト: '%s'\n", scrapedText)

	// タブを除去
	cleanText := strings.ReplaceAll(scrapedText, "\t", "")
	fmt.Printf("タブ除去後: '%s'\n", cleanText)

	// タブで分割して配列に格納
	parts := strings.Split(scrapedText, "\t")
	fmt.Println("タブで分割:")
	for i, part := range parts {
		fmt.Printf("  parts[%d]: '%s'\n", i, part)
	}
}
