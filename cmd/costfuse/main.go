// Command costfuse watches cloud spend velocity and trips a circuit breaker
// before a runaway bill happens.
//
// This is the v0 skeleton: it runs a single pass with stub data so the whole
// pipeline (collect -> estimate -> decide -> act -> notify) works end to end.
// Real GCP integration gets dropped into the collector and executor next.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sathwik714/costfuse/internal/collector"
	"github.com/sathwik714/costfuse/internal/executor"
	"github.com/sathwik714/costfuse/internal/notifier"
	"github.com/sathwik714/costfuse/internal/policy"
	"github.com/sathwik714/costfuse/internal/velocity"
)

func main() {
	dryRun := flag.Bool("dry-run", true, "report actions without performing them")
	flag.Parse()

	logger := log.New(os.Stdout, "costfuse ", log.LstdFlags)

	// 1. Collect leading signals from the cloud environment.
	col := collector.NewStubGCPCollector()
	signals, err := col.Collect()
	if err != nil {
		logger.Fatalf("collect: %v", err)
	}

	// 2. Turn raw signals into a spend-rate estimate.
	est := velocity.Compute(signals)
	logger.Printf("spend rate: $%.2f/hr, projected daily: $%.2f",
		est.SpendRatePerHour, est.ProjectedDaily)

	// 3. Decide which action tier (if any) the policy calls for.
	cfg := policy.DefaultConfig()
	tier, matched := cfg.Decide(est.ProjectedDaily)
	if !matched {
		logger.Println("within budget, nothing to do")
		return
	}

	// 4. Act (safely, dry-run by default) and notify a human.
	exec := executor.New(*dryRun, logger)
	dec := executor.Decision{
		Tier:   tier,
		Reason: fmt.Sprintf("projected daily $%.2f crossed a policy threshold", est.ProjectedDaily),
	}
	if err := exec.Execute(dec); err != nil {
		logger.Fatalf("execute: %v", err)
	}

	n := notifier.NewLogNotifier(logger)
	_ = n.Notify(fmt.Sprintf("tier=%s: %s", dec.Tier, dec.Reason))
}
