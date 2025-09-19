package todo

import (
	"strings"
	"testing"
)

func TestNewDescription_Validation(t *testing.T) {
	cases := []struct {
		name string
		in   string
		ok   bool
	}{
		{"empty ok", "", true},
		{"trim spaces ok", "  hello  ", true},
		{"max length", strings.Repeat("あ", 300), true},
		{"over max", strings.Repeat("x", 301), false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewDescription(tc.in)
			if tc.ok && err != nil {
				t.Fatalf("want ok, got err=%v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("want error, got=%v", got)
			}
		})
	}
}

func TestDescription_Ptr(t *testing.T) {
	d, _ := NewDescription("")
	if d.Ptr() != nil {
		t.Fatalf("empty description should return nil ptr")
	}
	d2, _ := NewDescription("x")
	if p := d2.Ptr(); p == nil || *p != "x" {
		t.Fatalf("want ptr x, got %+v", p)
	}
}
