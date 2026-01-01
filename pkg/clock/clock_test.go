package clock

import (
	"testing"
	"time"
)

func TestIsDay(t *testing.T) {
	tests := []struct {
		hour int
		want bool
	}{
		{6, false},
		{7, true},
		{12, true},
		{18, true},
		{20, false},
	}
	for _, tt := range tests {
		v := time.Date(2026, 1, 1, tt.hour, 0, 0, 0, time.UTC)
		got := IsDay(v)
		if got != tt.want {
			t.Fatalf("got: %t, want %t for %s", got, tt.want, v)
		}
	}
}
