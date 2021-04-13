package testdata

import (
	"github.com/oklog/ulid"

	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/common"
)

// ULID ULIDのテストデータ
var ULID = ulid.MustParse(tdString.ULID.Valid)
