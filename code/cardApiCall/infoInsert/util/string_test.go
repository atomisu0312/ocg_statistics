package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitByNewlinesAndTabs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "改行とタブで分割",
			input:    "\n\t\n\t\t\tアイツ\n\t\n\t\t\tアイツ\n\t\n\t\tAitsu\n\t\n\t",
			expected: []string{"アイツ", "アイツ", "Aitsu"},
		},
		{
			name:     "空文字列",
			input:    "",
			expected: []string{},
		},
		{
			name:     "空白のみ",
			input:    "\n\t\n\t",
			expected: []string{},
		},
		{
			name:     "通常のテキスト",
			input:    "Hello\nWorld\tTest",
			expected: []string{"Hello", "World", "Test"},
		},
		{
			name:     "連続する区切り文字",
			input:    "A\n\n\nB\t\t\tC",
			expected: []string{"A", "B", "C"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitByNewlinesAndTabs(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCleanText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "余分な空白を除去",
			input:    "  Hello   World  ",
			expected: "Hello World",
		},
		{
			name:     "改行とタブを空白に変換",
			input:    "Hello\nWorld\tTest",
			expected: "Hello World Test",
		},
		{
			name:     "連続する区切り文字",
			input:    "A\n\n\nB\t\t\tC",
			expected: "A B C",
		},
		{
			name:     "空文字列",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CleanText(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExtractNonEmptyLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "空でない行のみ抽出",
			input:    "Line1\n\nLine2\n  \nLine3",
			expected: []string{"Line1", "Line2", "Line3"},
		},
		{
			name:     "空文字列",
			input:    "",
			expected: []string{},
		},
		{
			name:     "空白のみの行",
			input:    "  \n\t\n  ",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractNonEmptyLines(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRemoveExtraWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "連続する空白を単一に",
			input:    "Hello   World    Test",
			expected: "Hello World Test",
		},
		{
			name:     "タブとスペースの混在",
			input:    "Hello\t\tWorld    Test",
			expected: "Hello World Test",
		},
		{
			name:     "前後の空白を除去",
			input:    "  Hello World  ",
			expected: "Hello World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveExtraWhitespace(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// ベンチマークテスト
func BenchmarkSplitByNewlinesAndTabs(b *testing.B) {
	input := "\n\t\n\t\t\tアイツ\n\t\n\t\t\tアイツ\n\t\n\t\tAitsu\n\t\n\t"
	for i := 0; i < b.N; i++ {
		_ = SplitByNewlinesAndTabs(input)
	}
}

func BenchmarkCleanText(b *testing.B) {
	input := "  Hello\n\n\nWorld\t\t\tTest  "
	for i := 0; i < b.N; i++ {
		_ = CleanText(input)
	}
}
