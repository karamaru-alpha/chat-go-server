package testdata

import (
	"github.com/oklog/ulid"
)

var Room = struct {
	ID roomID
}{
	ID: roomID{
		Valid: ulid.MustParse("01D0KDBRASGD5HRSNDCKA0AH53"),
	},
}

type roomID struct {
	Valid ulid.ULID
}
