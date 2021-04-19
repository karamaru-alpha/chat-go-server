package testdata

import (
	"strings"

	"github.com/karamaru-alpha/chat-go-server/domain/model/message"
)

// Body メッセージ本文に関連する文字列テストデータ
var Body = struct {
	Valid, Empty, TooLong string
}{
	Valid:   "valid_message_body",
	Empty:   "",
	TooLong: strings.Repeat("a", message.BODY_MAX_LEN+1),
}
