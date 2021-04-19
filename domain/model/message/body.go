package message

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

const BODY_MIN_LEN = 1
const BODY_MAX_LEN = 255

// Body メッセージ本文を表す値オブジェクト
type Body string

// NewBody メッセージ本文を表す値オブジェクトのコンストラクタ
func NewBody(v string) (Body, error) {
	if v == "" {
		return "", errors.New("MessageBody is empty")
	}
	if utf8.RuneCountInString(v) < BODY_MIN_LEN || utf8.RuneCountInString(v) > BODY_MAX_LEN {
		return "", fmt.Errorf("MessageBody should be %d to %d characters", BODY_MIN_LEN, BODY_MAX_LEN)
	}

	return Body(v), nil
}
