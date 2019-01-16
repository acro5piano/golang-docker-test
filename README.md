# golang-docker-test

test binary docker image


# Build

```
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w"
docker build -t acro5piano/golang-docker-test .
docker push acro5piano/golang-docker-test
```
