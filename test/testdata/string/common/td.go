package testdata

// ULID ULIDの元になる文字列テストデータ
var ULID = struct {
	Valid, Invalid string
}{
	Valid:   "01D0KDBRASGD5HRSNDCKA0AH53",
	Invalid: "invalid_ulid",
}
