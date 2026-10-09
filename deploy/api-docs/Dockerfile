# Serves the quotes API reference (openapi.yaml) with Swagger UI.
# Build from the repo root: docker build -f deploy/api-docs/Dockerfile .
FROM swaggerapi/swagger-ui:v5.17.14

ARG GIT_SHA=unknown
LABEL application_name="backend-api-docs"
LABEL description="Swagger UI for the pantalasa quotes API"
LABEL owner="earthly"
LABEL source_uri="https://github.com/earthly/pantalasa"
LABEL git_sha="${GIT_SHA}"

COPY openapi.yaml /api/openapi.yaml
ENV SWAGGER_JSON=/api/openapi.yaml
EXPOSE 8080
