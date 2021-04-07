package message

// IInputPort メッセージ一覧取得アプリケーションサービスのインターフェース
type IInputPort interface {
	Handle(InputData) OutputData
}
