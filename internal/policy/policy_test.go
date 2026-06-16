package policy

import "testing"

func TestDecideEscalatesToStop(t *testing.T) {
	c := DefaultConfig()
	tier, matched := c.Decide(1200)
	if !matched {
		t.Fatal("expected a rule to match at $1200/day")
	}
	if tier != TierStop {
		t.Fatalf("expected stop tier, got %q", tier)
	}
}

func TestDecideWithinBudget(t *testing.T) {
	c := DefaultConfig()
	if _, matched := c.Decide(10); matched {
		t.Fatal("did not expect any rule to match at $10/day")
	}
}
