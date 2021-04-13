package room

import (
	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// OutputData トークルームを全件取得するアプリケーションサービスの出力値
type OutputData struct {
	Rooms []domain.Room
	Err   error
}
