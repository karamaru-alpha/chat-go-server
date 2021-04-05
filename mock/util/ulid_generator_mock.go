package util

import (
	"github.com/oklog/ulid"

	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

// IULIDGenerator ULID生成処理のインターフェース
type IULIDGenerator interface {
	Generate() ulid.ULID
}

type ulidGenerator struct{}

// NewULIDGenerator ULID生成処理のコンストラクタ
func NewULIDGenerator() IULIDGenerator {
	return &ulidGenerator{}
}

// Generate ULID生成処理を固定値でモックした関数
func (ulidGenerator) Generate() ulid.ULID {
	return tdULID.Room.ID
}
