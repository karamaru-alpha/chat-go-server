package testdata

import (
	"strings"

	"github.com/oklog/ulid"
)

var Room = struct {
	ID    roomID
	Title roomTitle
}{
	ID: roomID{
		ValidPlain:   "01D0KDBRASGD5HRSNDCKA0AH53",
		InvalidPlain: "invalid_ulid",
		Valid:        ulid.MustParse("01D0KDBRASGD5HRSNDCKA0AH53"),
	},
	Title: roomTitle{
		Valid:    "valid_title",
		TooShort: ".",
		TooLong:  strings.Repeat("a", 100),
	},
}

type roomID struct {
	ValidPlain   string
	InvalidPlain string
	Valid        ulid.ULID
}

type roomTitle struct {
	Valid    string
	TooShort string
	TooLong  string
}
