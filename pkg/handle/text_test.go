package handle

import (
	"net/url"
	"testing"
)

func TestParseURL(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		s := "1"
		u, err := url.Parse(s)
		if err != nil {
			t.Error(err)
		}
		t.Log(u)
	})

	t.Run("url", func(t *testing.T) {
		s := "http://123"
		u, err := url.ParseRequestURI(s)
		if err != nil {
			t.Error(err)
		}
		t.Log(u.Host)
	})
}
