package testdata

import "strings"

// Body メッセージ本文に関連する文字列テストデータ
var Body = struct {
	Valid, Empty, TooLong string
}{
	Valid:   "valid_message_body",
	Empty:   "",
	TooLong: strings.Repeat("a", 256),
}
