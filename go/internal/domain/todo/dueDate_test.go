package todo

import "testing"

func TestNewDueDate(t *testing.T) {
	cases := []struct {
		name string
		in   string
		ok   bool
		out  *string
	}{
		{"empty ok (nil)", "", true, nil},
		{"valid yyyy-mm-dd", "2025-09-19", true, strPtr("2025-09-19")},
		{"invalid format", "2025/09/19", false, nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d, err := NewDueDate(tc.in)
			if tc.ok && err != nil {
				t.Fatalf("want ok, got err=%v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("want error, got=%v", d)
			}
			if tc.ok {
				if got := d.StringPtr(); (got == nil) != (tc.out == nil) || (got != nil && *got != *tc.out) {
					t.Fatalf("StringPtr mismatch: got=%v want=%v", got, tc.out)
				}
			}
		})
	}
}

func strPtr(s string) *string { return &s }
