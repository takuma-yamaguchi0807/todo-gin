package user

import "testing"

func TestNewPassword(t *testing.T) {
	cases := []struct {
		name string
		in   string
		ok   bool
	}{
		{"too short", "Ab1!", false},
		{"only lower", "abcdefgh", false},
		{"lower+digit ok", "abcd1234", true},
		{"upper+lower ok", "Abcdefgh", true},
		{"lower+symbol ok", "abcd!!!!", true},
		{"mixed ok", "Abcdef1!", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewPassword(tc.in)
			if tc.ok && err != nil {
				t.Fatalf("want ok, got %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("want error")
			}
		})
	}
}
