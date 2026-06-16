# Architecture

costfuse is a single pipeline. Read it top to bottom; the action loops back to
the start.

```
Cloud environment ──> Collectors ──> Velocity engine ──> Policy engine ──> Action executor ──> Notifier
       ^                                                                          │
       └──────────────────────────  hard stop  ───────────────────────────────────┘
```

## The one idea that matters: leading vs lagging signals

Cloud billing data lags by hours. If you watch the bill, you will always be too
late to stop a runaway loop. costfuse watches *leading* signals instead —
running instance count, egress rate, function invocation rate, autoscaler
events — which move the instant something goes wrong. That is the difference
between reacting in minutes and finding out from the invoice.

## Packages

| Package              | Job                                                          |
|----------------------|--------------------------------------------------------------|
| `internal/collector` | Pull leading signals from the cloud.                         |
| `internal/velocity`  | Convert signals into spend rate + projected daily spend.     |
| `internal/policy`    | Decide the action tier; protect production from auto-stop.   |
| `internal/executor`  | Perform the action; dry-run by default; audit every call.    |
| `internal/notifier`  | Alert a human (Slack, PagerDuty, webhook).                   |
| `cmd/costfuse`       | Wire the packages together. No business logic lives here.    |

## Build order

1. **v0.1** — real GCP collector (Cloud Monitoring time-series + a Pub/Sub
   listener for budget events). Alert-only. Useful and totally safe.
2. **v0.2** — the billing-disable action, behind a confirmation flag.
3. **v0.3** — load policy from YAML; implement protected-resource matching.
4. **v0.4** — Slack and PagerDuty notifiers.
5. **v0.5** — optional AI-API token proxy for real-time LLM spend.

Ship each step publicly before starting the next — real feedback should steer
what comes after.
