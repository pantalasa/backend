# notifier

Outbound webhook dispatcher. Picks up quote-match events from the backend
service and POSTs them to customer-registered webhook URLs.

Runs as its own container so it can be scaled independently of the
quotes API and take its own retry/backoff configuration.

## Usage

```bash
docker build -t notifier:latest .
docker run --rm -e NOTIFIER_QUEUE_URL=... notifier:latest
```
