package util

import (
	"github.com/karamaru-alpha/chat-go-server/test/testdata"
	"github.com/oklog/ulid"
)

// GenerateULID ULID生成処理を固定値でモックした関数
func GenerateULID() ulid.ULID {
	return testdata.Room.ID.Valid
}
