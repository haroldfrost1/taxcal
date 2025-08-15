# taxcal

A Tax Calculator on the commnad-line

# Try it!
```
docker run -it haroldfrost/taxcal:latest
```
 
# Assumptions & adjustment

- Edge cases like int overflowing have not been considered / addressed
- Slightly adjustment the output to be a table with more info

# Decisions

- Go
  - I always find it very pleasant to read and trace other people's code in Go
  - it's compiled to a binary, easy and portable to test
- Tax Rates as a file
  - I'm aware that IRL this'll be from a DB and, from my experience in HSBC, this usually will be updated through a migration and pipeline under heavy review and regulation. But the main idea is decouples the data / configuration from the code
- Tests
  - For ease of implementation, I re-used the tax_rates.json as the test data, which I think is fine because real tax rates are not subject to change
- Docker
  - Used golang:tip-alpine3.22 for 0 security vulns
- Buildkite
  - I took the liberty to add a CI/CD pipeline. Please see https://buildkite.com/harold-frost/taxcal-cicd

# Development

## Prerequisites

- Go 1.21 or higher
- Docker (optional)

## Building and Running

### Local Development

1. Build the binary:

   ```bash
   make build
   ```

2. Run the application:
   ```bash
   ./taxcal
   ```

### Docker

1. Build the Docker image:

   ```bash
   make docker-build
   ```

2. Run with Docker:
   ```bash
   make docker-run
   ```

## Testing

Run the test suite:

```
go test ./...
```
