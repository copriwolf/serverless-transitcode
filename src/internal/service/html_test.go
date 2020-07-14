package service

import (
	"testing"
)

var (
	ctxOriginHtml = `"https://or2.in" <root@or2.in>`
	ctxEncodeHtmlEntity = `&quot;https://or2.in&quot; &lt;root@or2.in&gt;`
	ctxEncodeHtmlDecimal = `&#34;https://or2.in&#34; &lt;root@or2.in&gt;`

	optHtmlEncode = 1
	optHtmlDecode = 2
)

func TestHtmlEncode(t *testing.T) {
	input := ctxOriginHtml
	wantEntity := ctxEncodeHtmlEntity
	wantDecimal := ctxEncodeHtmlDecimal
	opt := optHtmlEncode
	if got := Html(opt, input); got != wantEntity && got != wantDecimal {
		t.Errorf("Html(%d, %s) = %q, want %q", opt, input, got, wantDecimal)
	}
}

func TestHtmlDecode(t *testing.T) {
	input := ctxEncodeHtmlEntity
	want := ctxOriginHtml
	opt := optHtmlDecode
	if got := Html(opt, input); got != want {
		t.Errorf("Html(%d, %s) = %q, want %q", opt, input, got, want)
	}
}
