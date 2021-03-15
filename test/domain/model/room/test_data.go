package room

import (
	"strings"

	"github.com/oklog/ulid"
)

var testData = struct {
	id    idTestData
	title titleTestData
}{
	id: idTestData{
		valid: ulid.MustParse("01D0KDBRASGD5HRSNDCKA0AH53"),
	},
	title: titleTestData{
		valid:    "valid_title",
		tooShort: ".",
		tooLong:  strings.Repeat("a", 100),
	},
}

type idTestData struct {
	valid ulid.ULID
}

type titleTestData struct {
	valid    string
	tooShort string
	tooLong  string
}
