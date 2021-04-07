package message

// IInputPort メッセージ作成アプリケーションサービスのインターフェース
type IInputPort interface {
	Handle(InputData) OutputData
}
