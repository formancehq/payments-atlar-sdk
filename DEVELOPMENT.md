Download the latest API file from https://cdn.atlar.com/api/payments.oas30.yaml.

This client has been generated using the following commands

```sh
mkdir -p client
docker run --rm -w $PWD -v $PWD:$PWD hidori/oapi-codegen -generate types,client -package client ./2024-11-08_651fdb9c4db7e30048abd7a7_openapi-3-0.json > client/rest_client.go

cd .. && go mod tidy
```

The link to the current API specs can be obtained from https://docs.atlar.com/reference in the future.
