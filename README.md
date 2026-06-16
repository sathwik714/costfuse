# costfuse

![CI](https://github.com/sathwik714/costfuse/actions/workflows/ci.yml/badge.svg)

**A circuit breaker for cloud spend.** costfuse watches how fast your cloud
bill is growing and trips a hard stop *before* a runaway loop turns into a
five-figure invoice — not hours later when the bill finally updates.

> A misconfigured loop has been enough to run up a $34,000 bill in eight days
> with zero users. Cloud providers give you budget *alerts*, which fire after
> the money is already gone. costfuse is the thing that actually pulls the plug.

> ⚠️ Early development (v0). Today it runs the full pipeline on stub data in
> dry-run mode. The real GCP integration is being built in the open — follow
> along in the issues.

## Quickstart

```bash
git clone https://github.com/sathwik714/costfuse.git
cd costfuse
make run        # runs a single pass in dry-run mode (safe — touches nothing)
```

You'll see costfuse estimate a spend rate, decide an action tier, and log
exactly what it *would* do. Nothing is ever stopped in dry-run mode.

## How it works

costfuse is a pipeline. A signal flows down it; if spend is accelerating past
your policy, the action loops back to the cloud and stops it.

```
Cloud environment ──> Collectors ──> Velocity engine ──> Policy engine ──> Action executor ──> Notifier
       ^                                                                          │
       └──────────────────────────  hard stop  ───────────────────────────────────┘
```

- **Collectors** pull *leading* signals (running instances, egress rate,
  function invocations) — not the lagging bill. This is what makes reacting in
  minutes possible.
- **Velocity engine** turns those signals into a spend rate ($/hr) and a
  projection of daily spend.
- **Policy engine** maps the projection to an escalating tier
  (`notify -> throttle -> stop`) and guarantees protected resources are
  alert-only.
- **Action executor** carries it out, with a mandatory dry-run mode and an
  audit log of every decision.
- **Notifier** pages a human (Slack, PagerDuty, webhook).

## Safety first

- **Dry-run is the default.** Arming real actions is an explicit choice.
- **Protected resources are never auto-stopped.** Tag production and costfuse
  will only ever alert on it.
- **Everything is audited.** Every decision is logged with the reason it fired.

## Open core

The agent in this repo is free and self-hostable (Apache-2.0). A hosted control
plane — multi-account dashboards, ML anomaly detection, RBAC/SSO, per-team
budgets, audit reports — is planned as the paid layer on top.

## Roadmap

- [ ] v0.1 — GCP leading-signal collector + budget Pub/Sub listener (alert-only)
- [ ] v0.2 — billing-disable action behind a confirmation flag
- [ ] v0.3 — YAML policy loader + protected-resource matching
- [ ] v0.4 — Slack and PagerDuty notifiers
- [ ] v0.5 — AI-API token proxy for real-time LLM spend

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). Issues labeled `good first issue` are a
great place to start.

## License

Apache-2.0. See [LICENSE](LICENSE).
