package room

import "context"

// InputData メッセージ作成の入力データ
type InputData struct {
	Context context.Context
	RoomID  string
	Body    string
}
