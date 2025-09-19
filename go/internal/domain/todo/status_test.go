package todo

import "testing"

func TestNewStatus(t *testing.T) {
	valid := []string{"todo", "doing", "done"}
	for _, v := range valid {
		if _, err := NewStatus(v); err != nil {
			t.Fatalf("valid %q got err=%v", v, err)
		}
	}
	if _, err := NewStatus("waiting"); err == nil {
		t.Fatalf("invalid status should error")
	}
}
