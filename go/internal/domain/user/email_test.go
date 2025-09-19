package user

import "testing"

func TestNewEmail(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"not-mail", false},
		{"a@b", true},
		{"test@example.com", true},
	}
	for _, c := range cases {
		_, err := NewEmail(c.in)
		if c.ok && err != nil {
			t.Fatalf("want ok for %q, got %v", c.in, err)
		}
		if !c.ok && err == nil {
			t.Fatalf("want error for %q", c.in)
		}
	}
}
