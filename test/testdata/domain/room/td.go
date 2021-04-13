package testdata

import (
	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"

	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/room"
	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

// Entity トークルームエンティティのテストデータ
var Entity = domain.Room{
	ID:    ID,
	Title: Title,
}

// ID トークルームエンティティ識別子値オブジェクトのテストデータ
var ID = domain.ID(tdULID.ULID)

// Title トークルームタイトルのテストデータ
var Title = domain.Title(tdString.Title.Valid)
