package testdata

import (
	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"

	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/message"
	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

// Entity メッセージエンティティのテストデータ
var Entity = messageDomain.Message{
	ID:     ID,
	RoomID: RoomID,
	Body:   Body,
}

// ID メッセージ識別子値オブジェクトのテストデータ
var ID = messageDomain.ID(tdULID.ULID)

// RoomID メッセージエンティティが紐づくトークルーム識別子値オブジェクトのテストデータ
var RoomID = roomDomain.ID(tdULID.ULID)

// Body メッセージ本文値オブジェクトのテストデータ
var Body = messageDomain.Body(tdString.Body.Valid)
