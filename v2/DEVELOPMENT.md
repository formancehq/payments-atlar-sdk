The file `2024-11-08_651fdb9c4db7e30048abd7a7_openapi-3-1.json` was downloaded from the Atlar API docs. And renamed into what it is now in order to be able to better keep track of it.

This client has been generated using the following commands

```sh
# manual step: Pretty style the json document for nicer error messages

npx @apiture/openapi-down-convert --input 2024-11-08_651fdb9c4db7e30048abd7a7_openapi-3-1.json --output 2024-11-08_651fdb9c4db7e30048abd7a7_openapi-3-0.json

# manual step: Fix all the issues with the files like duplicate declaration of type names. Rerun the next command until no errors are being produced anymore.

docker run --rm -it --user $(id -u):$(id -g) -v ".:/workspace" -v "./ogen.yml:/ogen.yml" ghcr.io/ogen-go/ogen:latest --target workspace/api workspace/2024-11-08_651fdb9c4db7e30048abd7a7_openapi-3-0.json

# BEGIN Alternative client generator that did not work
# mkdir -p client
# docker run --rm -w $PWD -v $PWD:$PWD hidori/oapi-codegen -generate types,client -package client ./2024-11-08_651fdb9c4db7e30048abd7a7_openapi-3-0.json > client/rest_client.go
# END Alternative client generator that did not work

go mod tidy
```

The link to the current API specs can be obtained from https://docs.atlar.com/reference.
