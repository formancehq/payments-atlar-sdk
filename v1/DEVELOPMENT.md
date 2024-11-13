This client has been generated using the following commands

```sh
docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger generate client -f https://cdn.atlar.com/api/swagger.json -A .

rm ./go.mod ./go.sum

cd .. && go mod tidy # unlike pointed out by previous command
```
