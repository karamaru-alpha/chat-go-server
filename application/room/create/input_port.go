package room

// IInputPort トークルームを新規作成するアプリケーションサービスのインターフェース
type IInputPort interface {
	Handle(InputData) OutputData
}
