package message

import (
	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/message"
)

// OutputData トークルーム入室アプリケーションサービスの出力
type OutputData struct {
	Messages *[]domainModel.Message
	Err      error
}
