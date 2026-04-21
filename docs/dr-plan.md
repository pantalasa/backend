---
rto_minutes: 30
rpo_minutes: 5
last_reviewed: "2026-02-28"
approver: "shane@pantalasa.org"
tier: 1
extends: "github.com/pantalasa/compliance-docs/docs/dr-plan.md"
---

# Disaster Recovery Plan — Backend API

This plan is **service-scoped** and tightens the organization-wide DR plan for the Backend API service. Backend is an internet-accessible tier-1 service, so it commits to stricter recovery objectives than the org baseline.

See the [organization DR plan](https://github.com/pantalasa/compliance-docs/blob/main/docs/dr-plan.md) for cross-service procedures (cross-region failover, company-wide incident comms, etc.). This document covers only what is specific to the Backend API.

## Recovery Objectives

| Objective | Backend API | Organization baseline |
|-----------|-------------|-----------------------|
| Recovery Time Objective (RTO) | **30 minutes** | 60 minutes |
| Recovery Point Objective (RPO) | **5 minutes** | 15 minutes |

The tighter backend numbers reflect the customer-facing nature of the API and the revenue impact of downtime.

## Scope

This plan covers the Backend API service only, including:

- The `backend` deployment in the primary region (`us-east-1`).
- The backend-owned PostgreSQL primary + replicas.
- The backend-specific Redis cache.

Cross-region failover, shared infrastructure, and company-wide incident processes are governed by the organization DR plan linked above.

## Recovery Steps

1. **Detect** — PagerDuty alert fires on error-rate spike (>1% over 2 min) or health check failure. On-call engineer pages.
2. **Assess** — Within 5 minutes, the on-call confirms the scope via the backend dashboards and determines whether to activate DR or attempt an in-place fix.
3. **Activate** — If DR is warranted, follow the failover runbook: promote the read replica, update DNS via automation, drain the failed primary.
4. **Verify** — Confirm service health via synthetic checks hitting the customer-facing API surface. All four synthetic probes must pass for 5 consecutive minutes before declaring recovery.
5. **Communicate** — Update the status page and notify stakeholders via the `#incidents` Slack channel. For customer-visible impact, trigger the customer-comms template referenced in the org incident-response policy.

## Dependencies

- PostgreSQL (primary + read replica, managed by the platform team)
- Redis cache
- DNS automation tooling
- PagerDuty escalation to `cronos-backend` schedule

## Review Cadence

Reviewed twice yearly or after any backend DR exercise or real incident. Last review: 2026-02-28 by shane@pantalasa.org.
