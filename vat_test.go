package shop

import "testing"

func TestWithVAT(t *testing.T) {
	if got := WithVAT(1000); got != 1200 {
		t.Fatalf("WithVAT(1000) = %d, want 1200", got)
	}
}
