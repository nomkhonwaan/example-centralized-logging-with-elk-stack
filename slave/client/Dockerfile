FROM golang:1.7-alpine

ENV GOPATH=/go
WORKDIR $GOPATH/src/github.com/nomkhonwaan/stresser
ADD ./config $GOPATH/src/github.com/nomkhonwaan/stresser/config
ADD ./main.go $GOPATH/src/github.com/nomkhonwaan/stresser

# Install package dependencies
RUN apk -u --no-cache add git && \
    go get -u github.com/spf13/viper && \
    go build -o $GOPATH/bin/stresser main.go

ENTRYPOINT ["/go/bin/stresser"]