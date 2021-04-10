package room

// IInputPort メッセージ作成アプリケーションサービスのインターフェース
type IInputPort interface {
	Handle(InputData) OutputData
}
