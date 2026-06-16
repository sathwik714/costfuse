// Package policy decides which action tier to take based on the spend
// estimate, and guarantees that resources marked protected are never
// auto-stopped (alert-only).
package policy

// Tier is an escalating response.
type Tier string

const (
	TierNotify   Tier = "notify"
	TierThrottle Tier = "throttle"
	TierStop     Tier = "stop"
)

// Rule fires when projected daily spend crosses a threshold.
type Rule struct {
	Name         string
	OverDailyUSD float64
	Action       Tier
}

// Config is the loaded policy. ProtectedResources are alert-only and must
// never be auto-stopped, even if a rule would otherwise trip.
type Config struct {
	ProtectedResources []string
	Rules              []Rule
}

// DefaultConfig is used until the YAML loader lands.
// See configs/policy.example.yaml for the format it will read.
func DefaultConfig() Config {
	return Config{
		ProtectedResources: []string{"prod-*"},
		Rules: []Rule{
			{Name: "warn", OverDailyUSD: 100, Action: TierNotify},
			{Name: "throttle", OverDailyUSD: 500, Action: TierThrottle},
			{Name: "trip", OverDailyUSD: 1000, Action: TierStop},
		},
	}
}

// Decide returns the action for the highest threshold the projection exceeds.
// The bool is false when spend is within budget and no rule applies.
func (c Config) Decide(projectedDaily float64) (Tier, bool) {
	var chosen Tier
	var matched bool
	var highest float64
	for _, r := range c.Rules {
		if projectedDaily >= r.OverDailyUSD && r.OverDailyUSD >= highest {
			highest = r.OverDailyUSD
			chosen = r.Action
			matched = true
		}
	}
	return chosen, matched
}
