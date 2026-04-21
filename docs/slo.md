---
service: "backend"
tier: 1
last_reviewed: "2026-03-10"
next_review: "2026-09-10"
owner: "shane@pantalasa.org"
---

# Service Level Objectives — Backend API

## Availability

| Window | Target | Error Budget |
|--------|--------|--------------|
| 30-day rolling | 99.9% | 43m 12s |
| 90-day rolling | 99.9% | 2h 9m 36s |

## Latency

| Metric | Target |
|--------|--------|
| p50 request latency | < 80 ms |
| p99 request latency | < 500 ms |

## Error Rate

| Metric | Target |
|--------|--------|
| 5xx response rate | < 0.5% over 5-minute window |
| 4xx response rate | < 2.0% over 5-minute window (informational) |

## Error Budget Policy

- If 50% of the 30-day error budget is consumed, the team triggers a focused reliability review within the next sprint.
- If 75% is consumed, all non-reliability feature work pauses until budget recovers or a focused reliability effort closes the gap.
- If 100% is consumed before the window resets, an incident commander is assigned to drive a dedicated reliability push until budget is restored.

## Review Cadence

SLOs are reviewed every 6 months or after any major architectural change. Last review: 2026-03-10 by shane@pantalasa.org.
