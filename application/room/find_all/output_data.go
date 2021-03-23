package room

import (
	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// OutputData トークルームを全件取得するアプリケーションサービスの出力値
type OutputData struct {
	Rooms *[]domainModel.Room
	Err   error
}
