package room

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// Create Roomエンティティの生成処理を担うファクトリ
func Create(title *Title) (*Room, error) {

	ulid := generateULID()

	roomID, err := NewID(&ulid)
	if err != nil {
		return nil, err
	}

	return NewRoom(roomID, title)
}

func generateULID() ulid.ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}
