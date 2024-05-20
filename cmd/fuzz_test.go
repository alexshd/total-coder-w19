package cmd

import (
	"net/url"
	"reflect"
	"testing"
	"unicode/utf8"
)

func FuzzParseQuery(f *testing.F) {
	f.Add("x=1&y=2")
	f.Fuzz(func(t *testing.T, queryStr string) {
		query, err := url.ParseQuery(queryStr)
		if err != nil {
			t.Skip()
		}

		queryStr2 := query.Encode()
		query2, err := url.ParseQuery(queryStr2)
		if err != nil {
			t.Fatalf("ParseQuery failed to decode a valid encode query %s: %v", queryStr2, err)
		}
		if !reflect.DeepEqual(query, query2) {
			t.Errorf("ParseQuery gave different query after being encoded\nbefore: %v\nafter: %v", query, query2)
		}
	})
}

func FuzzFirstRune(f *testing.F) {
	f.Add("Hello")
	f.Add("world")
	f.Fuzz(func(t *testing.T, s string) {
		got := FirstRune(s)
		want, _ := utf8.DecodeRuneInString(s)
		if want == utf8.RuneError {
			t.Skip() // don't bother testing invalid runes
		}
		if want != got {
			t.Errorf("given %q (0x%[1]x): want '%c' (0x%[2]x)",
				s, want)
			t.Errorf("got '%c' (0x%[1]x)", got)
		}
	})
}

func FirstRune(s string) rune {
	for _, r := range s {
		return r
	}
	return utf8.RuneError
}
