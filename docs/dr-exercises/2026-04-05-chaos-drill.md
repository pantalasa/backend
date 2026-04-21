---
date: "2026-04-05"
exercise_type: "chaos-drill"
facilitator: "shane@pantalasa.org"
participants: 4
outcome: "successful"
---

# DR Exercise — Backend API Chaos Drill

## Scenario

Scheduled chaos engineering exercise on the backend API. The facilitator terminated 50% of backend pods in the primary region at 13:00 UTC while traffic was at typical daytime volume. Team was not given advance warning of the exact timing.

## Recovery Steps Tested

1. Detection latency (kubelet → readiness probe → load balancer removal → alert).
2. Horizontal pod autoscaler response to reduced capacity.
3. Error rate observed from the customer perspective during recovery.
4. Manual escalation paths if HPA had failed to recover within 10 minutes.

## Participants

- Facilitator: shane@pantalasa.org
- On-call engineer: dane@pantalasa.org
- SRE observer: mick@pantalasa.org
- Scribe: alma@pantalasa.org

## Timing

- **Chaos action:** 13:00 UTC (50% of pods terminated).
- **First alert fired:** 13:01 UTC.
- **Capacity fully restored via HPA:** 13:06 UTC.
- **Error rate back to baseline:** 13:07 UTC.
- **Total customer-visible impact:** ~6 minutes of elevated error rate (~2% vs baseline 0.3%).

## Action Items

All closed as of 2026-04-18:

- [x] Tune HPA cool-down period (was 30s, reduced to 15s for faster response).
- [x] Add a CPU-based scaling metric alongside the existing request-count metric.
- [x] Update the backend runbook to document expected HPA behavior during partial pod loss.

## Lessons Learned

- Customer-facing error rate recovery was within the 30-minute RTO target, well before HPA's worst-case projection.
- The alert for partial-capacity situations was noisy — several duplicate pages before we suppressed them. Follow-up: implement alert deduplication at the PagerDuty level.
