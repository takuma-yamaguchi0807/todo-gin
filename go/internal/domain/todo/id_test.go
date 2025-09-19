package todo

import "testing"

func TestNewId(t *testing.T) {
	if _, err := NewId("not-uuid"); err == nil {
		t.Fatalf("invalid uuid should error")
	}
	v := "11111111-1111-1111-1111-111111111111"
	id, err := NewId(v)
	if err != nil {
		t.Fatalf("want ok, got err=%v", err)
	}
	if id.String() != v {
		t.Fatalf("roundtrip mismatch: %s", id.String())
	}
}
