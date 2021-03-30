package config

import "os"

// Port ポート番号を取得
func Port() string {
	return os.Getenv("PORT")
}
