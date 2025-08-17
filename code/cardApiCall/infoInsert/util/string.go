package util

import (
	"regexp"
	"strings"
)

// SplitByNewlinesAndTabs は改行とタブでテキストを分割し、空の要素を除去します
func SplitByNewlinesAndTabs(text string) []string {
	// 改行とタブで分割
	parts := regexp.MustCompile(`[\n\t]+`).Split(text, -1)

	// 空の要素と空白のみの要素を除去
	var result []string
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	if len(result) == 0 {
		return []string{}
	}
	return result
}

// CleanText はテキストから余分な空白、改行、タブを除去します
func CleanText(text string) string {
	// 連続する空白、改行、タブを単一の空白に置換
	cleaned := regexp.MustCompile(`[\s\n\t]+`).ReplaceAllString(text, " ")
	// 前後の空白を除去
	return strings.TrimSpace(cleaned)
}

// ExtractNonEmptyLines は空でない行のみを抽出します
func ExtractNonEmptyLines(text string) []string {
	lines := strings.Split(text, "\n")
	var result []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	if len(result) == 0 {
		return []string{}
	}
	return result
}

// RemoveExtraWhitespace は連続する空白文字を単一の空白に置換します
func RemoveExtraWhitespace(text string) string {
	cleaned := regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")
	return strings.TrimSpace(cleaned)
}
