package testdata

import (
	"github.com/oklog/ulid"

	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

// Room ルームエンティティ生成に必要なULIDのテストデータ
var Room = struct {
	ID ulid.ULID
}{
	ID: ulid.MustParse(tdString.Room.ID.Valid),
}

// Message メッセージエンティティ生成に必要なULIDのテストデータ
var Message = struct {
	ID ulid.ULID
}{
	ID: ulid.MustParse(tdString.Message.ID.Valid),
}
