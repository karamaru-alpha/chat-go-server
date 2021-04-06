package testdata

import (
	"strings"
)

type ulid struct {
	Valid, Invalid string
}

type roomTitle struct {
	Valid, TooShort, TooLong string
}

type messageBody struct {
	Valid, Empty, TooLong string
}

// Room トークルームにまつわる文字列のテストデータ
var Room = struct {
	ID    ulid
	Title roomTitle
}{
	ID: ulid{
		Valid:   "01D0KDBRASGD5HRSNDCKA0AH53",
		Invalid: "invalid_ulid",
	},
	Title: roomTitle{
		Valid:    "valid_title",
		TooShort: ".",
		TooLong:  strings.Repeat("a", 100),
	},
}

// Message メッセージにまつわる文字列のテストデータ
var Message = struct {
	ID   ulid
	Body messageBody
}{
	ID: ulid{
		Valid:   "01D0KDBRASGD5HRSNDCKA0AH53",
		Invalid: "invalid_ulid",
	},
	Body: messageBody{
		Valid:   "valid_message_body",
		Empty:   "",
		TooLong: strings.Repeat("a", 256),
	},
}
