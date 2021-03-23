package room

// IInputPort トークルームを全件取得するアプリケーションサービスのインターフェース
type IInputPort interface {
	Handle() OutputData
}
