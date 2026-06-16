// Package executor performs the chosen action. In dry-run mode (the default)
// it logs what it would do without touching any resources, and it always
// writes an audit line for every decision so there is a record of what fired
// and why.
package executor

import (
	"log"

	"github.com/sathwik714/costfuse/internal/policy"
)

// Decision is what the policy engine resolved to.
type Decision struct {
	Tier   policy.Tier
	Reason string
}

// Executor carries out decisions against the cloud environment.
type Executor struct {
	DryRun bool
	logger *log.Logger
}

// New returns an executor. When dryRun is true, no real action is taken.
func New(dryRun bool, logger *log.Logger) *Executor {
	return &Executor{DryRun: dryRun, logger: logger}
}

// Execute performs (or, in dry-run, simulates) the action and audits it.
func (e *Executor) Execute(d Decision) error {
	e.logger.Printf("AUDIT tier=%s reason=%q dry_run=%t", d.Tier, d.Reason, e.DryRun)
	if e.DryRun {
		e.logger.Printf("[dry-run] would apply tier %q", d.Tier)
		return nil
	}
	// TODO: real actions, e.g. disable project billing, scale a managed
	// instance group to zero, or revoke an API key. Honour the policy's
	// ProtectedResources here before ever stopping anything.
	e.logger.Printf("applying tier %q", d.Tier)
	return nil
}
