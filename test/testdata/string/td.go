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
