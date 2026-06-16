// Package velocity converts raw leading signals into a spend-rate estimate
// and a projection of how fast the budget is being consumed.
package velocity

import "github.com/sathwik714/costfuse/internal/collector"

// Estimate is the velocity engine's output.
type Estimate struct {
	SpendRatePerHour float64
	ProjectedDaily   float64
}

// hourlyCost is a crude per-signal cost model. Replace with real cloud
// pricing (per-SKU rates) as the project matures.
var hourlyCost = map[string]float64{
	"running_instances":    0.50,              // ~$0.50 per instance-hour
	"egress_bytes_per_sec": 0.09 / 1e9 * 3600, // ~$0.09 per GB of egress
}

// Compute estimates current spend rate from the latest signals.
func Compute(signals []collector.Signal) Estimate {
	var perHour float64
	for _, s := range signals {
		if w, ok := hourlyCost[s.Name]; ok {
			perHour += s.Value * w
		}
	}
	return Estimate{
		SpendRatePerHour: perHour,
		ProjectedDaily:   perHour * 24,
	}
}
