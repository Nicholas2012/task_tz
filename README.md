# Test service

Service is a simple API which returns a random user data.

## Run

```bash
docker-compose up -d
```

## Test

```bash
go test ./...
```

For integration tests you need set `INTEGRATION_TEST` environment variable:

```bash
INTEGRATION_TEST=true go test ./...
```

Integration tests use local docker.

## Curl

Random user

```bash
curl -X GET http://localhost:8080/random
```

List of users

```bash
curl -X GET http://localhost:8080/list
```
