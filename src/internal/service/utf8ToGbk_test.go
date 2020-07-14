package service

import (
	"testing"
)

var (
	ctxGBK = "ת\xbb\xbb\xb2\xe2\xca\xd4"
	ctxUTF8 = "转换测试"

	optUtf8Encode = 1
	optUtf8Decode = 2
)

func TestUtf8ToGbkEncode(t *testing.T) {
	input := ctxUTF8
	want := ctxGBK
	opt := optUtf8Encode
	if got := Utf8ToGbk(opt, input); got != want {
		t.Errorf("Utf8ToGbk(%d, %s) = %q, want %q", opt, input, got, want)
	}
}

func TestUtf8ToGbkDecode(t *testing.T) {
	input := ctxGBK
	want := ctxUTF8
	opt := optUtf8Decode
	if got := Utf8ToGbk(opt, input); got != want {
		t.Errorf("Utf8ToGbk(%d, %s) = %q, want %q", opt, input, got, want)
	}
}
