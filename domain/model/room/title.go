package room

import (
	"errors"
	"unicode/utf8"
)

// Title トークルーム名を表現する値オブジェクト
type Title string

// NewTitle トークルーム名の値オブジェクトを生成するコンストラクタ
func NewTitle(v string) (Title, error) {
	if v == "" {
		return "", errors.New("RoomTitle is null")
	}
	if utf8.RuneCountInString(v) < 3 || utf8.RuneCountInString(v) > 50 {
		return "", errors.New("RoomTitle should be Three to twenty characters")
	}

	return Title(v), nil
}
