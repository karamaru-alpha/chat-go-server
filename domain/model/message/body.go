package message

import (
	"errors"
	"unicode/utf8"
)

// Body メッセージ本文を表す値オブジェクト
type Body string

// NewBody メッセージ本文を表す値オブジェクトのコンストラクタ
func NewBody(v string) (*Body, error) {

	if v == "" {
		return nil, errors.New("MessageBody is empty")
	}

	if utf8.RuneCountInString(v) == 0 || utf8.RuneCountInString(v) > 255 {
		return nil, errors.New("MessageBody should be 1 to 255 characters")
	}

	body := Body(v)
	return &body, nil
}