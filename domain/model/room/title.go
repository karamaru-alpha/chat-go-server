package room

import (
	"errors"
	"unicode/utf8"
)

// Title トークルーム名を表現する値オブジェクト
type Title string

// NewTitle トークルーム名の値オブジェクトを生成するコンストラクタ
func NewTitle(v string) (*Title, error) {

	if v == "" {
		return nil, errors.New("RoomTitle is null")
	}

	if utf8.RuneCountInString(v) < 3 || utf8.RuneCountInString(v) > 20 {
		return nil, errors.New("RoomTitle should be Three to twenty characters")
	}

	title := Title(v)
	return &title, nil
}
