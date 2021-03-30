package util

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// IULIDGenerator ULID生成のインターフェース
type IULIDGenerator interface {
	Generate() ulid.ULID
}

type ulidGenerator struct{}

// NewULIDGenerator ULID生成処理のコンストラクタ
func NewULIDGenerator() IULIDGenerator {
	return &ulidGenerator{}
}

// Generate ランダムなULIDを生成する処理
func (ulidGenerator) Generate() ulid.ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}
