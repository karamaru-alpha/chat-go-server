package testdata

import "strings"

// Title トークルームタイトルに関連する文字列のテストデータ
var Title = struct {
	Valid, TooShort, TooLong string
}{
	Valid:    "valid_title",
	TooShort: ".",
	TooLong:  strings.Repeat("a", 100),
}
