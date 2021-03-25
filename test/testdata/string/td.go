package testdata

import (
	"strings"
)

var Room = struct {
	ID    roomID
	Title roomTitle
}{
	ID: roomID{
		Valid:   "01D0KDBRASGD5HRSNDCKA0AH53",
		Invalid: "invalid_ulid",
	},
	Title: roomTitle{
		Valid:    "valid_title",
		TooShort: ".",
		TooLong:  strings.Repeat("a", 100),
	},
}

type roomID struct {
	Valid   string
	Invalid string
}

type roomTitle struct {
	Valid    string
	TooShort string
	TooLong  string
}
