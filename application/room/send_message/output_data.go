package room

import (
	domain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
)

// OutputData メッセージ作成のアプリケーションサービスの出力
type OutputData struct {
	Message domain.Message
	Err     error
}
