package room

import (
	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// OutputData トークルームを新規作成するアプリケーションサービスの出力値
type OutputData struct {
	Room domain.Room
	Err  error
}
