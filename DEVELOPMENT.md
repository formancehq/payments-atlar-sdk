Download the latest API file from https://cdn.atlar.com/api/payments.oas30.yaml.

This client has been generated using the following commands

```sh
mkdir -p client
docker run --rm -w $PWD -v $PWD:$PWD hidori/oapi-codegen -generate types,client -package client ./2024-11-08_651fdb9c4db7e30048abd7a7_openapi-3-0.json > client/rest_client.go

cd .. && go mod tidy
```

The link to the current API specs can be obtained from https://docs.atlar.com/reference in the future.

## Go compiler

The SDK keeps its Go 1.23.2 minimum. The `toolchain` directive selects Go
1.27.3 or newer for work in this module when `GOTOOLCHAIN=auto`. CI uses
Go 1.27.3 with `GOTOOLCHAIN=local` and verifies the compiler before building.

```sh
go version
go env GOTOOLCHAIN GOVERSION GOROOT
go build -mod=readonly ./...
go test -mod=readonly -race ./...
go vet -mod=readonly ./...
```
