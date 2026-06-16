// Package collector pulls leading spend signals (instance counts, egress
// rate, function invocations) from a cloud environment.
//
// Leading signals, NOT the billing API, are what let costfuse react in
// minutes instead of hours. The bill lags by hours; a runaway loop does its
// damage long before the invoice updates.
package collector

// Signal is one leading indicator sampled from the cloud environment.
type Signal struct {
	Name  string  // e.g. "running_instances", "egress_bytes_per_sec"
	Value float64 // current sampled value
	Unit  string  // e.g. "count", "B/s"
}

// Collector pulls the current set of leading signals.
type Collector interface {
	Collect() ([]Signal, error)
}

// StubGCPCollector returns fixed sample data so the pipeline runs end to end
// before the real GCP Cloud Monitoring + Billing integration is wired in.
type StubGCPCollector struct{}

// NewStubGCPCollector returns a collector that emits canned sample signals.
func NewStubGCPCollector() *StubGCPCollector { return &StubGCPCollector{} }

// Collect returns sample signals.
//
// TODO: replace with real Cloud Monitoring time-series queries and a Pub/Sub
// listener for budget threshold events.
func (c *StubGCPCollector) Collect() ([]Signal, error) {
	return []Signal{
		{Name: "running_instances", Value: 12, Unit: "count"},
		{Name: "egress_bytes_per_sec", Value: 4.5e6, Unit: "B/s"},
	}, nil
}
