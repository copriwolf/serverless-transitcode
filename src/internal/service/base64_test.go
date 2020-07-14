package service

import "testing"

var (
	ctxOriginBase64 = "aaaaaaaaaaaaaaaa"
	ctxEncodeBase64 = "YWFhYWFhYWFhYWFhYWFhYQ=="

	optBase64Encode = 1
	optBase64Decode = 2
)

func TestBase64Encode(t *testing.T) {
	input := ctxOriginBase64
	want := ctxEncodeBase64
	opt := optBase64Encode
	if got := Base64(opt, input); got != want {
		t.Errorf("Base64(%d, %s) = %q, want %q", opt, input, got, want)
	}
}

func TestBase64Decode(t *testing.T) {
	input := ctxEncodeBase64
	want := ctxOriginBase64
	opt := optBase64Decode
	if got := Base64(opt, input); got != want {
		t.Errorf("Base64(%d, %s) = %q, want %q", opt, input, got, want)
	}
}
