package util

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// GenerateULID ランダムなULIDを生成する処理
func GenerateULID() ulid.ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}
