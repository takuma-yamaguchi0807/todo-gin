package todo

import (
	"strings"
	"testing"
)

func TestNewTitle_Validation(t *testing.T) {
	cases := []struct {
		name string
		in   string
		ok   bool
	}{
		{"empty", "", false},
		{"spaces only", "   ", false},
		{"min length", "a", true},
		{"max length", strings.Repeat("あ", 100), true},
		{"over max", strings.Repeat("x", 101), false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewTitle(tc.in)
			if tc.ok && err != nil {
				t.Fatalf("want ok, got err=%v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("want error, got=%v", got)
			}
		})
	}
}
