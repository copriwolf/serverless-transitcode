package service

import (
	"strings"
	"testing"
)

var (
	ctxOriginUnicode = `转换测试`
	ctxEncodeUnicodeDis = `\u8f6c\u6362\u6d4b\u8bd5`
	ctxEncodeUnicode = "\u8f6c\u6362\u6d4b\u8bd5"

	optUniEncode = 1
	optUniDecode = 2
)

func TestUnicodeEncode(t *testing.T) {
	input := ctxOriginUnicode
	want := ctxEncodeUnicodeDis
	opt := optUniEncode
	if got := Unicode(opt, input); strings.ToLower(got) != want {
		t.Errorf("Url(%d, %s) = %q, want %q", opt, input, got, want)
	}
}

func TestUnicodeDecode(t *testing.T) {
	input := ctxEncodeUnicode
	want := ctxOriginUnicode
	opt := optUniDecode
	if got := Unicode(opt, input); got != want {
		t.Errorf("Url(%d, %s) = %q, want %q", opt, input, got, want)
	}
}
