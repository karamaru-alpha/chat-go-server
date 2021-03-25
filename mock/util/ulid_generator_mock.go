package util

import (
	"github.com/oklog/ulid"

	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

// GenerateULID ULID生成処理を固定値でモックした関数
func GenerateULID() ulid.ULID {
	return tdULID.Room.ID.Valid
}
