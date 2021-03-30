package config

import (
	"os"
)

// Port ポート番号を取得する
func Port() string {
	return os.Getenv("PORT")
}
