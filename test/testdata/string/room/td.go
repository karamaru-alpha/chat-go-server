package testdata

import (
	"strings"

	"github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// Title トークルームタイトルに関連する文字列のテストデータ
var Title = struct {
	Valid, TooShort, TooLong string
}{
	Valid:    "valid_title",
	TooShort: strings.Repeat("a", room.TITLE_MIN_LEN-1),
	TooLong:  strings.Repeat("a", room.TITLE_MAX_LEN+1),
}
