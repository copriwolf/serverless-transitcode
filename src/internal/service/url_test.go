package service

import (
	"strings"
	"testing"
)

var (
	ctxOriginUrl = "https://or2.in"
	ctxEncodeUrl = "https%3a%2f%2for2.in"

	optUrlEncode = 1
	optUrlDecode = 2
)

func TestUrlEncode(t *testing.T) {
	input := ctxOriginUrl
	want := ctxEncodeUrl
	opt := optUrlEncode
	if got := Url(opt, input); strings.ToLower(got) != want {
		t.Errorf("Url(%d, %s) = %q, want %q", opt, input, got, want)
	}
}

func TestUrlDecode(t *testing.T) {
	input := ctxEncodeUrl
	want := ctxOriginUrl
	opt := optUrlDecode
	if got := Url(opt, input); got != want {
		t.Errorf("Url(%d, %s) = %q, want %q", opt, input, got, want)
	}
}
