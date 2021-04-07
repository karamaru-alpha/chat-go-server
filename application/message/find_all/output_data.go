package message

import (
	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/message"
)

// OutputData メッセージ一覧取得のアプリケーションサービスの出力
type OutputData struct {
	Messages *[]domainModel.Message
	Err      error
}
