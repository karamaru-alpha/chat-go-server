package message

// IInputPort トークルーム入室アプリケーションサービスのインターフェース
type IInputPort interface {
	Handle(InputData)
}
