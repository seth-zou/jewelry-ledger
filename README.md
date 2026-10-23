# jewelry-ledger

This repository contains Go chaincode domain logic and Fabric network governance templates for the jewelry blockchain tracking system.

## Local Verification

```bash
go test ./...
go vet ./...
```

## Boundaries

The current chaincode skeleton validates asset registration, trace event appending, and state lookup in pure Go domain code. It is not production-ready Fabric deployment code.

## License

No license has been specified yet.
