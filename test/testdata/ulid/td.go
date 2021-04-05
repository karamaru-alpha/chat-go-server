package testdata

import (
	"github.com/oklog/ulid"

	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

var Room = struct {
	ID ulid.ULID
}{
	ID: ulid.MustParse(tdString.Room.ID.Valid),
}
