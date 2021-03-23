package room

import (
	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// OutputData トークルームを新規作成するアプリケーションサービスの出力値
type OutputData struct {
	Room *domainModel.Room
	Err  error
}
