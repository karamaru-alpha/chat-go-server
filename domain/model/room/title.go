package room

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

const TITLE_MIN_LEN = 3
const TITLE_MAX_LEN = 50

// Title トークルーム名を表現する値オブジェクト
type Title string

// NewTitle トークルーム名の値オブジェクトを生成するコンストラクタ
func NewTitle(v string) (Title, error) {
	if v == "" {
		return "", errors.New("RoomTitle is null")
	}
	if utf8.RuneCountInString(v) < TITLE_MIN_LEN || utf8.RuneCountInString(v) > TITLE_MAX_LEN {
		return "", fmt.Errorf("RoomTitle should be %d to %d characters", TITLE_MIN_LEN, TITLE_MAX_LEN)
	}

	return Title(v), nil
}
